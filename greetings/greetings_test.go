package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "James"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello("James")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("James") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
