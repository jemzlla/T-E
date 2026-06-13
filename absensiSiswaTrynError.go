package main

import "fmt"

const NMAX int = 100

type Siswa struct {
	nama             string
	kelas            string
	jumlahHadir      int
	jumlahTidakHadir int
}

type ArraySiswa struct {
	data [NMAX]Siswa
	n    int
}

type AbsensiHarian struct {
	tanggal string
	siswa   [NMAX]string
	status  [NMAX]bool
	n       int
}

type ArrayAbsensi struct {
	data [NMAX]AbsensiHarian
	n    int
}

// Prosedur insertionSortByNama untuk emngurutkan nama siswa
// IS: aeeSiswa terdefinisi, berisi data siswa yang belum terurut
// FS: arrSiswa.data terurut ascending berdasarkan nama siswa (A-Z)
func insertionSortByNama(arrSiswa *ArraySiswa) {
	for i := 1; i < arrSiswa.n; i++ {
		test := arrSiswa.data[i]
		j := i - 1
		for j >= 0 && arrSiswa.data[j].nama > test.nama {
			arrSiswa.data[j+1] = arrSiswa.data[j]
			j--
		}
		arrSiswa.data[j+1] = test
	}
}

// Prosedur selectionSortByPersentase
// IS: arrSiswa terdefinisi, berisi data siswa dengan jumlah hadir/tidak hadir
// FS: arrSiswa.data terurut berdasarkan persentase kehadiran; ascending = true, descending = false
func selectionSortByPersentase(arrSiswa *ArraySiswa, ascending bool) {
	for i := 0; i < arrSiswa.n-1; i++ {
		idxTarget := i
		for j := i + 1; j < arrSiswa.n; j++ {
			totalTarget := arrSiswa.data[idxTarget].jumlahHadir + arrSiswa.data[idxTarget].jumlahTidakHadir
			totalJ := arrSiswa.data[j].jumlahHadir + arrSiswa.data[j].jumlahTidakHadir

			pTarget := 0.0
			pJ := 0.0
			if totalTarget > 0 {
				pTarget = float64(arrSiswa.data[idxTarget].jumlahHadir) / float64(totalTarget)
			}
			if totalJ > 0 {
				pJ = float64(arrSiswa.data[j].jumlahHadir) / float64(totalJ)
			}
			if ascending {
				if pJ < pTarget {
					idxTarget = j
				}
			} else {
				if pJ > pTarget {
					idxTarget = j
				}
			}
		}
		if idxTarget != i {
			arrSiswa.data[i], arrSiswa.data[idxTarget] = arrSiswa.data[idxTarget], arrSiswa.data[i]
		}
	}
}

// Fungsi sequentialSearchSiswa
func sequentialSearchSiswa(arrSiswa ArraySiswa, nama string) int {
	hasil := -1
	for i := 0; i < arrSiswa.n; i++ {
		if arrSiswa.data[i].nama == nama {
			hasil = i
		}
	}
	return hasil
}

// Fungsi sequentialSearchTanggal
// IS: arrAbsensi terdefinisi, tanggal terdefinisi
// FS: mengembalikan index absensi jika tanggal ditemukan, jika tidak -1 dan pencarian dilakukan satu per satu dari index ke 0 hingga ke n-1
func sequentialSearchTanggal(arrAbsensi ArrayAbsensi, tanggal string) int {
	hasil := -1
	for i := 0; i < arrAbsensi.n; i++ {
		if arrAbsensi.data[i].tanggal == tanggal {
			hasil = i
		}
	}
	return hasil
}

// Fungsi bianrySearchSiswa
//
func binarySearchSiswa(arrSiswa ArraySiswa, nama string) int {
	low := 0
	high := arrSiswa.n - 1
	hasil := -1

	for low <= high && hasil == -1 {
		mid := (low + high) / 2

		if arrSiswa.data[mid].nama == nama {
			hasil = mid
		} else if arrSiswa.data[mid].nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return hasil
}

// Prosedur tambahSiswaAbsensi (gabungan Tambah Siswa + Input Absensi)
// IS: arrSiswa dan arrAbsensi terdefinisi
// FS: jika siswa belum ada -> siswa baru ditambahkan lalu langsung input absensi untuk hari itu
//	jika siswa sudah ada -> langsung input absensi untuk hari itu
//	arrAbsensi diperbarui dengan entri tanggal yang sesuai (baru atau yang sudah ada)
func tambahSiswaAbsensi(arrSiswa *ArraySiswa, arrAbsensi *ArrayAbsensi) {
	var nama string
	var tanggal string
	var statusInt int
	var kelas string

	fmt.Print("Nama siswa: ")
	fmt.Scan(&nama)

	// Cek apakah siswa sudah ada dengan Sequential Search
	idx := sequentialSearchSiswa(*arrSiswa, nama)

	// Jika siswa belum ada, daftarkan dulu
	if idx == -1 {
		if arrSiswa.n >= NMAX {
			fmt.Println("Data siswa sudah penuh!")
		} else {
			fmt.Print("Kelas: ")
			fmt.Scan(&kelas)

			arrSiswa.data[arrSiswa.n].nama = nama
			arrSiswa.data[arrSiswa.n].kelas = kelas
			arrSiswa.data[arrSiswa.n].jumlahHadir = 0
			arrSiswa.data[arrSiswa.n].jumlahTidakHadir = 0
			idx = arrSiswa.n
			arrSiswa.n++

			fmt.Println("Siswa baru berhasil didaftarkan!")
		}
	} else {
		fmt.Printf("Siswa %s ditemukan (Kelas: %s)\n", arrSiswa.data[idx].nama, arrSiswa.data[idx].kelas)
	}

	// Lanjut input absensi jika idx valid
	if idx != -1 {
		fmt.Print("Tanggal (DD-MM-YYYY): ")
		fmt.Scan(&tanggal)
		fmt.Print("Status kehadiran (1=Hadir, 0=Tidak Hadir): ")
		fmt.Scan(&statusInt)

		// Cari array absensi pada tanggal tersebut
		idxAbsen := sequentialSearchTanggal(*arrAbsensi, tanggal)

		// Jika tanggal belum ada, buat array baru
		if idxAbsen == -1 {
			if arrAbsensi.n >= NMAX {
				fmt.Println("Data absensi sudah penuh!")
			} else {
				arrAbsensi.data[arrAbsensi.n].tanggal = tanggal
				arrAbsensi.data[arrAbsensi.n].n = 0
				idxAbsen = arrAbsensi.n
				arrAbsensi.n++
			}
		}

		// Tambahkan siswa ke array absensi tanggal tersebut jika idxAbsen valid
		if idxAbsen != -1 {
			pos := arrAbsensi.data[idxAbsen].n
			arrAbsensi.data[idxAbsen].siswa[pos] = nama
			arrAbsensi.data[idxAbsen].status[pos] = (statusInt == 1)
			arrAbsensi.data[idxAbsen].n++

			if statusInt == 1 {
				arrSiswa.data[idx].jumlahHadir++
				fmt.Printf("Absensi %s tanggal %s: Hadir - berhasil disimpan!\n", nama, tanggal)
			} else {
				arrSiswa.data[idx].jumlahTidakHadir++
				fmt.Printf("Absensi %s tanggal %s: Tidak Hadir - berhasil disimpan!\n", nama, tanggal)
			}
		}
	}
}

// Prosedur untuk mengubah data siswa (menggunakan binary search)
// IS: arrSiswa terdefinisi, nama siswa yang ingin diubah terdefinisi
// FS: data siswa dengan nama yang dicari berhasil diubah (nama dan kelas), jika tidak ditemukan maka tidak ada perubahannya. Menggunakan Insertion Sort + Binary Search untuk pencarian
func ubahSiswa(arrSiswa *ArraySiswa) {
	var nama string
	fmt.Print("Nama siswa yang ingin diubah: ")
	fmt.Scan(&nama)

	//Sebelum Binary Search lakukan proses Insertion Sort terlebih dahulu
	insertionSortByNama(arrSiswa)

	// Binary Search untuk cari nama siswa
	idx := binarySearchSiswa(*arrSiswa, nama)
	if idx == -1 {
		fmt.Println("Siswa tidak ditemukan")
	} else {
		fmt.Println("Data baru:")
		fmt.Print("Nama: ")
		fmt.Scan(&arrSiswa.data[idx].nama)
		fmt.Print("Kelas: ")
		fmt.Scan(&arrSiswa.data[idx].kelas)
		fmt.Println("Data siswa berhasil diubah!")
	}
}

// Prosedure untuk menghapus data siswa
// IS: arrSiswa terdefinisi, nama siswa yang ingin dihapus telah terdefinisi
// FS: data siswa dengan nama yang dicari berhasil dihapus, ketika arrSiswa.n berkurang 1 maka data digeser ke kiri dan jika tidak ditemukan maka tidak ada perubahan.
// menggunakan Insertion Sort + Binary Search untuk pencarian
func hapusSiswa(arrSiswa *ArraySiswa) {
	var nama string
	fmt.Print("Nama siswa yang ingin dihapus: ")
	fmt.Scan(&nama)

	//Sebelum Binary Search lakukan proses Insertion Sort terlebih dahulu
	insertionSortByNama(arrSiswa)

	// Binary Search untuk cari nama siswa
	idx := binarySearchSiswa(*arrSiswa, nama)
	if idx == -1 {
		fmt.Println("Siswa tidak ditemukan")
	} else {
		// Menggeser data ke kiri
		for i := idx; i < arrSiswa.n-1; i++ {
			arrSiswa.data[i] = arrSiswa.data[i+1]
		}
		arrSiswa.n--

		fmt.Println("Data siswa berhasil dihapus!")
	}
}

// Prosedure cariKehadiranTanggal
// IS: arrAbsensi dan tanggal terdefinsisi
// FS: menampilkan daftar kehadiran siswa pada tanggal yang dicari, jika tidak tampilkan pesan error.
//	menggunakan Sequential Search untuk mencari tanggal
func cariKehadiranTanggal(arrAbsensi ArrayAbsensi) {
	var tanggal string
	fmt.Print("Tanggal (DD-MM-YYYY): ")
	fmt.Scan(&tanggal)

	if arrAbsensi.n == 0 {
		fmt.Println("Belum ada data absensi")
	} else {
		idx := sequentialSearchTanggal(arrAbsensi, tanggal)
		if idx == -1 {
			fmt.Println("Data absensi tanggal tersebut tidak ditemukan!")
		} else {
			fmt.Printf("\n=== Absensi tanggal %s ===\n", tanggal)
			for j := 0; j < arrAbsensi.data[idx].n; j++ {
				status := "Tidak Hadir"
				if arrAbsensi.data[idx].status[j] {
					status = "Hadir"
				}
				fmt.Printf("%s: %s\n", arrAbsensi.data[idx].siswa[j], status)
			}
		}
	}
}

// Prosedur cariSiswaMenu
// IS: arrSiswa, nama terdefinisi
// FS: menampilkan data siswa yang ditemukan beserta statistik kehadirannya,
//	menggunakan Insertion Sort + Binary Search untuk pencarian
func cariSiswaMenu(arrSiswa *ArraySiswa) {
	var nama string

	if arrSiswa.n == 0 {
		fmt.Println("Belum ada data siswa!")
	} else {
		fmt.Print("Nama siswa yang dicari: ")
		fmt.Scan(&nama)

		insertionSortByNama(arrSiswa)
		idx := binarySearchSiswa(*arrSiswa, nama)

		if idx == -1 {
			fmt.Println("Siswa tidak ditemukan!")
		} else {
			total := arrSiswa.data[idx].jumlahHadir + arrSiswa.data[idx].jumlahTidakHadir
			persentase := 0.0
			if total > 0 {
				persentase = float64(arrSiswa.data[idx].jumlahHadir) / float64(total) * 100
			}
			fmt.Println("\n=== DATA SISWA DITEMUKAN ===")
			fmt.Printf("Nama       : %s\n", arrSiswa.data[idx].nama)
			fmt.Printf("Kelas      : %s\n", arrSiswa.data[idx].kelas)
			fmt.Printf("Hadir      : %d kali\n", arrSiswa.data[idx].jumlahHadir)
			fmt.Printf("Tidak Hadir: %d kali\n", arrSiswa.data[idx].jumlahTidakHadir)
			fmt.Printf("Persentase kehadiran: %.2f%%\n", persentase)
		}
	}
}

// Prosedur tampilkanStatistik
// IS: arrSiswa terdefinisi dan tidak dapat kosong
// FS: menampilkan nama, kelas, jumlah hadir, tidak hadir, dan persentase kehadiran setiap siswa
func tampilkanStatistik(arrSiswa *ArraySiswa) {
	if arrSiswa.n == 0 {
		fmt.Println("Belum ada data siswa!")
	} else {
		fmt.Println("\n=== Statistik Kehadiran Per Siswa ===")
		fmt.Println("=====================================")

		for i := 0; i < arrSiswa.n; i++ {
			total := arrSiswa.data[i].jumlahHadir + arrSiswa.data[i].jumlahTidakHadir
			persentase := 0.0
			if total > 0 {
				persentase = float64(arrSiswa.data[i].jumlahHadir) / float64(total) * 100
			}

			fmt.Printf("\nNama        : %s\n", arrSiswa.data[i].nama)
			fmt.Printf("Kelas       : %s\n", arrSiswa.data[i].kelas)
			fmt.Printf("Hadir       : %d kali\n", arrSiswa.data[i].jumlahHadir)
			fmt.Printf("Tidak Hadir : %d kali\n", arrSiswa.data[i].jumlahTidakHadir)
			fmt.Printf("Persentase  : %.2f%%\n", persentase)
			fmt.Println("-------------------------------------")
		}
	}
}

// Prosedur tampilkanSiswaTerurut
// IS: arrSiswa terdefinisi dan tidak dapat kosong; ascending = 1, descending =2
// FS: arrSiswa.data terurt berdasarkan persentase kehadiran, menampilkan daftar siswa beserta persentase kehadiran, menampilkan daftar siswa beserta persentase kehadiran dan menggunakan Selection Sort untuk proses pengurutannya
func tampilkanSiswaTerurut(arrSiswa *ArraySiswa) {
	var urutan int
	valid := true

	if arrSiswa.n == 0 {
		fmt.Println("Belum ada data siswa!")
		valid = false
	} else {
		fmt.Println("Urutan:")
		fmt.Println("1. Ascending (terendah ke tertinggi)")
		fmt.Println("2. Descending (tertinggi ke terendah)")
		fmt.Print("Pilih: ")
		fmt.Scan(&urutan)
	}
	if urutan == 1 {
		selectionSortByPersentase(arrSiswa, true)
	} else if urutan == 2 {
		selectionSortByPersentase(arrSiswa, false)
	} else {
		fmt.Println("Pilihan tidak valid!")
		valid = false
	}

	if valid == true {
		fmt.Println("\n=== SISWA TERURUT (PERSENTASE KEHADIRAN) ===")
		fmt.Println("=============================================")

		for i := 0; i < arrSiswa.n; i++ {
			total := arrSiswa.data[i].jumlahHadir + arrSiswa.data[i].jumlahTidakHadir
			persentase := 0.0
			if total > 0 {
				persentase = float64(arrSiswa.data[i].jumlahHadir) / float64(total) * 100
			}
			fmt.Printf("%d. %s (%s) - %.2f%%\n", i+1, arrSiswa.data[i].nama, arrSiswa.data[i].kelas, persentase)
		}
	}
}

// Prosedur cariMinMaxKehadiran
// IS: arrSiswa terdefinisi dan tidak kosong
// FS: menampilkan siswa dengan persentase kehadiran tertinggi (max) dan terendah (min)
// menggunakan sequential scan untuk menemukan min dan max
func cariMinMaxKehadiran(arrSiswa *ArraySiswa) {
	if arrSiswa.n == 0 {
		fmt.Println("Belum ada data siswa!")
	} else {
		idxMin := 0
		idxMax := 0

		for i := 1; i < arrSiswa.n; i++ {
			totalI := arrSiswa.data[i].jumlahHadir + arrSiswa.data[i].jumlahTidakHadir
			totalMin := arrSiswa.data[idxMin].jumlahHadir + arrSiswa.data[idxMin].jumlahTidakHadir
			totalMax := arrSiswa.data[idxMax].jumlahHadir + arrSiswa.data[idxMax].jumlahTidakHadir

			pI := 0.0
			pMin := 0.0
			pMax := 0.0

			if totalI > 0 {
				pI = float64(arrSiswa.data[i].jumlahHadir) / float64(totalI)
			}
			if totalMin > 0 {
				pMin = float64(arrSiswa.data[idxMin].jumlahHadir) / float64(totalMin)
			}
			if totalMax > 0 {
				pMax = float64(arrSiswa.data[idxMax].jumlahHadir) / float64(totalMax)
			}

			if pI < pMin {
				idxMin = i
			}
			if pI > pMax {
				idxMax = i
			}
		}

		totalMax := arrSiswa.data[idxMax].jumlahHadir + arrSiswa.data[idxMax].jumlahTidakHadir
		totalMin := arrSiswa.data[idxMin].jumlahHadir + arrSiswa.data[idxMin].jumlahTidakHadir
		pMax := 0.0
		pMin := 0.0
		if totalMax > 0 {
			pMax = float64(arrSiswa.data[idxMax].jumlahHadir) / float64(totalMax) * 100
		}
		if totalMin > 0 {
			pMin = float64(arrSiswa.data[idxMin].jumlahHadir) / float64(totalMin) * 100
		}

		fmt.Println("\n=== MIN / MAX KEHADIRAN ===")
		fmt.Println("Kehadiran Tertinggi:")
		fmt.Printf("  Nama      : %s\n", arrSiswa.data[idxMax].nama)
		fmt.Printf("  Kelas     : %s\n", arrSiswa.data[idxMax].kelas)
		fmt.Printf("  Persentase: %.2f%%\n", pMax)
		fmt.Println("Kehadiran Terendah:")
		fmt.Printf("  Nama      : %s\n", arrSiswa.data[idxMin].nama)
		fmt.Printf("  Kelas     : %s\n", arrSiswa.data[idxMin].kelas)
		fmt.Printf("  Persentase: %.2f%%\n", pMin)
		fmt.Println("===========================")
	}
}

// Prosedur tampilkanMenu
// IS: tidak ada
// FS: menampilkan menu pilihan ke layar
func tampilkanMenu() {
	fmt.Println("\n=========================================")
	fmt.Println("Aplikasi Absensi Siswa")
	fmt.Println("===========================================")
	fmt.Println("1. Tambah Siswa + Input Absensi  Note: Dont use space use _ instead")
	fmt.Println("2. Ubah Data Siswa		[Binary Search]")
	fmt.Println("3. Hapus Siswa		    [Binary Search]")
	fmt.Println("4. Cari Kehadiran      [Sequential Search]")
	fmt.Println("5. Cari Siswa		    [Binary Search]")
	fmt.Println("6. Statistik Kehadiran")
	fmt.Println("7. Tampilkan Siswa Terurut [Selection Sort]")
	fmt.Println("8. Min/Max Persentase Kehadiran")
	fmt.Println("9. Keluar")
	fmt.Println("=========================================")
	fmt.Print("Pilih menu: ")
}

func main() {
	var arrSiswa ArraySiswa
	var arrAbsensi ArrayAbsensi
	var pilihan int
	running := true

	arrSiswa.n = 0
	arrAbsensi.n = 0

	for running {
		tampilkanMenu()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahSiswaAbsensi(&arrSiswa, &arrAbsensi)
		} else if pilihan == 2 {
			ubahSiswa(&arrSiswa) // Insertion Sort + Binary Search
		} else if pilihan == 3 {
			hapusSiswa(&arrSiswa) // Insertion Sort + Binary Search
		} else if pilihan == 4 {
			cariKehadiranTanggal(arrAbsensi) // Sequential Search
		} else if pilihan == 5 {
			cariSiswaMenu(&arrSiswa) // Binary Search
		} else if pilihan == 6 {
			tampilkanStatistik(&arrSiswa)
		} else if pilihan == 7 {
			tampilkanSiswaTerurut(&arrSiswa) // Selection Sort
		} else if pilihan == 8 {
			cariMinMaxKehadiran(&arrSiswa)
		} else if pilihan == 9 {
			fmt.Println("\nTerima kasih telah menggunakan aplikasi!")
			running = false
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
