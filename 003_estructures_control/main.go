package main

import "fmt"

func main() {
	var taula = [5]int {1,2,3,4,5}
	
	//While:
	fmt.Println("while:")
	i := 0
	for i < 5 {
		fmt.Println(taula[i])
		i++
	}

	//For
	fmt.Println("for:")
	for i:=4;i>=0; i--{
		fmt.Println(taula[i])
	}
	
	fmt.Println("for:")
	for index,valor:=range taula{
		fmt.Printf("Index %d conté el valor %d\n", index, valor)
	}

	/*
	fmt.Println("while(1): El programa es quedarà penjat aquí")
	//while(1)
	for{} 
	*/
}
	// (Paràmetres d'entrada) (Paràmetres de sortida)
	func suma_mult1(numero1, numero2 int) (suma, multiplicacio int) {
		suma = numero1 + numero2
		multiplicacio = numero1 * numero2
		return
	}

	
	func suma_mult2(numero1, numero2 int) (int, int) {
		return numero1 + numero2, numero1 * numero2
	}
	

