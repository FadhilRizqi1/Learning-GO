package main

import (
	"encoding/base64"
	"fmt"
)

func main() { 

	value := "Ahmad Fadhil Rizqi"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Decoded:", string(decoded))
}