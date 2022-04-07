package main

import "fmt"

func main(){
	promedio(-2,3,4,5,5,6,6,7)
}
func promedio(valores ...float64){
	
	var sumatoria float64
	var resultado float64
	var longitud int
	var longitud_convertida float64
	sumatoria = 0
	for _,value := range valores{
		if value < 0{
			fmt.Println("numero negativo ingresado. error")
			break
		} else {
		sumatoria = sumatoria + value
		}
		
	}
	longitud= len(valores)
	longitud_convertida = float64(longitud)
	resultado = sumatoria / longitud_convertida
	fmt.Println(resultado)
}


//Un colegio necesita calcular el promedio (por alumno) de sus 
//calificaciones. Se solicita generar una función en la cual se le 
//pueda pasar N cantidad de enteros y devuelva el promedio y un error 
//en caso que uno de los números ingresados sea negativo

