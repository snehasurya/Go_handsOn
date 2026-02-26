package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	result, err := Add(2, 3)
	if result != 6 {
		t.Errorf("result is %d, expected is %d", result, 6)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hey there error")
}
func TestAddFatal(t *testing.T) {
	result, err := Add(2, 3)
	if result != 6 {
		t.Fatalf("result is %d, expected is %d", result, 6)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hey there fatal")
}

func TestAddLog(t *testing.T) {
	result, err := Add(2, 3)
	if result != 6 {
		t.Logf("result is %d, expected is %d", result, 6)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hey there log")
}

func TestIsPalindrom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty String", "", true},
		{"single char", "a", true},
		{"simple palindrom", "mom", true},
		{"not a palindrom", "mommy", false},
		{"case insensitive", "Mom", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isPalindrom(test.input)
			if result != test.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestAddWithTestify(t *testing.T) {
	r, err := Add(4, -3)
	require.Equal(t, 7, r, "this is not equal")
	assert.Equal(t, 3, r, "add is not equal")
	assert.NotNil(t, err)

}
