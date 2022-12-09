package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func indexNilai(nilai int) string {
	if nilai < 50 {
		return "E"
	}
	if nilai < 60 {
		return "D"
	}
	if nilai < 70 {
		return "C"
	}
	if nilai < 80 {
		return "B"
	}
	return "A"
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}
		if uname == "anka" && pwd == "anka" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
	})
}

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataMahasiswa, err := json.Marshal(nilaiNilaiMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func SaveMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mahasiswa NilaiMahasiswa
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&Mahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if int(Mahasiswa.Nilai) > 100 {
			panic("Nilai hanya boleh diisi maksimal dengan angka 100")
		}

		Mahasiswa = NilaiMahasiswa{
			ID:          uint(len(nilaiNilaiMahasiswa) + 1),
			Nama:        Mahasiswa.Nama,
			MataKuliah:  Mahasiswa.MataKuliah,
			Nilai:       uint(Mahasiswa.Nilai),
			IndeksNilai: indexNilai(int(Mahasiswa.Nilai)),
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Mahasiswa)
		dataMahasiswa, _ := json.Marshal(Mahasiswa)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Success!")
	})

	http.HandleFunc("/mahasiswa", GetMahasiswa)
	http.Handle("/mahasiswa/save", Auth(http.HandlerFunc(SaveMahasiswa)))

	fmt.Println("starting web server at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
