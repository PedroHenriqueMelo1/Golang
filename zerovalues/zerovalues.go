package main 

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("Inteiro: %v\n", i) // Saída 0
	fmt.Printf("Float: %v\n", f) // Saída 0
	fmt.Printf("Bool: %v\n", b) // false
	fmt.Printf("String: %q\n", s) // string vazia
}

// Conceito de zerovalues são valores que são atribuidos ao declarar uma varíavel porém não inicializar com um valor explicito.