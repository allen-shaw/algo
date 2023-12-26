package stack

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	for _, ast := range asteroids {
		alive := true
		if ast < 0 && len(stack) > 0 {
			for stack[len(stack)-1] > 0 {
				if -ast == stack[len(stack)-1] {
					stack = stack[:len(stack)-1] // pop
					alive = false
					break
				} else if -ast < stack[len(stack)-1] {
					alive = false
					break
				} else if -ast > stack[len(stack)-1] {
					stack = stack[:len(stack)-1] // pop
					alive = true
				}
			}
		}
		if alive {
			stack = append(stack, ast)
		}
	}

	return stack
}
