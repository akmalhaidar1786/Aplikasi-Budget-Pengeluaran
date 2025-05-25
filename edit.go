func edit(app *Aplikasi) {
	var idx int
	var valid bool

	if app.jumlahData == 0 {
		fmt.Println("Belum ada pengeluaran!")
		return
	}
	
	daftar(app)
	
	for valid = false; !valid; {
		fmt.Print("pilih nomor (1-", app.jumlahData, "): ")
		fmt.Scan(&idx)

		if idx >= 1 && idx <= app.jumlahData {
			valid = true
		} else {
			fmt.Println("Inut tidak valid")
		}
	}
	idx--

	var ulangi bool = true
	var p Pengeluaran

	for ulangi {
		fmt.Print("Masukkan kategori baru: ")
		fmt.Scan(&p.kategori)

		valid = p.kategori == "transportasi" || p.kategori == "akomodasi" || 
		p.kategori == "makanan" || p.kategori == "Makanan" || p.kategori == "hiburan" || p.kategori == "Hiburan"

		if valid {
			ulangi = false
		} else {
			fmt.Println("Mohon maaf kategori tidak tersedia.")
		}
	}
	
	fmt.Print("Masukkan jumlah baru: ")
	fmt.Scan(&p.jumlah)

	app.data[idx] = p
	fmt.Println("Pengeluaran berhasil diedit!")
}
