//<Author> Diego Ernesto Puentes Matta</Author>
//Prueba técnica Examen Mercadolibre
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dna string

// type Post struct {
// 	ID string `json:"id"`
// 	Title string `json:"title"`
// 	Body string `json:"body"`
//   }

func mutant(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		dna = string(body)
		fmt.Println(dna)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	var response = true
	// dna := [6]string{"ATGCwA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"} //array solicitado en el ejemplo.
	// isMutand(dna)
	fmt.Println()
	if response {
		w.WriteHeader(200)
		w.Write([]byte("Mutant found"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Humand found"))
	}
}

func request() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/mutant", mutant).Methods("POST")
	myRouter.HandleFunc("/stats", mutant).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	request()
}
