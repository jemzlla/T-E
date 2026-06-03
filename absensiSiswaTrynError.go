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

func insertionSortByNama(arrSiswa *ArraySiswa) {
// IS: aeeSiswa terdefinisi, berisi data siswa yang belum terurut
// FS: arrSiswa.data terurut ascending berdasarkan nama siswa (A-Z)
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

func selectionSortByPersentase(arrSiswa *ArraySiswa, ascending bool) {
// IS: arrSiswa terdefinisi, berisi data siswa dengan jumlah hadir/tidak hadir
// FS: arrSiswa.data terurut berdasarkan persentase kehadiran; ascending = true, descending = false
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

func sequentialSearchSiswa(arrSiswa ArraySiswa, nama string) int {
// Fungsi sequentialSearchSiswa mengembalikan indeks siswa jika ditemukan atau -1 jika tidak ditemukan dengan pencarian berurutan dari 0 hingga n-1
	hasil := -1
	for i := 0; i < arrSiswa.n; i++ {
		if arrSiswa.data[i].nama == nama {
			hasil = i
		}
	}
	return hasil
}

func sequentialSearchTanggal(arrAbsensi ArrayAbsensi, tanggal string) int {
// Fungsi sequentialSearchTanggal mengembalikan index absensi jika tanggal ditemukan, 
// -1 jika tidak ditemukan, pencarian dilakukan satu per satu dari index 0 hingga n-1
	hasil := -1
	for i := 0; i < arrAbsensi.n; i++ {
		if arrAbsensi.data[i].tanggal == tanggal {
			hasil = i
		}
	}
	return hasil
}

func binarySearchSiswa(arrSiswa ArraySiswa, nama string) int {
// Fungsi binarySearchSiswa mengembalikan index siswa jika ditemukan, -1 jika tidak ditemukan dengan pencarian membagi array menjadi dua bagian
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

func tambahSiswa(arrSiswa *ArraySiswa) {
// IS: arrSiswa terdefinisi, n < NMAX
// FS: arrSiswa.data bertambah satu elemen siswa baru, arrSiswa.n bertambah 1
	if arrSiswa.n >= NMAX {
		fmt.Println("Data siswa sudah penuh!")
	} else {
		var siswa Siswa
		fmt.Print("Nama siswa: ")
		fmt.Scan(&siswa.nama)
		fmt.Print("Kelas: ")
		fmt.Scan(&siswa.kelas)

		siswa.jumlahHadir = 0
		siswa.jumlahTidakHadir = 0

		arrSiswa.data[arrSiswa.n] = siswa
		arrSiswa.n++

		fmt.Println("Siswa berhasil ditambahkan!")
	}
}

func ubahSiswa(arrSiswa *ArraySiswa) {
// IS: arrSiswa terdefinisi, nama siswa yang ingin diubah terdefinisi
// FS: data siswa dengan nama yang dicari berhasil diubah (nama dan kelas), 
//     jika tidak ditemukan maka tidak ada perubahannya. Menggunakan Insertion Sort + Binary Search untuk pencarian
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

func hapusSiswa(arrSiswa *ArraySiswa) {
// IS: arrSiswa terdefinisi, nama siswa yang ingin dihapus telah terdefinisi
// FS: data siswa dengan nama yang dicari berhasil dihapus, ketika arrSiswa.n berkurang 1 maka data digeser ke kiri 
//     dan jika tidak ditemukan maka tidak ada perubahan menggunakan Insertion Sort + Binary Search untuk pencarian
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

func inputAbsensi(arrSiswa *ArraySiswa, arrAbsensi *ArrayAbsensi) {
// IS: arrSiswa terdefinisi dan tidak kosong, arrAbsensi terdefinisi
// FS: arrAbsensi.data bertambah satu data absensi harian baru, jumlahHadir/jumlahTidakHadir setiap siswa diperbarui
	var absen AbsensiHarian
	absen.n = arrSiswa.n
	var status int

	if arrSiswa.n == 0 {
		fmt.Println("Belum ada data siswa!")
	} else if arrAbsensi.n >= NMAX {
		fmt.Println("Data absen	si sudah penuh!")
	} else {
		fmt.Print("Tanggal (DD-MM-YYYY): ")
		fmt.Scan(&absen.tanggal)

		fmt.Println("\n Input kehadiran (1=Hadir, 0=Tidak Hadir):")
		for i := 0; i < arrSiswa.n; i++ {
			fmt.Printf("%s: ", arrSiswa.data[i].nama)
			fmt.Scan(&status)

			absen.siswa[i] = arrSiswa.data[i].nama
			absen.status[i] = (status == 1)

			if status == 1 {
				arrSiswa.data[i].jumlahHadir++
			} else {
				arrSiswa.data[i].jumlahTidakHadir++
			}
		}
		arrAbsensi.data[arrAbsensi.n] = absen
		arrAbsensi.n++

		fmt.Println("Absensi berhasil disimpan")
	}
}

func cariKehadiranTanggal(arrAbsensi ArrayAbsensi) {
// IS: arrAbsensi dan tanggal terdefinsisi
// FS: menampilkan daftar kehadiran siswa pada tanggal yang dicari, jika tidak tampilkan pesan error.
//     menggunakan Sequential Search untuk mencari tanggal
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

func cariSiswaMenu(arrSiswa *ArraySiswa) {
// IS: arrSiswa, nama dan metode terdefinisi
// FS: menampilkan data siswa yang ditemukan beserta statistik kehadirannya,
//     jika Sequential Search -> cari satu per satu tanpa perlu data terurut
//     jika Binary Search -> pertama sort dengan Insertion sort kemudian menggunakan binary search
	var nama string
	var cara int

	if arrSiswa.n == 0 {
		fmt.Println("Belum ada data siswa!")
	} else {
		fmt.Print("Nama siswa yang dicari:")
		fmt.Scan(&nama)
		fmt.Println("Metode penacrian")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Print("Pilih:")
		fmt.Scan(&cara)
	}

	idx := -1
	valid := true

	if cara == 1 {
		// Sequential Search
		idx = sequentialSearchSiswa(*arrSiswa, nama)
		fmt.Println("Menggunakan Sequential Search")
	} else if cara == 2 {
		// Insertion Sort terlebih dahulu kemudian Binary Search
		insertionSortByNama(arrSiswa)
		idx = binarySearchSiswa(*arrSiswa, nama)
		fmt.Println("Menggunakan Binary Search")
	} else {
		fmt.Println("Metode tidak valid!")
		valid = false
	}
	if idx == -1 {
		fmt.Println("Siswa tidak ditemukan!")
	} else {
		valid = false
	}

	if valid {
		fmt.Println("Siswa tidak ditemukan!")
	} else {
		total := arrSiswa.data[idx].jumlahHadir + arrSiswa.data[idx].jumlahTidakHadir
		persentase := 0.0
		if total > 0 {
			persentase = float64(arrSiswa.data[idx].jumlahHadir) / float64(total) * 100
		}
		fmt.Println("\n=== DATA SISWA DITEMUKAN ===")
		fmt.Printf("Nama : %s\n", arrSiswa.data[idx].nama)
		fmt.Printf("Kelas : %s\n", arrSiswa.data[idx].kelas)
		fmt.Printf("Hadir : %d kali\n", arrSiswa.data[idx].jumlahHadir)
		fmt.Printf("Tidak : %d kali\n", arrSiswa.data[idx].jumlahTidakHadir)
		fmt.Printf("Persentase kehadiran: %.2f%%\n", persentase)
	}
}

func tampilkanStatistik(arrSiswa ArraySiswa) {
// IS: arrSiswa terdefinisi dan tidak dapat kosong
// FS: menampilkan nama, kelas, jumlah hadir, tidak hadir, dan persentase kehadiran setiap siswa
	if arrSiswa.n == 0 {
		fmt.Println("Belum ada data siswa!")
	} else {
		fmt.Println("\n=== Statistik kehadiran Per Siswa ===")
		fmt.Println("====================================")

		for i := 0; i < arrSiswa.n; i++ {
			total := arrSiswa.data[i].jumlahHadir + arrSiswa.data[i].jumlahTidakHadir
			persentase := 0.0
			if total > 0 {
				persentase = float64(arrSiswa.data[i].jumlahHadir) / float64(total) * 100
			}

			fmt.Printf("\nNama : %s\n", arrSiswa.data[i].nama)
			fmt.Printf("Kelas : %s\n", arrSiswa.data[i].kelas)
			fmt.Printf("Hadir : %d kali\n", arrSiswa.data[i].jumlahHadir)
			fmt.Printf("Tidak : %d kali\n", arrSiswa.data[i].jumlahTidakHadir)
			fmt.Printf("Persentase kehadiran: %.2f%%\n", persentase)
			fmt.Println("-------------------------------------")
		}
	}
}

func tampilkanSiswaTerurut(arrSiswa *ArraySiswa) {
// IS: arrSiswa terdefinisi dan tidak dapat kosong; ascending = 1, descending =2
// FS: arrSiswa.data terurt berdasarkan persentase kehadiran, menampilkan daftar siswa beserta persentase kehadiran, 
//     menampilkan daftar siswa beserta persentase kehadiran dan menggunakan Selection Sort untuk proses pengurutannya
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

func tampilkanMenu() {
	fmt.Println("\n=========================================")
	fmt.Println("Aplikasi Absensi Siswa")
	fmt.Println("===========================================")
	fmt.Println("1. Tambah Siswa		    Note : Dont use space use _ instead")
	fmt.Println("2. Ubah Data Siswa		[Binary Search]")
	fmt.Println("3. Hapus Siswa		    [Binary Search]")
	fmt.Println("4. Input Absensi Harian")
	fmt.Println("5. Cari Kehadiran      [Sequential Search]")
	fmt.Println("6. Cari Siswa		    [Sequential/Binary Search]")
	fmt.Println("7. Statistik Kehadiran")
	fmt.Println("8. Tampilkan Siswa Terurut [Selection Sort]")
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
			tambahSiswa(&arrSiswa)
		} else if pilihan == 2 {
			ubahSiswa(&arrSiswa) // Insertion Sort + Binary Search
		} else if pilihan == 3 {
			hapusSiswa(&arrSiswa) // Insertion Sort + Binary Search
		} else if pilihan == 4 {
			inputAbsensi(&arrSiswa, &arrAbsensi)
		} else if pilihan == 5 {
			cariKehadiranTanggal(arrAbsensi) // Sequential Search
		} else if pilihan == 6 {
			cariSiswaMenu(&arrSiswa) // Sequential/Binary Search
		} else if pilihan == 7 {
			tampilkanStatistik(arrSiswa)
		} else if pilihan == 8 {
			tampilkanSiswaTerurut(&arrSiswa) // Selection Sort
		} else if pilihan == 9 {
			fmt.Println("\nTerima kasih telah menggunakan aplikasi!")
			running = false
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
