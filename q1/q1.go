package q1

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	// package main

import "fmt"

func eli(valor float64, histórico []float64) (float64, error) {
	var soma, media float64
	var count, desconto float64
	if valor > 1000 {
		desconto = valor * 0.1
	}
	if valor <= 1000 && valor > 500 {
		desconto = valor * 0.05
	}
	if valor <= 500 {
		desconto = valor * 0.02
	}
	if len(histórico) == 0 {
		desconto = valor * 0.1
	}
	for x := 0; x < len(histórico); x++ {
		soma += histórico[x]
		count++
	}
	media = soma / count
	if media > 1000 {
		desconto = valor * 0.2
	}
	if valor <= 0 {
		return 0, fmt.Errorf("Não possui valor de compra atual")
	}
	return desconto, nil

}

func main() {
	var compra float64
	var histcomp []float64
	compra = 500
	result, err := eli(compra, histcomp)
	if err != nil {
		fmt.Println("Não possui valor de compra atual")
	} else {
		fmt.Printf("O valor do desconto é: %.2f", result)
	}
}

	return 0, nil
}
