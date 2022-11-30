package main

import (
	"fmt"
)

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

// soal 4
type movie struct {
	title, genre   string
	duration, year int
}

func main() {
	// soal 1
	fmt.Println("soal 1")
	type buah struct {
		Nama, Warna string
		AdaBijinya  bool
		Harga       int
	}
	var nanas buah
	nanas.Nama = "Nanas"
	nanas.Warna = "Kuning"
	nanas.AdaBijinya = false
	nanas.Harga = 9000
	fmt.Println(nanas)

	jeruk := buah{Nama: "Jeruk", Warna: "Orange", AdaBijinya: true, Harga: 8000}
	fmt.Println(jeruk)

	semangka := buah{"Semangka", "Hijau & Merah", true, 10000}
	fmt.Println(semangka)

	pisang := buah{Nama: "Pisang", Warna: "Kuning", AdaBijinya: false, Harga: 5000}
	fmt.Print(pisang, "\n\n")

	// soal 2
	segitiga_a := segitiga{alas: 10, tinggi: 10}
	fmt.Println("luas segitiga :", segitiga_a.luasSegitiga())

	persegi_a := persegi{sisi: 10}
	fmt.Println("luas persegi :", persegi_a.luasPersegi())

	persegipanjang_a := persegiPanjang{lebar: 10, panjang: 10}
	fmt.Print("luas persegi panjang :", persegipanjang_a.luaspersegiPanjang(), "\n\n")

	// soal 3
	fmt.Println("soal 3")
	nokia := phone{name: "Nokia C1", brand: "Nokia", year: 2020}
	nokia.colors = nokia.coloring("black", "white")
	fmt.Print(nokia, "\n\n")

	// soal 4
	fmt.Println("soal 4")
	var dataFilm = []movie{}

	tambahDataFilm := func(title string, minutes int, genre string, year int, data *[]movie) {
		m := movie{title, genre, minutes, year}
		*data = append(*data, m)
	}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, v := range dataFilm {
		fmt.Printf("%d. Title :%s\n", i+1, v.title)
		fmt.Println("   Genre :", v.genre)
		fmt.Println("   Duration :", v.duration)
		fmt.Print("   Year :", v.year, "\n\n")

	}
}

// soal 2
func (s segitiga) luasSegitiga() int {
	return s.alas * s.tinggi / 2
}
func (p persegi) luasPersegi() int {
	return p.sisi * p.sisi
}
func (pp persegiPanjang) luaspersegiPanjang() int {
	return pp.lebar * pp.lebar
}

// soal 3
func (p phone) coloring(color ...string) []string {
	return append(p.colors, color...)
}
