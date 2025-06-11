package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Round(9.7))
	fmt.Println(math.Round(9.5))
	fmt.Println(math.Round(9.4))

	fmt.Println(math.Ceil(9.2))
	fmt.Println(math.Floor(9.7))

	fmt.Println(math.Min(9.2, 9.7))
	fmt.Println(math.Max(9.2, 9.7))
}