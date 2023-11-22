package main

import (
	"josefcalderon/gocero/webserver/middleware"
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

	/*numero, texto := ejercicios.ConvNumerico("hgjh")
	fmt.Println(numero)
	fmt.Println(texto)*/

	//teclado.IngresoNumeros()

	//iteraciones.Iterar()

	//fmt.Println(ejercicios.TabladeMultiplicar())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClouse()
	//funciones.Exponencia(2)
	/*
		Pedro := new(models.Hombre)
		ejerinterfaces.HumanosRespirando(Pedro)

		Maria := new(models.Mujer)
		ejerinterfaces.HumanosRespirando(Maria)*/
	/*canal1 := make(chan bool)
	go gorutines.MiNombreLentooo("EJCALDERON", canal1)
	defer func() {
		<-canal1
	}()

	//go gorutines.MiNombreLentooo("Edwin Josef")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)*/

	middleware.MyMiddleware()
}
