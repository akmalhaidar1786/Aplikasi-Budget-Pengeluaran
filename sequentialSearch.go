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
