package main

import "fmt"

func main() {
    // 1. Declaración estándar con 'var' (útil para valores por defecto)
    // Si no asignas valor, Go asigna el "Zero Value" (0 para números, "" para strings)
    var edad int = 25
    var nombre string = "Sara"

    // 2. Inferencia de tipo (Go deduce que es un float64)
    var precio = 19.99

    // 3. Declaración corta (Walrus Operator ':=')
    // Es la más común dentro de funciones. Declara y asigna en un solo paso.
    ciudad := "Madrid" //ciutat queda tipat com a string
    puntuacion := 100

    // 4. Declaración de múltiples variables
    var (
        isProgrammer = true
        lenguaje     = "Go"
    )

    // Imprimimos los resultados
    fmt.Println("--- Perfil de Usuario ---")
    fmt.Printf("Nombre: %s, Edad: %d\n", nombre, edad)
    fmt.Printf("Ciudad: %s, Puntuación: %d\n", ciudad, puntuacion)
    fmt.Printf("¿Es programador?: %t en %s\n", isProgrammer, lenguaje)
    fmt.Println("Precio del curso:", precio)
}