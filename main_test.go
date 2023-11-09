package main

import (
	"testing"
)

func FuzzParseInput(f *testing.F) {
	testcases := []string{"", "a", "0", "5", "10", "100"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, path string) {
		n, err := parseInput(path)
		if err == nil && n <= 0 {
			t.Errorf("parseInput(%q) returned zero or negative int: %d", path, n)
		}
	})
}
