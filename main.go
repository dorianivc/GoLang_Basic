package main

import (
	"fmt"
)

type figuras2D interface{
	area() float64
}

type cuadradro struct{
	base float64
}
type rectangulo struct{
	base float64
	altura float64
}

func (c cuadradro) area() float64{
	return c.base* c.base
}

func (r rectangulo) area() float64{
	return r.base * r.altura
}

func calculate(f figuras2D){
	fmt.Println("Area: ", f.area())
}
func main() {
	myCuadrado:= cuadradro{base: 4}
	myRectamgulo:= rectangulo{base: 4, altura: 7}
	
	calculate(myCuadrado)
	calculate(myRectamgulo)

	//Lista de interfaces
	myInterface:= []interface{/*Interfaz Vacia*/}{"Hola", 12, 4.77}
	fmt.Println(myInterface...)

}
