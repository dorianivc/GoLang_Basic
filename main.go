package main
import (
	"fmt"
	"mypackage"
)

func main() {

var myCar mypackage.CarPublic
myCar.Brand="Ferrari"
myCar.Year=2020
fmt.Println(myCar)

mypackage.PrintMessage("Hola Platzi")
//Error por ser privado
mypackage.printMessage1("Hola Platzi")
}