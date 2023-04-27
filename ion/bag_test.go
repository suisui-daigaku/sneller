// Copyright 2023 Sneller, Inc.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package ion

import (
	"testing"
)

func TestBag(t *testing.T) {
	items := []Datum{
		Null,
		String("foo"),
		Int(-1),
		Uint(1000),
		Bool(true),
		Bool(false),
		NewStruct(nil,
			[]Field{
				{"foo", String("foo"), 0},
				{"bar", Null, 0},
				{"inner", NewList(nil, []Datum{
					Int(-1), Uint(0), Uint(1),
				}).Datum(), 0},
				{"name", String("should-come-first"), 0},
			},
		).Datum(),
	}

	var bag Bag
	for i := range items {
		bag.AddDatum(items[i])
	}
	if bag.Len() != len(items) {
		t.Fatalf("bag.Len=%d; expected %d", bag.Len(), len(items))
	}
	i := 0
	bag.Each(func(d Datum) bool {
		if !d.Equal(items[i]) {
			t.Errorf("item %d is %v", i, d)
		}
		i++
		return true
	})

	// transcode to a second symbol table
	var st Symtab
	for _, x := range []string{"baz", "bar", "foo", "quux"} {
		st.Intern(x)
	}
	var buf Buffer
	var bag2 Bag
	bag.Encode(&buf, &st)
	bag2.Add(&st, buf.Bytes())
	if !bag.Equals(&bag2) {
		t.Fatal("!bag.Equal(bag2)")
	}

	bag.Append(&bag2)
	if bag.Len() != len(items)*2 {
		t.Fatalf("bag.Len=%d, want %d", bag.Len(), len(items)*2)
	}
	i = 0
	n := 0
	bag.Each(func(d Datum) bool {
		if !d.Equal(items[i]) {
			t.Errorf("item %d is %v", i, d)
		}
		i++
		n++
		if i == len(items) {
			i = 0
		}
		return true
	})
	if n != bag.Len() {
		t.Fatalf("Each iterated %d times, but bag.Len()=%d", n, bag.Len())
	}

	bag.Reset()
	if bag.Len() != 0 {
		t.Fatalf("bag.Len = %d after reset?", bag.Len())
	}
	i = 0
	bag.Each(func(d Datum) bool {
		i++
		return true
	})
	if i > 0 {
		t.Fatalf("bag has contents (%d items) after reset?", i)
	}
	bag = bag2.Clone()
	if !bag.Equals(&bag2) {
		t.Errorf("cloned Bag not equal to itself")
	}

	var bag3 Bag
	buf.Reset()
	st.Reset()
	w := bag3.Writer()
	bag2.Encode(&buf, &st)
	stpos := buf.Size()
	st.Marshal(&buf, true)
	data := append(buf.Bytes()[stpos:], buf.Bytes()[:stpos]...)
	n, err := w.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(data) {
		t.Fatalf("ion.Bag.Writer().Write wrote %d instead of %d bytes", n, len(data))
	}
	if !bag3.Equals(&bag2) {
		t.Fatal("using bagWriter.Write did not produce an equivalent Bag")
	}
}
