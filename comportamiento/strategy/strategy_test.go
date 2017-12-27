package strategy

import "testing"

func TestStrategyAdd(t *testing.T) {
	var expect = 15
	contextAdd := Contect{OpearationAdd{}}
	result := contextAdd.ExecuteStrategy(10, 5)

	if result != expect {
		t.Errorf("Estrategia Add esperaba %d pero devolvio %d", expect, result)
	}
}

func TestStrategySubstract(t *testing.T) {
	var expect = 5
	contextAdd := Contect{OpearationSubstract{}}
	result := contextAdd.ExecuteStrategy(10, 5)

	if result != expect {
		t.Errorf("Estrategia Substract esperaba %d pero devolvio %d", expect, result)
	}
}

func TestStrategyMultiply(t *testing.T) {
	var expect = 50
	contextAdd := Contect{OpearationMultiply{}}
	result := contextAdd.ExecuteStrategy(10, 5)

	if result != expect {
		t.Errorf("Estrategia Multiply esperaba %d pero devolvio %d", expect, result)
	}
}
