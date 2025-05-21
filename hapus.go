package main

import "fmt"

func hapusPengeluaran() {
	tampilkanSemua()

	var index int
	fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
	fmt.Scanln(&index)
	index--

	for i := index; i < jumlahData-1; i++ {
		pengeluarans[i] = pengeluarans[i+1]
	}

	jumlahData--
	fmt.Println("Pengeluaran berhasil dihapus.")
}
