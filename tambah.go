func tambahPengeluaran(data *[MaksData]Pengeluaran, jumlah *int) {
	if *jumlah < MaksData {
		var kategori string
		var jumlahPengeluaran int

		fmt.Print("Masukkan kategori pengeluaran (transportasi/akomodasi/makanan/hiburan): ")
		fmt.Scanln(&kategori)
		fmt.Print("Masukkan jumlah pengeluaran: ")
		fmt.Scanln(&jumlahPengeluaran)

		data[*jumlah].Kategori = kategori
		data[*jumlah].Jumlah = jumlahPengeluaran
		*jumlah = *jumlah + 1

		fmt.Println("Pengeluaran berhasil ditambahkan.")
	} else {
		fmt.Println("Data pengeluaran penuh. Tidak bisa menambah lagi.")
	}
}
