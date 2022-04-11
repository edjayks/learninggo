package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {

	input := "The quick brown fox jumped over the lazy dog"
	reversedInput, revError := Reverse(input)
	doubleReversedInput, doubleRevError := Reverse(reversedInput)

	fmt.Printf("original string: %q\n", input)
	fmt.Printf("Reversed string: %q, error: %v\n", reversedInput, revError)
	fmt.Printf("Double Reversed string: %q, error: %v\n", doubleReversedInput, doubleRevError)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Input is not a valid UTF-8")
	}
	fmt.Printf("input: %q\n", s)
	b := []rune(s)
	fmt.Printf("runes: %q\n", b)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
