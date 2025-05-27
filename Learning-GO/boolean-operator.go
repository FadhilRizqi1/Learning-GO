package main

import "fmt"

func main() {

	var ujian = 85
	var absensi = 80

	var lulusUjian = ujian >= 85
	var lulusAbsensi = absensi >= 85
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 85 && absensi >= 85)

}
