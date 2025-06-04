package main

import "fmt"

func callName(name string) map[string]string {
	if name == " " {
		return nil
	} else {
		return map[string]string{
			"nama": name,
		}
	}
}

func main() {
	var contohNil map[string]string = nil
	fmt.Println(contohNil)

	var filled = callName("Fadhil")
	fmt.Println(filled)

	var empty = callName(" ")
	fmt.Println(empty)
}