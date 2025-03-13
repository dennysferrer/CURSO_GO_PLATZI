package main

import "fmt"

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

}
