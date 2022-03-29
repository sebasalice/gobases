package main

import "fmt"

func main(){
	//var 1nombre string  --> mal. el nombre de la variable no puede empezar con numeros
	var apellido string 
	// var int edad --> mal definido
	// 1apellido := 6 --> mal. empieza variable con numero
	// var licencia_de_conducir = true --> mal
	//var estatura de la persona int --> mal no puede tener espacios
	 cantidadDeHijos := 2
 
	
	var nombre string
	var edad int
	index_apellido := 6 
	var licencia_de_conducir bool
	var estatura_de_la_persona float64

	licencia_de_conducir = true
	apellido = "salice"
	nombre = "sebastian"
	estatura_de_la_persona = 1.75

	fmt.Println(apellido,cantidadDeHijos,nombre,edad,index_apellido,estatura_de_la_persona)
	fmt.Printf("%t",licencia_de_conducir)
}

