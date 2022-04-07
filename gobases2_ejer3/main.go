//Una empresa marinera necesita calcular el salario de sus empleados
//basándose en la cantidad de horas trabajadas por mes y la categoría.

//Si es categoría C, su salario es de $1.000 por hora
//Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

//Se solicita generar una función que reciba por parámetro la
//cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

package main

import "fmt"

const A, B, C = "A", "B", "C"

func main() {
	mensaje := "Su sueldo correspondiente a su categoria es: "
	fmt.Print(mensaje, salario(5389, A), "\n")
	fmt.Print(mensaje, salario(1325, B), "\n")
	fmt.Print(mensaje, salario(3000, C), "\n")

}

func salario(minutos_mensuales float64, categoria string) (salario float64) {

	switch categoria {
	case A:
		salario = ((minutos_mensuales / 60) * 3000) * 1.5
		return salario
	case B:
		salario = ((minutos_mensuales / 60) * 1500) * 1.2
		return salario

	case C:
		salario = ((minutos_mensuales / 60) * 1000) * 1.5
		return salario
	}
	return 0
}
