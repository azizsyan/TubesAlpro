package main

import "fmt"

type Menu struct {
	Nama      string
	Harga     float64
	Deskripsi string
	Jenis     string
}

type Order struct {
	MenuItems []Menu
	Total     float64
}

func main() {
	menu := []Menu{
		{
			Nama: "Paket Ayam Geprek",
			Harga: 16000,
			Deskripsi: "Nasi ayam geprek, tahu, tempe, sambal, es teh manis",
			Jenis: "Makanan + Minuman",
		},
		{
			Nama: "Mie Geprek Telur Ceplok",
			Harga: 14000,
			Deskripsi: "Mie dengan bumbu geprek ditambah telor ceplok dan sayuran",
			Jenis: "Makanan",
		},
		{
			Nama: "Es Teh Manis",
			Harga: 4000,
			Deskripsi: "Minuman segar penambah semangat",
			Jenis: "Minuman",
		},
	}

	order := Order{
		MenuItems: []Menu{},
		Total:     0,
	}

	fmt.Println("Selamat datang di Sistem Pemesanan Makanan Online!")

		fmt.Println("\n1. Pesan makanan")
		fmt.Println("2. Batalkan pesanan")
		fmt.Println("3. Hitung total harga")
		fmt.Println("4. Tampilkan histori pemesanan")
		fmt.Println("5. Tampilkan daftar toko makanan")
		fmt.Println("6. Tampilkan jenis makanan")
		fmt.Println("7. Tampilkan makanan yang ditawarkan secara terurut")
		fmt.Println("8. Keluar")
		fmt.Println("Pilih opsi (masukkan nomor opsi):")

		var input int
		fmt.Scanln(&input)

		
		 if input == 3 {
			// Hitung total harga
			fmt.Printf("Total harga pesanan Anda: Rp%.2f\n", order.Total)
		} 
		
	}
