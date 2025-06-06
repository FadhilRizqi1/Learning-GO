package database

var connection string

func init() {
	connection = "localhost:3306"
}

func GetConnection() string {
	return connection
}