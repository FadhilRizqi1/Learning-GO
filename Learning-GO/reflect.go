package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Nama string `required:"true" max:"10" json:"nama"`
}

func main() {
	sample := sample{Nama: "Fadhil"}

	var SampleType = reflect.TypeOf(sample)
	fmt.Println("Type:", SampleType.NumField())
	fmt.Println("Name:", SampleType.Field(0).Name)

	fmt.Println("Tag:", SampleType.Field(0).Tag)
	fmt.Println("Tag JSON:", SampleType.Field(0).Tag.Get("required"))
	fmt.Println("Tag XML:", SampleType.Field(0).Tag.Get("max"))
	fmt.Println("Tag YAML:", SampleType.Field(0).Tag.Get("json"))
}