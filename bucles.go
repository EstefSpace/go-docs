package main

import "fmt"

func main() {
	/*
	Estructura de un bluce for

	Se conforma de tres componentes que se separan con ;

	- Una instruccion inicial que se ejecuta antes de la primera interación (opcional)
	- Una expresion de condicion que se evalua antes de cada iteracion. El bucle se detien cuando esta condicion es false
	- Una instruccion posterior que se ejecuta al final de cada iteracion (opcional)


	 */

	 //ejemplo simple

	 numero := 0
	 for i := 1; i <= 100; i++ {
	 	numero += i
	 }

	 fmt.Println(numero)

}