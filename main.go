package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Messi": "Edad:35 \nNombre completo:leonel andres messi cuccittini \nApodo:pulga \nFecha de nacimiento:24/06/1987 \nEstatura:171 \nPeso:72 ",
		"gauto": "Edad:18 \nNombre completo:juan gauto \nApodo:juanchi \nFecha de nacimiento:02/06/2004 \nEstatura:170 \nPeso:70 ",
		//TODO agregar un jugador de la seleccion
	}

	// Verificar si se proporciona un argumento (la palabra a buscar).
	if len(os.Args) != 2 {
		fmt.Println("Uso: tp26.exe apellido_jugador")
		return
	}

	// Obtener el apellido proporcionado como argumento.
	apellido := os.Args[1]

	// Utilizar una declaración switch para buscar y mostrar el descripcion.
	switch descripcion := diccionario[apellido]; descripcion {
	case "":
		fmt.Printf("El jugador '%s' no se encuentra.\n", apellido)
	default:
		fmt.Printf("%s:\n%s\n", apellido, descripcion)
	}
}
