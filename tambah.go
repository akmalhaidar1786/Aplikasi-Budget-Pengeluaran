package main

import "fmt"

func tambahPengeluaran() {
	var kategori string
	var jumlah float64

	fmt.Print("Masukkan kategori: ")
	fmt.Scanln(&kategori)
	fmt.Print("Masukkan jumlah pengeluaran (Rp): ")
	fmt.Scanln(&jumlah)

	var p Pengeluaran
	p.Kategori = kategori
	p.Jumlah = jumlah
	pengeluarans[jumlahData] = p
	jumlahData++
	fmt.Println("Pengeluaran berhasil ditambahkan.")
}
