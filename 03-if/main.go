// if faz uma pergunta e executa um bloco de código se a resposta for verdadeira
package main

import "fmt"

func main() {
	var idade int
	idade = 37

	if idade <= 18 {
		//bloco executa se a condição for verdadeira
		fmt.Println("Você é menor de idade")
	} else if idade >= 18 && idade < 65 {
		//bloco executa se essa condição for verdadeira
		fmt.Println("Você é adulto")
	}else {
		//bloco executa se nenhuma das condições anteriores forem verdadeiras
		fmt.Println("Você é idoso")
	}

	
}