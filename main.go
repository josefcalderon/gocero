package main

import (
	"fmt"
	"josefcalderon/gocero/ejercicios"
)

func main() {
	//variables.RestoVariables()
	/*estado, texto := variables.ConviertoaTexto(1500)
	fmt.Println(estado)
	fmt.Println(texto)*/

	/*
		// Practicando el IF
		if os := runtime.GOOS; os == "linux" || os == "OS X" {
			fmt.Println("Es linux", "es : ", os)
		} else {
			fmt.Println("No es linux", "es : ", os)
		}

		// Practicando el Switch
		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es linux")
		case "darwin":
			fmt.Println("Esto es darwin")
		default:
			fmt.Printf("%s \n", os)
		}*/
	numero, texto := ejercicios.ConvNumerico("hgjh")
	fmt.Println(numero)
	fmt.Println(texto)

}
