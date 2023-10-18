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
		"Messi": "Edad:36 \nNombre completo:Lionel Andres Messi Cuccittini \nApodo:La pulga \nFecha de nacimiento:24 de junio de 1987  \nEstatura:1.7 \nPeso:65 ",
		"Basile": "Edad:79 \nNombre completo:Alfio Ruben Basile \nApodo:Coco Basile \nFecha de nacimiento: 1 de noviembre de 1943 \nEstatura:1.92 \nPeso:120 ",
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
