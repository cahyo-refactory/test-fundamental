package main

import "fmt"

func sayHello() {
	fmt.Println("Hallo fungsi")
}

func sayName(name string) {
	fmt.Println("Hallo ", name)
}

func sayCity(name string) string {
	return "Hello " + name
}

func getAddress() (string, string, string) {
	return "RT 008", "RW 001", "Kel. Maridan"
}

// func main() {
// 	RT, RW, Kelurahan := getAddress()
// 	fmt.Println(RT, RW, Kelurahan)
// 	fmt.Println(getAddress())
// }

func getAddressNamed() (RT, RW, Kel string) {
	RT = "RT 008"
	RW = "RW 001"
	Kel = "Kel. Maridan"
	return
}

// func main() {
// 	fmt.Println(getAddressNamed())
// }

func sumAll(pengali int, numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total * pengali
}

// func main() {
// 	fmt.Println(sumAll(2, 2, 3, 5))
// 	numbers := []int{
// 		2, 3, 5,
// 	}
// 	fmt.Println(sumAll(2, numbers...))
// }

func getCountry(name string) string {
	return "Halo " + name
}

// func main() {
// 	getNegara := getCountry
// 	fmt.Println(getCountry("Indonesia"))
// 	fmt.Println(getNegara("Indonesia"))
// }

type Filter func(string) string

func textFilter(name string, filter Filter) {
	filtered := filter(name)
	fmt.Println("hello", filtered)
}

func filter(text string) string {
	if text == "anjing" {
		return "..."
	} else {
		return text
	}

}

// func main() {
// 	textFilter("cahyo", filter)
// 	textFilter("anjing", filter)
// }

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("are you blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

// func main() {
// 	blacklist := func(name string) bool {
// 		return name == "admin"
// 	}
// 	registerUser("admin", blacklist)
// 	registerUser("eko", func(name string) bool {
// 		return name == "admin"
// 	})

// }

//menggunakan for loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//menggunakan recrusive
func factorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}

// func main() {
// 	loop := factorialLoop(5)
// 	fmt.Println(loop)
// 	fmt.Println(5 * 4 * 3 * 2 * 1)

// 	recrusive := factorialRecrusive(5)
// 	fmt.Println(recrusive)

// }

func main() {
	count := 0
	increment := func() {
		fmt.Println("Tambah")
		count++
	}

	increment()
	increment()
	fmt.Println(count)
}
