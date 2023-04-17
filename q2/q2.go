package q2

func AverageLettersPerWord(text string) (float64, error) {
	package main

import "fmt"

func medialetras(texto string) (float64, error) {
	var palavras, letras, media float64
	letras = 0
	palavras = 0
	for x := 0; x < len(texto); x++ {
		letras++
		if string(texto[x]) == " " {
			letras--
			palavras++
		}
	}
	palavras++
	media = letras / palavras
	if len(texto) == 0 {
		return 0, fmt.Errorf("O texto está vazio")
	}
	return media, nil
}

func main() {
	var str string
	str = ""
	result, err := medialetras(str)
	if err != nil {
		fmt.Print("O texto está vazio")
	} else {
		fmt.Printf("A média de letras por palavra é: %.2f", result)
	}
}

	return 0, nil
}
