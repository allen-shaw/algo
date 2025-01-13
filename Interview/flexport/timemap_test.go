package flexport

import (
	"fmt"
	"sort"
	"testing"
)

type entry struct {
	Timestamp int
	Value     string
}

type TimeMap struct {
	m map[string][]entry
}

func NewTimeMap() TimeMap {
	return TimeMap{
		m: make(map[string][]entry),
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	_, ok := m.m[key]
	if !ok {
		m.m[key] = make([]entry, 0)
	}
	e := entry{Timestamp: timestamp, Value: value}
	m.m[key] = append(m.m[key], e)
}

func (m *TimeMap) Get(key string, timestamp int) string {
	vals, ok := m.m[key]
	if !ok {
		return ""
	}
	idx := sort.Search(len(vals), func(i int) bool {
		return vals[i].Timestamp > timestamp
	})
	if idx > 0 {
		return vals[idx-1].Value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
func TestTimeMap(t *testing.T) {
	timeMap := NewTimeMap()
	timeMap.Set("foo", "bar", 1)       // 存储键 "foo" 和值 "bar" ，时间戳 timestamp = 1
	fmt.Println(timeMap.Get("foo", 1)) // 返回 "bar"
	fmt.Println(timeMap.Get("foo", 3)) // 返回 "bar", 因为在时间戳 3 和时间戳 2 处没有对应 "foo" 的值，所以唯一的值位于时间戳 1 处（即 "bar"） 。
	timeMap.Set("foo", "bar2", 4)      // 存储键 "foo" 和值 "bar2" ，时间戳 timestamp = 4
	fmt.Println(timeMap.Get("foo", 4)) // 返回 "bar2"
	fmt.Println(timeMap.Get("foo", 5)) // 返回 "bar2"
}
