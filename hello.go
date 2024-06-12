package greetings

import (
	"fmt"
)

// M - m test
func Hello(name string) string {
	message := fmt.Sprintf("Hello %v, go", name)
	return message
}

