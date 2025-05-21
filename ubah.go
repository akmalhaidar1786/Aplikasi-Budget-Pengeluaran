package main

import "fmt"

func ubahPengeluaran() {
	tampilkanSemua()

	var index int
	fmt.Print("Masukkan nomor pengeluaran yang ingin diubah: ")
	fmt.Scanln(&index)
	index-- 

	if index < 0 || index >= jumlahData {
		fmt.Println("Nomor tidak valid.")
		return
	}

	var kategoriBaru string
	var jumlahBaru float64
	fmt.Print("Masukkan kategori baru: ")
	fmt.Scanln(&kategoriBaru)
	fmt.Print("Masukkan jumlah baru (Rp): ")
	fmt.Scanln(&jumlahBaru)

	pengeluarans[index] = Pengeluaran{Kategori: kategoriBaru, Jumlah: jumlahBaru}
	fmt.Println("Pengeluaran berhasil diubah.")
}
