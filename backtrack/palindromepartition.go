package backtrack

func partition(s string) [][]string {
	return newPartitionSolution().Partition(s)
}

type partitionSolution struct {
	res [][]string
}

func newPartitionSolution() *partitionSolution {
	s := &partitionSolution{}
	s.res = make([][]string, 0)

	return s
}

func (s *partitionSolution) Partition(str string) [][]string {
	palindromes := NewLinkedList[string]()
	s.backtrack(palindromes, str, 0)
	return s.res
}

func (s *partitionSolution) backtrack(palindromes *LinkedList[string], str string, start int) {
	if start == len(str) {
		s.res = append(s.res, palindromes.ToList())
		return
	}

	for i := start; i < len(str); i++ {
		if !s.isPalindrome(str[start : i+1]) {
			continue
		}

		palindromes.PushBack(str[start : i+1])
		s.backtrack(palindromes, str, i+1)

		palindromes.RemoveLast()
	}
}

func (s *partitionSolution) isPalindrome(str string) bool {
	if len(str) == 0 {
		return false
	}

	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}

	return true
}
