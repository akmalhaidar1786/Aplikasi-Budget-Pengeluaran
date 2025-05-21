func menu() {
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Ubah Pengeluaran")
		fmt.Println("3. Hapus Pengeluaran")
		fmt.Println("4. Tampilkan Semua")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tambahPengeluaran()
		} else if pilihan == 2 {
			ubahPengeluaran()
		} else if pilihan == 3 {
			hapusPengeluaran()
		} else if pilihan == 4 {
			tampilkanSemua()
		} else if pilihan == 5 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		}
	}
}
