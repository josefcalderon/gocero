package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  // Slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] // Slice creado desde un vector, desde la posicion 6 hasta la 7

	fmt.Printf("Desde posicion 3 : %v\n", porcion)
	fmt.Printf("Desde posicion 0 a 5 : %d\n", porcion2)
	fmt.Printf("Desde posicion 6 a 7 : %d\n", porcion3)
}

// Los slice se pueden acortar o agrandar en tiempo de ejecucion.
func Capacidad() {
	// tiene 5 elementos pero su capacidad es de 20
	elemntos := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad: %d \n", len(elemntos), cap(elemntos))

	// Slices nos permite AUmentar o reducir en tiempo de ejecucion y reserva memoria.
	// Aumenta la reserva de capacidad arreglo y reserva capacidad de memoria
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo: %d, Capacidad: %d \n", len(nums), cap(nums))
}
