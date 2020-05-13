package main

//Importações
//Este grupo de códigos em parênteses da importação, é a declaração de importação "consignada".
import (
	"fmt"
	"math"
);


// Você também pode escrever várias declarações de importação assim:
//import "fmt"
//import "math"
//Mas é um bom estilo usar a declaração de importação consignada.

func main(){
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}