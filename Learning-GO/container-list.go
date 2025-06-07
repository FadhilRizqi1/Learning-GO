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
}
