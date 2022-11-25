package main

import (
	"fmt"
)

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := func(title string, time string, genre string, year string) {
		dataFilm = append(dataFilm, map[string]string{
			"genre": genre,
			"jam":   time,
			"tahun": year,
			"title": title,
		})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func luasPersegiPanjang(x int, y int) int {
	return x * y
}
func kelilingPersegiPanjang(x int, y int) int {
	return 2 * (x + y)
}
func volumeBalok(x int, y int, z int) int {
	return x * y * z
}
func introduce(name string, gender string, job string, age string) string {
	panggilan := "bu"
	if gender == "laki-laki" {
		panggilan = "pak"
	}
	return (panggilan + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun")
}
func buahFavorit(name string, fruits ...string) string {
	fruit := ``
	for _, f := range fruits {
		fruit += `"` + f + `", `
	}
	return "halo nama saya " + name + " dan buah favorit saya adalah " + fruit[:len(fruit)-2]
}
