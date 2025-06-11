package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])i")

	fmt.Println("Regex:", regex)
	fmt.Println("Match:", regex.MatchString("api"))
	fmt.Println("Match:", regex.MatchString("aLi"))
	fmt.Println("Match:", regex.MatchString("ade"))


	fmt.Println("FindString:", regex.FindAllString("api aLi ade aKi aTi adi ami", -1))
}