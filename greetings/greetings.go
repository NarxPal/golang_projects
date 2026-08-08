package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {
	// if no name given return an error with msg
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name is provided than
	// Return a greeting that embeds the name in a message.
	greets := "Nice meeting you"
	message := fmt.Sprintf("%s, %s", greets, name)
	return message, nil
}
