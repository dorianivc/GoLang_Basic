package main

import "fmt"

func main() {
	// For Condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Nuevo Ciclo")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++

	}

	//For forever
	counterForever:=0
	for{
		fmt.Println(counterForever)
		counterForever++
	}
}
