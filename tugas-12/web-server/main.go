package main

import (
	"fmt"
	"net/http"
)

var phi = 3.142857142857143
var jari2 = 7
var tinggi = 10

func Volume(r int, t int) float64 {
	return float64(r*r*t) * phi
}
func LuasAlas(r int, t int) float64 {
	return float64(r*r) * phi
}
func KelilingAlas(r int, t int) float64 {
	return float64(r) * phi * 2
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "jariJari : 7, tinggi: 10, volume : %f, luas alas: %f, keliling alas: %f", Volume(jari2, tinggi), LuasAlas(jari2, tinggi), KelilingAlas(jari2, tinggi))
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("starting web server at http://localhost:8080/")
	fmt.Printf("jariJari : 7, tinggi: 10, volume : %f, luas alas: %f, keliling alas: %f", Volume(jari2, tinggi), LuasAlas(jari2, tinggi), KelilingAlas(jari2, tinggi))

	http.ListenAndServe(":8080", nil)
}
