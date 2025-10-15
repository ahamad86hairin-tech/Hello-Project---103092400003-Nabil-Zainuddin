package main

import (
	"fmt"
	"strings"
)

const MAX = 100

type Lagu struct {
	Judul  string
	Artis  string
	Genre  string
	Rating int
}

type KoleksiMusik [MAX]Lagu

var daftar KoleksiMusik
var jumlah int

func tambahLagu(judul, artis, genre string, rating int) {
	if jumlah < MAX {
		daftar[jumlah] = Lagu{judul, artis, genre, rating}
		jumlah++
	} else {
		fmt.Println("Koleksi penuh.")
	}
}

func tampilkanSemua() {
	fmt.Println("Daftar Lagu:")
	for i := 0; i < jumlah; i++ {
		tampilkanLagu(daftar[i])
	}
}

func tampilkanLagu(l Lagu) {
	fmt.Printf("Judul : %s\nArtis : %s\nGenre : %s\nRating: %d\n\n", l.Judul, l.Artis, l.Genre, l.Rating)
}

func ubahLagu(idx int, judul, artis, genre string, rating int) {
	if idx >= 0 && idx < jumlah {
		daftar[idx] = Lagu{judul, artis, genre, rating}
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusLagu(idx int) {
	if idx >= 0 && idx < jumlah {
		for i := idx; i < jumlah-1; i++ {
			daftar[i] = daftar[i+1]
		}
		jumlah--
		fmt.Println("Lagu berhasil dihapus.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func sequentialSearchArtis(artis string) {
	found := false
	for i := 0; i < jumlah; i++ {
		if strings.ToLower(daftar[i].Artis) == strings.ToLower(artis) {
			tampilkanLagu(daftar[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Lagu dengan artis tersebut tidak ditemukan.")
	}
}

func insertionSortRating(asc bool) {
	for i := 1; i < jumlah; i++ {
		temp := daftar[i]
		j := i - 1
		if asc {
			for j >= 0 && daftar[j].Rating > temp.Rating {
				daftar[j+1] = daftar[j]
				j--
			}
		} else {
			for j >= 0 && daftar[j].Rating < temp.Rating {
				daftar[j+1] = daftar[j]
				j--
			}
		}
		daftar[j+1] = temp
	}
}

func selectionSortArtis(asc bool) {
	for i := 0; i < jumlah-1; i++ {
		idx := i
		for j := i + 1; j < jumlah; j++ {
			if asc {
				if strings.ToLower(daftar[j].Artis) < strings.ToLower(daftar[idx].Artis) {
					idx = j
				}
			} else {
				if strings.ToLower(daftar[j].Artis) > strings.ToLower(daftar[idx].Artis) {
					idx = j
				}
			}
		}
		daftar[i], daftar[idx] = daftar[idx], daftar[i]
	}
}

func binarySearchRating(target int) int {
	// Pastikan sudah diurutkan ascending dulu
	insertionSortRating(true)

	kiri := 0
	kanan := jumlah - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftar[tengah].Rating == target {
			return tengah
		} else if daftar[tengah].Rating < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func menu() {
	fmt.Println("=== Aplikasi Manajemen Koleksi Musik ===")
	fmt.Println("1. Tambah Lagu")
	fmt.Println("2. Tampilkan Semua Lagu")
	fmt.Println("3. Ubah Lagu")
	fmt.Println("4. Hapus Lagu")
	fmt.Println("5. Cari Lagu Berdasarkan Artis")
	fmt.Println("6. Cari Lagu Berdasarkan Rating (Binary Search)")
	fmt.Println("7. Urutkan Berdasarkan Rating (Insertion Sort)")
	fmt.Println("8. Urutkan Berdasarkan Artis (Selection Sort)")
	fmt.Println("0. Keluar")
}

func main() {
	var pilihan int
	for {
		menu()
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		if pilihan == 0 {
			break
		}
		switch pilihan {
		case 1:
			var judul, artis, genre string
			var rating int
			fmt.Print("Judul: ")
			fmt.Scan(&judul)
			fmt.Print("Artis: ")
			fmt.Scan(&artis)
			fmt.Print("Genre: ")
			fmt.Scan(&genre)
			fmt.Print("Rating (0-10): ")
			fmt.Scan(&rating)
			tambahLagu(judul, artis, genre, rating)
		case 2:
			tampilkanSemua()
		case 3:
			var idx int
			var judul, artis, genre string
			var rating int
			fmt.Print("Indeks lagu yang diubah: ")
			fmt.Scan(&idx)
			fmt.Print("Judul baru: ")
			fmt.Scan(&judul)
			fmt.Print("Artis baru: ")
			fmt.Scan(&artis)
			fmt.Print("Genre baru: ")
			fmt.Scan(&genre)
			fmt.Print("Rating baru: ")
			fmt.Scan(&rating)
			ubahLagu(idx, judul, artis, genre, rating)
		case 4:
			var idx int
			fmt.Print("Indeks lagu yang dihapus: ")
			fmt.Scan(&idx)
			hapusLagu(idx)
		case 5:
			var artis string
			fmt.Print("Masukkan nama artis: ")
			fmt.Scan(&artis)
			sequentialSearchArtis(artis)
		case 6:
			var rating int
			fmt.Print("Masukkan rating: ")
			fmt.Scan(&rating)
			idx := binarySearchRating(rating)
			if idx != -1 {
				tampilkanLagu(daftar[idx])
			} else {
				fmt.Println("Lagu dengan rating tersebut tidak ditemukan.")
			}
		case 7:
			var asc string
			fmt.Print("Ascending? (y/n): ")
			fmt.Scan(&asc)
			insertionSortRating(strings.ToLower(asc) == "y")
			fmt.Println("Berhasil diurutkan berdasarkan rating.")
		case 8:
			var asc string
			fmt.Print("Ascending? (y/n): ")
			fmt.Scan(&asc)
			selectionSortArtis(strings.ToLower(asc) == "y")
			fmt.Println("Berhasil diurutkan berdasarkan artis.")
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}
