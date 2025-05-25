func tambah(app *Aplikasi) {
	var p Pengeluaran
	
	if app.jumlahData >= 250 {
		fmt.Println("Maaf, sudah mencapai batas maksimal pengeluaran")
		return
	}


	var ulangi bool = true
	for ulangi {	
		fmt.Print("Kategori (Transportasi/Akomodasi/Makanan/Hiburan): ")
		fmt.Scan(&p.kategori)

		var valid bool 
		valid = p.kategori == "transportasi" || p.kategori == "akomodasi" || 
		p.kategori == "makanan" || p.kategori == "Makanan" || p.kategori == "hiburan" || p.kategori == "Hiburan"
			
		if valid {
			ulangi = false
		} else {
			fmt.Println("Mohon maaf kategori tidak tersedia.")
		}
	}

	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&p.jumlah)

	app.data[app.jumlahData] = p
	app.jumlahData++
	fmt.Println("Pengeluaran berhasil ditambahkan!")
}
