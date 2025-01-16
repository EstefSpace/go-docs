package main

import (
	"fmt"// para mostrar datos en la consola
	"strconv"// para convertir numero a texto y viceversa
	"os"// para recibir argumentos cuando pones go run funciones.go <arg1> <arg2>
)

/* Estructura de una funcion

Si en la funcion no hay return es innecesario poner tipoDeResultado (string, int, float32, etc)

func nombreFuncion(parametros) tipoDeResultado {
	codigo...
}
*/

func main() {

	resultado := sumar(os.Args[1], os.Args[2])// llama a la funcion sumar para realizar una serie de instrucciones
    fmt.Println("Resultado:", resultado)// muestra el resultado    
}

func sumar(numero1 string, numero2 string) int {
	// se recibe dos parametros de tipo string para luego convertirlos a numeros enteros y hacer la suma
	// el tipo de dato que se devolvera sera un int

	int1, _ := strconv.Atoi(numero1)
    int2, _ := strconv.Atoi(numero2)
    return int1 + int2// return es para devolver algo, en este caso el resultado de la suma
}