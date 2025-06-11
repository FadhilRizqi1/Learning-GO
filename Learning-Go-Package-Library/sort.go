package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (value Users) Len() int {
	return len(value)
}

func (value Users) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value Users) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := Users{
		{Name: "Ahmad", Age: 20},
		{Name: "Fadhil", Age: 19},
		{Name: "Rizqi", Age: 21},
	}

	sort.Sort(users)

	for _, user := range users {
		fmt.Println(user.Name, user.Age)
	}
}