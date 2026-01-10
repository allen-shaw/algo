package presum

import "math/rand/v2"

type RandomPicker struct {
	presum []int
}

func newRandomPicket(weights []int) RandomPicker {
	presum := make([]int, len(weights)+1)
	for i, w := range weights {
		presum[i+1] = presum[i] + w
	}

	return RandomPicker{
		presum: presum,
	}
}

func (p *RandomPicker) PickIndex() int {
	r := rand.IntN(p.presum[len(p.presum)-1])
	l, r := 0, len(p.presum)-1
	for l <= r {
		mid := (l + r) / 2
		if p.presum[mid] <= r {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
