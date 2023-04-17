package main

import "fmt"

func precoprod(precbase float64, estado string, imposto int) (float64, error) {
	if precbase <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}
	var perfrete, aliquota float64
	if estado == "SP" {
		perfrete = 0.1
	} else if estado == "RJ" {
		perfrete = 0.15
	} else if estado == "MG" {
		perfrete = 0.2
	} else if estado == "ES" {
		perfrete = 0.25
	} else {
		perfrete = 0.3
	}
	if imposto == 1 {
		aliquota = 0.05
	} else if imposto == 2 {
		aliquota = 0.1
	} else if imposto == 3 {
		aliquota = 0.15
	} else {
		return 0, fmt.Errorf("imposto não encontrado")
	}
	precof := precbase + precbase*aliquota + precbase*perfrete
	return precof, nil
}

func main() {
	base := 300.00
	state := "RJ"
	codimposto := 2
	fmt.Print(precoprod(base, state, codimposto))

}
