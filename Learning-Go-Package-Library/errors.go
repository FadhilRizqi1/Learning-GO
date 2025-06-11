package main

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
)

func getByID(id int) (string, error) {
	if id <= 0 {
		return "", ErrInvalidInput
	}
	if id != 1 {
		return "", ErrNotFound
	}
	return "Item 1", nil
}

func main() {
	item, err := getByID(2)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			println("Error:", err.Error())
		} else if errors.Is(err, ErrInvalidInput) {
			println("Error:", err.Error())
		}
		return
	}
	println("Retrieved item:", item)
}