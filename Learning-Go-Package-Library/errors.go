package main

import (
	"errors"
	"fmt"
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
	var itemID int

	fmt.Println("Input item ID (1 for valid item, any other number for not found):")
    fmt.Scan(&itemID)

	item, err := getByID(itemID) 
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