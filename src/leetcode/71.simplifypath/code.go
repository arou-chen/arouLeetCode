package _1_simplifypath

import "strings"

func simplifyPath(path string) string {
	list := strings.Split(path, "/")
	result := make([]string, 0)
	for i := 0; i < len(list); i++ {
		str := list[i]
		if str == " " || str == "." {
			continue
		} else if str == ".." {
			if length := len(result); length > 0 {
				result = result[:length - 1]
			}
		} else {
			result = append(result, str)
		}
	}
	return "/" + strings.Join(result, "/")
}