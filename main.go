package main

import (
	"fmt"
	"os"
)

func main() {
	diccionario := map[string]string{
		"Reus":         "Edad: 34\nNombre completo: Marco Reus \nApodo:Woody \nFecha de nacimiento:31 de mayo de 1989   \nEstatura:1.8 \nPeso:71 kg ",
		"Foden":        "Edad: 23\nNombre completo: Philip Walter Foden \nApodo: foden \nFecha de nacimiento: 28 de mayo de 2000 \nEstatura:1.71 \nPeso:70 kg ",
		"Aimar":        "Edad: 43\nNombre completo: Pablo César Aimar Giordano \nApodo: El payaso Aimar \nFecha de nacimiento:  3 de noviembre de 1979 \nEstatura:  168 \nPeso: 69 kg",
		"Tagliafico":   "Edad: 31\nNombre completo: Nicolás Alejandro \nApellido:Tagliafico\nApodo: taglia\nFecha de nacimiento: 31 de agosto de 1992\nEstatura: 172 cm\nPeso: 72 kg",
		"Francescoli":  "Edad: 61\nNombre completo: Enzo Francescoli \n Apodo: El Principe \n Fecha de nacimiento: 12 de noviembre de 1961 \n Estatura: 1.81 \n Peso: 73kg",
		"Messi":        "Edad: 36\nNombre completo: Lionel Andres Messi Cuccittini \nApodo:La Pulga \nFecha de nacimiento:24 de junio de 1987  \nEstatura:1.7 \nPeso:65 ",
		"Mac Allister": "Edad: 24\nNombre completo: Alexis Mac Allister\nApodo: Macca\nFecha de nacimiento: 24 de diciembre de 1998\nEstatura: 1,76 m\nPeso: 72kg",
		"Álvarez":      "Edad: 26\nNombre completo: Julian\nApellido: Álvarez\nApodo: araña\nFecha de nacimiento:  30 de enero de 2000\nEstatura:  170 cm\nPeso: 71 kg",
		"Martínez":     "Edad: 31\nNombre completo: Damián Emiliano \nApellido: Martínez\nApodo: Dibu\nFecha de nacimiento: 2 de septiembre de 1992\nEstatura: 195 cm\nPeso: 88 kg",
		"Mbappé":       "Edad: 24\nNombre completo: Kylian \nMbappé Lottin \nApodo: Mbappé \nFecha de nacimiento: 20 de diciembre de 1998 \nEstatura: 1,78 m \nPeso: 75 kg",
		"Di María":     "Edad: 35\nNombre completo: Angel Fabian \nApellido: Di Maria\nApodo: Fideo\nFecha de nacimiento: 14 de febrero de 1988\nEstatura: 178cm\nPeso: 75kg",
		"Ronaldo":      "Edad: 36\nNombre completo: Cristiano Ronaldo \nApellido: Cristiano Ronaldo\nApodo: CR7\nFecha de nacimiento: 13 de febrero de 1987\nEstatura: 179cm\nPeso: 78kg",
		"Modric":       "Edad: 38\nNombre completo: Luka Modric \nApellido:luka modric\nApodo:modric\nfecha de nacimiento:9 de septiembre de 1985\nEstatura:1,72\nPeso: 66kg",
		"Ibarra":       "Edad: 49\nNombre completo: Hugo Benjamín Ibarra \nApodo: El Negro \nFecha de Nacimiento: 1 de abril de 1974 \nEstatura: 1,71 m \nPeso: 74 kg",
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
