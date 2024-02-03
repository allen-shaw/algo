package doubleheap

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// meeting[start, end]
func mostBooked(n int, meetings [][]int) int {
	// 按照开始时间排序
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	idle := &hp{make([]*pair, 0, n)}
	using := &hp{}

	for i := 0; i < n; i++ {
		heap.Push(idle, &pair{i: i})
	}

	ans := make([]int, n)
	for idx, meeting := range meetings {
		// 回收已经完成的会议室
		for len(using.data) > 0 && using.data[0].end <= meeting[0] {
			p := heap.Pop(using).(*pair)
			p.end = 0
			heap.Push(idle, p)
		}

		if len(idle.data) > 0 {
			p := heap.Pop(idle).(*pair)
			p.end = meeting[1]
			heap.Push(using, p)
			ans[p.i]++
		} else {
			if len(using.data) == 0 {
				fmt.Println(idx, using.data, idle.data)
			}
			p := heap.Pop(using).(*pair)
			mstart, mend := meeting[0], meeting[1]
			p.end = p.end + (mend - mstart)
			heap.Push(using, p)
			ans[p.i]++
		}
	}

	mid, m := 0, 0
	for i, n := range ans {
		if n > m {
			m = n
			mid = i
		}
	}

	return mid
}

func TestMostBooked(t *testing.T) {
	n := 2
	meetings := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}, {8, 11}, {9, 12}}
	fmt.Println(mostBooked(n, meetings))
}

func assignTasks(servers []int, tasks []int) []int {
	idle := &hp{data: make([]*pair, 0, len(servers))}
	using := &hp{data: make([]*pair, 0)}

	for i, w := range servers {
		heap.Push(idle, &pair{i: i, w: w})
	}

	ans := make([]int, len(tasks))

	for tid, task := range tasks {
		// 回收using
		for len(using.data) > 0 {
			if using.data[0].end <= tid {
				p := heap.Pop(using).(*pair)
				p.end = 0
				heap.Push(idle, p)
			} else {
				break
			}
		}

		if len(idle.data) != 0 {
			srv := heap.Pop(idle).(*pair)
			heap.Push(using, &pair{i: srv.i, end: tid + task, w: srv.w})
			ans[tid] = srv.i
		} else {
			srv := heap.Pop(using).(*pair)
			heap.Push(using, &pair{i: srv.i, end: srv.end + task, w: srv.w})
			ans[tid] = srv.i
		}
	}

	return ans
}

func TestAssignTasks(t *testing.T) {
	servers := []int{3, 3, 2}
	tasks := []int{1, 2, 3, 2, 1, 2}
	ans := assignTasks(servers, tasks)
	fmt.Println(ans)
}

type pair struct{ i, end, w int }
type hp struct {
	data []*pair
}

// Len implements heap.Interface.
func (hp *hp) Len() int {
	return len(hp.data)
}

// Less implements heap.Interface.
func (hp *hp) Less(i int, j int) bool {
	if hp.data[i].end == hp.data[j].end {
		if hp.data[i].w == hp.data[j].w {
			return hp.data[i].i < hp.data[j].i
		}
		return hp.data[i].w < hp.data[j].w
	}

	return hp.data[i].end < hp.data[j].end
}

// Pop implements heap.Interface.
func (hp *hp) Pop() any {
	e := hp.data[len(hp.data)-1]
	hp.data = hp.data[:len(hp.data)-1]
	return e
}

// Push implements heap.Interface.
func (hp *hp) Push(x any) {
	e := x.(*pair)
	hp.data = append(hp.data, e)
}

// Swap implements heap.Interface.
func (hp *hp) Swap(i int, j int) {
	hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}

// func busiestServers(k int, arrival []int, load []int) []int {

// }
