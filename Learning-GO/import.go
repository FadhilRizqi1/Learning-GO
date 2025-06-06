package main

import (
	"Learning-GO/helper"
	"Learning-GO/other"
	"fmt"
)

func main() {
	helper.SayName("Fadhil")
	other.SayName("Rizqi")

	// fmt.Print(helper.unexported) // error: helper.version is not exported
	fmt.Print(helper.Exported) // exported variable can be accessed
}