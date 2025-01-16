package main

import (
	"fmt"// para mostrar datos en la consola
	"strconv"// para convertir numero a texto y viceversa
)

func main() {
	var numero int = 52// Numero entero
	var decimal float32 = 52.5// Numero flotante
	var funcionando bool = true// Booleanos (representan si algo es verdadero o falso, true o false)
	var cadena string = "Esto es una cadena de texto"// Cadena de texto

	// si no se inicializa una variable, golang le proporciona un valor preterminado
	fmt.Println(numero, decimal, funcionando, cadena)

	// seguro tendras una situacion donde tienes que convertir de numero a cadena de texto y viceversa

	i, _ := strconv.Atoi("-42")// convierte este string a un int
	s := strconv.Itoa(-42)// convierte este int a un string

	// go run tipos_de_datos.go para verificar esto
	fmt.Println(i, s)
}