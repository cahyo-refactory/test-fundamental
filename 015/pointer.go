package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// func main() {
// 	address1 := Address{"Subang", "JABAR", "Indonesia"}
// 	address2 := address1

// 	address2.City = "Bandung"
// 	fmt.Println("# add1", address1)
// 	fmt.Println("# add2", address2)

// 	var address3 *Address = &address1
// 	// address3 := &address1
// 	address3.City = "Majalengka"
// 	fmt.Println("$ add1", address1)
// 	fmt.Println("$ add3", address3)

// 	address3 = &Address{"BERAU", "KALTIM", "Indonesia"}
// 	fmt.Println(">> add1", address1)
// 	fmt.Println(">> add3", address3)

// 	address4 := &address1
// 	*address4 = Address{"Balikpapan", "KALTIM", "Indonesia"}

// 	fmt.Println(">>>> add1", address1)
// 	fmt.Println(">>>> add2", address2)
// 	fmt.Println(">>>> add3", address3)
// 	fmt.Println(">>>> add4", address4)

// 	// var address5 *Address = new(Address)
// 	// fmt.Println(address5)

// }

func ChangeCountry(address *Address) {
	address.Country = "Indonesia"
}

// func main() {
// 	var alamat = Address{
// 		City:     "Klaten",
// 		Province: "Jawa Tengah",
// 		Country:  "",
// 	}
// 	ChangeCountry(&alamat)
// 	fmt.Println(alamat)
// }

type Man struct {
	Name string
}

func (man *Man) School() {
	man.Name = "Adik " + man.Name
}

func main() {
	cahyo := Man{"cahyo"}
	// cahyo.School()
	fmt.Println(cahyo.Name)
}
