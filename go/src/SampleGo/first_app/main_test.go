package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	type testCase struct {
		Name string
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			main()
		})
	}

	validate(t, &testCase{})
}
