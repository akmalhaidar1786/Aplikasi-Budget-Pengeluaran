package main

import "fmt"

type Pengeluaran struct {
	kategori string
	jumlah float64
}

type Aplikasi struct {
	totalBudget float64
	data [250]Pengeluaran
	jumlahData int
}

func menu() {
	fmt.Println("\nMenu Utama:")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Edit Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Tampilkan Laporan")
	fmt.Println("5. Cari Pengeluaran")
	fmt.Println("6. Urutkan Pengeluaran")
	fmt.Println("7. Keluar")
	fmt.Print("Pilih menu: ")
}
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

func daftar(app *Aplikasi) {
    fmt.Println("\nDaftar Pengeluaran:")
    var i int
    for i = 0; i < app.jumlahData; i++ {
        fmt.Printf("%d. [%s] Rp%.2f\n", 
            i+1, app.data[i].kategori, app.data[i].jumlah)
    }
}

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

func selectionSort(app *Aplikasi, jenis string, asc bool) {
	var i, extIdx, pass int
	var temp Pengeluaran

	for pass = 0; pass < app.jumlahData - 1; pass++ {
		extIdx = pass
		for i = pass + 1; i < app.jumlahData; i++ {
			var kondisi bool

			switch jenis {
				case "kategori":
					if asc {
						kondisi = app.data[i].kategori < app.data[extIdx].kategori
					}
			}
		}
	}
}


// func insertionSort(){

// }

func main(){
	var app Aplikasi
	var pilihan int
	var berjalan bool = true

	fmt.Print("Masukkan total budget perjalanan: ")
	fmt.Scan(&app.totalBudget)
	
	for berjalan {
		menu()
		fmt.Scan(&pilihan)

		switch pilihan {
			case 1:
				tambah(&app)
			case 2:
				edit(&app)
			case 3:
				hapus(&app)
			case 4:
				laporan(&app)
			case 5:
				cari()
			case 6:
				urut()
			case 7:
				berjalan = false
			default:
				fmt.Println("Mohon maaf pilihan tidak tersedia.")
		}

	}
}
