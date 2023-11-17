package deferpanic

import (
	"fmt"
	"log"
)

// Defer permite indicar al sistema cuales van a ser la instrucciones finales.
func VemosDefer() {
	fmt.Println("Este es un primer mensaje.")
	defer fmt.Println("Este es el mensaje final.")
	fmt.Println("Este es el segundo mensaje.")
}

func VemosPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1.")
	}
}
