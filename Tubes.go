package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxAdmin int = 1
const maxSoal int = 10
const maxPeserta int = 10
const maxPilihan int = 4

type soal struct {
	pertanyaan string
	pilihan    [maxPilihan]string
	jawaban    string
	benar      int
	salah      int
}

type peserta struct {
	nama     string
	password string
	skor     int
}

type admin struct {
	username string
	password string
}

type arrAdmin [maxAdmin]admin
type arrSoal [maxSoal]soal
type arrPeserta [maxPeserta]peserta

var jumSoal, jumPeserta, jumAdmin, iPeserta int

func header() {
	fmt.Println("=============================================================")
	fmt.Println("==========|     Who One to Be a Millionaire      | ==========")
	fmt.Println("==========|           Created By :               | ==========")
	fmt.Println("==========|   Ragadhitya Janatantra Koeshutama   | ==========")
	fmt.Println("==========|       Muhammad Yudha Pratama         | ==========")
	fmt.Println("==========|     Tubes Algoritma Pemograman       | ==========")
	fmt.Println("=============================================================")
}

func menu(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var pilih string
	fmt.Println("=== Menu ===")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)
	for pilih != "1" && pilih != "2" && pilih != "3" {
		fmt.Println("Input Tidak Valid")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
	}
	if pilih == "1" {
		menu1(Q, A, P)
	} else if pilih == "2" {
		menu2(Q, A, P)
	}
}

func menu1(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var pilih string
	fmt.Println("=== Silahkan Pilih ===")
	fmt.Println("1. Admin")
	fmt.Println("2. Peserta")
	fmt.Println("3. Kembali")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)
	for pilih != "1" && pilih != "2" && pilih != "3" {
		fmt.Println("Input Tidak Valid")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
	}
	if pilih == "1" {
		regisAdmin(Q, A, P)
	} else if pilih == "2" {
		regisPeserta(Q, A, P)
	} else {
		menu(Q, A, P)
	}
}

func menu2(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var pilih string
	fmt.Println("=== Silahkan Pilih ===")
	fmt.Println("1. Admin")
	fmt.Println("2. Peserta")
	fmt.Println("3. Kembali")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)
	for pilih != "1" && pilih != "2" && pilih != "3" {
		fmt.Println("Input Tidak Valid")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
	}
	if pilih == "1" {
		loginAdmin(Q, A, P)
	} else if pilih == "2" {
		loginPeserta(Q, A, P)
	} else {
		menu(Q, A, P)
	}
}

func regisAdmin(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var username, pass string
	var simpan int
	if jumAdmin < maxAdmin {
		fmt.Println("=== Silahkan Registrasi ===")
		fmt.Print("Username :")
		fmt.Scan(&username)
		fmt.Print("Password :")
		fmt.Scan(&pass)
		fmt.Println("========================================")
		fmt.Println("Pilihlah")
		fmt.Println("1. Simpan 2. Batal")
		fmt.Scan(&simpan)
		if simpan == 1 {
			A[jumAdmin].username = username
			A[jumAdmin].password = pass
			fmt.Println("Registrasi akun telah berhasil")
			jumAdmin++
		} else {
			fmt.Println("Registrasi di batalkan")
		}
	} else {
		fmt.Println("Jumlah admin sudah penuh")
	}
	menu(Q, A, P)
}

func regisPeserta(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var username, pass string
	var simpan, iNama int
	if jumPeserta < maxPeserta {
		fmt.Println("=== Silahkan Registrasi ===")
		fmt.Print("Username :")
		fmt.Scan(&username)
		fmt.Print("Password :")
		fmt.Scan(&pass)
		iNama = cariPesertaQuiz(P, username)
		for iNama != -1 {
			fmt.Println("Username sudah digunakan, ganti username lain")
			fmt.Print("Username :")
			fmt.Scan(&username)
			fmt.Print("Password :")
			fmt.Scan(&pass)
			iNama = cariPesertaQuiz(P, username)
		}
		if iNama == -1 {
			fmt.Println("========================================")
			fmt.Println("Pilihlah")
			fmt.Println("1. Simpan 2. Batal")
			fmt.Scan(&simpan)
			if simpan == 1 {
				P[jumPeserta].nama = username
				P[jumPeserta].password = pass
				fmt.Println("Registrasi akun telah berhasil")
				jumPeserta++
			} else {
				fmt.Println("Registrasi di batalkan")
			}
		}

	} else {
		fmt.Println("Jumlah Peserta telah penuh")
	}

	menu(Q, A, P)
}

func cariAdmin(A *arrAdmin, username, pass string) bool {
	var found bool = false
	var i int = 0
	for i < jumAdmin && !found {
		found = A[i].username == username && A[i].password == pass
		i++
	}
	return found
}

func urutPeserta(A *arrAdmin, P *arrPeserta) {
	var i, j, idx int
	var t peserta
	i = 1
	for i <= jumPeserta-1 {
		idx = i - 1
		j = i
		for j < jumPeserta {
			if P[idx].nama > P[j].nama {
				idx = j
			}
			j = j + 1
		}
		t = P[idx]
		P[idx] = P[i-1]
		P[i-1] = t
		i = i + 1
	}
}

func cariPeserta(P *arrPeserta, username, pass string) bool {
	var kr int = 0
	var kn int = jumPeserta - 1
	var med int

	for kr <= kn {
		med = (kr + kn) / 2
		if P[med].nama == username && P[med].password == pass {
			return true
		} else if P[med].nama < username {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	return false
}

func cariPesertaQuiz(P *arrPeserta, nama string) int {
	var found int = -1
	var i int = 0
	for i < jumPeserta && found == -1 {
		if P[i].nama == nama {
			found = i
		}
		i++
	}
	return found
}

func loginAdmin(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var username, pass string
	fmt.Println("=== Silahkan Login ===")
	fmt.Print("Username :")
	fmt.Scan(&username)
	fmt.Print("Password :")
	fmt.Scan(&pass)
	if cariAdmin(A, username, pass) {
		menuAdmin(Q, A, P)
	} else {
		fmt.Println("Username atau Password SALAH")
		menu(Q, A, P)
	}
}

func loginPeserta(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var username, pass string
	fmt.Println("=== Silahkan Login ===")
	fmt.Print("Username :")
	fmt.Scan(&username)
	fmt.Print("Password :")
	fmt.Scan(&pass)
	urutPeserta(A, P)
	if cariPeserta(P, username, pass) {
		iPeserta = cariPesertaQuiz(P, username)
		menuPeserta(Q, A, P)
	} else {
		fmt.Println("Username atau Password SALAH")
		menu(Q, A, P)
	}
}

func menuAdmin(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var pilih string
	fmt.Println("=== Menu Admin ===")
	fmt.Println("1. Tambah Soal")
	fmt.Println("2. Edit Soal")
	fmt.Println("3. Hapus Soal")
	fmt.Println("4. Lihat Soal")
	fmt.Println("5. Leaderboard")
	fmt.Println("6. Top 5 Soal")
	fmt.Println("7. Kembali")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)
	for pilih != "1" && pilih != "2" && pilih != "3" && pilih != "4" && pilih != "5" && pilih != "6" && pilih != "7" {
		fmt.Println("Input Tidak Valid")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
	}
	if pilih == "1" {
		addSoal(Q, A, P)
	} else if pilih == "2" {
		editSoal(Q, A, P)
	} else if pilih == "3" {
		DeleteSoal(Q, A, P)
	} else if pilih == "4" {
		showSoal(Q, A, P)
	} else if pilih == "5" {
		leaderboardAdmin(Q, A, P)
	} else if pilih == "6" {
		menuTop5(Q, A, P)
	} else if pilih == "7" {
		menu(Q, A, P)
	}
}

func menuTop5(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var pilih, back string
	fmt.Println("=== Silahkan Pilih ===")
	fmt.Println("1. Top 5 Soal paling banyak BENAR")
	fmt.Println("2. Top 5 Soal paling banyak SALAH")
	fmt.Println("3. Kembali")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)
	for pilih != "1" && pilih != "2" && pilih != "3" {
		fmt.Println("Input Tidak Valid")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
	}
	if pilih == "1" {
		benarDescen(Q)
		for i := 0; i < 5; i++ {
			fmt.Println(i+1, Q[i].pertanyaan, "\nTotal Benar : ", Q[i].benar)
		}
		fmt.Println("1. Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&back)
		for back != "1" {
			fmt.Println("Input Tidak Valid")
			fmt.Print("Pilihan : ")
			fmt.Scan(&back)
		}
	} else if pilih == "2" {
		salahDescen(Q)
		for i := 0; i < 5; i++ {
			fmt.Println(i+1, Q[i].pertanyaan, "\nTotal Salah : ", Q[i].salah)
		}
		fmt.Println("1. Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&back)
		for back != "1" {
			fmt.Println("Input Tidak Valid")
			fmt.Print("Pilihan : ")
			fmt.Scan(&back)
		}
	}
	menuAdmin(Q, A, P)
}

func benarDescen(Q *arrSoal) {
	var i, j, idx int
	var temp soal
	i = 1
	for i <= jumSoal-1 {
		idx = i - 1
		j = i
		for j < jumSoal {
			if Q[idx].benar < Q[j].benar {
				idx = j
			}
			j++
		}
		temp = Q[idx]
		Q[idx] = Q[i-1]
		Q[i-1] = temp
		i++
	}
}

func salahDescen(Q *arrSoal) {
	var i, j, idx int
	var temp soal
	i = 1
	for i <= jumSoal-1 {
		idx = i - 1
		j = i
		for j < jumSoal {
			if Q[idx].salah < Q[j].salah {
				idx = j
			}
			j++
		}
		temp = Q[idx]
		Q[idx] = Q[i-1]
		Q[i-1] = temp
		i++
	}
}

func menuPeserta(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var pilih string
	fmt.Println("=== Menu Peserta===")
	fmt.Println("1. Mainkan Quiz")
	fmt.Println("2. Leaderboard")
	fmt.Println("3. Kembali")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)
	for pilih != "1" && pilih != "2" && pilih != "3" {
		fmt.Println("Input Tidak Valid")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
	}
	if pilih == "1" {
		acakSoal(Q)
		mainQuiz(Q, A, P)
	} else if pilih == "2" {
		leaderboardPeserta(Q, A, P)
	} else if pilih == "3" {
		menu(Q, A, P)
	}
}

func leaderboardPeserta(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var back string
	SkorDescen(P)
	fmt.Println("===== Leaderboard =====")
	for i := 0; i < jumPeserta; i++ {
		fmt.Println(P[i].nama, "Skor : ", P[i].skor)
	}
	fmt.Println("1. Kembali")
	fmt.Print("Pilihan : ")
	fmt.Scan(&back)
	for back != "1" {
		fmt.Println("Input Tidak Valid")
		fmt.Scan(&back)
	}
	menuPeserta(Q, A, P)
}

func leaderboardAdmin(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var back string
	SkorDescen(P)
	if jumPeserta != 0 {
		fmt.Println("===== Leaderboard =====")
		for i := 0; i < jumPeserta; i++ {
			fmt.Println(P[i].nama, "Skor : ", P[i].skor)
		}
	} else {
		fmt.Println("Belum ada peserta yang melakukan registrasi")
	}
	fmt.Println("1. Kembali")
	fmt.Print("Pilihan : ")
	fmt.Scan(&back)
	for back != "1" {
		fmt.Println("Input Tidak Valid")
		fmt.Scan(&back)
	}
	menuAdmin(Q, A, P)
}

func cekSoal(Q *arrSoal, soal string) bool {
	var found bool = false
	var i int = 0
	for i < jumSoal && !found {
		found = Q[i].pertanyaan == soal
		i++
	}
	return found
}

func addSoal(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var soal string
	var simpan int
	if jumSoal < maxSoal {
		fmt.Println("Ketikan Soal (Gunakan tanda underscore sebagai pengganti spasi)")
		fmt.Scan(&soal)
		if !cekSoal(Q, soal) {
			for i := 0; i < maxPilihan; i++ {
				fmt.Printf("Opsi %c: ", 'A'+i)
				fmt.Scan(&Q[jumSoal].pilihan[i])
			}
			fmt.Print("Kunci Jawaban (A/B/C/D) : ")
			fmt.Scan(&Q[jumSoal].jawaban)
			for Q[jumSoal].jawaban != "A" && Q[jumSoal].jawaban != "B" && Q[jumSoal].jawaban != "C" && Q[jumSoal].jawaban != "D" {
				fmt.Print("Kunci Jawaban (A/B/C/D) : ")
				fmt.Scan(&Q[jumSoal].jawaban)
			}
			fmt.Println("========================================")
			fmt.Println("Pilihlah")
			fmt.Println("1. Simpan 2. Batal")
			fmt.Scan(&simpan)
			if simpan == 1 {
				Q[jumSoal].pertanyaan = soal
				fmt.Println("Soal berhasil ditambahkan")
				jumSoal++
			} else {
				fmt.Println("Penambahan soal dibatalkan")
			}
		} else {
			fmt.Println("Soal sudah tersedia di Bank Soal")
		}
	} else {
		fmt.Println("Bank Soal Sudah Penuh")
	}
	menuAdmin(Q, A, P)
}

func editSoal(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var edit string
	var idx, simpan int
	displaySoal(Q, A, P)
	fmt.Println("Masukan nomor soal yang akan di edit")
	fmt.Scan(&idx)
	fmt.Println("Masukan Soal yang baru")
	fmt.Scan(&edit)
	for i := 0; i < maxPilihan; i++ {
		fmt.Printf("Opsi %c: ", 'A'+i)
		fmt.Scan(&Q[idx-1].pilihan[i])
	}
	fmt.Println("Kunci Jawaban (A/B/C/D)")
	fmt.Scan(&Q[idx-1].jawaban)
	fmt.Println("========================================")
	fmt.Println("Pilihlah")
	fmt.Println("1. Simpan 2. Batal")
	fmt.Scan(&simpan)
	if simpan == 1 {
		Q[idx-1].pertanyaan = edit
		fmt.Println("Soal berhasil diedit")
	} else {
		fmt.Println("Perubahan soal dibatalkan")
	}

	menuAdmin(Q, A, P)
}

func DeleteSoal(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var idx int
	displaySoal(Q, A, P)
	fmt.Println("Masukan nomor soal yang akan di hapus")
	fmt.Scan(&idx)
	if idx != -1 {
		for i := idx - 1; i < jumSoal-1; i++ {
			Q[i] = Q[i+1]
		}
		jumSoal--
		fmt.Println("Soal berhasil dihapus")
	} else {
		fmt.Println("Soal tidak ditemukan")
	}
	menuAdmin(Q, A, P)
}

func displaySoal(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	for i := 0; i < jumSoal; i++ {
		fmt.Println("Soal", i+1, "\n", Q[i].pertanyaan)
		for j := 0; j < maxPilihan; j++ {
			fmt.Printf("Opsi %c: %s\n", 'A'+j, Q[i].pilihan[j])
		}
		fmt.Println("Kunci Jawaban : ", Q[i].jawaban)
	}
}

func showSoal(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	for i := 0; i < jumSoal; i++ {
		fmt.Println("Soal", i+1, "\n", Q[i].pertanyaan)
		for j := 0; j < maxPilihan; j++ {
			fmt.Printf("Opsi %c: %s\n", 'A'+j, Q[i].pilihan[j])
		}
		fmt.Println("Kunci Jawaban : ", Q[i].jawaban)
	}
	menuAdmin(Q, A, P)
}

func mainQuiz(Q *arrSoal, A *arrAdmin, P *arrPeserta) {
	var skor int
	var jawaban string
	skor = 0
	if jumSoal != 0 {
		if iPeserta != -1 {
			for i := 0; i < jumSoal; i++ {
				fmt.Println("Soal", i+1, Q[i].pertanyaan)
				for j := 0; j < maxPilihan; j++ {
					fmt.Printf("Opsi %c: %s\n", 'A'+j, Q[i].pilihan[j])
				}
				fmt.Println("Masukan Jawaban anda (A/B/C/D)")
				fmt.Scan(&jawaban)
				if jawaban == Q[i].jawaban {
					skor += 10
					Q[i].benar++
				} else {
					Q[i].salah++
				}
				fmt.Println("Jawaban yang benar adalah : ", Q[i].jawaban)
			}
			P[iPeserta].skor = skor
			fmt.Println("Skor anda : ", P[iPeserta].skor)
		}
	} else {
		fmt.Println("Soal belum tersedia. Silahkan tunggu admin membuat soal")
	}
	menuPeserta(Q, A, P)
}

func SkorDescen(P *arrPeserta) {
	var i, j int
	var temp peserta
	i = 1
	for i <= jumPeserta-1 {
		j = i
		temp = P[j]
		for j > 0 && temp.skor > P[j-1].skor {
			P[j] = P[j-1]
			j--
		}
		P[j] = temp
		i++
	}
}

func acakSoal(Q *arrSoal) {
	rand.Seed(time.Now().UnixNano())
	for i := jumSoal - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		Q[i], Q[j] = Q[j], Q[i]
	}
}

func main() {
	var Q arrSoal
	var P arrPeserta
	var A arrAdmin
	var x string
	fmt.Print("Press anything to Continue...")
	fmt.Scan(&x)
	header()
	menu(&Q, &A, &P)

}
