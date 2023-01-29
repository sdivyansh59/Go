package calc

import (
	"testing"
)

func TestAdd(t *testing.T){
	if result := Add(10,2) ; result != 12 {
		t.Error("Expected 25, got ", result)
	}
	
}

func TestSubstract(t *testing.T){
	if result := Substract(10,2) ; result != 8 {
		t.Error("Expected 8, got ", result)
	}
	
}