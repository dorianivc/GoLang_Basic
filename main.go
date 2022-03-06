package main

import (
	"fmt"
)

func main() {
	//Defer -- Ejecuta la ultima funcion antes de que el programa muera
	//Funciona para cerrar archivos o conexiones a bases de datos

	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y Break

	for i:=0; i<10; i++{
		fmt.Println(i)

		//Continue
		if i==2{
			fmt.Println("Es 2")
			continue
		}

		if i==7{
			fmt.Println("Break, Bye")
			break
		}
	}
}