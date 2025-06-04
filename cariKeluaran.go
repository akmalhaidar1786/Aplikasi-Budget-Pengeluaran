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

func sequentialSearch(data *[MaksData]Pengeluaran, jumlah int, kategoriCari string) {
	found := false
	i := 0
	for i < jumlah {
		if data[i].Kategori == kategoriCari {
			fmt.Printf("Data ke-%d: %s - %d\n", i+1, data[i].Kategori, data[i].Jumlah)
			found = true
		}
		i = i + 1
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func binarySearch(data *[MaksData]Pengeluaran, jumlah int, kategoriCari string) {
	low := 0
	high := jumlah - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if data[mid].Kategori == kategoriCari {
			fmt.Printf("Data ditemukan di index ke-%d: %s - %d\n", mid+1, data[mid].Kategori, data[mid].Jumlah)
			found = true
			low = high + 1
		} else if data[mid].Kategori < kategoriCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}
