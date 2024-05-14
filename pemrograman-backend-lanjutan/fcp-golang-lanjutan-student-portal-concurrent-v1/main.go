package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	sync.Mutex
	students             []model.Student
	studentStudyPrograms map[string]string
	failedLoginAttempts  map[string]int
	//add map for tracking login attempts here
	// TODO: answer here
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
		//inisialisasi failedLoginAttempts di sini:
		failedLoginAttempts: map[string]int{},
		// TODO: answer here
	}
}

func ReadStudentsFromCSV(filename string) ([]model.Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // ID, Name and StudyProgram

	var students []model.Student
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		student := model.Student{
			ID:           record[0],
			Name:         record[1],
			StudyProgram: record[2],
		}
		students = append(students, student)
	}
	return students, nil
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	return sm.students // TODO: replace this
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	if id == "" || name == "" {//Pengecekan jika id dan name kosong
        return "", errors.New("ID or Name is undefined!")
    }
	for _, student := range sm.students {//Pengecekan jika id dan name ada
		if student.ID == id && student.Name == name {
			delete(sm.failedLoginAttempts, id)
			return fmt.Sprintf("Login berhasil: Selamat datang %s! Kamu terdaftar di program studi: %s", name, sm.studentStudyPrograms[student.StudyProgram]), nil
		}
	}

	sm.failedLoginAttempts[id]++ //penambahan percobaan login
	if sm.failedLoginAttempts[id] > 3 {//pengecekan jika percobaan login > 3
		delete(sm.failedLoginAttempts, id)//maka menghapus id/percobaan dari failedLoginAttempts
		return "", errors.New("Login gagal: Batas maksimum login terlampaui")
	}
	return "", errors.New("Login gagal: data mahasiswa tidak ditemukan") // TODO: replace this
}

func (sm *InMemoryStudentManager) RegisterLongProcess() {
	// 30ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	// 30ms delay to simulate slow processing. DO NOT REMOVE THIS LINE
	sm.RegisterLongProcess()

	// Below lock is needed to prevent data race error. DO NOT REMOVE BELOW 2 LINES
	sm.Lock()
	defer sm.Unlock()

	if id == "" || name == "" || studyProgram == "" {//Pengecekan jika id, name dan studyProgram kosong
		return "", errors.New("ID, Name or StudyProgram is undefined!")

	}
	_, fouund := sm.studentStudyPrograms[studyProgram]//Pengecekan jika studyProgram tidak ada
	if !fouund {
		return "", errors.New(fmt.Sprintf("Study program %s is not found", studyProgram))

	}
	for _, student := range sm.students {//Pengecekan jika id sudah digunakan
		if student.ID == id {
			return "", errors.New("Registrasi gagal: id sudah digunakan")
		}
	}

	sm.students = append(sm.students, model.Student{//Menambahkan data ke dalam slice students
		ID:           id,
		Name:         name,
		StudyProgram: studyProgram,
	})
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil // TODO: replace this
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if code == "" {//Pengecekan jika code/program studi kosong
		return "", errors.New("Code is undefined!")
	}

	if program, exists := sm.studentStudyPrograms[code]; exists {//Pengecekan jika code/program studi ada
		return program, nil
	}
	return "", errors.New("Kode program studi tidak ditemukan") // TODO: replace this
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	for i, student := range sm.students {//pengecekan dalam slice students
		if student.Name == name {//Pengecekan jika name sama dengan name yang diberikan
			err := fn(&sm.students[i])
			if err != nil {
				return "", errors.New("Mahasiswa tidak ditemukan.")
			}
			return "Program studi mahasiswa berhasil diubah.", nil
		}
	}
	return "", errors.New("Mahasiswa tidak ditemukan.") // TODO: replace this
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		newProgram, found := sm.studentStudyPrograms[programStudi]//Pengecekan jika programStudi ada
		if !found {//jika program studi tidak ada
			return errors.New("Kode program studi tidak ditemukan")
		}
		s.StudyProgram = newProgram
		return nil// TODO: replace this
	}
}


func (sm *InMemoryStudentManager) ImportStudents(filenames []string) error {

	errCh := make(chan error)
	var finalErr error// menyimpan error terakhir

	var totalStudents int //menyimpan total mahasiswa yang diimport

	var wg sync.WaitGroup
	wg.Add(len(filenames)) //menambahkan jumlah goroutine yang akan dijalankan
	for _, filename := range filenames {//mengulang sebanyak len(filenames) 
		//menjalankan goroutine untuk setiap file dalam slice filenames
		go func(filename string) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			students, err := ReadStudentsFromCSV(filename)
			if err != nil {
				fmt.Println(err.Error())
				errCh <- err
			}
			//menambahkan data ke dalam slice students
			sm.Lock()
			sm.students = append(sm.students, students...)
			sm.Unlock()

			totalStudents += len(students)
		}(filename)
	}

	wg.Wait() //menunggu semua goroutine selesai
	select {
		case err := <-errCh:
		return err
		default:
			finalErr = nil
	}

	if totalStudents < 3000 {//Pengecekan jika total mahasiswa kurang dari 3000
		finalErr = errors.New("Jumlah mahasiswa kurang dari 3000")
	}
	return finalErr // TODO: replace this
}

func (sm *InMemoryStudentManager) SubmitAssignmentLongProcess() {
	// 3000ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) SubmitAssignments(numAssignments int) {
	start := time.Now()
	c := make(chan int)
	var wg sync.WaitGroup
	//mengirimkan data ke dalam channel c sebanyak numAssignments
	go func() {
		for i := 1; i <= numAssignments; i++ {
			c <- i
		}
		close(c)//menutup channel
	}()

	// menjalankan goroutine untuk setiap tugas yang diberikan, kemudian menunggu menerima tugas dari channel
	for i := 1; i <= 3; i++{
		wg.Add(1)
		go func() {
			for taskNumber := range c {
				fmt.Printf("worker %d started task %d\n", i, taskNumber)
				sm.SubmitAssignmentLongProcess()
				fmt.Printf("worker %d completed task %d\n", i, taskNumber)				
			}
			defer wg.Done()
		}()
	}

	// TODO: answer here
	wg.Wait()
	elapsed := time.Since(start) //menghitung waktu yang diperlukan
	fmt.Printf("Submitting %d assignments took %s\n", numAssignments, elapsed)
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Bulk Import Student")
		fmt.Println("6. Submit assignment")
		fmt.Println("7. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "5":
			helper.ClearScreen()
			fmt.Println("=== Bulk Import Student ===")

			// Define the list of CSV file names
			csvFiles := []string{"students1.csv", "students2.csv", "students3.csv"}

			err := manager.ImportStudents(csvFiles)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Import successful!")
			}

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')

		case "6":
			helper.ClearScreen()
			fmt.Println("=== Submit Assignment ===")

			// Enter how many assignments you want to submit
			fmt.Print("Enter the number of assignments you want to submit: ")
			numAssignments, _ := reader.ReadString('\n')

			// Convert the input to an integer
			numAssignments = strings.TrimSpace(numAssignments)
			numAssignmentsInt, err := strconv.Atoi(numAssignments)

			if err != nil {
				fmt.Println("Error: Please enter a valid number")
			}

			manager.SubmitAssignments(numAssignmentsInt)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "7":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
