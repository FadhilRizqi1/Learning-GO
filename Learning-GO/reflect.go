package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Nama string `required:"true" max:"10"`
}

type sample2 struct {
	Nama string
	Umur int
}

func validate(data interface{}) bool {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("required")
		value := v.Field(i)

		if tag == "true" && value.String() == "" {
			return false
		}

		maxTag := field.Tag.Get("max")
		if maxTag != "" {
			maxLength := 0
			fmt.Sscanf(maxTag, "%d", &maxLength)
			if len(value.String()) > maxLength {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := sample{Nama: "Fadhil"}

	var SampleType = reflect.TypeOf(sample)
	fmt.Println("Type:", SampleType.NumField())
	fmt.Println("Name:", SampleType.Field(0).Name)

	fmt.Println("Tag:", SampleType.Field(0).Tag)
	fmt.Println("Tag JSON:", SampleType.Field(0).Tag.Get("required"))
	fmt.Println("Tag XML:", SampleType.Field(0).Tag.Get("max"))

	sample.Nama = ""
	fmt.Println("Valid:", validate(sample))

	fmt.Println("Valid:", validate(sample2{Nama: "Fadhil", Umur: 19}))
}