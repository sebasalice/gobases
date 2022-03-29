package main

import "fmt"

func main (){
	salario(5389,"A")
}

func salario(minutos_mensuales int, categoria string)(salario_a,salario_b,salario_c float64){
	var salario_a,salario_b,salario_c float64
	var categoria string
	fmt.Println(minutos_mensuales)
	if categoria = "A" {
		salario_a = ((minutos_mensuales / 60) * 3000) * 1.5
		return salario_a
	}
	if categoria = "B" {
		salario_b = ((minutos_mensuales / 60) * 1500) * 1.2
		return salario_b
	}
	if categoria = "C" {
		salario_c = ((minutos_mensuales / 60) * 1000) 
		return salario_c
	}
}
//Una empresa marinera necesita calcular el salario de sus empleados 
//basándose en la cantidad de horas trabajadas por mes y la categoría.

//Si es categoría C, su salario es de $1.000 por hora
//Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

//Se solicita generar una función que reciba por parámetro la 
//cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
