package _3__restoreipaddresses

import "strconv"

func restoreIpAddresses(s string) []string {
	return dfs(s, 4)
}

func dfs(s string, count int) []string {
	if len(s) == 0 {
		return []string{}
	}
	if count == 1 {
		if isAvail(s) {
			return []string{s}
		}
		return []string{}
	}
	ret := make([]string, 0)
	for i := 0; i < len(s); i++ {
		str := s[:i+1]
		if len(str) > 3 {
			break
		}
		if isAvail(str) {
			result := dfs(s[i+1:], count-1)
			for j := 0; j < len(result); j++ {
				ret = append(ret, str + "." + result[j])
			}
		}
	}
	return ret
}

func isAvail(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '0' {
		return len(s) == 1
	}
	num, _ := strconv.Atoi(s)
	return num <= 255
}