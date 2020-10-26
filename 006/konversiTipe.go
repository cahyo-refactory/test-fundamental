package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)
	var nama = "cahyo"
	var c = nama[0]
	var cString string = string(c)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println(cString)

	str1 := "1234"

	i1, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println(i1)
	}

	str2 := "567"
	i2, err := strconv.ParseInt(str2, 10, 64)
	if err == nil {
		fmt.Println(i2)
	} else {
		fmt.Println(err)
	}

}
