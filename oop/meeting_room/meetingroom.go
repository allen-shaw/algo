package meetingroom

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// ===== 错误定义 =====

var (
	ErrNoRoomAvailable  = errors.New("没有可用的会议室")
	ErrMeetingNotFound  = errors.New("会议不存在")
	ErrInvalidTimeRange = errors.New("无效的时间范围：结束时间必须大于开始时间")
)

// ===== 会议 =====

// Meeting 表示一次会议预约。
type Meeting struct {
	ID     int
	Start  int // 开始时间（可用整数表示时间戳或小时等）
	End    int // 结束时间（左闭右开：[Start, End)）
	RoomID int // 分配的会议室编号
}

// ===== 会议室 =====

// MeetingRoom 表示一个会议室及其已预订的会议列表。
type MeetingRoom struct {
	ID       int
	meetings []*Meeting // 按开始时间排序
}

// IsAvailable 检查该会议室在 [start, end) 时间段是否空闲。
//
// 利用会议列表按 start 排序的特性，通过二分查找实现 O(log n)。
func (r *MeetingRoom) IsAvailable(start, end int) bool {
	// 找到第一个 Start >= end 的会议索引
	idx := sort.Search(len(r.meetings), func(i int) bool {
		return r.meetings[i].Start >= end
	})
	// 检查前一个会议是否与 [start, end) 重叠
	if idx > 0 && r.meetings[idx-1].End > start {
		return false
	}
	return true
}

// book 将会议插入到已排序的会议列表中。
func (r *MeetingRoom) book(m *Meeting) {
	pos := sort.Search(len(r.meetings), func(i int) bool {
		return r.meetings[i].Start >= m.Start
	})
	// 在 pos 位置插入：前半部分 + 新元素 + 后半部分
	r.meetings = append(r.meetings[:pos], append([]*Meeting{m}, r.meetings[pos:]...)...)
}

// cancel 从会议列表中移除指定会议。
// 利用列表按 Start 排序的特性，二分查找定位，O(log M)。
func (r *MeetingRoom) cancel(m *Meeting) {
	// 二分找到 Start 所在位置
	pos := sort.Search(len(r.meetings), func(i int) bool {
		return r.meetings[i].Start >= m.Start
	})
	// 从 pos 开始找到匹配的 meetingID（同一 Start 可能有多个会议）
	for i := pos; i < len(r.meetings) && r.meetings[i].Start == m.Start; i++ {
		if r.meetings[i].ID == m.ID {
			r.meetings = append(r.meetings[:i], r.meetings[i+1:]...)
			return
		}
	}
}

// ===== 会议预约系统 =====

// MeetingScheduler 是会议预约系统的核心结构。
//
// 核心数据结构：
//   - rooms: 会议室列表
//   - meetings: meetingID → Meeting 的映射，O(1) 查找和取消
//   - nextID: 自增的会议 ID 生成器
//   - mu: 读写锁，保证并发安全
type MeetingScheduler struct {
	rooms    []*MeetingRoom
	meetings map[int]*Meeting // meetingID → Meeting
	nextID   int
	mu       sync.RWMutex
}

// NewMeetingScheduler 创建一个拥有 numRooms 个会议室的预约系统。
func NewMeetingScheduler(numRooms int) *MeetingScheduler {
	rooms := make([]*MeetingRoom, numRooms)
	for i := 0; i < numRooms; i++ {
		rooms[i] = &MeetingRoom{ID: i}
	}
	return &MeetingScheduler{
		rooms:    rooms,
		meetings: make(map[int]*Meeting),
		nextID:   1,
	}
}

// Book 预约会议。给定 [start, end) 时间段，分配一个空闲会议室。
// 成功返回 meetingID；没有空闲会议室则返回错误。
//
// 时间复杂度: O(R · log M)，R = 会议室数量，M = 单个会议室的会议数
func (s *MeetingScheduler) Book(start, end int) (int, error) {
	if start >= end {
		return 0, ErrInvalidTimeRange
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, room := range s.rooms {
		if room.IsAvailable(start, end) {
			meeting := &Meeting{
				ID:     s.nextID,
				Start:  start,
				End:    end,
				RoomID: room.ID,
			}
			s.nextID++
			room.book(meeting)
			s.meetings[meeting.ID] = meeting
			return meeting.ID, nil
		}
	}

	return 0, ErrNoRoomAvailable
}

// Cancel 取消指定 meetingID 的会议，释放对应会议室的时间段。
//
// 时间复杂度: O(M)，M = 该会议室的会议数
func (s *MeetingScheduler) Cancel(meetingID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	meeting, exists := s.meetings[meetingID]
	if !exists {
		return ErrMeetingNotFound
	}

	room := s.rooms[meeting.RoomID]
	room.cancel(meeting)
	delete(s.meetings, meetingID)
	return nil
}

// GetMeeting 查询指定 meetingID 的会议信息。
//
// 时间复杂度: O(1)
func (s *MeetingScheduler) GetMeeting(meetingID int) (*Meeting, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	meeting, exists := s.meetings[meetingID]
	if !exists {
		return nil, ErrMeetingNotFound
	}
	return meeting, nil
}

// ===== 示例 =====

// Example 演示预约系统的基本用法。
func Example() {
	scheduler := NewMeetingScheduler(2) // 2 个会议室

	// 预约会议
	id1, err := scheduler.Book(9, 10)
	fmt.Printf("预约会议 1: id=%d, err=%v\n", id1, err) // id=1, err=<nil>

	id2, err := scheduler.Book(9, 10)
	fmt.Printf("预约会议 2: id=%d, err=%v\n", id2, err) // id=2, err=<nil>

	// 两个会议室都在 9-10 点被占用，第三个预约失败
	id3, err := scheduler.Book(9, 10)
	fmt.Printf("预约会议 3: id=%d, err=%v\n", id3, err) // id=0, err=没有可用的会议室

	// 不冲突的时间段可以预约
	id4, err := scheduler.Book(10, 11)
	fmt.Printf("预约会议 4: id=%d, err=%v\n", id4, err) // id=3, err=<nil>

	// 取消会议后释放时间段
	_ = scheduler.Cancel(id1)
	id5, err := scheduler.Book(9, 10)
	fmt.Printf("取消后重新预约: id=%d, err=%v\n", id5, err) // id=4, err=<nil>
}
