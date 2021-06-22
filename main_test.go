package main

import "testing"

func TestHandler(t *testing.T) {
	h := stats()
	if h.StatusCode != 200 {
		t.Error("Expected", 200, "Got", h.StatusCode)
	}

}

func TestMutantHttpr(t *testing.T) {
	h := returnMutantHttp("test", 200)
	if h.StatusCode != 200 {
		t.Error("Expected", 200, "Got", h.StatusCode)
	}

}
func TestReturnMutantHttp(t *testing.T) {

	a := returnMutantHttp("Es un mutante", 400)
	if a.StatusCode != 400 {
		t.Error("Expected", 400, "Got", a.StatusCode)
	}
}
