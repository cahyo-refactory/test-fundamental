package main

import "fmt"

func main() {

	var result = 10 + 20
	fmt.Println(result)

	result /= 2
	fmt.Println(result)

	result++
	fmt.Println(result)

	tes := 'd'
	var angka int
	angka = 'd'

	fmt.Println("ini hasil tes ", tes, angka)

	var (
		angka1 = 10
		angka2 = 21
	)

	var nilai bool = angka1 == angka2
	fmt.Println(nilai)

	var (
		ujian        = 90
		absensi      = 90
		lulusUjian   = ujian >= 90
		lulusabsensi = absensi >= 70
		lulus        = lulusUjian && lulusabsensi
	)

	fmt.Println(lulus)

}
