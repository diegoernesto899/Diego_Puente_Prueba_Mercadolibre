package main

//<Author> Diego Ernesto Puentes Matta</Author>
//Prueba técnica Examen Mercadolibre
import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	D "github.com/diegoernesto899/Diego_Puentes_Prueba_Mercadolibre/Data"
	L "github.com/diegoernesto899/Diego_Puentes_Prueba_Mercadolibre/pkg"
	"github.com/gorilla/mux"
)

type jsonStruct struct {
	RequetADN [6]string `json:"dna"`
}

func mutant(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body) //Obtener json param dna form r request
	var t jsonStruct
	err := decoder.Decode(&t) //mapear el parametro adn  al objeto jsonStruct
	D.ErrorCheck(err)

	adnString := strings.Join(t.RequetADN[:], ",")

	isMutant := L.IsMutand(t.RequetADN)
	if isMutant == "true" {
		D.AddADNRegistration(true, adnString)
		w.WriteHeader(200)
		_, e := w.Write([]byte("Mutant found"))
		D.ErrorCheck(e)
	} else if isMutant == "false" {
		D.AddADNRegistration(false, adnString)
		w.WriteHeader(http.StatusForbidden)
		_, e := w.Write([]byte("Humand found"))
		D.ErrorCheck(e)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, e := w.Write([]byte(isMutant))
		D.ErrorCheck(e)
	}
}

type Repond struct {
	Count_mutant_dna int     `json:"count_mutant_dna"`
	Count_human_dna  int     `json:"count_human_dna"`
	Ratio            float64 `json:"ratio"`
}

var Response Repond

func stats(w http.ResponseWriter, _ *http.Request) {
	mutan, human, ratio := D.GetRegistrationCount()
	Response = Repond{Count_mutant_dna: mutan, Count_human_dna: human, Ratio: ratio}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(Response)
	D.ErrorCheck(err)
}

func request() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/mutant", mutant).Methods("POST")
	myRouter.HandleFunc("/stats", stats).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	request()
}
