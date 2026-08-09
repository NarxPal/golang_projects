package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {
	// if no name given return an error with msg
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name is provided than
	// Return a greeting that embeds the name in a message.
	// create msg using random Format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomly return msg from set of greeting msgs
func randomFormat() string {
	// a slice of msg formats
	formats := []string{
		"Hi, %v Welcome!",
		"What's going on %v",
		"everything fine %v bro",
	}

	// use random index for slice of formats

	return formats[rand.Intn(len(formats))]

}
