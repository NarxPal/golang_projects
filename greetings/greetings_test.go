package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName fn calls greetings.Hello with a name, checking
// for a valid return value.

// fn name must start with "Test"
// t *testing.T is used to report test failures
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want) // t.Errorf marks test as failed
	}
}

// TestHelloEmpty fn calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
