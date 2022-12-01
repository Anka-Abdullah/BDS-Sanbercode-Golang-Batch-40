package tugas8

import (
	"fmt"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}
type Print interface {
	PrintPhone() map[string]interface{}
}

func LuasPersegi(numb int, boo bool) interface{} {
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
func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}
func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3

}
func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}
func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}
func (t Tabung) Volume() float64 {
	return 3.14 * t.JariJari * t.Tinggi
}
func (t Tabung) LuasPermukaan() float64 {
	return 2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi)
}
func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}
func (b Balok) LuasPermukaan() float64 {
	return float64(2 * ((b.Panjang * b.Lebar) + (b.Panjang * b.Tinggi) + (b.Lebar * b.Tinggi)))
}
func (p Phone) PrintPhone() map[string]interface{} {
	return map[string]interface{}{
		"Name":   p.Name,
		"Brand":  p.Brand,
		"Year":   p.Year,
		"Colors": p.Colors,
	}
}
