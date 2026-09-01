// basico do for: sintaxe: for init; condition; post { }
package main

import "fmt"

func main() {
	for i:= 0; i < 10; i++ {
		fmt.Println(i)
	}

// fazendo iteração por índice. O texto[i] pega o byte naquela posição, por isso ainda precisa converter para string.
	texto := "Golang"
	for i :=0; i < len(texto); i++ {
		fmt.Println(string(texto[i]))
	}

// for range: itera sobre os elementos de um array, slice, string ou map. Retorna o índice e o valor do elemento.
	for i, letra := range texto {
		fmt.Println(i, string(letra))
	}

// fazendo uma tabuada
	for numBase := 1; numBase <= 10; numBase++ {
		for numMultiplicado := 1; numMultiplicado <= 10; numMultiplicado++ {
			fmt.Printf("%d x %d = %d\n", numBase, numMultiplicado, numBase*numMultiplicado)
		}
		fmt.Println()
	}

}

