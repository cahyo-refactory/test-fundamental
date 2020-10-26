package main

import "fmt"

type nilai struct {
	Nama map[string]interface{}
}

func helper(sq *nilai) {
	sq.Nama = make(map[string]interface{})
	sq.Nama["data"] = "ss"
	fmt.Println(sq)
}

func main() {

	// var countedData map[string][]ChartElement

	// countedData := make(map[string][]ChartElement)

	var tesNilai nilai

	helper(&tesNilai)
	fmt.Println(tesNilai)
}
