package backtrack

var (
	mapping = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'u', 'v', 'w'},
		'9': {'x', 'y', 'z'},
	}
	lctrack []byte
	lcans   []string
)

func letterCombinations(digits string) []string {
	lctrack = make([]byte, 0)
	lcans = make([]string, 0)

	lcdfs(digits, 0)
	return lcans
}

// func spiltDigists(digits string) [][]byte {
// 	res := make([][]byte, 0)
// 	for i := 0; i < len(digits); i++ {
// 		c := digits[i]
// 		chars := mapping[c]
// 		res = append(res, chars)
// 	}

// 	return res
// }

func lcdfs(digits string, n int) {
	if n == len(digits) {
		lcans = append(lcans, string(lctrack))
		return
	}

	nums := mapping[digits[n]]
	for i := 0; i < len(nums); i++ {
		lctrack = append(lctrack, nums[i])
		lcdfs(digits, n+1)
		lctrack = lctrack[:len(lctrack)-1]
	}
}
