# 1.给定一个整数数组和一个滑动窗口大小，求在这个窗口的滑动过程中，每个时刻其包含的最大值。
# Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
# Output: [3,3,5,5,6,7]


from collections import deque
from typing import List

# 单调队列
def max_sliding_window(nums: List[int], k: int) -> List[int]:
    if not nums or k < 0 or k > len(nums):
        return []

    queue = deque()
    result = []

    for i, num in enumerate(nums):
        while queue and queue[0] <= i-k:  # 先确定滑动窗口大小
            queue.popleft()
        while queue and nums[queue[-1]] <= num:
            queue.pop()
        queue.append(i)
        if i >= k - 1:
            result.append(nums[queue[0]])

    return result


# 给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
 
# 示例 1：
# 输入: s = "leetcode"
# 输出: 0
# 示例 2:
# 输入: s = "loveleetcode"
# 输出: 2
# 示例 3:
# 输入: s = "aabb"
# 输出: -1
 
# 提示:
# ●1 <= s.length <= 105
# ●s 只包含小写字母



def find_first_uniq_char(s: str) -> int:
    counter = [0] * 26
    for c in s:
        counter[ord(c) - ord('a')] += 1
    for i, c in enumerate(s):
        if counter[ord(c) - ord('a')] == 1:
            return i
    return -1
 

result1 = max_sliding_window([1,3,-1,-3,5,3,6,7],3)
print(result1)


result2 = find_first_uniq_char("leetcode")
result3 = find_first_uniq_char("loveleetcode")
result4 = find_first_uniq_char("aabb")

print(result2, result3, result4)


