package main

import (
	"bufio"
	"io"
	"os"
)

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

func readFile(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		message += string(line) + "\n"
	}
	return message, nil
}


func main () {
	err := createNewFile("example.txt", "Ini adalah contoh file.\n")
	if err != nil {
		panic(err)
	}

	result, err := readFile("example.txt")
	if err != nil {
		panic(err)
	}

	println(result)
}