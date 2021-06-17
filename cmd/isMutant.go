//<Author> Diego Ernesto Puentes Matta</Author>
//Prueba técnica Examen Mercadolibre
package main

import (
	"fmt"
	"regexp"
)

func main() {
	dna := [6]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"} //array solicitado en el ejemplo.
	//dna := [6]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"} //Ejemplo array humano
	//dna := [6]string{"ATGCGA", "CAGTGC", "TTATGT", "AAAAGG", "CCCCTA", "TCACTG"} //Ejemplo doble cadena horizontal

	if validRequest(dna) {
		isMutand := isMutand(dna)
		fmt.Println(isMutand)
	} else {
		fmt.Println("La cadena de ADN contiene valores no validos")
	}
}

//<Description>Esta función recorre el array y busca si hay mas de 1 secuencia de 4 letras iguales (ADN Mutante)</descripcion>
//<Params>Recibe array de string con secuencia de adn</params>
//<Return>Booleano si es o no Mutante</return>
func isMutand(dna [6]string) bool {

	var countHorizontal = validMutantADN(dna) //retorna cantidad de adn mutante horizontal

	if countHorizontal > 1 {
		return true // Retorna verdadero si la funcion validMutantADN es > 1
	}
	var verticalADN = mappingStringVertical(dna) //Convierte string horizontal en vertical
	var vertical = validMutantADN(verticalADN)   //retorna cantidad de adn mutante verticalmente
	horVert := countHorizontal + vertical

	if horVert > 1 { //Valida si en el recorrido horizontal mas el vertical es mutante
		return true
	} else {
		var diagonal = mappingStringOblicua(dna, verticalADN) //retorna cantidad de adn mutante de forma oblicua
		resultado := horVert + diagonal
		if resultado > 1 {
			return true
		}
	}
	return false //si no se encuentra mas de una secuencia retorna false
}

//<Description>Esta función recorre el array de manera horizontal</descripcion>
//<Params>recibe array de string con secuencia de adn</params>
//<Return>Cantidad de patrones encontrados de adn mutante</return>
func validMutantADN(dna [6]string) int {
	var contadorADN int
	var digitoActual rune
	var digitoAnt rune
	var contadorIguales int
	for _, value := range dna {
		//Recorre los digitos horizontalmente para comparar si son iguales
		for _, digito := range value {
			if digitoAnt == 0 {
				digitoAnt = digito
			}
			digitoActual = digito

			if digitoAnt == digitoActual { //Compara el digito anterior con el actual si son iguales agrega al contador
				contadorIguales++
				if contadorIguales == 4 {
					//  fmt.Println("hay uno" + value) //esta variable muestra la cadena de ADN Mutante
					contadorADN++
				}
			} else {
				digitoAnt = digitoActual
				contadorIguales = 0 //Se reinicia el contador si los digitos no son iguales
			}

		}
		digitoAnt = 0
	}
	// fmt.Println("contadorADN = " + strconv.Itoa(contadorADN))
	return contadorADN
}

//<Description>Esta función transforma el array horizontal lo formatea de manera vertical</descripcion>
//<Params>recibe array de string con secuencia de adn</params>
//<Return>Cantidad de patrones encontrados de adn mutante</return>
func mappingStringVertical(dna [6]string) [6]string {
	var newArrayADN [6]string
	var p1 string
	var p2 string
	var p3 string
	var p4 string
	var p5 string
	var p6 string
	for _, value := range dna {
		p1 = p1 + string(value[0])
		p2 = p2 + string(value[1])
		p3 = p3 + string(value[2])
		p4 = p4 + string(value[3])
		p5 = p5 + string(value[4])
		p6 = p6 + string(value[5])
	}
	newArrayADN[0] = p1
	newArrayADN[1] = p2
	newArrayADN[2] = p3
	newArrayADN[3] = p4
	newArrayADN[4] = p5
	newArrayADN[5] = p6

	return newArrayADN
}

//<Description>Esta función transforma el array horizontal lo formatea de manera Oblicua</descripcion>
//<Params>recibe array de string con secuencia de dna y string dna vertical</params>
//<Return>Cantidad de patrones encontrados de adn mutante en forma oblicua</return>
func mappingStringOblicua(dna [6]string, verticalADN [6]string) int {
	diagonal1 := validMutantADN(mappingOblicuaIzq(dna))
	diagonal2 := validMutantADN(mappingOblicuaDer(verticalADN))
	a := reverseArray(dna)
	b := reverseArray(verticalADN)
	diagonalInvertida1 := validMutantADN(mappingOblicuaIzq(a))
	diagonalInvertida2 := validMutantADN(mappingOblicuaDer(b))
	return diagonal1 + diagonal2 + diagonalInvertida1 + diagonalInvertida2
}

//<Description>Esta función transforma el array horizontal lo formatea de manera Oblicua del centro a la izquierda</descripcion>
//<Params>recibe array de string con secuencia de dna</params>
//<Return>string array de forma oblicuadel centro a la izquierda</return>
func mappingOblicuaIzq(dna [6]string) [6]string {
	var newArrayADN [6]string
	var dato1 string
	arrayLengt := len(dna)
	for i := 0; i < arrayLengt; i++ {
		dna := dna[i:]
		for d1, value := range dna {
			dato1 = dato1 + string(value[d1])
		}
		if len(dato1) >= 4 {
			newArrayADN[i] = dato1
		}
		dato1 = ""
	}
	// fmt.Println("diagonal", newArrayADN)
	return newArrayADN
}

//<Description>Esta función transforma el array horizontal lo formatea de manera Oblicua del centro a la izquierda</descripcion>
//<Params>recibe array de string con secuencia de dna</params>
//<Return>string array de forma oblicuadel centro a la derecha</return>
func mappingOblicuaDer(dna [6]string) [6]string {
	var newArrayADN [6]string
	var dato1 string
	msg := make([]byte, 0)
	arrayLengt := len(dna)
	for i := 0; i < arrayLengt; i++ {
		dna := dna[i:]
		for d1, value := range dna {
			dato1 = dato1 + string(value[d1])
			msg = append(msg, []byte(string(value[d1]))[0])
		}
		if len(dato1) >= 4 && len(dato1) < 6 {
			newArrayADN[i] = dato1
		}
		dato1 = ""
	}
	// fmt.Println("diagonal", newArrayADN)
	return newArrayADN
}

//<Description>Esta función recorre el array y lo retorna de manera inversa</descripcion>
//<Params>recibe array de string con secuencia de adn</params>
//<Return>string array cadena inversa</return>
func reverseArray(dna [6]string) [6]string {
	last := len(dna) - 1
	for i := 0; i < len(dna)/2; i++ {
		dna[i], dna[last-i] = dna[last-i], dna[i]
	}
	return dna
}

//<Description>Esta función valid mediante regex si los parametros ingresados cumplen el requerimiento</descripcion>
//<Params>recibe array de string con secuencia de adn</params>
//<Return>parametro Booleano si cumple las reglas</return>
func validRequest(dna [6]string) bool {
	var re = regexp.MustCompile(`(?m)[aAtTcCgG]$`)
	var valid bool
	for _, value := range dna {
		valid = re.MatchString(value)
		if !valid {
			return false
		}
	}
	return true
}
