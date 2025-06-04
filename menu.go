func main() {
	var dataPengeluaran [MaksData]Pengeluaran
	var jumlahData int
	var totalBudget int

	fmt.Print("Masukkan total budget perjalanan: ")
	fmt.Scanln(&totalBudget)

	var pilihan int = 0

	for pilihan != 7 {
		fmt.Println("\n=== Aplikasi Pengelolaan Budget Traveling ===")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Ubah Pengeluaran")
		fmt.Println("3. Hapus Pengeluaran")
		fmt.Println("4. Lihat Laporan")
		fmt.Println("5. Cari Pengeluaran")
		fmt.Println("6. Urutkan Pengeluaran")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tambahPengeluaran(&dataPengeluaran, &jumlahData)
		} else if pilihan == 2 {
			ubahPengeluaran(&dataPengeluaran, jumlahData)
		} else if pilihan == 3 {
			hapusPengeluaran(&dataPengeluaran, &jumlahData)
		} else if pilihan == 4 {
			tampilkanLaporan(&dataPengeluaran, jumlahData, totalBudget)
		} else if pilihan == 5 {
			cariPengeluaranMenu(&dataPengeluaran, jumlahData)
		} else if pilihan == 6 {
			urutkanPengeluaranMenu(&dataPengeluaran, jumlahData)
		} else if pilihan == 7 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
