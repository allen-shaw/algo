package saleforce

// items map[id]num
// package 10

type Knapsack struct {
	slot []int 		// index, 剩余空间 
	m map[int]int   
	slotIdx int
}


func (k *Knapsack) put(id int, num int) bool {
	idx, ok := k.m[id]
	if !ok {
		k.slotIdx++
		idx = k.slotIdx
	}

	if num > len(k.slot) * 200 + k.slot[k.slotIdx] {
		return false
	}

	for num > k.slot[idx] {
		num -= k.slot[idx]
		k.slot[idx] = 0
		idx = k.newIndex(id)
	}
	k.slot[idx] -= num
	return true
}

func (k *Knapsack) newIndex(item int) int {
	k.slotIdx++
	k.m[item] = k.slotIdx
	return k.slotIdx
}


func pack(items map[int]int, p int) bool {
	curIdx := 0
	curSize := 200

	for _, num := range items {
		for num > curSize {
			num -= curSize
			curIdx++
			if curIdx >= p {
				return false
			}
		}
		curSize -= num

		curIdx++
		curSize = 200
		if curIdx >= p {
			return false
		}
	}

	return true
}
