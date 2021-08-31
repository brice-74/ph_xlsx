package main

import (
	"testing"
	"reflect"
)

// Const Initializers
func _RowsTest() (*[][]string) {
	return &[][]string{
		{"Code analytique 1", "%", "Code analytique 2", "%", "Code analytique 3", "%", "rubrique 1", "rubrique 2", "rubrique 3",},
		{"T1", "100,00", "", "", "", "", "200,00", "0", "50,00",},
		{"T2", "50,00", "T1", "25,00", "T3", "25,00", "0", "-1000,00", "100,00",},
	}
}

func _MapRowsTest() (*map[int]map[string]string) {
	return &map[int]map[string]string{
		1:{"Code analytique 1":"T1", "%1":"100,00", "rubrique 1":"200,00", "rubrique 3":"50,00",},
		2:{"Code analytique 1":"T2", "%1":"50,00", "Code analytique 2":"T1", "%2":"25,00", "Code analytique 3":"T3", "%3":"25,00", "rubrique 2":"-1000,00", "rubrique 3":"100,00",},
	}
}

func _MapSumTest() (*map[string]map[string]float64) {
	return &map[string]map[string]float64{
		"T1":{"rubrique 1":200.00, "rubrique 2":-250.00, "rubrique 3":75.00,},
		"T2":{"rubrique 2":-500.00, "rubrique 3":50.00,},
		"T3":{"rubrique 2":-250.00, "rubrique 3":25.00,},
	}
} 

func _HeaderCod() (*[]string) {
	return &[]string{"Code analytique 1", "%1", "Code analytique 2", "%2", "Code analytique 3", "%3",}
}

func _HeaderRub() (*[]string) {
	return &[]string{"rubrique 1", "rubrique 2", "rubrique 3",}
}

func _HeaderGlob() (*[]string) {
	return &[]string{"Code analytique 1", "%1", "Code analytique 2", "%2", "Code analytique 3", "%3", "rubrique 1", "rubrique 2", "rubrique 3",}
}

// Tests
func TestExtractHeaders(t *testing.T) {
	headers := extractHeaders(_RowsTest())

	isArrEq(t, *_HeaderCod(), headers.codification)
	isArrEq(t, *_HeaderRub(), headers.rubrique_paie)
	isArrEq(t, *_HeaderGlob(), headers.global)
} 

func isArrEq(t *testing.T, want []string, got []string) {
	if ok := reflect.DeepEqual(want, got); !ok {
		t.Errorf("want %v;\ngot %v", want, got)
	}
	return
}

func TestMapRows(t *testing.T) {
	mapRows := mapRows(_RowsTest(), &Headers{
		codification:*_HeaderCod(),
		rubrique_paie:*_HeaderRub(),
		global:*_HeaderGlob(),
	})

	if ok := reflect.DeepEqual(*_MapRowsTest(), *mapRows); !ok {
		t.Errorf("want %v;\ngot %v", *_MapRowsTest(), *mapRows)
	}
}

func TestMapSum(t *testing.T) {
	sumMapping := mapSum(_MapRowsTest(), &Headers{
		codification:*_HeaderCod(),
		rubrique_paie:*_HeaderRub(),
		global:*_HeaderGlob(),
	})

	if ok := reflect.DeepEqual(*_MapSumTest(), *sumMapping); !ok {
		t.Errorf("want %v;\ngot %v", *_MapSumTest(), *sumMapping)
	}
}