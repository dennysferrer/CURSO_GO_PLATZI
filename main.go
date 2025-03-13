package main

import "fmt"

func normalFunction(mensaje string) {
	fmt.Println(mensaje)
}

func suma(a int, b int) int {
	c := a + b
	return c
}

func ciclos() {
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}

func main() {
	const pi float64 = 3.141592653589793238
	const pi2 = 3.1415
	fmt.Println(pi, pi2)

	//Declaración de variables
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Calcular el área de un cuadrado

	const baseCuadrado = 10
	var areaCuadrado = baseCuadrado * baseCuadrado
	fmt.Println("El área del cuadrado es: ", areaCuadrado)

	//Operadores Aritméticos

	x := 10
	y := 50

	//Suma
	resultSuma := x + y
	fmt.Println("Suma: ", resultSuma)

	//Resta
	resultResta := y - x
	fmt.Println("Resta: ", resultResta)

	//Multiplicación
	resultMultiplicacion := x * y
	fmt.Println("Multiplicación: ", resultMultiplicacion)

	//División
	resulDivision := y / x
	fmt.Println("División: ", resulDivision)

	//Módulo
	resultModulo := y % x
	fmt.Println("Módulo: ", resultModulo)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Uso de la funcion fmt

	nombre := "Dennys Ferrer"
	edad := 37

	fmt.Printf("Mi nombre es %s y tengo %d años\n", nombre, edad)

	normalFunction("Hola soy un mensaje de parametro del a funcion normal")

	//Uso de la funcion suma
	resultadoFunctionSuma := suma(10, 20)
	fmt.Println("El resultado de la suma es: ", resultadoFunctionSuma)

}
