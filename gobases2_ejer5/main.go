//Un refugio de animales necesita calcular cuánto alimento
//debe comprar para las mascotas. Por el momento solo tienen
//tarántulas, hamsters, perros, y gatos, pero se espera que
//puedan haber muchos más animales que refugiar.
//
//perro necesitan 10 kg de alimento
//gato 5 kg
//Hamster 250 gramos.
//Tarántula 150 gramos./
//
//Se solicita:
//Implementar una función Animal que reciba como parámetro
//un valor de tipo texto con el animal especificado y que
//retorne una función y un mensaje (en caso que no exista el animal)
//Una función para cada animal que calcule la cantidad de
//alimento en base a la cantidad del tipo de animal especificado.
package main

import "fmt"

const (
	tarantulas, hamsters, perros, gatos = "tarantula", "hamsters", "perros", "gatos"
)

func main() {
	fmt.Println(Animal(tarantulas, 3))
}
func Animal(valor string, amount int) (msg string) {
	mensajes := "el animal no existe"
	switch valor {
	case tarantulas:
		fmt.Println(alimento_tarantulas(amount))
	case hamsters:
		fmt.Println(alimento_hamsters(amount))
	case perros:
		fmt.Println(alimento_perros(amount))
	case gatos:
		fmt.Println(alimento_gatos(amount))

	default:
		fmt.Println(mensajes)
	}
	return
}
func alimento_tarantulas(amount int) (cantidades int, mensajes string) {
	mensaje := "gr, es la cantidad a comprar."
	cantidad := amount * 150
	return cantidad, mensaje
}
func alimento_hamsters(amount int) (cantidades int, mensajes string) {
	mensaje := "gr, es la cantidad a comprar."
	cantidad := amount * 250
	return cantidad, mensaje
}
func alimento_perros(amount int) (cantidades int, mensajes string) {
	mensaje := "kg, es la cantidad a comprar."
	cantidad := amount * 10
	return cantidad, mensaje
}
func alimento_gatos(amount int) (cantidades int, mensajes string) {
	mensaje := "kg, es la cantidad a comprar."
	cantidad := amount * 5
	return cantidad, mensaje
}
