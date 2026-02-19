package main

import "fmt"

// "Clase" base
type Animal struct {
    Nombre string
}

func (a Animal) Respirar() {
    fmt.Printf("%s está respirando...\n", a.Nombre)
}

// "Subclase" que embebe a Animal
type Perro struct {
    Animal   // <--- Composición: Perro "tiene un" Animal
    Raza   string
}

func (p Perro) Ladrar() {
    fmt.Printf("%s dice: ¡Guau!\n", p.Nombre)
}

func main() {
    miPerro := Perro{
        Animal: Animal{Nombre: "Toby"},
        Raza:   "Labrador",
    }

    // Puedes acceder a métodos de Animal directamente
    miPerro.Respirar() 
    miPerro.Ladrar()
}