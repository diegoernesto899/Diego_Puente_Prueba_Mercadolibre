package isMutant

import "testing"

func TestIsMutand(t *testing.T) {
	dna := [6]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	h := IsMutand(dna)
	if h != "true" {
		t.Error("Expected", "true", "Got", IsMutand(dna))
	}

}

func TestIsMutand2(t *testing.T) {
	dna := [6]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
	h := IsMutand(dna)
	if h != "false" {
		t.Error("Expected", "false", "Got", IsMutand(dna))
	}

}

func TestValidMutantADN(t *testing.T) {
	dna := [6]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	h := validMutantADN(dna)
	if h < 1 {
		t.Error("Expected", "res > 1", "Got", validMutantADN(dna))
	}

}
