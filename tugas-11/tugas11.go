package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// soal 1
	fmt.Println("soal 1")
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup
	for i, p := range phones {
		wg.Add(1)
		time.Sleep(time.Second)
		go printPhone(i+1, p, &wg)
	}
	wg.Wait()

	// soal 2
	time.Sleep(time.Second)
	fmt.Println("\nsoal 2")

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}

func printPhone(idx int, phone string, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Printf("%d. %s\n", idx, phone)
}

func getMovies(ch chan string, data ...string) {
	for _, d := range data {
		ch <- d
	}
	close(ch)
}
