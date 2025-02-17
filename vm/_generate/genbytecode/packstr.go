// Copyright (C) 2023 Sneller, Inc.
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
	"math/rand"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

// packstrings finds a single string that contains all strings
// from the set.
//
// The goal is to represent the strings as a slices to that
// final string.  The procedure finds the string which is
// shorter than trivial concatenation of all inputs.
func packstrings(set map[string]struct{}) string {
	uniq := make([]string, 0, len(set))
	n := 0
	for s := range set {
		uniq = append(uniq, s)
		n += len(s)
	}

	sort.Strings(uniq)
	const tries = 35

	// Run packing with different input ordering, as it
	// makes difference from time to time. The value
	// of tries was set experimantally.
	s := rand.NewSource(int64(n))
	r := rand.New(s)
	var compacted string
	for i := 0; i < tries; i++ {
		tmp := packstringsaux(slices.Clone(uniq))
		if len(tmp) < n {
			compacted = tmp
			n = len(tmp)
		}

		shuffle(uniq, r)
	}

	return compacted
}

func packstringsaux(uniq []string) string {
	compacted := ""
	for len(uniq) > 0 {
		bestStr := ""
		bestCost := 0

		for _, word := range uniq {
			ss := allsubstrings(word)
			for _, s := range ss {
				tmp := compacted + s
				cost := costfunc(tmp, uniq)
				if cost > bestCost {
					bestCost = cost
					bestStr = tmp
				}

				tmp = s + compacted
				cost = costfunc(tmp, uniq)
				if cost > bestCost {
					bestCost = cost
					bestStr = tmp
				}
			}
		}

		if bestCost == 0 {
			break
		}

		compacted = bestStr
		uniq = prune(bestStr, uniq)
	}

	return compacted
}

func prune(compacted string, set []string) []string {
	tmp := []string{}
	for _, s := range set {
		if !strings.Contains(compacted, s) {
			tmp = append(tmp, s)
		}
	}

	return tmp
}

func costfunc(compacted string, set []string) int {
	n := 0
	for _, s := range set {
		if strings.Contains(compacted, s) {
			n += 1
		}
	}

	return n
}

func allsubstrings(s string) []string {
	n := len(s)
	c := n * n / 2
	if c == 0 {
		c = n
	}
	r := make([]string, 0, c)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			r = append(r, s[i:j+1])
		}
	}

	return r
}
