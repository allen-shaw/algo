package flexport

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func randStringGenerate(intput string, N int) string {
	m := make(map[string][]string, 0)

	tokens := strings.Split(intput, " ")
	n := len(tokens)
	for i, token := range tokens {
		next := (i + 1) % n
		m[token] = append(m[token], tokens[next])
	}

	result := make([]string, 0)
	curStr := pick(tokens)
	result = append(result, curStr)
	for i := 1; i < N; i++ {
		condidates := m[curStr]
		curStr = pick(condidates)
		result = append(result, curStr)
	}

	return strings.Join(result, " ")
}

func pick(ss []string) string {
	idx := rand.Intn(len(ss))
	return ss[idx]
}

func pickN(ss []string, n int) []string {
	idx := rand.Intn(len(ss))
	res, _ := pickFrom(ss, idx, n)
	return res
}

func pickFrom(ss []string, idx int, n int) ([]string, int) {
	if idx+n < len(ss) {
		return ss[idx : idx+n], idx + n
	}
	result := make([]string, 0)
	result = append(result, ss[idx:]...)
	result = append(result, ss[:n-(len(ss)-idx)]...)
	return result, n - (len(ss) - idx)
}

func Test_pickFrom(t *testing.T) {
	testStr := []string{"a", "b", "c", "d", "e"}
	out, next := pickFrom(testStr, 1, 3)
	fmt.Println(out, next) // b, c, d; 4

	out, next = pickFrom(testStr, 3, 4) // d, e, a, b; 2
	fmt.Println(out, next)
}

func Test_randStringGenerate(t *testing.T) {
	input := "it is a good sentence but is also a bad one"
	N := 5
	out := randStringGenerate(input, N)
	fmt.Println(out)
}

// M < N

func randStringGenerate2(intput string, N, M int) string {
	m := make(map[string][]string)

	tokens := strings.Split(intput, " ")

	for i := range tokens {
		picked, next := pickFrom(tokens, i, M)
		key := strings.Join(picked, "")
		m[key] = append(m[key], tokens[next])
	}

	tokenpool := pickN(tokens, M)

	result := make([]string, 0)
	result = append(result, tokenpool...)

	for i := 0; i < N-M; i++ {
		key := strings.Join(tokenpool, "")
		candidates := m[key]
		nextWord := pick(candidates)
		result = append(result, nextWord)
		tokenpool = append(tokenpool, nextWord)
		tokenpool = tokenpool[1:]
	}

	return strings.Join(result, " ")
}

func Test_randStringGenerate2(t *testing.T) {
	input := "it is a good sentence but is also a bad one"
	N := 5
	M := 2
	out := randStringGenerate2(input, N, M)
	fmt.Println(out)
}
