// Copyright 2014 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package btree

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/google/btree"
)

func init() {
	seed := time.Now().Unix()
	fmt.Println(seed)
	rand.Seed(seed)
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm(n int) (out []*IntItem) {
	for _, v := range rand.Perm(n) {
		out = append(out, &IntItem{v: v})
	}
	return
}

// rang returns an ordered list of Int items in the range [0, n).
func rang(n int) (out []IntItem) {
	for i := 0; i < n; i++ {
		out = append(out, IntItem{v: i})
	}
	return
}

var _ btree.Item = &IntItem{}

type IntItem struct {
	v int
}

func (a *IntItem) Less(b btree.Item) bool {
	return a.v < b.(*IntItem).v
}

const benchmarkTreeSize = 10000

var btreeDegree = flag.Int("degree", 32, "B-Tree degree")

func BenchmarkInsertGoogle(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		tr := btree.New(*btreeDegree)
		for _, item := range insertP {
			tr.ReplaceOrInsert(btree.Item(item))
			i++
			if i >= b.N {
				return
			}
		}
	}
}

func BenchmarkSeekGoogle(b *testing.B) {
	b.StopTimer()
	size := 100000
	insertP := perm(size)
	tr := btree.New(*btreeDegree)
	for _, item := range insertP {
		tr.ReplaceOrInsert(item)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tr.AscendGreaterOrEqual(&IntItem{v: i % size}, func(i btree.Item) bool { return false })
	}
}

func BenchmarkDeleteGoogle(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	removeP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		b.StopTimer()
		tr := btree.New(*btreeDegree)
		for _, v := range insertP {
			tr.ReplaceOrInsert(v)
		}
		b.StartTimer()
		for _, item := range removeP {
			tr.Delete(item)
			i++
			if i >= b.N {
				return
			}
		}
		if tr.Len() > 0 {
			panic(tr.Len())
		}
	}
}

func BenchmarkGetGoogle(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	removeP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		b.StopTimer()
		tr := btree.New(*btreeDegree)
		for _, v := range insertP {
			tr.ReplaceOrInsert(v)
		}
		b.StartTimer()
		for _, item := range removeP {
			tr.Get(item)
			i++
			if i >= b.N {
				return
			}
		}
	}
}

type byInts []*IntItem

func (a byInts) Len() int {
	return len(a)
}

func (a byInts) Less(i, j int) bool {
	return a[i].v < a[j].v
}

func (a byInts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func BenchmarkAscendGreaterOrEqualGoogle(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := 100
		k := 0
		tr.AscendGreaterOrEqual(&IntItem{v: 100}, func(item btree.Item) bool {
			if item.(*IntItem).v != arr[j].v {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].v, item.(*IntItem).v)
			}
			j++
			k++
			return true
		})
		if j != len(arr) {
			b.Fatalf("expected: %v, got %v", len(arr), j)
		}
		if k != len(arr)-100 {
			b.Fatalf("expected: %v, got %v", len(arr)-100, k)
		}
	}
}

func BenchmarkDescendLessOrEqualGoogle(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := len(arr) - 100
		k := len(arr)
		tr.DescendLessOrEqual(arr[len(arr)-100], func(item btree.Item) bool {
			if item.(*IntItem).v != arr[j].v {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].v, item.(*IntItem).v)
			}
			j--
			k--
			return true
		})
		if j != -1 {
			b.Fatalf("expected: %v, got %v", -1, j)
		}
		if k != 99 {
			b.Fatalf("expected: %v, got %v", 99, k)
		}
	}
}
