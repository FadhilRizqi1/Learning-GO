package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Halo\n")
	writer.WriteString("Ini adalah writer.\n")
	writer.Flush()
}