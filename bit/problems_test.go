package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/
// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

func singleNumber(nums []int) int {
	r := 0
	for _, num := range nums {
		r = r ^ num
	}
	return r
}

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	nums3 := []int{1}
	r1 := singleNumber(nums)
	r2 := singleNumber(nums2)
	r3 := singleNumber(nums3)
	fmt.Println(r1, r2, r3)
}

// 137. 只出现一次的数字 II
// https://leetcode.cn/problems/single-number-ii/
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法且不使用额外空间来解决此问题
func singleNumber2(nums []int) int {
	fmt.Println(nums)
	if len(nums) == 1 {
		return nums[0]
	}

	// 随机选一个数n，遍历整个数组，将小于n的移到左边，大于等于n放到右边，返回n的位置
	idx := partition(nums)
	// 如果idx % 3 == 0，说明唯一的数在右边, 如果idx % 3 == 1说明唯一的数在左边
	if idx%3 == 0 {
		return singleNumber2(nums[idx+1:])
	} else {
		return singleNumber2(nums[0 : idx+1])
	}
}

func partition(nums []int) int {
	// idx := rand.Intn(len(nums) - 1)
	// nums[0], nums[idx] = nums[idx], nums[0]
	r := len(nums) - 1
	n := nums[r]

	i := 1
	j := r - 1

	for i <= j {
		for i < r && nums[i] <= n {
			i++
		}
		for j >= 0 && nums[j] > n {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	nums[r], nums[i] = nums[i], nums[r]
	return i
}

func TestPartition(t *testing.T) {
	nums1 := []int{0, 1, 0, 1, 0, 1, 99}
	nums2 := []int{30000, 500, 100, 30000, 100, 30000, 100}
	nums3 := []int{2, 2, 3, 2}
	r1 := partition(nums1)
	r2 := partition(nums2)
	r3 := partition(nums3)

	fmt.Println(nums1, r1)
	fmt.Println(nums2, r2)
	fmt.Println(nums3, r3)
}

// NO PASS
func TestSingleNumber2(t *testing.T) {
	// nums := []int{0, 1, 0, 1, 0, 1, 99}
	nums := []int{2, 2, 3, 2}
	// nums := []int{30000, 500, 100, 30000, 100, 30000, 100}
	r1 := singleNumber2(nums)
	fmt.Println(r1)
}

// 260. 只出现一次的数字 III
// https://leetcode.cn/problems/single-number-iii
// 给你一个整数数组 nums，其中恰好有两个元素（a,b）只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
// 你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。
func singleNumber3(nums []int) []int {
	// 整个数组异或，得到n = a ^ b
	n := xor(nums)

	// 找到最后一位为1的位，假设第k位，a和b必然有一个数在第k位为1，另一个在第k位为0
	k := n & (^n + 1)

	// 遍历所有数，如果这
	a := 0
	for _, num := range nums {
		if num&k == 0 {
			// 假设a的第k位为0，那a和所以第k位为0的值异或（肯定不包含b,而且肯定是偶数个）= a
			a = a ^ num
		}
	}

	b := n ^ a // n=a^b => n^a=a^b^a=b
	return []int{a, b}
}

func xor(nums []int) int {
	r := 0
	for _, num := range nums {
		r = r ^ num
	}
	return r
}

func TestSingleNumber3(t *testing.T) {
	nums1 := []int{1, 2, 1, 3, 2, 5}
	r1 := singleNumber3(nums1)
	fmt.Println(r1)

	nums2 := []int{0, 1}
	r2 := singleNumber3(nums2)
	fmt.Println(r2)

	nums3 := []int{0, -1}
	r3 := singleNumber3(nums3)
	fmt.Println(r3)
}
