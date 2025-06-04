func cariPengeluaranMenu(data *[MaksData]Pengeluaran, jumlah int) {
	fmt.Print("Masukkan kategori yang dicari: ")
	var kategoriCari string
	fmt.Scanln(&kategoriCari)

	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (Data harus sudah diurutkan)")

	var metode int = 0
	for metode != 1 && metode != 2 {
		fmt.Print("Pilih metode (1/2): ")
		fmt.Scanln(&metode)
		if metode != 1 && metode != 2 {
			fmt.Println("Metode tidak valid, coba lagi.")
		}
	}

	if metode == 1 {
		sequentialSearch(data, jumlah, kategoriCari)
	} else {
		binarySearch(data, jumlah, kategoriCari)
	}
}
