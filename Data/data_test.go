package Data

import (
	"testing"
)

func TestGetRegistrationCount(t *testing.T) {

	GetRegistrationCount()

}
func TestCalRatio(t *testing.T) {

	a := CalRatio(1, 2)
	if a != 0.5 {
		t.Error("Expected", 0.5, "Got", CalRatio(1, 2))
	}

}
func TestCalRatio1(t *testing.T) {

	a := CalRatio(40, 100)
	if a != 0.4 {
		t.Error("Expected", 0.4, "Got", CalRatio(1, 2))
	}

}

func TestAddADNRegistration(t *testing.T) {

	AddADNRegistration(true, "[ATGCG,CAGTGC,TTATGT,AGAAGG,CCCCTA,TCACTG]")

}

func TestObtenerBaseDeDatos(t *testing.T) {

	obtenerBaseDeDatos()

}
