package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Martínez": "Edad:27 \nNombre completo: Lautaro Martínez\nApodo: El toro\nFecha de nacimiento: 14/07/1991\nEstatura: 174cm\nPeso: 54kg ",
		//TODO agregar un jugador de la seleccion
                "Julian": "Edad: 23\nNombre completo: Julian Alvarez\nApodo: La Araña\nfecha de nacimiento: 31/01/2000\nEstatura: 170cm\nPeso: 71 ",
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
