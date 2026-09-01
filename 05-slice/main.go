// lista é u ma estrutura de dados que armazena uma colação de elementos do mesmo tipo. Em Go, as listas são implamentadas como slices

package main

import "fmt"
func main() {
	// existem 4 maneiras de criar um slice e popular com elementos.
	// 1. slice literal, declarção com os valores iniciais
	lista1 := []int{1, 2, 3, 4, 5}
	// resultado: [1 2 3 4 5]
	fmt.Println(lista1)

	// 2. Slice Vazio + append (Crescimento Dinâmico) 
	var numeros []int //slice nulo len=0, cap=0
	numeros = append(numeros, 6, 7, 8, 9, 10) // adicionando elementos ao slice
	// resultado: [6 7 8 9 10]
	fmt.Println(numeros)

	// 3. make com Tamanho Zero + append
	lista3 := make([]int, 0) 
	// resultado: []
	lista3 = append(lista3, 11, 12, 13, 14, 15)
	fmt.Println(lista3)

	// 4. make com Tamanho Fixo + Atribuição Direta por Índice
	lista4 := make([]int, 5) // cria um slice com tamnho 5 e capcidade 5
	// resultado: [0 0 0 0 0]
	// atribuindo valores aos 5 elementos do slice
	lista4[0] = 16
	lista4[1] = 17
	lista4[2] = 18
	lista4[3] = 19
	lista4[4] = 20
	fmt.Println(lista4)

 // Tentar atribuir em um índice que não existe causa um erro em tempo de execução (panic: runtime error: index out of range), exemplo:
  lista5 := []int{10, 20, 30} // indices validos: 0, 1, 2
  // lista5[3] = 40 // isso causaria um erro, pois o índice 3 não existe no slice lista5
  lista5 = append(lista5, 40) // forma correta de adicionar um elemento ao slice
  fmt.Println(lista5) // resultado: [10 20 30 40]	

  // percorrer um slice com for
  for i := 0; i < len(lista5); i++ {

	fmt.Println(lista5[i])
  }

  // criando sublsitas apartir de um slice existente. A sintaxe é: slice[índiceInicial:índiceFinal]
  lista6 := []int{100, 200, 300, 400, 500}
  segundaLista := lista6[:3] // cria um novo slice com os 3 primeiros elementos de lista6
  terceiraLista := lista6[2:] // cria um novo slice com os elementos a partir do índice 2 de lista6
  ultimoElemento := lista6[len(lista6)-1] // pega o último elemento do slice lista6
  fmt.Println("Segunda Lista:", segundaLista)
  fmt.Println("Terceira Lista:", terceiraLista)
  fmt.Println("Último Elemento:", ultimoElemento)

  // O fatiamento não cria uma cópia inteiramente nova na memória. Ele cria uma view (uma janela) que aponta para o mesmo array subjacente. então se a lista original for elterada as sublistas tambem serão alteradas.
}