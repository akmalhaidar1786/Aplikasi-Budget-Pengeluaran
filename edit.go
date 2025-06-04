func ubahPengeluaran(data *[MaksData]Pengeluaran, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		fmt.Println("Daftar pengeluaran:")
		i := 0
		for i < jumlah {
			fmt.Printf("%d. %s - %d\n", i+1, data[i].Kategori, data[i].Jumlah)
			i = i + 1
		}

		var index int
		fmt.Print("Masukkan nomor pengeluaran yang ingin diubah: ")
		fmt.Scanln(&index)

		if index >= 1 && index <= jumlah {
			var kategoriBaru string
			var jumlahBaru int

			fmt.Print("Masukkan kategori baru: ")
			fmt.Scanln(&kategoriBaru)
			fmt.Print("Masukkan jumlah baru: ")
			fmt.Scanln(&jumlahBaru)

			data[index-1].Kategori = kategoriBaru
			data[index-1].Jumlah = jumlahBaru

			fmt.Println("Pengeluaran berhasil diubah.")
		} else {
			fmt.Println("Nomor tidak valid.")
		}
	}
}
