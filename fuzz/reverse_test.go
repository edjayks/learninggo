package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testCases {
		reversedString, revError := Reverse(tc.in)
		if revError != nil {
			return
		}
		if reversedString != tc.want {
			t.Errorf("Reverse: %q, want: %q", reversedString, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, oriString string) {
		t.Logf("Fuzz Testing %q\n", oriString)
		reversed, revError := Reverse(oriString)
		if revError != nil {
			return
		}
		doubleReversed, doubleRevError := Reverse(reversed)
		if doubleRevError != nil {
			return
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(oriString), utf8.RuneCountInString(reversed), utf8.RuneCountInString(doubleReversed))
		if oriString != doubleReversed {
			t.Errorf("Before: %q, after: %q", oriString, doubleReversed)
		}
		if utf8.ValidString(oriString) && !utf8.ValidString(reversed) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", reversed)
		}
	})
}
