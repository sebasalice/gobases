package main

import "fmt"

var temperatura,presionAtmosferica float64
var humedad int

func main(){
	temperatura,presionAtmosferica,humedad= 25.5,1007.2,74
	fmt.Println("Temperatura: ", temperatura,"\n","Presion atmosferica: ",presionAtmosferica,"\n","Humedad :",humedad)

}