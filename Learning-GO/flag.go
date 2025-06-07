package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host ", "localhost", "host server")
	var user *string = flag.String("user", "root", "user database")
	var password *string = flag.String("password", "root", "password database")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
}