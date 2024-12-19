package main

import (
	"fmt"
	"os"
	"os/exec"
)

type arrmkn [100]Menu
type arrpesan [100]Pesan

type Menu struct {
	makanan string
	harga   int
	jenis   string
	stok    int
}

type Pesan struct {
	nama          string
	menuPesan     [100]Jumlah
	status        bool
	banyakpesanan int
}

type Jumlah struct {
	namaMakanan string
	stok        int
}

func menu(data arrmkn, n int, order arrpesan, npesan int) { //tampilan utama
	var pilih int
	fmt.Println("--------------------------------------------------")
	fmt.Println("Selamat datang di Sistem Pemesanan Makanan Online!")
	fmt.Println("--------------------------------------------------")
	fmt.Println("1. Admin")
	fmt.Println("2. Customer")
	fmt.Println("3. Exit")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Pilih (1/2/3)? ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		bersihLayar()
		adm(data, n, order, npesan)
		bersihLayar()
	} else if pilih == 2 {
		bersihLayar()
		Order(data, n, order, npesan)
		bersihLayar()
	} else if pilih == 3 {
		bersihLayar()
	}
}

func Order(data arrmkn, n int, order arrpesan, npesan int) { // tampilan customer
	var pilih int
	fmt.Println("--------------------------------------------------")
	fmt.Println("Selamat datang di Sistem Pemesanan Makanan Online!")
	fmt.Println("--------------------------------------------------")
	fmt.Println("1. Pesan makanan")
	fmt.Println("2. Tampilkan histori pemesanan")
	fmt.Println("3. Tampilkan menu makanan")
	fmt.Println("4. Tampilkan makanan secara terurut berdasarkan harga yang termahal")
	fmt.Println("5. Tampilkan makanan secara terurut berdasarkan harga yang termurah")
	fmt.Println("6. Keluar")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Pilih (1/2/3/4/5/6) ?")
	fmt.Scan(&pilih)

	if pilih == 1 {
		bersihLayar()
		tambahPesan(&order, &npesan, &data, n)
		Order(data, n, order, npesan)
	} else if pilih == 2 {
		bersihLayar()
		historiPesanan(&data, &npesan, &order, n)
		Order(data, n, order, npesan)
	} else if pilih == 3 {
		bersihLayar()
		tampilMenu(data, n)
		Order(data, n, order, npesan)
	} else if pilih == 4 {
		urutMahal(&data, n)
		fmt.Println("Daftar menu setelah diurutkan berdasarkan harga termahal:")
		tampilMenu(data, n)
		Order(data, n, order, npesan)
	} else if pilih == 5 {
		urutMurah(&data, n)
		fmt.Println("Daftar menu setelah diurutkan berdasarkan harga termurah:")
		tampilMenu(data, n)
		Order(data, n, order, npesan)
	} else if pilih == 6 {
		menu(data, n, order, npesan)
	}
}

func historiPesanan(A *arrmkn, np *int, data *arrpesan, n int) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("                 Histori Pemesanan                ")
	fmt.Println("--------------------------------------------------")
	if *np == 0 {
		fmt.Println("Tidak ada pesanan yang ditemukan.")
		return
	}
	for i := 0; i < *np; i++ {
		var totalHarga int
		fmt.Printf("Pesanan ke-%d\n", i+1)
		fmt.Printf("Nama Pemesan : %s\n", data[i].nama)
		fmt.Println("Menu yang dipesan:")
		for j := 0; j < data[i].banyakpesanan; j++ {
			fmt.Printf("  Nama Makanan: %s, Jumlah: %d\n", data[i].menuPesan[j].namaMakanan, data[i].menuPesan[j].stok)
		}
		for i := 0; i < *&data[*np-1].banyakpesanan; i++ {
			for j := 0; j < n; j++ {
				if A[j].makanan == data[*np-1].menuPesan[i].namaMakanan {
					totalHarga += A[j].harga * data[*np-1].menuPesan[i].stok
				}
			}
		}
		fmt.Println("Total Harga : ", totalHarga)
	}
}

func adm(data arrmkn, n int, order arrpesan, npesan int) { // tampilan admin
	var pilih int
	fmt.Println("--------------------------------------------------")
	fmt.Println("Selamat datang di Sistem Pemesanan Makanan Online!")
	fmt.Println("--------------------------------------------------")
	fmt.Println("1. Penambahan Menu")
	fmt.Println("2. Penghapusan Menu")
	fmt.Println("3. Daftar Menu")
	fmt.Println("4. Edit Menu")
	fmt.Println("5. Exit")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Pilih (1/2/3/4/5) ?")
	fmt.Scan(&pilih)

	if pilih == 1 {
		bersihLayar()
		tambahMenu(&data, &n)
		adm(data, n, order, npesan)
	} else if pilih == 2 {
		bersihLayar()
		fmt.Print("Tulis Nama Menu: ")
		hapusMenu(&data, &n)
		adm(data, n, order, npesan)
	} else if pilih == 3 {
		bersihLayar()
		tampilMenu(data, n)
		adm(data, n, order, npesan)
	} else if pilih == 4 {
		bersihLayar()
		editMenu(&data, &n)
		adm(data, n, order, npesan)
	} else if pilih == 5 {
		bersihLayar()
		menu(data, n, order, npesan)
	}
}

func tambahMenu(data *arrmkn, n *int) { 
	fmt.Print("Nama Makanan : ")
	fmt.Scan(&data[*n].makanan)
	fmt.Print("Harga : ")
	fmt.Scan(&data[*n].harga)
	fmt.Print("Jenis : ")
	fmt.Scan(&data[*n].jenis)
	fmt.Print("Stok : ")
	fmt.Scan(&data[*n].stok)
	*n++
}

func tampilMenu(data arrmkn, n int) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("                 Menu Makanan                     ")
	fmt.Println("--------------------------------------------------")
	for i := 0; i < n; i++ {
		if data[i].stok > 0 {
			fmt.Print("Nama Makanan : ", data[i].makanan, " Harga : ", data[i].harga, " Jenis : ", data[i].jenis, " Stok : ", data[i].stok, "\n")
		}
	}
}

func hapusMenu(data *arrmkn, n *int) { 
	var nama string
	fmt.Scan(&nama)
	index := sequentialSearch(*data, *n, nama)
	if index != -1 {
		for i := index; i < *n-1; i++ {
			(*data)[i] = (*data)[i+1]
		}
		*n--
	} else {
		fmt.Println("Menu Tidak Ditemukan")
	}
}

func sequentialSearch(data arrmkn, n int, makanan string) int {
	for i := 0; i < n; i++ {
		if data[i].makanan == makanan {
			return i
		}
	}
	return -1
}

func editMenu(data *arrmkn, n *int) { 
	var makanan string
	fmt.Print("Pilih Nama Makanan: ")
	fmt.Scan(&makanan)
	index := sequentialSearch(*data, *n, makanan)
	if index != -1 {
		fmt.Print("Edit Nama Makanan: ")
		fmt.Scan(&(*data)[index].makanan)
		fmt.Print("Edit Harga: ")
		fmt.Scan(&(*data)[index].harga)
		fmt.Print("Edit Jenis: ")
		fmt.Scan(&(*data)[index].jenis)
		fmt.Print("Edit Stok: ")
		fmt.Scan(&(*data)[index].stok)
		tampilMenu(*data, *n)
	} else {
		fmt.Println("Menu Tidak Ditemukan")
	}
}

func urutMahal(A *arrmkn, n int) {
	for pass := 0; pass < n-1; pass++ {
		idx := pass
		for i := pass + 1; i < n; i++ {
			if A[i].harga > A[idx].harga {
				idx = i
			}
		}
		temp := A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}
}

func urutMurah(A *arrmkn, n int) {
	for pass := 0; pass < n-1; pass++ {
		idx := pass
		for i := pass + 1; i < n; i++ {
			if A[i].harga < A[idx].harga {
				idx = i
			}
		}
		temp := A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}
}

func bersihLayar() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func tambahPesan(A *arrpesan, np *int, data *arrmkn, n int) {
	var pesanan string
	var stok int
	tampilMenu(*data, n)
	fmt.Print("Nama Pemesan : ")
	fmt.Scan(&A[*np].nama)
	fmt.Print("Pesanan (Nama Makanan dan Jumlah Stok, ketik 'cukup' untuk berhenti): ")
	fmt.Scan(&pesanan, &stok)
	for i := 0; i < 100 && pesanan != "cukup"; i++ {
		A[*np].menuPesan[i].namaMakanan = pesanan
		A[*np].menuPesan[i].stok = stok
		A[*np].banyakpesanan++
		fmt.Print("Pesanan (Nama Makanan dan Jumlah Stok, ketik 'cukup' untuk berhenti): ")
		fmt.Scan(&pesanan, &stok)
	}
	*np++
	confrim(&*A, &*np, data, n)
}

func confrim(A *arrpesan, np *int, data *arrmkn, n int) {
	var pilih int
	var conf string
	var totalHarga int
	fmt.Println("----------------------------")
	fmt.Println("List Menu Yang Dipesan")
	fmt.Println("Nama Pemesan : ", A[*np-1].nama)
	fmt.Println("----------------------------")
	for i := 0; i < *&A[*np-1].banyakpesanan; i++ {
		for j := 0; j < n; j++ {
			if data[j].makanan == A[*np-1].menuPesan[i].namaMakanan {
				fmt.Printf("Nama Makanan : %s, Harga : Rp. %d, Stok: %d => Total: Rp. %d\n", A[*np-1].menuPesan[i].namaMakanan, data[j].harga, A[*np-1].menuPesan[i].stok, (data[j].harga * A[*np-1].menuPesan[i].stok))
				totalHarga += data[j].harga * A[*np-1].menuPesan[i].stok
			}
		}
	}
	fmt.Println("Total Harga : ", totalHarga)
	fmt.Println("Apakah menu sudah benar? Y/N")
	fmt.Scan(&conf)
	if conf == "N" {
		fmt.Println("1. Tambah menu Pesanan")
		fmt.Println("2. Kurangi stok Pesanan")
		fmt.Println("3. hapus menu Pesanan")
		fmt.Print("Pilih Nomor :")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahMenuPesanan(&*A, &*np, *data, n)
			confrim(&*A, &*np, data, n)
		} else if pilih == 2 {
			editPesanan(&*A, &*np, *data, n)
			confrim(&*A, &*np, data, n)
		} else if pilih == 3 {
			deleteMenuPesanan(&*A, &*np, *data, n)
			confrim(&*A, &*np, data, n)

		}
	} else {
		for i := 0; i < A[*np-1].banyakpesanan; i++ {
			for j := 0; j < n; j++ {
				if A[*np-1].menuPesan[i].namaMakanan == data[j].makanan {
					data[j].stok -= A[*np-1].menuPesan[i].stok
				}
			}
		}
	}
}

func tambahMenuPesanan(A *arrpesan, np *int, data arrmkn, n int) {
	var pesanan string
	var stok int
	fmt.Print("Pesanan (Nama Makanan dan Jumlah Stok, ketik 'cukup' untuk berhenti): ")
	fmt.Scan(&pesanan, &stok)
	for i := A[*np-1].banyakpesanan; i < 100 && pesanan != "cukup"; i++ {
		A[*np-1].menuPesan[i].namaMakanan = pesanan
		A[*np-1].menuPesan[i].stok = stok
		A[*np-1].banyakpesanan++
		fmt.Print("Pesanan (Nama Makanan dan Jumlah Stok, ketik 'cukup' untuk berhenti): ")
		fmt.Scan(&pesanan, &stok)
	}
}

func editPesanan(A *arrpesan, np *int, data arrmkn, n int) {
	var pesanan string
	var stok int
	fmt.Print("Nama menu yang ingin diedit stok: ")
	fmt.Scan(&pesanan)
	fmt.Print("Masukan jumlah stok: ")
	fmt.Scan(&stok)
	for i := 0; i < A[*np-1].banyakpesanan; i++ {
		if A[*np-1].menuPesan[i].namaMakanan == pesanan {
			A[*np-1].menuPesan[i].stok = stok
		}
	}
}

func deleteMenuPesanan(A *arrpesan, np *int, data arrmkn, n int) {
	var hapusPesanan string
	fmt.Print("Nama menu yang ingin dihapus: ")
	fmt.Scan(&hapusPesanan)
	found := false
	for i := 0; i < (*A)[*np-1].banyakpesanan; i++ {
		if (*A)[*np-1].menuPesan[i].namaMakanan == hapusPesanan {
			found = true
		}
		if found && i < (*A)[*np-1].banyakpesanan-1 {
			(*A)[*np-1].menuPesan[i] = (*A)[*np-1].menuPesan[i+1]
		}
	}
	if found {
		(*A)[*np-1].banyakpesanan--
		(*A)[*np-1].menuPesan[(*A)[*np-1].banyakpesanan] = Jumlah{}
	}
}

func main() {
	var n int
	var data arrmkn
	var nPesanan int
	var order arrpesan
	menu(data, n, order, nPesanan)

}
