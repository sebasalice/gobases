//Los profesores de una universidad de Colombia necesitan calcular
//algunas estadísticas de calificaciones de los alumnos de un curso,
//requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.
//Se solicita generar una función que indique qué tipo de cálculo
//se quiere realizar (mínimo, máximo o promedio) y que devuelva
//otra función ( y un mensaje en caso que el cálculo no esté definido)
//que se le puede pasar una cantidad N de enteros y devuelva el
//cálculo que se indicó en la función anterior
//Ejemplo:

package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {

	fmt.Println(minFunc(2, 3, 3, 4, 1, 2, 4, 5, -1))
	fmt.Println(averageFunc(2, 3, 3, 4, 1, 2, 4, 5))
	fmt.Println(maxFunc(2, 3, 3, 4, 1, 2, 4, 5))
}

func minFunc(valores ...int) (message string, min int) {
	message = "El valor minimo es: "
	min = valores[0]
	for _, valor := range valores {
		if valor > 0 {
			{
				if valor < min {
					min = valor
				}
			}
		}
	}

	return message, min
}
func averageFunc(valores ...int) (message string, min int) {
	message = "El valor promedio es: "
	var average int
	longitud := len(valores)
	for _, valor := range valores {
		if valor > 0 {
			average = average + valor
		}

	}

	return message, (average / longitud)
}
func maxFunc(valores ...int) (message string, min int) {
	message = "El valor maximo es: "
	var max int
	max = valores[0]
	for _, valor := range valores {
		if valor > 0 {
			{
				if valor > max {
					max = valor
				}
			}
		}
	}

	return message, max
}
