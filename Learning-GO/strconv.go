package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Boolean value:", boolean)
	}

	intValue, err := strconv.ParseInt("1000000", 10, 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Integer value:", intValue)
	}

	stringValue := strconv.FormatInt(200000, 8)
	fmt.Println("String value:", stringValue)

	intValue2, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Integer value from string:", intValue2)
	}
}