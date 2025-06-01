package main

import "fmt"

func sumAll(number ...int) int {
	var total int

	for _, value := range number {
		total += value
	}

	return total
}

func main() {

	var numbers = []int{5, 5, 5, 5, 5}

	fmt.Println(sumAll(numbers...)) 
	fmt.Println(sumAll(10, 20, 10, 10, 5)) 
	fmt.Println(sumAll()) 
}

