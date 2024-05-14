package main

import (
	"fmt"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {

	var message string
	if id == "" || name == "" {
		message = "ID or Name is undefined!"
	} else if len(id) != 5 {
		message = "ID must be 5 characters long!"
	} else if name == "Aditira" && id == "A1234" {
		message = "Login berhasil: Aditira (TI)"
	} else if name == "Dito" && id == "B2131" {
		message = "Login berhasil: Dito (TK)"
	} else if name == "Afis" && id == "A3455" {
		message = "Login berhasil: Afis (MI)"
	} else {
		message = "Login gagal: data mahasiswa tidak ditemukan"
	}
	
	return message // TODO: replace this
}



func Register(id string, name string, major string) string {
	var pesan string
	for i := 0; i < len(Students); i++ {
		if id == "" || name == "" || major == "" {
			pesan = "ID, Name or Major is undefined!"
		}else if len(id) != 5 {
			pesan = "ID must be 5 characters long!"
		}else if id == "A1234" || id == "B2131" || id == "A3455" {
			pesan = "Registrasi gagal: id sudah digunakan"
		}else {
			pesan = "Registrasi berhasil: " + name + " (" + major + ")"
		}
	}
	return pesan // TODO: replace this
}

func GetStudyProgram(code string) string {
	var studyProgram string

	for i := 0; i < len(StudentStudyPrograms); i++ {
		if code == "TI" {
			studyProgram = "Teknik Informatika"
		} else if code == "TK" {
			studyProgram = "Teknik Komputer"
		} else if code == "SI" {
			studyProgram = "Sistem Informasi"
		} else if code == "MI" {
			studyProgram = "Manajemen Informasi"
		}else if code == "" {
			studyProgram = "Code is undefined!"
		}
	}
	return studyProgram  // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
