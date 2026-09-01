package main
import "fmt"
func main() {
	numeros := make([]int, 2)
	numeros[0] = 10
	numeros[1] = 20

	var soma int
	for i:=0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
		soma += numeros[i]
	} 

	fmt.Println("a soma do itens da lista é", soma)
}
