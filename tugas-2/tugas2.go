package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()

	// soal 1
	a := "Bootcamp"
	b := "Digital"
	c := "Skill"
	d := "Sanbercode"
	e := "Golang"
	fmt.Printf("%s %s %s %s %s \n", a, b, c, d, e)

	// soal 2
	halo := "Halo Dunia"
	fmt.Println(strings.ReplaceAll(halo, "Dunia", "GOLANG"))

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	kataKedua = strings.ToUpper(kataKedua[0:1]) + kataKedua[1:]
	// kataKetiga = kataKetiga[:len(kataKetiga)-1] + strings.ToUpper(kataKetiga[len(kataKetiga)-1:]) -->> cara panjang
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)
	fmt.Printf("%s %s %s %s \n", kataPertama, kataKedua, kataKetiga, kataKeempat)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	pertama, _ := strconv.ParseInt(angkaPertama, 10, 64)
	kedua, _ := strconv.ParseInt(angkaKedua, 10, 64)
	ketiga, _ := strconv.ParseInt(angkaKetiga, 10, 64)
	keempat, _ := strconv.ParseInt(angkaKeempat, 10, 64)
	fmt.Println(pertama + kedua + ketiga + keempat)

	// ssoal 5
	kalimat := "halo halo bandung"
	angka := 2021
	kalimat = strings.ReplaceAll(kalimat, "halo", "HI")
	fmt.Printf(`"%s" - %d`, kalimat, angka)

	fmt.Printf("\n\n")
}
