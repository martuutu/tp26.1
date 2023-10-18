package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Messi": "Edad:36 \nNombre completo:Lionel Andres Messi Cuccittini \nApodo:La pulga \nFecha de nacimiento:24 de junio de 1987  \nEstatura:1.7 \nPeso:65 ",
		"MacAllister": "Edad: 24\nNombre completo:Alexis Mac Allister\nApodo: Macca\nFecha de nacimiento: 24 de diciembre de 1998\nEstatura: 1,76 m\nPeso: 72kg",
		"Tagliafico": "Edad: 31\nNombre completo: Nicolás Alejandro Tagliafico\nApodo: Taglia\nFecha de nacimiento: 31 de agosto de 1992 \nEstatura: 1,72 m\nPeso: 65 kg",
		"Julian":     "Edad: 26\nNombre completo: Julian\nApellido: Álvarez\nApodo: araña\nFecha de nacimiento:  30 de enero de 2000\nEstatura:  170 cm\nPeso: 71 kg",
		"Tagliafico": "Edad: 31\nNombre completo: Nicolás Alejandro \nApellido: Tagliafico\nApodo: taglia\nFecha de nacimiento: 31 de agosto de 1992\nEstatura: 172 cm\nPeso: 72 kg",
		"Di Maria": "Edad: 35\nNombre completo: Angel Fabian \nApellido: Di Maria\nApodo: fideo\nFecha de nacimiento: 14 de febrero de 1988\nEstatura: 178cm\nPeso: 75kg",
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
