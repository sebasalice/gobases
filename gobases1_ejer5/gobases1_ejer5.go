package main

import "fmt"

func main (){
	var palabra string
	var longitud int
	palabra = "Palabrota"
	longitud = (len(palabra))
	fmt.Println("La longitud de la palabra es: ", longitud)
	for i := 0; i < longitud; i++ {
		fmt.Println(string(palabra[i]))
	}
	
}