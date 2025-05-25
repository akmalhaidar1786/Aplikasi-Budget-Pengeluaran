func hapus(app *Aplikasi) {
	var idx int
	var valid bool

	if app.jumlahData == 0 {
		fmt.Println("Belum ada pengeluaran!")
		return
	}

	daftar(app)
	for valid = false; !valid; {
		fmt.Print("Pilioh nomor 1-", app.jumlahData, ": ")
		fmt.Scan(&idx)
		if idx >= 1 && idx <= app.jumlahData {
			valid = true
		} else {
			fmt.Println("Input tidak valid!")
		}
	}
	idx--

	var i int
	for i = idx; i < app.jumlahData - 1; i++ {
		app.data[i] = app.data[i + 1]
	}
	app.jumlahData--
	fmt.Println("Pengeluaran berhasil dihapus!")
}
