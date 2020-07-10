package _3__restoreipaddresses

import "strconv"

func restoreIpAddress(s string) []string {
	length := len(s)
	if length == 0 {
		return []string{}
	}
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