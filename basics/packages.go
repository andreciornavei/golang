/*
Pacotes
Cada programa Go é composto de pacotes.
Programas começam rodando pelo pacote main.
Este programa está usando os pacotes com caminhos de importação "fmt" e "math/rand".
Por convenção, o nome do pacote é o mesmo que o último elemento do caminho de importação. 
Por exemplo, o pacote "math/rand" compreende arquivos que começam com package rand.
Nota: o ambiente em que esses programas são executados é determinístico, então rand.Intn 
sempre retornará o mesmo número.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
