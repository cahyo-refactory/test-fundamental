package main

import "fmt"

func main() {

	type nomorIdentitas string

	var noKTP nomorIdentitas = "647103"
	var noSIM nomorIdentitas = "123456"

	fmt.Println(noKTP)
	fmt.Println(noSIM)

}
