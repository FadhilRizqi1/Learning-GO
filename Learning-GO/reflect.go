package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Name string
}

func main() {
	sample := sample{Name: "Fadhil"}

	var SampleType = reflect.TypeOf(sample)
	fmt.Println("Type:", SampleType.NumField())
	fmt.Println("Name:", SampleType.Field(0).Name)
}