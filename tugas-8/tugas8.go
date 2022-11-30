package main

import (
	"fmt"
	"reflect"
	"strings"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}
type print interface {
	printPhone() map[string]interface{}
}

func main() {
	// soal 1
	fmt.Println("soal 1")
	var hitungBangunDatar hitungBangunDatar

	hitungBangunDatar = segitigaSamaSisi{10, 8}
	fmt.Println("luas segitiga sama sisi :", hitungBangunDatar.luas())
	fmt.Println("keliling segitiga sama sisi :", hitungBangunDatar.keliling())

	hitungBangunDatar = persegiPanjang{10, 8}
	fmt.Println("luas peregi panjang :", hitungBangunDatar.luas())
	fmt.Println("keliling persegi panjang", hitungBangunDatar.keliling())

	var hitungBangunRuang hitungBangunRuang

	hitungBangunRuang = tabung{5.0, 10}
	// fmt.Println("volume tabung :", hitungBangunRuang.volume())
	// fmt.Println("luas permukaan tabung :", hitungBangunRuang.luasPermukaan())

	hitungBangunRuang = balok{10, 8, 5}
	fmt.Println("volume balok :", hitungBangunRuang.volume())
	fmt.Println("luas permukaan balok :", hitungBangunRuang.luasPermukaan())
	fmt.Println()

	// soal 2
	fmt.Println("soal 2")
	var print print

	color := []string{"Mystic Bronze", "Mystic White", "Mystic Black"}

	print = phone{"samsung Galaxy Note 20", "samsung Galaxy Note 20", 2020, color}

	for k, v := range print.printPhone() {
		typeV := fmt.Sprintf("%s", reflect.TypeOf(v))
		if typeV == "[]string" {
			v = strings.Join(v.([]string), " , ")
		}
		fmt.Println(k, ":", v)
	}
	fmt.Println()

	// soal 3
	fmt.Println("soal 3")
	luasPersegi := func(numb int, boo bool) interface{} {
		var result interface{}
		if boo == true {
			result = "Maaf anda belum menginput sisi dari persegi"
			if numb > 0 {
				result = fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", numb, numb*numb)
			}
		} else {
			if numb > 0 {
				result = numb
			}
		}
		return result
	}
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
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

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}
func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3

}
func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}
func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}
func (t tabung) volume() float64 {
	return 3.14 * t.jariJari * t.tinggi
}
func (t tabung) luasPermukaan() float64 {
	return 2 * 3.14 * t.jariJari * (t.jariJari + t.tinggi)
}
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}
func (b balok) luasPermukaan() float64 {
	return float64(2 * ((b.panjang * b.lebar) + (b.panjang * b.tinggi) + (b.lebar * b.tinggi)))
}
func (p phone) printPhone() map[string]interface{} {
	return map[string]interface{}{
		"Name":   p.name,
		"Brand":  p.brand,
		"Year":   p.year,
		"Colors": p.colors,
	}
}
