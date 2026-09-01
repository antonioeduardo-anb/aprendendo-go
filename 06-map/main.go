// map é uma estrutura de dados que armazena uma coleção de pares chave-valor. Em go os maps são implementados como hash tables.

package main

import "fmt"

func main() {
	m := make(map[string]int) // cria um map vazio com chave do tipo string e valor do tipo int
	m["um"] = 1 // adiciona um par chave-valor ao map
	m["dois"] = 2
	m["tres"] = 3

	fmt.Println(m) // resultado: map[dois:2 tres:3 um:1]

	// acessando um valor do map
	valor, existe := m["dois"]  // verifica se a chave "dois" existe no map e retorna o valor associado a ela

	if existe {
		fmt.Println("O valor da chave \"dois\" é:", valor)
	} else {
		fmt.Println("A chave \"dois\" não existe no map.")
	}

	// range em maps: itera sobre os pares chave-valor do map. A ordem de iteração não é garantida.
	for chave, valor := range m {
		fmt.Printf("Chave: %s, Valor: %d\n", chave, valor)
	}

	// deletando um elemento do map
	delete(m, "um") // remove o par chave-valor com a chave "um"
	fmt.Println("Map após deletar a chave \"um\":", m) // resultado: map[dois:2 tres:3]
}