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
