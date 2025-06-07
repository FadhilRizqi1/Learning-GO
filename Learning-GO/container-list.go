package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()
	data.PushBack("Ahmad")
	data.PushBack("Fadhil")
	data.PushBack("Rizqi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for list := data.Front(); list != nil; list = list.Next() {
		fmt.Println(list.Value)
	}

	for list := data.Back(); list != nil; list = list.Prev() {
		fmt.Println(list.Value)
	}
}
