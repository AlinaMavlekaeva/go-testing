package greet

import "errors"

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}
	return "Hello, " + name, nil
}
