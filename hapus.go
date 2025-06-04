func hapusPengeluaran(data *[MaksData]Pengeluaran, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		fmt.Println("Daftar pengeluaran:")
		i := 0
		for i < *jumlah {
			fmt.Printf("%d. %s - %d\n", i+1, data[i].Kategori, data[i].Jumlah)
			i = i + 1
		}

		var index int
		fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
		fmt.Scanln(&index)

		if index >= 1 && index <= *jumlah {
			i := index - 1
			for i < *jumlah-1 {
				data[i] = data[i+1]
				i = i + 1
			}
			*jumlah = *jumlah - 1
			fmt.Println("Pengeluaran berhasil dihapus.")
		} else {
			fmt.Println("Nomor tidak valid.")
		}
	}
}
