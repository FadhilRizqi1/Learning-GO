package main

import "os"

func createNewFile(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}

func main () {
	err := createNewFile("example.txt", "Halo!")
	if err != nil {
		panic(err)
	}

	println("File created successfully")
}