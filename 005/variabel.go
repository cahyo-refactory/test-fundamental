package main

import "fmt"

func main() {
	var name string
	name = "rafactory"
	fmt.Println(name)

	name = "refactory.dom"
	fmt.Println(name)

	var city = "jogja"
	fmt.Println(city)

	age := 48
	fmt.Println(age)

	var (
		firstName = "Cahyo"
		lastName  = "Dwi Raharjo"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	var fullName = firstName + " " + lastName
	fmt.Println(fullName)

	const country string = "indonesia"
	const village = "niten"

	//multiple variabel constant
	const (
		namaDepan string = "cahyo"
		umur             = 30
	)

}

// func NewMap(name string) map[string]string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return map[string]string{
// 			"name": name,
// 		}
// 	}
// }

// func main() {
// 	person := NewMap("halo halo")
// 	if person == nil {
// 		fmt.Println("Data Kosong")
// 	} else {
// 		fmt.Println(person)

// 	}

// }
