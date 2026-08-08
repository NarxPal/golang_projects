package greetings

import "fmt"

// Hello returns a greeting for a named person
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	greets := "Nice meeting you"
	return fmt.Sprintf("%s, %s", greets, name)
}
