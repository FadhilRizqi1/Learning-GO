package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Nanosecond())

	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, err := time.Parse(layout, "2025-06-09")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(parse)
}