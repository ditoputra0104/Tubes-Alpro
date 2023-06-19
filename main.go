package main

import (
	"fmt"
)

const NMAX = 3

type Mobil struct {
	Nama            string
	Tahun           int
	Penjualan       int
	Pabrikan        string
	JumlahPenjualan int
}

type Pabrikan struct {
	Nama   string
	Mobil  [NMAX]Mobil
	Jumlah int
}

func tambahMobil(p *Pabrikan, mobil Mobil) bool {
	if p.Jumlah >= NMAX {
		return false
	}
	fmt.Print("Jumlah Penjualan: ")
	var jumlahPenjualan int
	fmt.Scan(&jumlahPenjualan)
	mobil.JumlahPenjualan = jumlahPenjualan

	p.Mobil[p.Jumlah] = mobil
	p.Jumlah++
	return true
}

func tambahPabrikan(pabrikans *[NMAX]Pabrikan, pabrikan Pabrikan) bool {
	for i := 0; i < NMAX; i++ {
		if pabrikans[i].Nama == "" {
			pabrikans[i] = pabrikan
			return true
		}
	}
	return false
}

func cariMobil(p *Pabrikan, namaMobil string) Mobil {
	for i := 0; i < p.Jumlah; i++ {
		if p.Mobil[i].Nama == namaMobil {
			return p.Mobil[i]
		}
	}
	return Mobil{}
}

func cariMobilPabrikan(p *[NMAX]Pabrikan, namaPabrikan string) [NMAX]Mobil {
	var mobilPabrikan [NMAX]Mobil
	var index int

	for i := 0; i < len(p); i++ {
		if p[i].Nama == namaPabrikan {
			for j := 0; j < p[i].Jumlah; j++ {
				mobilPabrikan[index] = p[i].Mobil[j]
				index++
			}
		}
	}

	return mobilPabrikan
}

func selectionSortPabrikan(p *[NMAX]Pabrikan, n int) {
	var idx_max int
	i := 0
	for i < n-1 {
		idx_max = i
		j := i + 1
		for j < n {
			if p[idx_max].Jumlah < p[j].Jumlah {
				idx_max = j
			}
			j = j + 1
		}
		if idx_max != i {
			t := p[idx_max]
			p[idx_max] = p[i]
			p[i] = t
		}
		i = i + 1
	}
}

func selectionSortMobil(P *Pabrikan, n int) {
	var idx_min int
	i := 0
	for i < n-1 {
		idx_min = i
		j := i + 1
		for j < n {
			if P.Mobil[idx_min].Tahun < P.Mobil[j].Tahun {
				idx_min = j
			}
			j = j + 1
		}
		if idx_min != i {
			t := P.Mobil[idx_min]
			P.Mobil[idx_min] = P.Mobil[i]
			P.Mobil[i] = t
		}
		i = i + 1
	}
}

func selectionSortMobilByPenjualan(P *Pabrikan, n int) {
	var idx_min int
	i := 0
	for i < n-1 {
		idx_min = i
		j := i + 1
		for j < n {
			if P.Mobil[idx_min].JumlahPenjualan < P.Mobil[j].JumlahPenjualan {
				idx_min = j
			}
			j = j + 1
		}
		if idx_min != i {
			t := P.Mobil[idx_min]
			P.Mobil[idx_min] = P.Mobil[i]
			P.Mobil[i] = t
		}
		i = i + 1
	}
}

func tampilkanPenjualanTertinggi(p *[NMAX]Pabrikan, jumlah int) {
	selectionSortPabrikan(p, jumlah)

	fmt.Println("Daftar Pabrikan dengan Jumlah Penjualan Tertinggi:")
	for i := 0; i < jumlah; i++ {
		if p[i].Nama != "" {
			fmt.Printf("Pabrikan: %s\n", p[i].Nama)
			fmt.Println("Daftar Mobil:")
			selectionSortMobilByPenjualan(&p[i], p[i].Jumlah)
			for j := 0; j < p[i].Jumlah; j++ {
				fmt.Printf("  - %s (%d) - %d penjualan\n", p[i].Mobil[j].Nama, p[i].Mobil[j].Tahun, p[i].Mobil[j].JumlahPenjualan)
			}
		}
	}
}


func menghapusDataPabrikanMobil(pabrikans *[NMAX]Pabrikan, jumlah int) {

	if jumlah < 1 || jumlah > len(pabrikans) {
		fmt.Println("Nomor pabrikan tidak valid")
		return
	}

	// Menghapus pabrikan dan mobil dari daftar
	pabrikans[jumlah-1] = Pabrikan{}

	fmt.Println("Data pabrikan dan mobil berhasil dihapus")
}

func main() {
	var pabrikans [NMAX]Pabrikan

	// Input data pabrikan dan mobil
	fmt.Println("Masukkan data pabrikan dan mobil:")

	i := 0
	shouldContinue := true
	for i < NMAX && shouldContinue {
		fmt.Printf("Pabrikan %d:\n", i+1)
		fmt.Print("Nama Pabrikan: ")
		var namaPabrikan string
		fmt.Scan(&namaPabrikan)

		pabrikan := Pabrikan{Nama: namaPabrikan}

		j := 0
		for j < NMAX {
			fmt.Printf("Mobil %d:\n", j+1)
			fmt.Print("Nama Mobil: ")
			var namaMobil string
			fmt.Scan(&namaMobil)

			fmt.Print("Tahun: ")
			var tahun int
			fmt.Scan(&tahun)

			mobil := Mobil{Nama: namaMobil, Tahun: tahun, Pabrikan: namaPabrikan}
			tambahMobil(&pabrikan, mobil)

			j++
		}

		if !tambahPabrikan(&pabrikans, pabrikan) {
			fmt.Println("Batas maksimum pabrikan telah tercapai.")
			shouldContinue = false
		} else {
			var lanjut string
			fmt.Print("Tambah pabrikan lain? (y/n): ")
			fmt.Scan(&lanjut)
			if lanjut != "y" {
				shouldContinue = false
			}
		}

		i++
	}

	// Pilihan Menu
	fmt.Println("Menu:")
	fmt.Println("1. Tambah Data Pabrikan dan Mobil")
	fmt.Println("2. Edit Data Pabrikan dan Mobil")
	fmt.Println("3. Hapus Data Pabrikan dan Mobil")
	fmt.Println("4. Cari Mobil berdasarkan Nama")
	fmt.Println("5. Cari Mobil berdasarkan Nama Pabrikan")
	fmt.Println("6. Tampilkan Daftar Penjualan Tertinggi")
	fmt.Println("7. Keluar")

	for {
		fmt.Print("Pilih Menu (1-7): ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Tambah Data Pabrikan dan Mobil:")
			fmt.Print("Nama Pabrikan: ")
			var namaPabrikan string
			fmt.Scan(&namaPabrikan)

			pabrikan := Pabrikan{Nama: namaPabrikan}

			j := 0
			for j < NMAX {
				fmt.Printf("Mobil %d:\n", j+1)
				fmt.Print("Nama Mobil: ")
				var namaMobil string
				fmt.Scan(&namaMobil)

				fmt.Print("Tahun: ")
				var tahun int
				fmt.Scan(&tahun)

				fmt.Print("Penjualan: ")
				var penjualan int
				fmt.Scan(&penjualan)

				mobil := Mobil{Nama: namaMobil, Tahun: tahun, Penjualan: penjualan, Pabrikan: namaPabrikan}
				tambahMobil(&pabrikan, mobil)

				j++
			}

			if !tambahPabrikan(&pabrikans, pabrikan) {
				fmt.Println("Batas maksimum pabrikan telah tercapai.")
			} else {
				fmt.Println("Data pabrikan dan mobil berhasil ditambahkan.")
			}

		case 2:
			fmt.Println("Edit Data Pabrikan dan Mobil:")
			fmt.Print("Pilih Nomor Pabrikan: ")
			var nomorPabrikan int
			fmt.Scan(&nomorPabrikan)

			if nomorPabrikan <= 0 || nomorPabrikan > i {
				fmt.Println("Nomor pabrikan tidak valid")
			}

			pabrikan := &pabrikans[nomorPabrikan-1]

			fmt.Print("Nama Pabrikan: ")
			var namaPabrikan string
			fmt.Scan(&namaPabrikan)

			pabrikan.Nama = namaPabrikan

			j := 0
			for j < NMAX {
				fmt.Printf("Mobil %d:\n", j+1)
				fmt.Print("Nama Mobil: ")
				var namaMobil string
				fmt.Scan(&namaMobil)

				fmt.Print("Tahun: ")
				var tahun int
				fmt.Scan(&tahun)

				mobil := Mobil{Nama: namaMobil, Tahun: tahun, Pabrikan: namaPabrikan}
				pabrikan.Mobil[j] = mobil

				j++
			}

			fmt.Println("Data pabrikan dan mobil berhasil diubah.")

		case 3:
			fmt.Println("Hapus Data Pabrikan dan Mobil:")
			fmt.Print("Pilih Nomor Pabrikan: ")
			var nomorPabrikan int
			fmt.Scan(&nomorPabrikan)

			menghapusDataPabrikanMobil(&pabrikans, nomorPabrikan)

		case 4:
			fmt.Println("Cari Mobil berdasarkan Nama:")
			fmt.Print("Masukkan nama mobil yang ingin dicari: ")
			var namaMobil string
			fmt.Scan(&namaMobil)

			mobil := cariMobil(&pabrikans[0], namaMobil)
			if mobil.Nama == "" {
				fmt.Println("Mobil tidak ditemukan.")
			} else {
				fmt.Printf("Mobil ditemukan: Nama = %s, Tahun = %d, Pabrikan = %s\n", mobil.Nama, mobil.Tahun, mobil.Pabrikan)
			}

		case 5:
			fmt.Println("Cari Mobil berdasarkan Nama Pabrikan:")
			fmt.Print("Masukkan nama pabrikan: ")
			var namaPabrikan string
			fmt.Scan(&namaPabrikan)

			mobilPabrikan := cariMobilPabrikan(&pabrikans, namaPabrikan)
			if mobilPabrikan[0].Nama == "" {
				fmt.Println("Mobil tidak ditemukan.")
			} else {
				fmt.Println("Daftar Mobil:")
				for i := 0; i < len(mobilPabrikan); i++ {
					if mobilPabrikan[i].Nama != "" {
						fmt.Printf("Mobil: %s, Tahun: %d\n", mobilPabrikan[i].Nama, mobilPabrikan[i].Tahun)
					}
				}
			}

		case 6:
			tampilkanPenjualanTertinggi(&pabrikans, i)

		case 7:
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
		}
	}
}