package Calculator

import (
	"testing"
)

func TestSum(t *testing.T) {
	Result := Sum(10, 12)
	Expected := 22.0
	if Result != Expected {
		t.Errorf("The sum function retutned wrong answer : %f and %f was expected", Result, Expected)
	}
}

func TestSubtract(t *testing.T) {
	Result := Subtract(20, 2)
	Expected := 18.0
	if Result != Expected {
		t.Errorf("The Subtract function retutned wrong answer : %f and %f was expected", Result, Expected)
	}
}

func TestDivide(t *testing.T) {
	Result := Divide(90, 3)
	Expected := 30.0
	if Result != Expected {
		t.Errorf("The divide function retutned wrong answer : %f and %f was expected", Result, Expected)
	}
}

func TestDivideByZero(t *testing.T) {
	Result := Divide(90, 0)
	Expected := 0.0
	if Result != Expected {
		t.Errorf("The divide by zero function retutned wrong answer : %f and %f was expected", Result, Expected)
	}
}
