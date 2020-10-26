package main

import (
	"errors"
	"fmt"
)

func logging() {
	fmt.Println("Selesai memanggil function")
	message := recover()
	if message != nil {
		fmt.Println("terjadi error", message)
	}
}

func runApplication(value int) {
	defer logging()
	fmt.Println("run application")
	result := 10 / value
	fmt.Println(result)

}

func main() {
	runApplication(2)
	fmt.Println("--------------")
	runApplication(0)
	fmt.Println("ini main akhir")

}

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

// func main() {
// 	var contohError error = errors.New("Ups Error")
// 	fmt.Println(contohError.Error())

// 	hasil, err := Pembagi(100, 0)
// 	if err == nil {
// 		fmt.Println("Hasil", hasil)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}
// }
