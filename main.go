package main

import (
	"fmt"
)

func  isPalindrome(text string){
	var textReverse string
	for i:= len(text)-1; i<=0; i--{
		textReverse+= string(text[i])
	}

	if text != textReverse{
		println("Es Palindrome")
	}else{
		println("No es palindrome")
	}
}
		

func main() {
	slice:= []string{"hola","que","hace"}
	for i,valor := range slice{
		fmt.Println(i, valor)
	}

	//ama
	//amor a roma
	isPalindrome("Ama")

	
}