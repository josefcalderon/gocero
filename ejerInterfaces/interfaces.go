package ejerinterfaces

import (
	"fmt"
	"josefcalderon/gocero/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("soy un/a %s, y estoy respirando \n", hu.Sexo())
	//hu.Comer()
	//hu.Pensar()

}
