package main

import "fmt"
import "net/http"
import "html/template"
import "path"

func main() {
    // Buat handler dengan rute `/`
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        // Fungsi path.Join() digunakan untuk menggabungkan folder, file atau keduanya menjadi sebuah path.
        var filepath = path.Join("views", "index.html") // menghasilkan `views/index.html`

        // Fungsi `template.ParseFiles()`, digunakan untuk parsing file template
        var tmpl, err = template.ParseFiles(filepath)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Data map yang berisikan key `title` dan `name` yang akan di sisipkan ke dalam template yang sudah di parsing.
        var data = map[string]interface{}{
            "title": "Learning Golang Web",
            "name":  "Student",
        }

        // Method `Execute()` milik `*template.Template`, digunakan untuk menyisipkan data pada template, untuk kemudian ditampilkan ke browser. Data bisa disipkan dalam bentuk `struct`, `map`, atau `interface{}`.
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

	http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("assets"))))

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}