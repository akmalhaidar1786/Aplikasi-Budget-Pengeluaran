func urutkanPengeluaranMenu(data *[MaksData]Pengeluaran, jumlah int) {
	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Selection Sort (berdasarkan kategori)")
	fmt.Println("2. Insertion Sort (berdasarkan jumlah)")

	var metode int = 0
	for metode != 1 && metode != 2 {
		fmt.Print("Pilih metode (1/2): ")
		fmt.Scanln(&metode)
		if metode != 1 && metode != 2 {
			fmt.Println("Metode tidak valid, coba lagi.")
		}
	}

	if metode == 1 {
		selectionSortKategori(data, jumlah)
		fmt.Println("Data sudah diurutkan berdasarkan kategori.")
	} else {
		insertionSortJumlah(data, jumlah)
		fmt.Println("Data sudah diurutkan berdasarkan jumlah.")
	}
}

func selectionSortKategori(data *[MaksData]Pengeluaran, jumlah int) {
	i := 0
	for i < jumlah-1 {
		minIndex := i
		j := i + 1
		for j < jumlah {
			if data[j].Kategori < data[minIndex].Kategori {
				minIndex = j
			}
			j = j + 1
		}
		if minIndex != i {
			temp := data[i]
			data[i] = data[minIndex]
			data[minIndex] = temp
		}
		i = i + 1
	}
}

func insertionSortJumlah(data *[MaksData]Pengeluaran, jumlah int) {
	i := 1
	for i < jumlah {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Jumlah > key.Jumlah {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
		i = i + 1
	}
}
