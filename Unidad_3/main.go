package main

import (
	"fmt"
)
const precio = 1100000 //$
var (
	total int 
	descuento int
)
func main() {
	fmt.Println("Bienvenido a Easyred")
	fmt.Printf("Ingrese cantidad: ")
	fmt.Scanf("%d",&total)
	if total < 5 {
		precioTotal := total * precio
		porciento := (precioTotal * 10)/100
		descuento = precioTotal - porciento
	}else if total >= 5 && total < 10 {
		precioTotal := total * precio
		porciento := (precioTotal * 20)/100
		descuento = precioTotal - porciento
	}else {
		precioTotal := total * precio
		porciento := (precioTotal * 40)/100
		descuento = precioTotal - porciento
	}

	fmt.Printf("Total a pagar %d\n",descuento)

}