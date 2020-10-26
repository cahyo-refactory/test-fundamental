package main

import "fmt"

func main() {
	//array

	// var cities [3]string
	// cities[0] = "balikpapan"
	// cities[1] = "solo"
	// cities[2] = "klaten"

	// fmt.Println(cities)
	// fmt.Println(cities[0])
	// fmt.Println(len(cities))

	// var ages = [3]int{
	// 	30,
	// 	23,
	// 	29,
	// }
	// fmt.Println(ages)
	// fmt.Println(ages[1])

	// fmt.Println(len(ages))

	//slice
	var angka = [...]string{
		"nol",
		"satu",
		"dua",
		"tiga",
		"empat",
		"lima",
		"enam",
		"tujuh",
		"delapan",
		"sembilan",
		"sepuluh",
		"sebelas",
	}
	var slice1 = angka[10:]
	fmt.Println(angka)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, "nilai_baru1")
	fmt.Println(angka)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "10"
	fmt.Println(angka)
	fmt.Println(slice1)

	// fmt.Println(angka)
	// fmt.Println(len(angka))
	// var slice1 = angka[4:7]
	// fmt.Println(slice1)
	// fmt.Println("panjang array", len(slice1))
	// fmt.Println("Kapasitas ", cap(slice1))

	// slice1[0] = "nilai4"
	// fmt.Println("nilai array :: ", angka)

	// slice1 = append(slice1, "nilai6")
	// fmt.Println(slice1)
	// fmt.Println("panjang array slice1", len(slice1))
	// fmt.Println("Kapasitas slice1", cap(slice1))

	// fmt.Println(angka)
	// fmt.Println(len(angka))

	// var slice2 = angka[10:]
	// fmt.Println(slice2)
	// fmt.Println("panjang array", len(slice2))
	// fmt.Println("Kapasitas ", cap(slice2))

	// var slice3 = append(slice2, "bukan angka")
	// fmt.Println(slice3)
	// fmt.Println(len(slice3))
	// fmt.Println(cap(slice3))

	// slice3[1] = "bukan nilai"
	// fmt.Println(slice3)
	// fmt.Println(slice2)

	newSlice := make([]string, 3, 6)

	newSlice[0] = "Klaten"
	newSlice[1] = "Solo"
	newSlice[2] = "Pacitan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	iniArray := [...]int{
		1, 2, 3, 4, 5,
	}
	iniSlice := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	//map
	var kota map[string]string = map[string]string{}

	user := map[string]string{
		"nama": "cahyo",
		"kota": "balikpapan",
	}
	fmt.Println(user)
	user["umur"] = "30"
	fmt.Println(user)
	delete(user, "umur")
	fmt.Println(user)
	fmt.Println(kota)

}
