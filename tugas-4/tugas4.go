package main

import "fmt"

func main() {
	// soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else if i%2 != 0 && i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else {
			fmt.Println(i, "- Santai")
		}
	}

	// soal 2
	count := "#"
	for i := 0; i < 7; i++ {
		fmt.Println(count)
		count += "#"
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	for i, s := range sayuran {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	for k, v := range satuan {
		fmt.Println(k, "=", v)
	}
}
