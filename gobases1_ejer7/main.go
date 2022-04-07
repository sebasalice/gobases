package main

import "fmt"

func main (){
	mes_actual := 5
	switch mes_actual {
	case 1:
		fmt.Println("1, Enero")
	case 2:
		fmt.Println("2, Febrero")
	case 3:
		fmt.Println("3, Marzo")
	case 4:
		fmt.Println("4, Abril")
	case 5:
		fmt.Println("5, Mayo")
	case 6:
		fmt.Println("6, Junio")
	case 7:
		fmt.Println("7, Julio")
	case 8:
		fmt.Println("8, Agosto")
	case 9:
		fmt.Println("9, Septiembre")
	case 10:
		fmt.Println("10, Octubre")
	case 11:
		fmt.Println("11, Noviembre")
	case 12:
		fmt.Println("12, Diciembre")
	}
}