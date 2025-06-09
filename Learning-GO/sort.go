package main

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