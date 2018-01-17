package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func holaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo!")
}

func holaNosotros(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Nosotros!")
}

func holaEllos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Ellos!")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.msg)
}

func main() {
	msg := mensaje{
		msg: "Hola Universo",
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	mux.HandleFunc("/nosotros", holaNosotros)
	mux.HandleFunc("/ellos", holaEllos)
	mux.Handle("/hola", msg)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())

}
