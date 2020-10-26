package main

import "fmt"

type Mahasiswa struct {
	Nama, Asal string
	Age        int
	Ipk        float32
}

func (mahasiswa Mahasiswa) halloMahasiswa(name string) {
	fmt.Println("hello", name, " nama saya ", mahasiswa.Nama)
}

func main() {
	var andi Mahasiswa
	andi.Nama = "Andi Budi"
	andi.Asal = "Balikpapan"
	andi.Age = 9
	andi.Ipk = 3.5
	fmt.Println(andi)
	fmt.Println(andi.Age)

	tono := Mahasiswa{
		Nama: "Tonolanto",
		Asal: "Balikpapan",
		Age:  35,
		Ipk:  2.5,
	}
	fmt.Println(tono)
	fmt.Println(tono.Nama)

	budi := Mahasiswa{"Budi Raharjo", "Bandung", 60, 4}
	fmt.Println(budi)

	budi.halloMahasiswa("siska")
}
