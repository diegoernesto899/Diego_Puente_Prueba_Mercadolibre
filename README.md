# Diego_Puentes_Prueba_Mercadolibre
Prueba mercado libre magneto

Se realiza prueba en lenguage GOLANG 
Implementando Funciones Lambda y servicio REST API GATEWAY de AMAZON WEB SERVICES

URLs Productivas Publicadas
Para el consumo de los metodos relacionados a continuacion se recomienda sean consumidos desde postman,
(puede usar cualquier software para peticiones de metodos HTTP)

Nivel 2:
API REST en donde se pueda detectar si un humano es mutante enviando la secuencia de ADN

Metodo POST 
URL:  https://q458a479i4.execute-api.us-east-2.amazonaws.com/prod/mutant
request Object

{
“dna”:["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}

return  HTTP 200-OK si es mutante y 403-Forbidden si es humano


Nivel 3: 
Metodo GET 
Servicio extra “/stats” que devuelva un Json con las estadísticas de las verificaciones de ADN

URL:  https://q458a479i4.execute-api.us-east-2.amazonaws.com/prod/stats

Return Json Object 
{
    "count_mutant_dna": 6,
    "count_human_dna": 5,
    "ratio": 1.2
}

Se integra base de datos RDS My SQL de aws 

Configuración base de datos
Se creó una base de datos de mysql alojada en Amazon Web Service

Base de datos : DBMercadoLibre
Tabla: TB_VerificacionADN

Diccionario de datos

id_dna	int	 auto_increment  id registro

isMutant_dna 	tinyint  valor boleano si es adn mutante

input_dna   varchar(45)   string adn ingresado

registrationDate_dna  datetime  fecha del registro


Se realizaron testing COVERAGE > 80 %
