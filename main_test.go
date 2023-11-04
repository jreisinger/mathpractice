package main

import (
	"testing"
)

func FuzzParseURLPath(f *testing.F) {
	testcases := []string{"5", "10", "0", "a"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, path string) {
		upto, err := parseURLPath(path)
		if err == nil && upto <= 0 {
			t.Errorf("parseURLPath(%q) returned zero or negative int: %d", path, upto)
		}
	})
}
