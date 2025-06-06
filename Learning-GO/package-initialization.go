package main

import (
	"Learning-GO/database"
	"fmt"
)

func main() {
	connection := database.GetConnection()
	fmt.Println(connection)
}