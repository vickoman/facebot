package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)
var p = fmt.Println
var pf = fmt.Printf

// Estructura Configuration
type Config struct {
	Port 	string `json:"port"`
	CertPem string `json:"cert_pem"`
	KeyPem 	string `json:"key_pem"`
}

var config Config

func main() {
	readConfig()

	http.HandleFunc("/", saludar)
	pf("Servidor iniciado https://knowblyhost:%s", config.Port)
	err := http.ListenAndServeTLS(config.Port, config.CertPem, config.KeyPem, nil)
	if err != nil {
		p(err)
	}

}

func saludar(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HTtp service with golang"))
}

func readConfig() {
	p("Leyendo el archivo de configuracion....")
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("No se pudo leer la configuracion inicial: %v", err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatalf("Error al convertir el archivo:  %v", err)
	}
	p("Archivo leido")
}
