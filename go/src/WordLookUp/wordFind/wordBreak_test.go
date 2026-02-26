package wordBreakLookup

import "testing"

type tests struct {
	s        string
	wordDict []string
	expected bool
}

func TestWordBreak(t *testing.T) {
	tests := []tests{
		{"applepenapple", []string{"apple", "pen"}, true},                    // Example 1
		{"catsanddog", []string{"cats", "dog", "sand", "and", "cat"}, false}, // Example 2
		{"leetcode", []string{"leet", "code"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for _, tt := range tests {
		result := WordBreak(tt.s, tt.wordDict)
		if result != tt.expected {
			t.Errorf("%v wordBreak wants %v got %v", tt.s, tt.expected, result)
		}
	}
}
