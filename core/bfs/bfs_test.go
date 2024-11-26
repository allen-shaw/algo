package bfs

import (
	"fmt"
	"testing"
)

func openLock(deadends []string, target string) int {
	m := make(map[string]struct{})
	for _, d := range deadends {
		m[d] = struct{}{}
	}

	deep := 0
	q := []string{"0000"}
	visited := make(map[string]struct{})

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			str := q[0]
			q = q[1:]
			if str == target {
				return deep
			}
			if _, ok := visited[str]; ok {
				continue
			}
			if _, ok := m[str]; ok {
				continue
			}

			visited[str] = struct{}{}

			for j := 0; j < 4; j++ {
				sup := up(str, j)
				q = append(q, sup)
				sdown := down(str, j)
				q = append(q, sdown)
			}
		}
		deep++
	}

	return -1
}

func up(src string, index int) string {
	target := []byte(src)
	if target[index] == '9' {
		target[index] = '0'
	} else {
		target[index]++
	}
	return string(target)
}

func down(src string, index int) string {
	target := []byte(src)
	if target[index] == '0' {
		target[index] = '9'
	} else {
		target[index]--
	}
	return string(target)
}

func Test_openLock(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	out := openLock(deadends, target)
	fmt.Println(out)
}
