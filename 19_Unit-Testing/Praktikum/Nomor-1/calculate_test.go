package calculate

import "testing"

func TestAddition(t *testing.T){
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3");
	}
	if Addition(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3");
	}
}

func TestSubstract(t *testing.T){
	if Substract(5, 2) != 3 {
		t.Error("Expected 5 (-) 2 to equal 3");
	}
	if Substract(2, 5) != -3 {
		t.Error("Expected 2 (-) 5 to equal -3");
	}
}

func TestMultiplication(t *testing.T){
	if Multiplication(2, 5) != 10 {
		t.Error("Expected 2 (*) 5 to equal 10");
	}
	if Multiplication(-3, -4) != 12 {
		t.Error("Expected -3 (*) -4 to equal 12");
	}
}

func TestDivision(t *testing.T){
	if Division(6, 3) != 2 {
		t.Error("Expected 6 (/) 3 to equal 2");
	}
	if Division(10, 2) != 5 {
		t.Error("Expected 10 (/) 2 to equal 5");
	}
}