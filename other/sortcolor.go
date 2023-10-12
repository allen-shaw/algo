package other

import "fmt"

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1

	for i := 0; i <= p2; {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
			fmt.Println("p0", p0, "i:", i)
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			for nums[i] == 2 && i <= p2 {
				nums[i], nums[p2] = nums[p2], nums[i]
				p2--
			}
			fmt.Println("p2:", p2, "i:", i)
		}
	}
}
