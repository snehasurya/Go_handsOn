package wordBreakLookup

func WordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)

	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	//s1 := "catsandog"
	//word := []string{"cats", "dog", "sand", "and", "cat"}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
