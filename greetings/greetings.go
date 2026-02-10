package greetings

import (
	"fmt"

	"errors"

	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[message] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{"Hi, %v. Welcome!", "Hey there, %v! Howdy?", "What's up, %v? All good?"}
	return formats[rand.Intn(len(formats))]
}
