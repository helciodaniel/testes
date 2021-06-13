package main

import (
	"log"
	"net/http"
	"html/template"
)

type Chave struct {
	nome string
}

func main() {
	http.HandleFunc("/", homeHandle)
	http.HandleFunc("/sala", salaHandle)
	log.Println("Ouvindo na porta 8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo!"))
}
func salaHandle(w http.ResponseWriter, r *http.Request) {
	ch := Chave{nome: "Luz 01"}

	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, ch)
}