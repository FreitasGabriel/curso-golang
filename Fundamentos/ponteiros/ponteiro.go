package main



func main(){
	i := 3

	//go não tem aritmética de ponteiro
	var p *int = nil

	p = &i //pegando o endereço da variável i e atribuindo ao ponteiro

	fmt.Println()
}