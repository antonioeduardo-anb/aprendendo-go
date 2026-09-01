package main
import (
	"fmt" 
	"reflect"
)

func main()  {
	var nome string
	nome = "Eduardo"
	var idade int
	idade = 37
	//omitindo tipo
	nacionalidade := "brasileiro"
	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(nacionalidade)

	//tipos primitivos e seus valore padrão
	var numero int
	var texto string
	var metro float32
	var ehfeminino bool

	fmt.Println(numero) // 0
	fmt.Println(texto) // vazio
	fmt.Println(metro) // 0
	fmt.Println(ehfeminino)// false

	//ver o tipo da variavel
	fmt.Println("o tipo é", reflect.TypeOf(numero))

	//constantes para vaalores que nunca mundam
	const taxa = 1.5
	fmt.Println("a taxa é", taxa)
	fmt.Println(" o tipo da cosntante é", reflect.TypeOf(taxa))

	
	
}