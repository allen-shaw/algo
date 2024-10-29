package started

import (
	"fmt"
	"testing"
)

type emptyStruct struct {
	val int
}

// go语言允许使用结构体对象作为key，只要所有字段都是可比较的。没有限制整个对象时不可变的
// 如果修改了对象中的成员，会导致key失效，找不到原来的value
func TestHashMap(t *testing.T) {
	m := make(map[emptyStruct]int, 0)

	a := emptyStruct{val: 1}
	b := emptyStruct{val: 2}

	m[a] = 1
	m[b] = 2
	fmt.Println(m[a], m[b])

	a.val = 100
	fmt.Println(m[a], m[b])
}

func TestHashMapRange(t *testing.T) {
	m := make(map[string]int)

	for i := 0; i < 5; i++ {
		m[fmt.Sprintf("key-%v", i)] = i
	}

	fmt.Println(m)

	i := 0
	for k, v := range m {
		if i < 5 {
			m[fmt.Sprintf("key-%v", 10+i)] = 10 + i
			i++
		}
		fmt.Println(k, v, m)
	}
}

func TestRange(t *testing.T) {
	for i := range 10 {
		fmt.Println(i)
	}
}
