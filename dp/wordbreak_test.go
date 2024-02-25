package dp

import (
	"fmt"
	"testing"
)

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{})
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
	}

	var dp func(s string, i int) bool
	dp = func(s string, i int) bool {
		if i == len(s) {
			return true
		}

		for l := 1; i+l <= len(s); l++ {
			prefix := s[i : i+l]
			_, ok := wordMap[prefix]
			if ok {
				if dp(s, i+l) {
					return true
				}
			}
		}

		return false
	}

	return dp(s, 0)
}

func TestWordBreak(t *testing.T) {
	// s := "leetcode"
	// wordDict := []string{"leet", "code"}
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}

	ans := wordBreak(s, wordDict)
	fmt.Println(ans)
}

func wordBreak2(s string, wordDict []string) []string {
	workMap := make(map[string]struct{})
	for _, w := range wordDict {
		workMap[w] = struct{}{}
	}

	memo := make([][]string, len(s))

	// 返回s[i:]的所有可能句子
	var dp func(s string, i int) []string
	dp = func(s string, i int) []string {
		res := make([]string, 0)
		if i == len(s) {
			res = append(res, "")
			return res
		}

		if len(memo[i]) != 0 {
			return memo[i]
		}

		for l := 1; i+l <= len(s); l++ {
			prefix := s[i : i+l]
			_, ok := workMap[prefix]
			if ok {
				ss := dp(s, i+l)
				for _, s := range ss {
					if s == "" {
						res = append(res, prefix)
					} else {
						res = append(res, prefix+" "+s)
					}
				}
			}
		}

		memo[i] = res
		return res
	}

	return dp(s, 0)
}

func TestWorkBreak2(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	ans := wordBreak2(s, wordDict)
	for i := 0; i < len(ans); i++ {
		fmt.Println(i, ans[i])
	}
}
