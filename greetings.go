package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome to team Apty", name)
	return message
}
