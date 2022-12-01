package main

import (
	. "BDS-Sanbercode-Golang-Batch-40/tugas-9/tugas-8"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// soal 1
	fmt.Println("soal 1")
	var hitungBangunDatar HitungBangunDatar

	hitungBangunDatar = SegitigaSamaSisi{Alas: 10, Tinggi: 8}
	fmt.Println("luas segitiga sama sisi :", hitungBangunDatar.Luas())
	fmt.Println("keliling segitiga sama sisi :", hitungBangunDatar.Keliling())

	hitungBangunDatar = PersegiPanjang{Panjang: 10, Lebar: 8}
	fmt.Println("luas peregi panjang :", hitungBangunDatar.Luas())
	fmt.Println("keliling persegi panjang", hitungBangunDatar.Keliling())

	var hitungBangunRuang HitungBangunRuang

	hitungBangunRuang = Tabung{JariJari: 5.0, Tinggi: 10}
	fmt.Println("volume Tabung :", hitungBangunRuang.Volume())
	fmt.Println("luas permukaan Tabung :", hitungBangunRuang.LuasPermukaan())

	hitungBangunRuang = Balok{Panjang: 10, Lebar: 8, Tinggi: 5}
	fmt.Println("volume Balok :", hitungBangunRuang.Volume())
	fmt.Println("luas permukaan Balok :", hitungBangunRuang.LuasPermukaan())
	fmt.Println()

	// soal 2
	fmt.Println("soal 2")
	var print Print

	color := []string{"Mystic Bronze", "Mystic White", "Mystic Black"}

	print = Phone{Name: "samsung Galaxy Note 20", Brand: "samsung Galaxy Note 20", Year: 2020, Colors: color}

	for k, v := range print.PrintPhone() {
		typeV := fmt.Sprintf("%s", reflect.TypeOf(v))
		if typeV == "[]string" {
			v = strings.Join(v.([]string), " , ")
		}
		fmt.Println(k, ":", v)
	}
	fmt.Println()

	// soal 3
	fmt.Println("soal 3")
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))
	fmt.Println()

	// soal 4
	fmt.Println("soal 4")

	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}
	w := kumpulanAngkaPertama.([]int)[0]
	x := kumpulanAngkaPertama.([]int)[1]
	y := kumpulanAngkaKedua.([]int)[0]
	z := kumpulanAngkaKedua.([]int)[1]
	fmt.Printf("%s %d + %d + %d + %d  = %d\n\n", prefix, w, x, y, z, (w + x + y + z))
}
