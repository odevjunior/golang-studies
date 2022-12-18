package greetings

import (
	"errors"
	"fmt"
)

// Hello function returns a greeting to the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//return the greeting
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message, nil

}
