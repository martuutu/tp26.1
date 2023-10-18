package main

import (
	"fmt"
	"os"
)

func main() {
	diccionario := map[string]string{
		"Messi": "Edad:36 \nNombre completo:Lionel Andres Messi Cuccittini \nApodo:La pulga \nFecha de nacimiento:24 de junio de 1987  \nEstatura:1.7 \nPeso:65 ",
		"Basile": "Edad:79 \nNombre completo:Alfio Ruben Basile \nApodo:Coco Basile \nFecha de nacimiento: 1 de noviembre de 1943 \nEstatura:1.92 \nPeso:120 ",
		"Julian":     "Edad: 26\nNombre completo: Julian\nApellido: Álvarez\nApodo: araña\nFecha de nacimiento:  30 de enero de 2000\nEstatura:  170 cm\nPeso: 71 kg",
		"Tagliafico": "Edad: 31\nNombre completo: Nicolás Alejandro \nApellido: Tagliafico\nApodo: taglia\nFecha de nacimiento: 31 de agosto de 1992\nEstatura: 172 cm\nPeso: 72 kg",
	}

	if len(os.Args) != 2 {
		fmt.Println("Uso: tp26.exe apellido_jugador")
		return
	}

	apellido := os.Args[1]
	switch descripcion := diccionario[apellido]; descripcion {
	case "":
		fmt.Printf("El jugador '%s' no se encuentra.\n", apellido)
	default:
		fmt.Printf("%s:\n%s\n", apellido, descripcion)
	}
}
