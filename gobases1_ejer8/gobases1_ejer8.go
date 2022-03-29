package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])
	fmt.Println("Los empleados mayores de 21 son: ")
	for key,element := range employees {
		if element > 21{
			fmt.Println(key, element)
				}
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}


//Saber cuántos de sus empleados son mayores de 21 años.
//Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
//Eliminar a Pedro del mapa.
