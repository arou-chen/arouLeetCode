/**
 *Create by chensr on 2019/9/27
*/

package dynamicProgramming

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	start, end, length, max := 0, 0, len(s), 0
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	for i := 1; i < length; i++ {
		for j := i -1; j >= 0; j-- {
			if s[i] == s[j] && (i - j < 2 || dp[i-1][j+1]) {
				dp[i][j] = true
				if max < i - j + 1 {
					max = i - j + 1
					start = j
					end = i
				}
			}
		}
	}

	return s[start:end + 1]
}