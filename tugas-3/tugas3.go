package main

import (
	"fmt"
	"strconv"
)

func main() {

	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	panjang, err := strconv.Atoi(panjangPersegiPanjang)
	if err != nil {
		panic(err)
	}
	lebar, err := strconv.Atoi(lebarPersegiPanjang)
	if err != nil {
		panic(err)
	}

	alas, err := strconv.Atoi(alasSegitiga)
	if err != nil {
		panic(err)
	}
	tinggi, err := strconv.Atoi(tinggiSegitiga)
	if err != nil {
		panic(err)
	}

	luasPersegiPanjang = panjang * lebar
	kelilingPersegiPanjang = 2 * (panjang + lebar)
	luasSegitiga = (alas * tinggi) / 2

	fmt.Printf("\nLuas Persegi Panjang : %d\nKeliling Persegi Panjang : %d\nLuas Segitiga : %d\n\n", luasPersegiPanjang, kelilingPersegiPanjang, luasSegitiga)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	fmt.Println("indeks nilai John :", index(nilaiJohn))
	fmt.Println("indeks nilai Doe :", index(nilaiDoe))
	fmt.Println()

	// soal 3
	var tanggal = 23
	var bulan = 7
	var tahun = 1996
	var bulanLahir string
	switch bulan {
	case 1:
		bulanLahir = "Januari"
	case 2:
		bulanLahir = "Februari"
	case 3:
		bulanLahir = "Maret"
	case 4:
		bulanLahir = "April"
	case 5:
		bulanLahir = "Mei"
	case 6:
		bulanLahir = "Juni"
	case 7:
		bulanLahir = "Juli"
	case 8:
		bulanLahir = "Agustus"
	case 9:
		bulanLahir = "September"
	case 10:
		bulanLahir = "Oktober"
	case 11:
		bulanLahir = "November"
	case 12:
		bulanLahir = "Desember"
	}
	fmt.Printf("%d %s %d \n\n", tanggal, bulanLahir, tahun)

	// soal 4
	generasi := "Generasi Z"

	if tahun < 1965 {
		generasi = "Baby boomer"
	}
	if tahun < 1980 {
		generasi = "Generasi X"
	}
	if tahun < 1995 {
		generasi = "Generasi Y (Milenials)"
	}
	fmt.Println(generasi)

}

func index(val int) string {
	if val < 50 {
		return "E"
	}
	if val < 60 {
		return "D"
	}
	if val < 70 {
		return "C"
	}
	if val < 80 {
		return "B"
	}
	return "A"
}
