package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Ahmad Fadhil Rizqi\n" + "Fadhil Rizqi\n" + "MAFALQI\n"

	reader := csv.NewReader(strings.NewReader((csvString)))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println("Record:", record)
	}

}