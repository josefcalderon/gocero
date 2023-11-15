package main

import (
	"fmt"
	"josefcalderon/gocero/variables"
)

func main() {
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1500)
	fmt.Println(estado)
	fmt.Println(texto)

}
