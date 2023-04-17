package q3

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	package main

import "fmt"

// Você está desenvolvendo um programa que recebe uma lista de números inteiros como entrada e precisa encontrar o
// maior valor e o menor valor da lista. Além disso, você precisa calcular a média dos valores da lista, desconsiderando o maior e o menor valor na média
// Escreva uma função que recebe uma lista de números inteiros como parâmetro e retorna uma tupla contendo o maior valor,
// o menor valor e a média dos valores da lista, desconsiderando o maior e o menor valor.
// A função deve retornar um erro caso a lista esteja vazia. O erro deve conter a mensagem "lista vazia".
func medialist(list []int) (int, int, float64, error) {
	var soma, maximo, min int
	var media float64
	var count float64
	soma = 0
	maximo = 0
	min = 10000
	for x := 0; x < len(list); x++ {
		if list[x] > maximo {
			maximo = list[x]
		}
		if list[x] < min {
			min = list[x]
		}
		if list[x] < maximo && list[x] > min {
			soma += list[x]
		}
		count++
	}
	if len(list) == 0 {
		return 0, 0, 0, fmt.Errorf("Lista vazia")
	}
	count -= 2
	media = float64(soma) / float64(count)
	return maximo, min, media, nil
}
func main() {
	var lista = []int{2, 6, 7, 5, 4, 3, 8, 9}
	fmt.Print(medialist(lista))
}

	return 0, 0, 0, nil
}
