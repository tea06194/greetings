package greetings

import (
	"errors"
	"fmt"
)

// H - h test
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hello %v, go", name)
	
	return message, nil
}

