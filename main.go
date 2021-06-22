package main

//<Author> Diego Ernesto Puentes Matta</Author>
//Prueba técnica Examen Mercadolibre
//goos=linux go build -o main main.go
import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	D "github.com/diegoernesto899/Diego_Puentes_Prueba_Mercadolibre/Data"
	L "github.com/diegoernesto899/Diego_Puentes_Prueba_Mercadolibre/pkg"
)

type jsonStruct struct {
	RequetADN [6]string `json:"dna"`
}
type ResponseStats struct {
	Count_mutant_dna int     `json:"count_mutant_dna"`
	Count_human_dna  int     `json:"count_human_dna"`
	Ratio            float64 `json:"ratio"`
}

var Stats ResponseStats

// Esta funcion concatena la respuestas
func returnMutantHttp(respuestaString string, code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       string(respuestaString),
		StatusCode: http.StatusBadRequest,
	}
}

//Esta funcion valida si es humano o mutante y retorna la respuesta http
func mutant(r events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	t := jsonStruct{}
	json.Unmarshal([]byte(r.Body), &t) //mapear el parametro adn  al objeto jsonStruct

	adnString := strings.Join(t.RequetADN[:], ",") // formar array de adn
	isMutant := L.IsMutand(t.RequetADN)            //validacion si es mutante o humano
	if isMutant == "true" {                        //asignacion de respuestas
		D.AddADNRegistration(true, adnString)
		return returnMutantHttp("Mutant found", http.StatusOK)
	} else if isMutant == "false" {
		D.AddADNRegistration(false, adnString)
		return returnMutantHttp("Humand found", http.StatusForbidden)
	} else {
		return returnMutantHttp(isMutant, http.StatusBadRequest)
	}
}

//Obtiene estadisticas de registros adn
func stats() events.APIGatewayProxyResponse {
	mutan, human, ratio := D.GetRegistrationCount()
	Stats = ResponseStats{Count_mutant_dna: mutan, Count_human_dna: human, Ratio: ratio}

	b, err := json.Marshal(Stats)
	D.ErrorCheck(err)
	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}
}

//Controlador que captura el formato http y asigna los submetodos
func Handler(c context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(request)
	log.Println("PathResult" + request.Path)

	if request.HTTPMethod == "GET" {
		return stats(), nil
	} else if request.HTTPMethod == "POST" {
		return mutant(request), nil
	}
	return returnMutantHttp("El metodo no ha sido encontrado", http.StatusNotFound), nil
}
func main() {
	lambda.Start(Handler)
}
