package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Martínez": "Edad: 26  \nNombre completo: Lautaro Martinez \nApodo: El Toro \nFecha de nacimiento: 22 de agosto de 1997 \nEstatura: 1.74 \nPeso:72 ",
		//TODO agregar un jugador de la seleccion
		"María": "Edad: 35  \nNombre completo: Ángel Di María \nApodo: Fideo  \nFecha de nacimiento: 14 de febrero de 1988 \nEstatura: 1.78 \nPeso:75 ",
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
