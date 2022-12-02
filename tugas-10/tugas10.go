package main

import (
	"fmt"
	"strconv"
)

func main() {
	// angka := 1
	// soal 1
	defer tryDefer("Golang Backend Development", 2022)

	// soal 2
	fmt.Println("\nsoal 2")
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	// soal 3
	// defer cetakAngka(&angka)

	// tambahAngka(7, &angka)

	// tambahAngka(6, &angka)

	// tambahAngka(-1, &angka)

	// tambahAngka(9, &angka)

}

func tryDefer(kalimat string, tahun int) {
	fmt.Println("\nsoal 1 \n", kalimat, tahun)
}

func kelilingSegitigaSamaSisi(num int, bo bool) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r)
		}
	}()
	result := ""
	if bo {
		result = "Maaf anda belum menginput sisi dari segitiga sama sisi"
		if num > 0 {
			result = fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", num, num*3)
		}
	} else {
		if num == 0 {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
		result = strconv.Itoa(num)
	}
	return result
}
