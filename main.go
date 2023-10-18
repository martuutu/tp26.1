package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Marco": "Edad:34 \nNombre completo:Marco Reus \nApodo:Woody \nFecha de nacimiento:31 de mayo de 1989   \nEstatura:1.8 \nPeso:71 kg ",
		"Philip": "Edad:23 \nNombre completo:Philip Walter Foden \nApodo: foden \nFecha de nacimiento: 28 de mayo de 2000 \nEstatura:1.71 \nPeso:70 kg ",
		"Pablo":     "Edad: 43\nNombre completo: Pablo César Aimar Giordano \nApodo: El payaso Aimar \nFecha de nacimiento:  3 de noviembre de 1979 \nEstatura:  168 \nPeso: 69 kg",
		"": "Edad: 31\nNombre completo: Nicolás Alejandro \nApellido: Tagliafico\nApodo: taglia\nFecha de nacimiento: 31 de agosto de 1992\nEstatura: 172 cm\nPeso: 72 kg",
		""
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
