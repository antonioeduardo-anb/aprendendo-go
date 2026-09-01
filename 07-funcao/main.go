package main
import "fmt"

func main() {
	// escutando a função saudacao
	saudacao("Antonio")
	saudacao("Eduardo")
	
	//a função multiplica retorna um inteiro
	resultado := multiplica(10, 5)
	println(resultado)

	// armazenando os valores retornados pela função
	soma, subtracao, multiplicacao, divisao := operacoes(10, 5)
	println(soma, subtracao, multiplicacao, divisao)
}
// função, bloco de codigo reultilizavel
func saudacao( nome string) {
	nome += ", Bom Dia."
	fmt.Println(nome)
}
// a funçao returna um valor com a palavra reservada return
func multiplica(numero1 int, numero2 int) int{
	resultado := numero1 * numero2
	return resultado
}

// retornando varios valores
func operacoes(numero1 int, numero2 int) (int, int, int, int) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	multiplicacao := numero1 * numero2
	divisao := numero1 / numero2
	return soma, subtracao, multiplicacao, divisao
}

// exemplo de função com retornos nomeados, 
func operacoes2(numero1 int, numero2 int) ( soma int,  subtracao int,  multiplicacao int, divisao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	multiplicacao = numero1 * numero2
	divisao = numero1 / numero2
	return
}