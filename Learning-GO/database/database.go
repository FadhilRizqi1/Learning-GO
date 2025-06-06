package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init dipanggil")
	connection = "localhost:3306"
}

func GetConnection() string {
	return connection
}