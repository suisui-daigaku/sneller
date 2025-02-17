// Copyright (C) 2022 Sneller, Inc.
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	var outpath string
	var inpath string
	flag.StringVar(&inpath, "i", "", "input asm file path")
	flag.StringVar(&outpath, "o", "", "output header file path")
	flag.Parse()
	if outpath == "" || inpath == "" {
		flag.Usage()
		return
	}

	consts := newconstpool()
	parser := AsmParser{
		consts: consts,
	}
	parser.skipPaths(outpath)
	parser.skipPaths(systemincludes...)

	parser.addPath(inpath)
	err := parser.parseAll()
	checkErr(err)

	err = consts.postprocess()
	checkErr(err)

	buf := bytes.NewBuffer(nil)
	stdout = buf
	consts.print()
	old, _ := os.ReadFile(outpath)
	if !slices.Equal(old, buf.Bytes()) {
		fmt.Printf("Creating %q\n", outpath)
		err := os.WriteFile(outpath, buf.Bytes(), 0644)
		checkErr(err)
	}
}

type AsmParser struct {
	paths  []string            // stack of paths
	seen   map[string]struct{} // already seen paths
	consts *constpool
}

func (a *AsmParser) parseAll() error {
	for len(a.paths) > 0 {
		n := len(a.paths)
		path := a.paths[n-1]
		a.paths = a.paths[:n-1]

		err := a.parse(path)
		if err != nil {
			return fmt.Errorf("%q: %s", path, err)
		}
	}

	return nil
}

var systemincludes = []string{"go_asm.h", "funcdata.h", "textflag.h"}

func (a *AsmParser) skipPaths(paths ...string) {
	if a.seen == nil {
		a.seen = make(map[string]struct{})
	}

	for _, p := range paths {
		a.seen[p] = struct{}{}
	}
}

func (a *AsmParser) addPath(path string) {
	if a.seen == nil {
		a.seen = make(map[string]struct{})
	}

	if _, ok := a.seen[path]; ok {
		return
	}

	a.paths = append(a.paths, path)
	a.seen[path] = struct{}{}
}

var constpat = `CONST[^(_]+_[0-9A-Za-z_]+\(\)`
var constdef = fmt.Sprintf(`(%s)\s=\s([0-9a-zA-Z.e+()-]+)`, constpat)
var repat = regexp.MustCompile(constpat)
var redef = regexp.MustCompile(constdef)

func (a *AsmParser) parse(path string) error {
	f, err := os.Open(path)
	checkErr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if include, ok := strings.CutPrefix(line, "#include "); ok {
			n := len(include)
			if n > 2 && include[0] == '"' && include[n-1] == '"' {
				a.addPath(include[1 : n-1])
			} else {
				return fmt.Errorf("malformed include: %s", line)
			}
			continue
		}

		s := redef.FindStringSubmatch(line)
		if len(s) == 3 {
			name := s[1][:len(s[1])-2]
			value := s[2]
			err := a.consts.add(name, value)
			if err != nil {
				return err
			}
			continue
		}

		s = repat.FindStringSubmatch(line)
		for i := range s {
			token := s[i][:len(s[i])-2]
			err := a.consts.parse(token)
			if err != nil {
				a.consts.unresolved[token] = struct{}{}
			}
		}
	}

	return scanner.Err()
}

type constpool struct {
	// maps value -> list of constants

	qword      map[uint64][]string
	dword      map[uint32][]string
	bytes      map[byte][]string
	f64        map[uint64][]string
	f32        map[uint32][]string
	unresolved map[string]struct{}

	u64keys []uint64
	u32keys []uint32
	u8keys  []byte
	f64keys []uint64
	f32keys []uint32
}

func newconstpool() *constpool {
	return &constpool{
		qword:      make(map[uint64][]string),
		dword:      make(map[uint32][]string),
		bytes:      make(map[byte][]string),
		f64:        make(map[uint64][]string),
		f32:        make(map[uint32][]string),
		unresolved: make(map[string]struct{}),
	}
}

var stdout io.Writer

func (c *constpool) print() {
	offset := 0

	writeln(autogenerated)
	if len(c.qword) > 0 {
		offset = c.printu64(offset)
	}

	if len(c.dword) > 0 {
		offset = c.printu32(offset)
	}

	if len(c.bytes) > 0 {
		offset = c.printu8(offset)
	}

	if len(c.f32) > 0 {
		offset = c.printf32(offset)
	}

	if len(c.f64) > 0 {
		offset = c.printf64(offset)
	}

	writeln("")
	writeln("CONST_GLOBAL(constpool, $%d)", offset)
}

func define(names []string, offset int) {
	for _, name := range uniq(names) {
		writeln("#define %s() CONST_GET_PTR(constpool, %d)", name, offset)
	}
}

func (c *constpool) printu64(offset int) int {
	f64alias := func(u64 uint64) (uint64, int) {
		keys := maps.Keys(c.f64)
		if slices.Index(keys, u64) >= 0 {
			return u64, 0
		}

		return 0, -1
	}

	f32alias := func(u64 uint64) (uint32, int) {
		keys := maps.Keys(c.f32)
		key := uint32(u64)
		if slices.Index(keys, key) >= 0 {
			return key, 0
		}
		key = uint32(u64 >> 32)
		if slices.Index(keys, key) >= 0 {
			return key, 4
		}
		return 0, -1
	}

	u32alias := func(u64 uint64) (uint32, int) {
		keys := maps.Keys(c.dword)
		key := uint32(u64)
		if slices.Index(keys, key) >= 0 {
			return key, 0
		}
		key = uint32(u64 >> 32)
		if slices.Index(keys, key) >= 0 {
			return key, 4
		}
		return 0, -1
	}

	u8alias := func(u64 uint64) (byte, int) {
		keys := maps.Keys(c.bytes)
		for i := 0; i < 8; i++ {
			key := byte(u64)
			u64 >>= 8
			if slices.Index(keys, key) >= 0 {
				return key, i
			}
		}

		return 0, -1
	}

	writeln("")
	writeln("// uint64 constants")
	for i, u64 := range c.u64keys {
		if i > 0 {
			writeln("")
		}
		{
			u8, off := u8alias(u64)
			if off != -1 {
				define(c.bytes[u8], offset+off)
				delete(c.bytes, u8)
			}
		}
		{
			u32, off := u32alias(u64)
			if off != -1 {
				define(c.dword[u32], offset+off)
				delete(c.dword, u32)
			}
		}
		{
			u32, off := f32alias(u64)
			if off != -1 {
				define(c.f32[u32], offset+off)
				delete(c.f32, u32)
			}
		}
		{
			f64key, off := f64alias(u64)
			if off != -1 {
				define(c.f64[f64key], offset+off)
				delete(c.f64, f64key)
			}
		}

		u64def := c.qword[u64]
		define(u64def, offset)

		writeln("CONST_DATA_U64(constpool, %d, $%d) // 0x%016x", offset, u64, u64)
		offset += 8
	}

	return offset
}

func (c *constpool) printu32(offset int) int {
	u8alias := func(u32 uint32) (byte, int) {
		keys := maps.Keys(c.bytes)
		for i := 0; i < 4; i++ {
			key := byte(u32)
			u32 >>= 8
			if slices.Index(keys, key) >= 0 {
				return key, i
			}
		}

		return 0, -1
	}

	writeln("")
	writeln("// uint32 constants")
	first := true
	for _, u32 := range c.u32keys {
		if _, ok := c.dword[u32]; !ok {
			continue
		}
		if !first {
			writeln("")
		}
		first = false

		{
			u8, off := u8alias(u32)
			if off != -1 {
				define(c.bytes[u8], offset+off)
				delete(c.bytes, u8)
			}
		}

		define(c.dword[u32], offset)
		writeln("CONST_DATA_U32(constpool, %d, $%d) // 0x%08x", offset, u32, u32)

		offset += 4
	}
	return offset
}

func (c *constpool) printu8(offset int) int {
	writeln("")
	writeln("// uint8 constants")
	first := true
	for _, u8 := range c.u8keys {
		if _, ok := c.bytes[u8]; !ok {
			continue
		}
		if !first {
			writeln("")
		}
		first = false

		define(c.bytes[u8], offset)
		writeln("CONST_DATA_U8(constpool, %d, $%d) // 0x%02x", offset, u8, u8)

		offset += 1
	}
	return offset
}

func (c *constpool) printf32(offset int) int {
	writeln("")
	writeln("// float32 constants")
	first := true
	for _, u32 := range c.f32keys {
		if _, ok := c.f32[u32]; !ok {
			continue
		}
		if !first {
			writeln("")
		}
		first = false

		define(c.f32[u32], offset)
		writeln("CONST_DATA_U32(constpool, %d, $0x%016x) // float32(%f)",
			offset, u32, math.Float32frombits(u32))

		offset += 4
	}
	return offset
}

func (c *constpool) printf64(offset int) int {
	writeln("")
	writeln("// float64 constants")
	first := true
	for _, u64 := range c.f64keys {
		if _, ok := c.f64[u64]; !ok {
			continue
		}
		if !first {
			writeln("")
		}
		first = false

		define(c.f64[u64], offset)
		writeln("CONST_DATA_U64(constpool, %d, $0x%016x) // float64(%f)",
			offset, u64, math.Float64frombits(u64))

		offset += 8
	}
	return offset
}

func (c *constpool) addaux(name, typ, value string) error {
	switch typ {
	case "CONSTB":
		u64, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return err
		}
		if u64 >= 1<<8 {
			log.Fatalf("%d can't fit in u8", u64)
		}
		b := byte(u64)
		c.bytes[b] = append(c.bytes[b], name)

	case "CONSTD":
		u64, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return err
		}
		if u64 >= 1<<32 {
			log.Fatalf("%d can't fit in u32", u64)
		}
		u32 := uint32(u64)
		c.dword[u32] = append(c.dword[u32], name)

	case "CONSTQ":
		u64, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return err
		}
		c.qword[u64] = append(c.qword[u64], name)

	case "CONSTF32":
		u32, err := parsefloat32(value)
		if err != nil {
			return err
		}
		c.f32[u32] = append(c.f32[u32], name)

	case "CONSTF64":
		u64, err := parsefloat64(value)
		if err != nil {
			return err
		}
		c.f64[u64] = append(c.f64[u64], name)

	default:
		return fmt.Errorf("%q is unsupported", typ)
	}

	return nil
}

func parsefloat64(value string) (uint64, error) {
	var u64 uint64
	var f64 float64
	var err error

	u64, err = strconv.ParseUint(value, 0, 64)
	if err == nil {
		f64 = float64(u64)
		return math.Float64bits(f64), nil
	}

	f64, err = strconv.ParseFloat(value, 64)
	if err == nil {
		return math.Float64bits(f64), nil
	}

	val, ok := stripendings(value, "uint64(", ")")
	if !ok {
		return 0, err
	}

	return strconv.ParseUint(val, 0, 64)
}

func parsefloat32(value string) (uint32, error) {
	u64, err := strconv.ParseUint(value, 0, 32)
	if err == nil {
		return math.Float32bits(float32(u64)), nil
	}

	f64, err := strconv.ParseFloat(value, 32)
	if err == nil {
		return math.Float32bits(float32(f64)), nil
	}

	val, ok := stripendings(value, "uint32(", ")")
	if !ok {
		return 0, err
	}

	u64, err = strconv.ParseUint(val, 0, 32)
	if err != nil {
		return 0, err
	}

	if u64 > uint64(0xffffffff) {
		return 0, fmt.Errorf("uint32 const greater than the allowed maximum")
	}

	return uint32(u64), nil
}

func (c *constpool) add(macroname, value string) error {
	typ, _, ok := strings.Cut(macroname, "_")
	if !ok {
		return nil
	}

	return c.addaux(macroname, typ, value)
}

func (c *constpool) parse(macroname string) error {
	typ, value, ok := strings.Cut(macroname, "_")
	if !ok {
		return nil
	}

	return c.addaux(macroname, typ, value)
}

func (c *constpool) postprocess() error {
	// sort keys
	c.u64keys = maps.Keys(c.qword)
	slices.Sort(c.u64keys)

	c.u32keys = maps.Keys(c.dword)
	slices.Sort(c.u32keys)

	c.u8keys = maps.Keys(c.bytes)
	slices.Sort(c.u8keys)

	c.f64keys = maps.Keys(c.f64)
	slices.Sort(c.f64keys)

	c.f32keys = maps.Keys(c.f32)
	slices.Sort(c.f32keys)

	// find aliases
	unresolved := maps.Keys(c.unresolved)
	slices.Sort(unresolved)

	n := 0
	for _, name := range unresolved {
		if !c.resolve(name) {
			if n == 0 {
				fmt.Printf("Unresolved macro names:\n")
			}

			fmt.Printf("- %q\n", name)
			n += 1
		}
	}

	if n > 0 {
		return fmt.Errorf("%d unresolved macros", n)
	}
	return nil
}

func (c *constpool) resolve(name string) bool {
	typ, _, ok := strings.Cut(name, "_")
	if !ok {
		return false
	}

	switch typ {
	case "CONSTB":
		for _, u8 := range c.u8keys {
			if has(c.bytes[u8], name) {
				return true
			}
		}

	case "CONSTD":
		for _, u32 := range c.u32keys {
			if has(c.dword[u32], name) {
				return true
			}
		}

	case "CONSTQ":
		for _, u64 := range c.u64keys {
			if has(c.qword[u64], name) {
				return true
			}
		}

	case "CONSTF64":
		for _, u64 := range c.u64keys {
			if has(c.qword[u64], name) {
				return true
			}
		}
		for _, u64 := range c.f64keys {
			if has(c.f64[u64], name) {
				return true
			}
		}

	case "CONSTF32":
		for _, u32 := range c.u32keys {
			if has(c.dword[u32], name) {
				return true
			}
		}
		for _, u32 := range c.f32keys {
			if has(c.f32[u32], name) {
				return true
			}
		}

	default:
		log.Fatalf("const type %q not supported", typ)
	}

	return false
}

func has(list []string, needle string) bool {
	return slices.Index(list, needle) >= 0
}

func uniq(list []string) []string {
	slices.Sort(list)
	return slices.Compact(list)
}

func stripendings(s, prefix, suffix string) (string, bool) {
	s1, ok1 := strings.CutPrefix(s, prefix)
	if !ok1 {
		return "", false
	}

	return strings.CutSuffix(s1, suffix)
}

const autogenerated = "// Code generated automatically; DO NOT EDIT"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeln(s string, args ...any) {
	_, err := stdout.Write([]byte(fmt.Sprintf(s, args...)))
	checkErr(err)
	_, err = stdout.Write([]byte{'\n'})
	checkErr(err)
}
