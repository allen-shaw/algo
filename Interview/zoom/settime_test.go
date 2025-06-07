package zoom

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func setAlarm(setTime, timeToSet string) int {
	setHour, setMinute := spiltTime(setTime)
	toSetHour, toSetMinute := spiltTime(timeToSet)

	ans := 0

	mih, mxh := min(setHour, toSetHour), max(setHour, toSetHour)
	ans += min(abs(mxh-mih), abs(mih-mxh+24))

	mim, mxm := min(setMinute, toSetMinute), max(setMinute, toSetMinute)
	ans += min(abs(mxm-mim), abs(mim-mxm+60))

	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func spiltTime(t string) (int, int) {
	ss := strings.Split(t, ":")
	if len(ss) != 2 {
		panic("invaild time")
	}

	h, m := ss[0], ss[1]
	hour, e := strconv.Atoi(h)
	if e != nil {
		panic(e)
	}
	minute, e := strconv.Atoi(m)
	if e != nil {
		panic(e)
	}
	return hour, minute
}

func Test_setAlarm(t *testing.T) {
	ans := setAlarm("7:30", "8:00")
	fmt.Println(ans)

	ans = setAlarm("23:45", "8:00")
	fmt.Println(ans)
}
