package main

import "fmt"

func main (){
	var edad int
	var sueldo float64
	var antiguedad int
	var empleado bool
	var mensaje_rechazado,mensaje_aprobado,impuestos_si,impuestos_no string

	edad = 23
	sueldo = 100000
	antiguedad = 2
	empleado = true
	mensaje_aprobado = "prestamo aprobado"
	mensaje_rechazado = "prestamo rechazado"
	impuestos_si = "pagara impuestos"
	impuestos_no = "no pagara impuestos"

	if edad > 22 {
			if antiguedad > 1 {
				if empleado {
					fmt.Println(mensaje_aprobado)
					if sueldo > 100000{
					fmt.Println(impuestos_no)
					}else {fmt.Println(impuestos_si)}
				}
			}
		} else{
			fmt.Println(mensaje_rechazado)
		}
	
	
}




