package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Nama string
}

func main() {
	sample := sample{Nama: "Fadhil"}

	var SampleType = reflect.TypeOf(sample)
	fmt.Println("Type:", SampleType.NumField())
	fmt.Println("Name:", SampleType.Field(0).Name)
}