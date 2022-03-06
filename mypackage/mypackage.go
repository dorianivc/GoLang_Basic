package mypackage

import "fmt"

//CarPublic car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type CarPrivate struct {
	brand string
	year  int
}

//PrintMessage Imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

//PrintMessage Imprimir un mensaje -- Privado
func printMessage1(text string) {
	fmt.Println(text)
}
