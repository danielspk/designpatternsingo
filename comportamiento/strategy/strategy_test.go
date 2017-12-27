package strategy

import "testing"

func TestStrategyAdd(t *testing.T) {
	var expect = 15
	contextAdd := Context{OpearationAdd{}}
	result := contextAdd.ExecuteStrategy(10, 5)

	if result != expect {
		t.Errorf("Estrategia Add esperaba %d pero devolvió %d", expect, result)
	}
}

func TestStrategySubstract(t *testing.T) {
	var expect = 5
	contextAdd := Context{OpearationSubstract{}}
	result := contextAdd.ExecuteStrategy(10, 5)

	if result != expect {
		t.Errorf("Estrategia Substract esperaba %d pero devolvió %d", expect, result)
	}
}

func TestStrategyMultiply(t *testing.T) {
	var expect = 50
	contextAdd := Context{OpearationMultiply{}}
	result := contextAdd.ExecuteStrategy(10, 5)

	if result != expect {
		t.Errorf("Estrategia Multiply esperaba %d pero devolvió %d", expect, result)
	}
}
