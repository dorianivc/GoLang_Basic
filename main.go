package main

import (
	"fmt"
)

type pc struct{
	ram int
	disk int
	brand string

}

func (myPC pc)ping(){
	fmt.Println(myPC.brand, "Pong")

}

func(myPC *pc) duplicateRam(){
	myPC.ram= myPC.ram*2

}

func main() {

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b=100
	fmt.Println(a)
	myPC:= pc{ram:32, disk: 1000, brand: "Lenovo"}
	fmt.Println(myPC)
	myPC.ping()
	myPC.duplicateRam()
	fmt.Println(myPC)


}
