package main

import "fmt"

func main() {

	suhu := 20
	if suhu <= 0 {
		fmt.Println("beku")
	} else if suhu > 0 && suhu <= 100 {
		fmt.Println("cair")
	} else {
		fmt.Println("uap")
	}

	nama := "cahyo dwi raharjo"

	if length := len(nama); length <= 5 {
		angka := 5
		fmt.Println("nama terlalu pendek", length, angka)
	} else {
		fmt.Println("nama sudah benar", length)
		// fmt.Println(angka) ini error
	}
	nama = "cahyo"
	switch nama {
	case "Budi":
		fmt.Println("Halo budi")
	case "cahyo":
	case "tini":
		fmt.Println("Halo tini")
	case "Andi":
		fmt.Println("Halo andi")
	default:
		fmt.Println("Nama Belum terdaftar")
	}

	switch length := len(nama); length <= 5 {
	case true:
		fmt.Println("nama terlalu pendek")
	case false:
		fmt.Println("nama sudah benar")
	}

	length := len(nama)
	switch {
	case length <= 5:
		fmt.Println("nama terlalu pendek")
	case length > 5:
		fmt.Println("nama sudah benar 1")
	case length > 10:
		fmt.Println("nama sudah benar 2")
	}

	count := 1
	for count <= 10 {
		fmt.Println("Perulangan ke ", count)
		count++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	slice := []string{"solo", "sleman", "klaten"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index ", i, " value ", value)
	}

	for i := 0; i < 10; i++ {

		if i == 7 {
			break
		}

		if i == 3 {
			continue
		}

		fmt.Println("perulangan ke ", i)

	}

}
