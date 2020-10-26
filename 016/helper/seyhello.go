package helper

import "fmt"

var versionFile = "1.0.0" // tidak bisa diakses dari luar package
//Application comment.
var Application = "golang" // dapat diakses diluar package

//SayHello dapat dikases diluar package
func SayHello(name string) {
	fmt.Println("hello", name)
}

//tidak dapat diakses dari luar package
func sayHello(name string) {
	fmt.Println("hello", name)
}

func init() {
	fmt.Println("ini akan di eksekusi ketika package di akses")
	fmt.Println("koneksi database")
}
