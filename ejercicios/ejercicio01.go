package ejercicios

import (
	"strconv"
)

func ConvNumerico(value string) (int, string) {
	// usamos _ para indicar que el argumento devuelto no lo estare usando
	//valueInt, _ := strconv.Atoi(value, 10, 64)
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, "Hubo error" + err.Error()
	}
	if valueInt > 100 {
		return valueInt, "Es mayor a 100"
	} else {
		return valueInt, "Es menor a 100"
	}
}
