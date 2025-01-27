package main

import "fmt"

func recibirEstadoDeLaPizza(cliente string) string {
	var estado string
	switch(cliente) {
	case "Pepe", "pepe": // puede que haya escrito mal el nombre del cliente
		estado = "no-lista"
	case "Juan", "juan":
		estado = "lista"
	default:
		estado = "no-es-cliente"
	}

	return estado
}


func dameUnNumero(numero int) int {
	return numero * 2
}

func main() {
	// vamos a aprender los if, else, switch, etc

	/* CONDICIONALES IF/ELSE

	se puede agregar los else if que se necesiten

	if -> Si pasa esto ejecuta esto
	else if -> Si pasa algo distinto a esto ejecuta esto
	else -> Si pasa lo contrario ejecuta esto

	Estructura:

	if condicion {
		codigo..
	} else if condicion {
		codigo..
	} else {
		codigo..
	}
	*/

	num1 := dameUnNumero(2)

	if num1 < 0 {// si num1 es menor que 0
		fmt.Println(num1, "Es un numero negativo")
	} else if num1 < 10 {// pero tambien si num1 es menor que 10
		fmt.Println(num1, "Solo tiene un digito")
	} else {// si ninguna de las dos anteriores sucede es porque num1 tiene mas de 1 digito
		fmt.Println(num1, "Tiene mas de 1 digito")
	}


	// tambien se pueden declarar variables en un bloque if y solo se pueden usar ahi

	if num2 := dameUnNumero(4); num2 < 0 {// si num2 es menor que 0
		fmt.Println(num2, "Es numero negativo")// solo se puede usar la variable num2 en este bloque if
	} else if num2 < 10 {// pero tambien si num2 es menor que 10
		fmt.Println(num2, "Solo tiene un digito")
	} else {// si ninguna de las dos anteriores sucede es porque num2 tiene mas de 1 digito
		fmt.Println(num2, "Tiene mas de 1 digito")
	}

	// fmt.Println(num2) si se quiere usar la variable num2 fuera del bloque if dara un error


	/* INSTRUCCIONES SWITCH
	En go se usan para no repetir tanto los if

	estructura:

	switch(variable) {
	case 1:
		codigo..
	case 2: 
		codigo..
	default:
		codigo..
	}
	 */

	cliente := "Emilio"// Juan, Pepe, Emilio
	pizzaEstado := recibirEstadoDeLaPizza(cliente)// Necesito ver en que estado esta la pizza ¡ya se ha tardado mucho!

	switch pizzaEstado {
	case "lista":
		fmt.Println("¡La pizza esta lista!")
	case "no-lista":
		fmt.Println("¡Aun no esta lista la pizza!")
	case "no-es-cliente":
		fmt.Println(cliente, "¡No es un cliente de nuestra pizzeria!")
	default: // default se usa para cuando no se cumple ninguno de los casos anteriores
		fmt.Println("¡No sabemos en que estado esta la pizza!")
	}

	}