package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome",
		"Great to see you, %v",
		"Hail, %v! Well met",
	}

	return formats[rand.Intn(len(formats))]
}

// Hello function returns a greeting to the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//return the greeting
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintf(randomFormat())
	return message, nil

}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	//obs: the "_" in the next line is a representation of index of the loop.
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
