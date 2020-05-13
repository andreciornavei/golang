/*
Variáveis
A instrução var declara uma lista de variáveis, como em listas de argumentos de função, o tipo é o último passado.
A declaração var pode estar num pacote ou a nível de função. Nós vemos ambos neste exemplo.
*/

package main

import "fmt"

var c, python, java bool

func main(){
	var i int
	fmt.Print(i, c, python, java)
}