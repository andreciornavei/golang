/*
Resultados Múltiplos
Uma função pode retornar qualquer número de resultados.
A função swap retorna duas strings.
*/

package main

import "fmt"

func swap(x, y string) (string, string){
	return x, y
}

func main(){
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}