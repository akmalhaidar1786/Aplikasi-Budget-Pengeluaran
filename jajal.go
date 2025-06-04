package main

import (
	"fmt"
)

const MaksData = 100

type Pengeluaran struct {
	Kategori string
	Jumlah   int
}

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
		fmt.Print("Pilih menu (1-7): ")
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

func tambahPengeluaran(data *[MaksData]Pengeluaran, jumlah *int) {
	if *jumlah < MaksData {
		var kategori string
		var jumlahPengeluaran int

		fmt.Print("Masukkan kategori pengeluaran (transportasi/akomodasi/makanan/hiburan): ")
		fmt.Scanln(&kategori)
		fmt.Print("Masukkan jumlah pengeluaran: ")
		fmt.Scanln(&jumlahPengeluaran)

		data[*jumlah].Kategori = kategori
		data[*jumlah].Jumlah = jumlahPengeluaran
		*jumlah = *jumlah + 1

		fmt.Println("Pengeluaran berhasil ditambahkan.")
	} else {
		fmt.Println("Data pengeluaran penuh. Tidak bisa menambah lagi.")
	}
}

func ubahPengeluaran(data *[MaksData]Pengeluaran, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		fmt.Println("Daftar pengeluaran:")
		i := 0
		for i < jumlah {
			fmt.Printf("%d. %s - %d\n", i+1, data[i].Kategori, data[i].Jumlah)
			i = i + 1
		}

		var index int
		fmt.Print("Masukkan nomor pengeluaran yang ingin diubah: ")
		fmt.Scanln(&index)

		if index >= 1 && index <= jumlah {
			var kategoriBaru string
			var jumlahBaru int

			fmt.Print("Masukkan kategori baru: ")
			fmt.Scanln(&kategoriBaru)
			fmt.Print("Masukkan jumlah baru: ")
			fmt.Scanln(&jumlahBaru)

			data[index-1].Kategori = kategoriBaru
			data[index-1].Jumlah = jumlahBaru

			fmt.Println("Pengeluaran berhasil diubah.")
		} else {
			fmt.Println("Nomor tidak valid.")
		}
	}
}

func hapusPengeluaran(data *[MaksData]Pengeluaran, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		fmt.Println("Daftar pengeluaran:")
		i := 0
		for i < *jumlah {
			fmt.Printf("%d. %s - %d\n", i+1, data[i].Kategori, data[i].Jumlah)
			i = i + 1
		}

		var index int
		fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
		fmt.Scanln(&index)

		if index >= 1 && index <= *jumlah {
			i := index - 1
			for i < *jumlah-1 {
				data[i] = data[i+1]
				i = i + 1
			}
			*jumlah = *jumlah - 1
			fmt.Println("Pengeluaran berhasil dihapus.")
		} else {
			fmt.Println("Nomor tidak valid.")
		}
	}
}

func tampilkanLaporan(data *[MaksData]Pengeluaran, jumlah int, totalBudget int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		totalPengeluaran := 0
		i := 0
		for i < jumlah {
			totalPengeluaran = totalPengeluaran + data[i].Jumlah
			i = i + 1
		}

		fmt.Println("\n--- Laporan Pengeluaran ---")
		fmt.Printf("Total Budget: %d\n", totalBudget)
		fmt.Printf("Total Pengeluaran: %d\n", totalPengeluaran)
		fmt.Printf("Sisa Budget: %d\n", totalBudget-totalPengeluaran)

		fmt.Println("\nRincian pengeluaran per kategori:")
		sumTransport := 0
		sumAkomodasi := 0
		sumMakanan := 0
		sumHiburan := 0

		i = 0
		for i < jumlah {
			if data[i].Kategori == "transportasi" {
				sumTransport = sumTransport + data[i].Jumlah
			} else if data[i].Kategori == "akomodasi" {
				sumAkomodasi = sumAkomodasi + data[i].Jumlah
			} else if data[i].Kategori == "makanan" {
				sumMakanan = sumMakanan + data[i].Jumlah
			} else if data[i].Kategori == "hiburan" {
				sumHiburan = sumHiburan + data[i].Jumlah
			}
			i = i + 1
		}

		fmt.Printf("Transportasi: %d\n", sumTransport)
		fmt.Printf("Akomodasi: %d\n", sumAkomodasi)
		fmt.Printf("Makanan: %d\n", sumMakanan)
		fmt.Printf("Hiburan: %d\n", sumHiburan)
	}
}

func cariPengeluaranMenu(data *[MaksData]Pengeluaran, jumlah int) {
	fmt.Print("Masukkan kategori yang dicari: ")
	var kategoriCari string
	fmt.Scanln(&kategoriCari)

	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (Data harus sudah diurutkan)")

	var metode int = 0
	for metode != 1 && metode != 2 {
		fmt.Print("Pilih metode (1/2): ")
		fmt.Scanln(&metode)
		if metode != 1 && metode != 2 {
			fmt.Println("Metode tidak valid, coba lagi.")
		}
	}

	if metode == 1 {
		sequentialSearch(data, jumlah, kategoriCari)
	} else {
		binarySearch(data, jumlah, kategoriCari)
	}
}

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

func binarySearch(data *[MaksData]Pengeluaran, jumlah int, kategoriCari string) {
	low := 0
	high := jumlah - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if data[mid].Kategori == kategoriCari {
			fmt.Printf("Data ditemukan di index ke-%d: %s - %d\n", mid+1, data[mid].Kategori, data[mid].Jumlah)
			found = true
			low = high + 1 // keluar loop
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

func urutkanPengeluaranMenu(data *[MaksData]Pengeluaran, jumlah int) {
	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Selection Sort (berdasarkan kategori)")
	fmt.Println("2. Insertion Sort (berdasarkan jumlah)")

	var metode int = 0
	for metode != 1 && metode != 2 {
		fmt.Print("Pilih metode (1/2): ")
		fmt.Scanln(&metode)
		if metode != 1 && metode != 2 {
			fmt.Println("Metode tidak valid, coba lagi.")
		}
	}

	if metode == 1 {
		selectionSortKategori(data, jumlah)
		fmt.Println("Data sudah diurutkan berdasarkan kategori.")
	} else {
		insertionSortJumlah(data, jumlah)
		fmt.Println("Data sudah diurutkan berdasarkan jumlah.")
	}
}

func selectionSortKategori(data *[MaksData]Pengeluaran, jumlah int) {
	i := 0
	for i < jumlah-1 {
		minIndex := i
		j := i + 1
		for j < jumlah {
			if data[j].Kategori < data[minIndex].Kategori {
				minIndex = j
			}
			j = j + 1
		}
		if minIndex != i {
			temp := data[i]
			data[i] = data[minIndex]
			data[minIndex] = temp
		}
		i = i + 1
	}
}

func insertionSortJumlah(data *[MaksData]Pengeluaran, jumlah int) {
	i := 1
	for i < jumlah {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Jumlah > key.Jumlah {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
		i = i + 1
	}
}
