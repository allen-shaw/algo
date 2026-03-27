"""会议预约系统

给定开始和结束时间，分配空闲会议室并返回 meetingId。
没有空闲会议室则抛出异常。
"""

import bisect
import threading
from dataclasses import dataclass, field


# ===== 异常定义 =====

class NoRoomAvailableError(Exception):
    """没有可用的会议室"""
    pass


class MeetingNotFoundError(Exception):
    """会议不存在"""
    pass


class InvalidTimeRangeError(Exception):
    """无效的时间范围"""
    pass


# ===== 会议 =====

@dataclass
class Meeting:
    """一次会议预约。"""
    id: int
    start: int  # 开始时间（左闭）
    end: int    # 结束时间（右开）：[start, end)
    room_id: int


# ===== 会议室 =====

class MeetingRoom:
    """管理单个会议室及其已预订的会议列表（按 start 排序）。"""

    def __init__(self, room_id: int):
        self.id = room_id
        self._meetings: list[Meeting] = []
        # 辅助数组，与 _meetings 一一对应，用于 bisect 二分查找
        self._starts: list[int] = []

    def is_available(self, start: int, end: int) -> bool:
        """检查该会议室在 [start, end) 时间段是否空闲。

        利用 _starts 排序特性，通过二分查找实现 O(log n)。
        """
        # 找到第一个 start >= end 的会议索引
        idx = bisect.bisect_left(self._starts, end)
        # 检查前一个会议是否与 [start, end) 重叠
        if idx > 0 and self._meetings[idx - 1].end > start:
            return False
        return True

    def book(self, meeting: Meeting) -> None:
        """将会议插入到已排序的会议列表中。O(log n) 定位 + O(n) 插入。"""
        pos = bisect.bisect_left(self._starts, meeting.start)
        self._meetings.insert(pos, meeting)
        self._starts.insert(pos, meeting.start)

    def cancel(self, meeting: Meeting) -> None:
        """从会议列表中移除指定会议。

        利用列表按 start 排序的特性，二分查找定位，O(log n)。
        """
        pos = bisect.bisect_left(self._starts, meeting.start)
        # 从 pos 开始找到匹配的 meetingID（同一 start 可能有多个会议）
        i = pos
        while i < len(self._meetings) and self._meetings[i].start == meeting.start:
            if self._meetings[i].id == meeting.id:
                self._meetings.pop(i)
                self._starts.pop(i)
                return
            i += 1


# ===== 会议预约系统 =====

class MeetingScheduler:
    """会议预约系统的核心结构。

    核心数据结构：
      - rooms: 会议室列表
      - meetings: meetingID → Meeting 的映射，O(1) 查找和取消
      - _next_id: 自增的会议 ID 生成器
      - _lock: 读写锁，保证并发安全
    """

    def __init__(self, num_rooms: int):
        self.rooms = [MeetingRoom(i) for i in range(num_rooms)]
        self.meetings: dict[int, Meeting] = {}
        self._next_id = 1
        self._lock = threading.Lock()

    def book(self, start: int, end: int) -> int:
        """预约会议。给定 [start, end) 时间段，分配一个空闲会议室。

        成功返回 meetingID；没有空闲会议室则抛出 NoRoomAvailableError。

        时间复杂度: O(R · log M)，R = 会议室数量，M = 单个会议室的会议数
        """
        if start >= end:
            raise InvalidTimeRangeError("结束时间必须大于开始时间")

        with self._lock:
            for room in self.rooms:
                if room.is_available(start, end):
                    meeting = Meeting(
                        id=self._next_id,
                        start=start,
                        end=end,
                        room_id=room.id,
                    )
                    self._next_id += 1
                    room.book(meeting)
                    self.meetings[meeting.id] = meeting
                    return meeting.id

            raise NoRoomAvailableError("没有可用的会议室")

    def cancel(self, meeting_id: int) -> None:
        """取消指定 meetingID 的会议，释放对应会议室的时间段。

        时间复杂度: O(log M)
        """
        with self._lock:
            meeting = self.meetings.get(meeting_id)
            if meeting is None:
                raise MeetingNotFoundError(f"会议 {meeting_id} 不存在")

            room = self.rooms[meeting.room_id]
            room.cancel(meeting)
            del self.meetings[meeting_id]

    def get_meeting(self, meeting_id: int) -> Meeting:
        """查询指定 meetingID 的会议信息。O(1)。"""
        meeting = self.meetings.get(meeting_id)
        if meeting is None:
            raise MeetingNotFoundError(f"会议 {meeting_id} 不存在")
        return meeting


# ===== 示例 =====

if __name__ == "__main__":
    scheduler = MeetingScheduler(num_rooms=2)  # 2 个会议室

    # 预约会议
    id1 = scheduler.book(9, 10)
    print(f"预约会议 1: id={id1}")  # id=1

    id2 = scheduler.book(9, 10)
    print(f"预约会议 2: id={id2}")  # id=2

    # 两个会议室都在 9-10 点被占用，第三个预约失败
    try:
        scheduler.book(9, 10)
    except NoRoomAvailableError as e:
        print(f"预约会议 3: {e}")  # 没有可用的会议室

    # 不冲突的时间段可以预约
    id4 = scheduler.book(10, 11)
    print(f"预约会议 4: id={id4}")  # id=3

    # 取消会议后释放时间段
    scheduler.cancel(id1)
    id5 = scheduler.book(9, 10)
    print(f"取消后重新预约: id={id5}")  # id=4
