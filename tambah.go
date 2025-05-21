package main

import "fmt"

func tambahPengeluaran() {
	var kategori string
	var jumlah float64

	fmt.Print("Masukkan kategori: ")
	fmt.Scanln(&kategori)
	fmt.Print("Masukkan jumlah pengeluaran (Rp): ")
	fmt.Scanln(&jumlah)

	pengeluarans[jumlahData] = Pengeluaran{Kategori: kategori, Jumlah: jumlah}
	jumlahData++
	fmt.Println("Pengeluaran berhasil ditambahkan.")
}
