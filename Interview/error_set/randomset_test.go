package errorset

import (
	"fmt"
	"math/rand"
	"testing"
)

const cap = 20000

type RandomizedSet struct {
	m    map[int]int // val -> index
	arr  []int
	size int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    make(map[int]int),
		arr:  make([]int, cap),
		size: 0,
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.m[val]
	if ok {
		return false
	}

	idx := rs.size
	rs.m[val] = idx
	rs.arr[idx] = val

	rs.size++
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.m[val]
	if !ok {
		return false
	}
	last := rs.size - 1
	rs.arr[last], rs.arr[idx] = rs.arr[idx], rs.arr[last]
	rs.m[rs.arr[idx]] = idx
	delete(rs.m, val)

	rs.size--
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	idx := rand.Int31n(int32(rs.size))
	return rs.arr[idx]
}

func Test_RandomizedSet(t *testing.T) {
	rs := Constructor()
	rs.Insert(0)
	rs.Insert(1)
	rs.Remove(0)
	rs.Insert(2)
	rs.Remove(1)
	fmt.Println(rs.GetRandom())
}

func Test_RandomizedSet2(t *testing.T) {
	rs := Constructor()
	rs.Remove(0)
	rs.Remove(0)
	rs.Insert(0)
	fmt.Println(rs.GetRandom())
	rs.Remove(0)
	fmt.Println(rs.Insert(0))
}
