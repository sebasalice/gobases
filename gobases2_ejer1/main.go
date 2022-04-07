package main

import "fmt"

func main(){
	descuento()
}
func descuento(){
	var neto float64
	var bruto float64

	bruto = 75000
	
	if bruto > 50000{
		if bruto > 150000{
			neto = bruto * 0.73
			fmt.Println("cobrara ", neto)
		} else{
			neto = bruto * 0.83
			fmt.Println("cobrara ", neto)
		}
	}
		
}


//Una empresa de chocolates necesita calcular el impuesto de sus empleados 
//al momento de depositar el sueldo, para cumplir el objetivo es necesario crear 
//una función que devuelva el impuesto de un salario. 
//Teniendo en cuenta que si la persona gana más de $50.000 
//se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
