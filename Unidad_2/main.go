package main

import (
	"fmt"
)


func main(){
	fmt.Println("Bienvenido al sistema")
	var name string
	fmt.Print("Ingrese su nombre: ")
	fmt.Scan(&name)
	fmt.Printf("mi nombre es: %s\n",name)
}