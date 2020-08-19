package exercises

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// dp[[i][j] means if the s[i,j] is palindrome
	// dp[i][j] == dp[i-1]dp[j-1] && s[i] == s[j]
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	theBegin, theLen := 0, 1
	for j := 1; j < len(s); j++ {
		for i := 0; i < len(s); i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if (j-1)-(i+1)+1 < 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				} // else>>>>
			} // else>>>

			// if ap[i][j] == true,
			// then s[i,j] is palindrome.
			// theBegin is the begin index of palindrome,
			// theLen is the length of palindrome.
			if dp[i][j] && j-i+1 > theLen {
				theBegin = i
				theLen = j - i + 1
			}
		} // for>>
	} // for>
	return s[theBegin : theBegin+theLen]
}
