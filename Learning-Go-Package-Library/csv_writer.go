package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Name", "Age", "City"})
	_ = writer.Write([]string{"Ahmad Fadhil Rizqi", "19", "Palembang"})
	_ = writer.Write([]string{"Fadhil Rizqi", "21", "Jakarta"})

	writer.Flush()
}