package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Martínez": "Edad: \nNombre completo: \nApodo: \nFecha de nacimiento: \nEstatura: \nPeso: ",
		"Alvarez":  "Edad: 23 \nNombre completo: Julián Álvarez \nApodo: La Araña \nFecha de nacimiento: 31 de enero de 2000 \nEstatura:  1.7 m \nPeso: 71 kg",
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
