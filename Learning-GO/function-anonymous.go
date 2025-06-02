package main

import "fmt"

type Blacklist = func(name string) bool
var name string

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User " + name + " is blacklisted and cannot be registered.")
	} else {
		fmt.Println("User " + name + " has been successfully registered.")
	}
}

func main(){
	blacklist := func(name string) bool {
		return name == "admin" || name == "root" 
	}

	registerUser("Fadhil", blacklist)
	registerUser("admin", blacklist)

	
	fmt.Println("Masukkan nama anda:")
	fmt.Scan(&name)
	registerUser(name, blacklist)
}
