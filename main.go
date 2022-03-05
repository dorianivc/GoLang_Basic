package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)
	
	//declaracion de variables enteras
	base :=12
	var altura int = 14
	var area int

	fmt.Println(base)
	fmt.Println(altura)
	fmt.Println(area)
	//Zero Values
	var a int
	var b float64
	var c string 
	var d bool
	fmt.Println(a,b,c,d)

	//Area cuadrado
	const baseCuadrado= 10
	areaCuadrado := baseCuadrado*baseCuadrado
	fmt.Println("Area Cuadrado: ", areaCuadrado)

	x:=10
	y:=50

	// suma
	result:= x+y
	fmt.Println("Suma: ", result)

	//Resta
	result =  y-x
	fmt.Println("Resta: ", result)
	
	//Multiplicacion
	result =  x*y
	fmt.Println("Multiplicacion: ", result)

	//Divison
	result =  x/y
	fmt.Println("Division: ", result)

	//Modulo
	result =  x%y
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	helloMessage := "Hello"
	worldMessage := "world"
	third:= helloMessage+" "+worldMessage

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)
	
	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
	
	fmt.Println(third)

}

