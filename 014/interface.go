package main

import "fmt"

type HasName interface {
	GetName() string
	GetAge() int
}

func SayHello(person HasName) {
	fmt.Println("hello", person.GetName(), person.GetAge())
}

type Person struct {
	Name string
	Age  int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}

// func main() {
// 	var person Person
// 	person.Name = "cahyo"
// 	person.Age = 30

// 	SayHello(person)
// }

//interface kosong

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

// func main() {
// 	var data interface{} = Ups(1)
// 	fmt.Println(data)
// }

func random() interface{} {
	return "Ups"
}

func main() {

	//tanpa switch
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	//menggunakan switch
	var result2 interface{} = random()
	switch value := result2.(type) {
	case int:
		fmt.Println("value", value, "is int")
	case string:
		fmt.Println("value", value, "is string")
	case bool:
		fmt.Println("value", value, "is bool")
	default:
		fmt.Println("value uknown type")
	}
}
