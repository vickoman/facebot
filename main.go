package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", saludar)
	log.Println("Servidor iniciado http://knowblyhost:9087")
	log.Println(http.ListenAndServe(":8087", nil))
}

func saludar(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HTtp service with golang"))
}
