func laporan(app *Aplikasi) {
	var total float64
	var transport, akomodasi, makanan, hiburan float64
	var i int

	for i = 0; i < app.jumlahData; i++ {
		total += app.data[i].jumlah
		switch app.data[i].kategori {
			case "transportasi":
				transport += app.data[i].jumlah
			case "akomodasi":
				akomodasi += app.data[i].jumlah
			case "makanan":
				makanan += app.data[i].jumlah
			case "hiburan":
				hiburan += app.data[i].jumlah
		}
	}

	fmt.Printf("\nTotal budget: Rp.%.2f\n", app.totalBudget)
	fmt.Printf("Total pengeluaran: Rp.%.2f\n", total)
	fmt.Printf("Selisih: Rp.%.2f\n", app.totalBudget - total)

	fmt.Println("\nRincian per Kategori: ")
	fmt.Printf("- Transportasi: Rp.%.2f\n", transport)
	fmt.Printf("- Akomodasi: Rp.%.2f\n", akomodasi)
	fmt.Printf("- Makanan: Rp.%.2f\n", makanan)
	fmt.Printf("- Hiburan: Rp.%.2f\n", hiburan)

	var saran bool
	var batas float64

	batas = app.totalBudget * 0.3

	fmt.Println("\nSaran Penghematan:")
	if transport > batas {
		fmt.Printf("Kurangi pengeluaran transport sebesar Rp%.2f!\n", transport - batas)
		saran = true
	}
	if akomodasi > batas {
		fmt.Printf("Kurangi pengeluaran akomodasi sebesar Rp.%.2f!\n", akomodasi - batas)
		saran = true
	}
	if makanan > batas {
		fmt.Printf("Kurangi pengeluaran makanan sebesar Rp.%.2f!\n", makanan - batas)
		saran = true
	}
	if hiburan > batas {
		fmt.Printf("Kurangi pengeluaran hiburan sebesar Rp.%.2f!\n", hiburan - batas)
		saran = true
	}

	if !saran {
		fmt.Println("Pengeluaran anda sudah efisien!\n")
	}
}
