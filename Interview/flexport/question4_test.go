package flexport

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// 1.4 Meeting Room Scheduling Problem
// Given a set of meeting time intervals, determine the minimum number of
// conference rooms required to accommodate all the meetings.
// Each interval includes the start and end time of a meeting,
// and no two meetings can occur in the same room at the same time.
// Discuss your approach to solving this problem.

func canAttendMeetings(intervals [][]int) bool {
	// sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			return false
		}
	}
	return true
}

func Test_canAttendMeetings(t *testing.T) {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	ans := canAttendMeetings(intervals)
	fmt.Println(ans)
}

func minMeetingRooms(intervals [][]int) int {
	sortEndTime := intervals
	sortBeginTime := cloneMetrix(intervals)
	sort.Slice(sortBeginTime, func(i, j int) bool {
		return sortBeginTime[i][0] < sortBeginTime[j][0]
	})
	sort.Slice(sortEndTime, func(i, j int) bool {
		return sortEndTime[i][1] < sortEndTime[j][1]
	})

	fmt.Println(sortBeginTime, sortEndTime)

	n := len(intervals)
	maxCount, count := 0, 0
	beginIdx, endIdx := 0, 0
	for beginIdx < n && endIdx < n {
		if sortBeginTime[beginIdx][0] <= sortEndTime[endIdx][1] {
			count++
			beginIdx++
		} else {
			count--
			endIdx++
		}
		maxCount = max(maxCount, count)
	}

	return maxCount
}

func Test_minMeetingRooms(t *testing.T) {
	intervals := [][]int{{6, 17}, {8, 9}, {11, 12}, {6, 9}}
	ans := minMeetingRooms(intervals)
	fmt.Println(ans)
}

func cloneMetrix(src [][]int) [][]int {
	target := make([][]int, len(src))
	copy(target, src)
	return target
}

type UsingHeep [][]int

// Len implements heap.Interface.
func (u UsingHeep) Len() int {
	return len(u)
}

// Less implements heap.Interface.
func (u UsingHeep) Less(i int, j int) bool {
	return u[i][0] < u[j][0]
}

// Pop implements heap.Interface.
func (u *UsingHeep) Pop() any {
	x := (*u)[len(*u)-1]
	*u = (*u)[:len(*u)-1]
	return x
}

// Push implements heap.Interface.
func (u *UsingHeep) Push(x any) {
	*u = append(*u, x.([]int))
}

// Swap implements heap.Interface.
func (u UsingHeep) Swap(i int, j int) {
	u[i], u[j] = u[j], u[i]
}

// [endtime, id]
type IdleHeap []int

// Len implements heap.Interface.
func (h IdleHeap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h IdleHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Pop implements heap.Interface.
func (h *IdleHeap) Pop() any {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

// Push implements heap.Interface.
func (h *IdleHeap) Push(x any) {
	(*h) = append((*h), x.(int))
}

// Swap implements heap.Interface.
func (h IdleHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func mostBooked(n int, meetings [][]int) int {
	rooms := make([]int, n)
	using := UsingHeep{}
	idle := IdleHeap{}
	for i := 0; i < n; i++ {
		idle = append(idle, i)
	}
	heap.Init(&idle)
	heap.Init(&using)

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0] // sort by startTime
	})

	fmt.Println("meetings:", meetings)

	for _, meeting := range meetings {
		st, end := meeting[0], meeting[1]
		// clean the using heap
		for len(using) > 0 && using[0][0] <= st {
			r := heap.Pop(&using).([]int)
			fmt.Println("clean", r)
			heap.Push(&idle, r[1])
		}

		if len(idle) == 0 {
			r := heap.Pop(&using).([]int)
			endtime := end + (r[0] - st)
			r[0] = endtime
			heap.Push(&using, r)
			fmt.Println("rooms:", rooms, " id:", r[1], "meeting:", meeting)

			rooms[r[1]]++
		} else {
			rid := heap.Pop(&idle).(int)
			r := []int{end, rid}
			heap.Push(&using, r)

			fmt.Println("rooms:", rooms, " id:", rid, "meeting:", meeting)
			rooms[rid]++
		}
	}

	fmt.Println("rooms:", rooms)
	idx, maxCnt := 0, 0
	for i, cnt := range rooms {
		if cnt > maxCnt {
			idx = i
			maxCnt = cnt
		}
	}
	return idx
}

func Test_mostBooked(t *testing.T) {
	n := 4
	meetings := [][]int{{10, 11}, {13, 15}, {9, 19}, {0, 12}, {12, 20}}

	fmt.Println(mostBooked(n, meetings))
}
