package main

//<Author> Diego Ernesto Puentes Matta</Author>
//Prueba técnica Examen Mercadolibre
import (
	"encoding/json"
	D "github.com/diegoernesto899/Diego_Puentes_Prueba_Mercadolibre/Data"
	L "github.com/diegoernesto899/Diego_Puentes_Prueba_Mercadolibre/pkg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type jsonStruct struct {
	Test [6]string `json:"dna"`
}

func mutant(w http.ResponseWriter, r *http.Request) {
	D.TestCon()
	decoder := json.NewDecoder(r.Body) //Obtener json param dna form r request
	var t jsonStruct
	err := decoder.Decode(&t) //mapear el parametro adn  al objeto jsonStruct
	if err != nil {
		panic(err)
	}

	isMutant := L.IsMutand(t.Test)
	if isMutant == "true" {
		w.WriteHeader(200)
		w.Write([]byte("Mutant found"))
	} else if isMutant == "false" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Humand found"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(isMutant))
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
