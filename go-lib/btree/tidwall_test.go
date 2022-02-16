package btree

import (
	"sort"
	"testing"

	"github.com/tidwall/btree"
)

func ByIntItem(a, b interface{}) bool {
	i1, i2 := a.(*IntItem), b.(*IntItem)
	return i1.Less(i2)
}
func BenchmarkInsertTiD(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		tr := btree.New(ByIntItem)
		for _, item := range insertP {
			tr.Set(item)
			i++
			if i >= b.N {
				return
			}
		}
	}
}

func BenchmarkSeekTiD(b *testing.B) {
	b.StopTimer()
	size := 100000
	insertP := perm(size)
	tr := btree.New(ByIntItem)
	for _, item := range insertP {
		tr.Set(item)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tr.Ascend(&IntItem{v: i % size}, func(i interface{}) bool { return false })
	}
}

func BenchmarkDeletTiD(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	removeP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		b.StopTimer()
		tr := btree.New(ByIntItem)
		for _, v := range insertP {
			tr.Set(v)
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

func BenchmarkGetTiD(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	removeP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		b.StopTimer()
		tr := btree.New(ByIntItem)
		for _, v := range insertP {
			tr.Set(v)
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

func BenchmarkAscendGreaterOrEqualTiD(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(ByIntItem)
	for _, v := range arr {
		tr.Set(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := 100
		k := 0
		tr.Ascend(&IntItem{v: 100}, func(item interface{}) bool {
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

func BenchmarkDescendLessOrEqualTiD(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(ByIntItem)
	for _, v := range arr {
		tr.Set(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := len(arr) - 100
		k := len(arr)
		tr.Descend(arr[len(arr)-100], func(item interface{}) bool {
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
