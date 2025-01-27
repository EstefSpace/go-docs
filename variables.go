package main 

import "fmt"// para mostrar datos en la consola

func main() {
	/* Variables */

	//las variables siempre deben usarse, si no habra un error al compilar

	var ocupacion string // se declara una nueva variable llamada primerNombre y se indica que es un string (Cadena de Texto)
	ocupacion = "programador"// a esto se le llama "inicializar una variable" es decir darle un valor
	fmt.Println(ocupacion)// lo mostramos por consola

	// declarando dos variables en una misma linea (es posible si las dos variables son del mismo tipo)
	var primerNombre, segundoNombre string
	primerNombre = "Estefan"
	segundoNombre = "Emilio"
	fmt.Println("Me llamo " + primerNombre + " " + segundoNombre)
	// declarando una variable de tipo int (entero)
	var edad int
	edad = 14
	fmt.Println(edad)
	// tambien se puede declarar e inicializar varias variables de diferentes tipos asi
	var (
		miTexto string = "¡Hola mundo!"
		miNumero int = 14
	)

	fmt.Println(miTexto, miNumero)
	// NOTA: No es necesario poner el tipo de variable porque golang ya sabe que tipo de dato es

	// declara variables en una sola linea
	var (
    	unTexto, otroTexto, unNumero = "estef", "dev", 14
	)

	fmt.Println(unTexto + otroTexto)
	fmt.Println(unNumero)


	/* Manera habitual de declarar e inicializar variables */
	texto := "¡Como te va!"
	edadNueva := 15

	fmt.Println(texto, edadNueva)


	/* Constantes */

	const numeroIncreible = 200// un numero y no hace falta indicar el tipo!


	// igual que con var se puede hacer varias constantes asi
	const (
		estadoDiscord = "Jugando a Roblox"
		tiempo = "20 horas"
	)

	
}