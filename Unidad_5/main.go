package main

import (
	"fmt"
	
)



// Persona struct
type Persona struct {
	nombre string
	edad,id int
}

// SetterNombre method
func (p *Persona) SetterNombre(nombre string) {
	if nombre == "" {
		fmt.Println("Nombre vacio")
	}else{
		p.nombre = nombre
	}
}

// SetterEdad method
func (p *Persona) SetterEdad(edad int) {
	if edad > 0 {
		p.edad = edad
	}else {
		fmt.Println("La edad debe ser mayor 0")
	}
}

// SetterID method
func (p *Persona) SetterID(id int) {
	if id > 0 {
		p.id = id
	}else{
		fmt.Println("El id debe ser mayor que 0")
	}

}

// GetterNombre method
func (p Persona) GetterNombre() string {
	return p.nombre
}
// GetterEdad method
func (p Persona) GetterEdad() int {
	return p.edad
}
// GetterID mehtod
func (p Persona) GetterID() int {
	return p.id
}
// Mostrar method
func (p Persona) Mostrar() {
	fmt.Printf("Nombre: %s Edad: %d Id: %d\n",p.GetterNombre(),p.GetterEdad(),p.GetterID())
}

// NewPersona constructor
func NewPersona(args ...Persona) *Persona {
	persona := new(Persona)
	persona.SetterNombre(args[0].nombre)
	persona.SetterEdad(args[0].edad)
	persona.SetterID(args[0].id)
	return persona
}

func main(){
	user := NewPersona(Persona{"Ann",30,645646123454})
	user.Mostrar()
	

}