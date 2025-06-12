package main

import (
	"encoding/base64"
	"fmt"
)

func main() { 

	value := "Ahmad Fadhil Rizqi"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encoded)
}