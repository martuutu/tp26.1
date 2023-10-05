package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Lautaro": "Edad: 26\nNombre completo: Lautaro Javier\nApellido: Martinez\nApodo: toro\nFecha de nacimiento:  22 de agosto de 1997\nEstatura:  174 cm\nPeso: 72 kg",
		"Messi":   "Edad: 34\nNombre completo: Lionel Andres Messi \nApellido: Cuccittini\nApodo: Leo\nFecha de nacimiento: 24 de junio de 1987\nEstatura: 170 cm\nPeso: 72 kg",
	}

	// Verificar si se proporciona un argumento (la palabra a buscar).
	if len(os.Args) != 2 {
		fmt.Println("Uso: tp26.exe apellido_jugador")
		return
	}

	// Obtener el apellido proporcionado como argumento.
	apellido := os.Args[1]

	// Utilizar una declaración switch para buscar y mostrar la descripción.
	switch descripcion := diccionario[apellido]; descripcion {
	case "":
		fmt.Printf("El jugador '%s' no se encuentra.\n", apellido)
	default:
		fmt.Printf("%s:\n%s\n", apellido, descripcion)
	}
}
