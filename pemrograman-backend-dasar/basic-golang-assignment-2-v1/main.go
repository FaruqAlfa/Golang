package main

import (
	"fmt"
	"strings"
	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {
	result := "Login gagal: ID or Name is undefined!"

	if id != "" && name != "" {
		found := false
		for i := 0; i < len(Students); i++ {//melakukan pengecekan dalam setiap nilai yang ada di slice student
			students := strings.Split(Students[i], "_") //memisahkan id dan name dari Students berdasarkan "_"
			if students[0] == id && students[1] == name {//membandingkan id dan name
				found = true
				break
			}
		}
		if found {
			result = "Login berhasil: " + name
		} else {
			result = "Login gagal: data mahasiswa tidak ditemukan"
		}
	} else if id == "" || name == "" {
		result = "ID or Name is undefined!"
	}
	return result // TODO: replace this
}


func Register(id string, name string, major string) string {
	var result string
	if id != "" && name != "" && major != "" {
		found := false
		for i := 0; i < len(Students); i++ {
			students := strings.Split(Students[i], "_")
			if students[0] == id {
				result = "Registrasi gagal: id sudah digunakan"
				found = true
				break
			}
		}
		if !found {
			Students = append(Students, id+"_"+name+"_"+major)
			result = "Registrasi berhasil: " + name + " (" + major + ")"
		}
	}else if id == "" || name == "" || major == "" {
		result = "ID, Name or Major is undefined!"
	}
	return result // TODO: replace this
}

func GetStudyProgram(code string) string {
	studyProgram := "Kode program studi tidak ditemukan"
	for i := 0; i < len(StudentStudyPrograms); i++ {
		studyProgram := StudentStudyPrograms[code]
		if studyProgram != "" {
			return studyProgram
		}
	}
	return studyProgram // TODO: replace this
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	var result string

	found := false
	for i, student := range Students {
		studentInfo := strings.Split(student, "_")
		if studentInfo[1] == nama {
			found = true
			fn(programStudi, &Students[i])
			result = "Program studi mahasiswa berhasil diubah."
		}
	}

	if !found {
		result = "Mahasiswa tidak ditemukan."
	}
	return result // TODO: replace this
}


func UpdateStudyProgram(programStudi string, students *string) {
	studentInfo := strings.Split(*students, "_")
	if len(studentInfo) == 3 {
		studentInfo[2] = programStudi
		*students = strings.Join(studentInfo, "_")
		fmt.Println(*students)
	}
	// TODO: answer here
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

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
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
