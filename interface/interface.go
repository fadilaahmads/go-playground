package main

import (
	"fmt"
	"math"
)

// Interface definition
type hitung interface {
	luas() float64
	keliling() float64
}

// Lingkaran
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// Persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func (p persegi) volumeKubus(tinggi float64) float64 {
	return p.luas() * tinggi
}

func main() {
	// Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja),
	// yang dibungkus dengan nama tertentu.
	// Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah nil.
	// Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki definisi method
	// minimal sama dengan yang ada di interface-nya

	var bangunDatar hitung

	bangunDatar = persegi{10}
	fmt.Println("=== Persegi ===")
	fmt.Println("Persegi keliling:", bangunDatar.keliling())
	fmt.Println("Persegi luas:", bangunDatar.luas())
	fmt.Println("Persegi volume:", bangunDatar.(persegi).volumeKubus(5))

	bangunDatar = lingkaran{14}
	fmt.Println("=== Lingkaran ===")
	fmt.Println("Lingkaran keliling", bangunDatar.keliling())
	fmt.Println("Lingkaran luas:", bangunDatar.luas())
	// Untuk mengakses method yang tidak ter-definisi di interface,
	// variabel-nya harus di-casting terlebih dahulu ke tipe asli variabel konkritnya (pada kasus ini tipenya lingkaran),
	// setelahnya method akan bisa diakses.
	// cara casting menuliskan nama tipe tujuan dalam kurung,
	// ditempatkan setelah nama interface dengan menggunakan notasi titik (seperti cara mengakses property, hanya saja ada tanda kurung nya).
	fmt.Println("Lingkaran jari-jari:", bangunDatar.(lingkaran).jarijari())

}
