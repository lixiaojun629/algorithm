package recursive

import (
	"errors"
	"strings"
)

//https://leetcode-cn.com/problems/restore-ip-addresses/
//回溯，
func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return nil
	}
	res := make([]string, 0)
	path := make([]string, 0)

	var dfs func(int, int)
	dfs = func(start, depth int) {
		if depth == 4 && start >= n {
			res = append(res, strings.Join(path, "."))
			return
		}
		if start >= n {
			return
		}

		if s[start] == '0' {
			path = append(path, "0")
			dfs(start+1, depth+1)
			path = path[:len(path)-1] //回溯完成后，撤销对path的更改，下同
			return
		}

		for i := start; i < start+3 && i < n; i++ {
			seg, err := validIpSegment(s[start : i+1])
			if err == nil {
				path = append(path, seg)
				dfs(i+1, depth+1)
				path = path[:len(path)-1] //
			}
		}
	}
	dfs(0, 0)
	return res
}

func validIpSegment(str string) (string, error) {
	res := 0
	for i := 0; i < len(str); i++ {
		res = 10*res + int(str[i]-'0')
	}
	if res >= 0 && res <= 255 {
		return str, nil
	}
	return "", errors.New("invalid ip segment " + str)
}
