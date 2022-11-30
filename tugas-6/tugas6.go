package main

import (
	"fmt"
	"math"
)

func main() {
	// >>>>>>>>>>>>>>> soal 1 <<<<<<<<<<<<<<<
	var luasLigkaran float64
	var kelilingLingkaran float64

	jari_jari := 10.00
	jari2 := &jari_jari

	countingCircles := func(r *float64) {
		phi := 3.142857142857143
		luasLigkaran = phi * math.Pow(*r, 2)
		kelilingLingkaran = phi * 2 * *r
		fmt.Printf("\nsoal1\nluas lingkaran : %f\nkeliling lingkaran : %f\n\n", luasLigkaran, kelilingLingkaran)
	}
	countingCircles(jari2)

	// >>>>>>>>>>>>>>> soal 2 <<<<<<<<<<<<<<<
	var sentence string
	fmt.Println("soal 2")
	introduce := func(str *string, name string, gender string, job string, age string) {
		*str = ""
		panggilan := "pak"
		if gender == "perempuan" {
			panggilan = "bu"
		}
		*str += fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", panggilan, name, job, age)
	}
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence)

	// >>>>>>>>>>>>>>> soal 3 <<<<<<<<<<<<<<<
	var buah = []string{}
	fmt.Println("\nsoal 3")
	addFruit := func(fruit *[]string) {
		*fruit = append(*fruit, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	}
	addFruit(&buah)
	for i, v := range buah {
		fmt.Printf("%d.%s\n", i+1, v)
	}

	// >>>>>>>>>>>>>>> soal 4 <<<<<<<<<<<<<<<
	var dataFilm = []map[string]string{}
	fmt.Println("\nsoal 4")

	tambahDataFilm := func(title string, time string, genre string, year string, data *[]map[string]string) {
		*data = append(*data, map[string]string{
			"title":    title,
			"duration": time,
			"genre":    genre,
			"year":     year,
		})
	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	count := 1
	for _, v := range dataFilm {
		str := fmt.Sprintf("%d.", count)
		for x, y := range v {
			fmt.Println(str, x, ":", y)
			str = "  "
		}
		count += 1
		fmt.Println()
	}
}
