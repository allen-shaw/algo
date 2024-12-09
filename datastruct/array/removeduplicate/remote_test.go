package removeduplicate

import (
	"fmt"
	"strings"
	"testing"
)

// 0, 0, 1, 1, 2, 3, 3|, 4, 3, 4
//
//	s    f
func removeDuplicates(nums []int) int {
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func Test_removeDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 4}
	out := removeDuplicates(nums)
	fmt.Println(nums, nums[:out])
	fmt.Println(out)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	ss := make([]rune, 0)
	for _, c := range s {
		if !isLetterOrDigit(c) {
			continue
		}
		ss = append(ss, c)
	}
	for l, r := 0, len(ss)-1; l < r; l, r = l+1, r-1 {
		if ss[l] != ss[r] {
			return false
		}
	}
	return true
}

func isLetterOrDigit(s rune) bool {
	if ('a' <= s && s <= 'z') || ('0' <= s && s <= '9') {
		return true
	}
	return false
}

func Test_isPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	ans := isPalindrome(s)
	fmt.Println(ans)
}

// i
// 0, 0, 2, 1, 1, 2
//
//	p0       p2
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	i := p0
	for i <= p2 {
		n := nums[i]
		if n < 1 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
			if i < p0 {
				i = p0
			}
		} else if n == 1 {
			i++
		} else if n > 1 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		}
	}
}

func Test_sortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := len(nums1) - 1
	p1, p2 := m-1, n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] <= nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}
	fmt.Println("nums1:", nums1, ",p1:", p1, ",p2:", p2)

	if p2 >= 0 {
		copy(nums1, nums2[:p2+1])
	}
}

func Test_merge(t *testing.T) {
	nums1 := []int{4, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 3, 5, 6}
	m, n := 1, 5
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
