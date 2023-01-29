package calc

import "testing"

func  TestAdd(t * testing.T) {
	result := Add(10,12)
	if result != 22 {
		t.Error("Expeted value is 25, got ", result)
	}
}


func TestSubstract(t *testing.T){
	result := Substract(100,48)
	if result != 52 {
		t.Error("Expected valu
		
		e is 52, got ", result)
	}
}

