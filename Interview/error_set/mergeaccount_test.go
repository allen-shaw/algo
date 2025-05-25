package errorset

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	emailToAcct := make(map[string][]int)
	for acct, emails := range accounts {
		for i := 1; i < len(emails); i++ {
			email := emails[i]
			emailToAcct[email] = append(emailToAcct[email], acct)
		}
	}

	n := len(accounts)
	visited := make([]bool, n)
	visitedEmails := make(map[string]struct{})

	var dfs func(acct int, emailList []string)
	dfs = func(acct int, emailList []string) {
		if visited[acct] {
			return
		}

		visited[acct] = true
		emails := accounts[acct][1:]

		for _, email := range emails {
			if _, ok := visitedEmails[email]; ok {
				continue
			}
			visitedEmails[email] = struct{}{}
			emailList = append(emailList, email)

			accts := emailToAcct[email]
			for _, acct := range accts {
				dfs(acct, emailList)
			}
		}
	}

	ans := make([][]string, 0)
	for acct := range accounts {
		if visited[acct] {
			continue
		}
		emails := make([]string, 0)
		dfs(acct, emails)

		sort.Strings(emails)
		user := accounts[acct][0]
		res := []string{user}
		res = append(res, emails...)
		ans = append(ans, res)
	}

	return ans
}
