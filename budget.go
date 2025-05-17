package main(){

}

import "fmt"

type pengeluaran struct {
	kategori string
	jumlah float64
}

totalBudget float64
pengeluaran [250]pengeluaran
jumlahData int


func selectionSort() {
	var i, j int
	var minIdx int
	var temp pengeluaran

	for i = 0; i < jumlahData - 1; i++ {

		
	}


}

func tambah() {
	var p pengeluaran
	
	if app.jumlahData >= 250 {
		fmt.Println("Maaf, sudah mencapai batas maksimal pengeluaran")
		return
	}

	fmt.Print("Kategori (transportasi/akomodasi/makanan/hiburan): ")
	fmt.Scan(&p.kategori)

	var valid bool = false
	if p.kategori == "transportasi" || p.kategori == "akomodasi" ||
	p.kategori == "makanan" || p.kategori == "hiburan" {
		valid = true
	}

	if !valid {
		fmt.Println("Mohon maaf kategori tidak tersedia.")
		return
	}

	fmt.Print("Jumlah")
	fmt.Scan(&p.Jumlah)

	pengeluaran[jumlahData] = p
	jumlahData++
	fmt.Println("Pengeluaran berhasil ditambahkan!")
}


func edit() {
	var idx int

	if app.jumlahData == 0 {
		fmt.Println("Belum ada pengeluaran!")
		return
	}
	
	
}

func main(){
	var app aplikasi
	var pilihan int

	fmt.Println("Masukkan total budget perjalanan: ")
	fmt.Scan(&totalBudget)

	for {
		fmt.Println("\nMenu Utama:")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Edit Pengeluaran")
		fmt.Println("3. Hapus Pengeluaran")
        fmt.Println("4. Tampilkan Laporan")
        fmt.Println("5. Cari Pengeluaran")
        fmt.Println("6. Urutkan Pengeluaran")
        fmt.Println("7. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&pilihan)
	}

	switch pilihan {
		case 1:
			tambahPengeluaran()
		case 2:
			edit()
		case 3:

		case 4:

		case 5:

		case 6:

		case 7:

		default:
			fmt.Println("Mohon maaf pilihan tidak tersedia.")
	}
}