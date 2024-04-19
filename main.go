package main

import (
	"fmt"
	"os"
	"time"

)

var (
	adminUsername = "admin"
	adminPassword = "root"
	//memberUsername = "seeker"
	//memberPassword = "seekr"

	adminActions = []string{"Crear archivo", "Ver cantidad de teclas pulsadas"}

// actionCounter = make(map[string]int)
)

func main() {
	mainMenu()
}

func mainMenu() {
	for {
		fmt.Println("Bienvenido a Labora")
		fmt.Println("1. Iniciar sesión como Administrador")
		fmt.Println("2. Iniciar sesión como Miembro")
		fmt.Println("3. Salir")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			adminLogin()
		case 2:
			//memberLogin()
		case 3:
			fmt.Println("¡Hasta luego!")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func adminLogin() {
	fmt.Println("Ingrese las credenciales de Administrador")
	fmt.Print("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == adminUsername && password == adminPassword {
		fmt.Println("Inicio de sesión exitoso como Administrador")
		adminMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo.")
	}
}

func adminMenu() {
	for {
		fmt.Println("Menú de Administrador")
		for i, action := range adminActions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Println("0. Cerrar sesión")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createFile()
		case 2:
			//showFile()
		case 0:
			fmt.Println("Cerrando sesión de Administrador")
			mainMenu()
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func createFile() {
	fmt.Println("Escriba el contenido del archivo")
	var content string
	fmt.Scanln(&content)
	generateTextFile(content)
}

func generateTextFile(content string) {
	fileName := fmt.Sprintf("mensajes_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir los mensajes en el archivo
	// TODO: Verificar por qué sólo se escribe una palabra
	_, err = file.WriteString(fmt.Sprintf("%s\n",content))
	if err != nil {
		fmt.Printf("Error al escribir en el archivo: %v\n", err)
		return
	}

	fmt.Printf("Archivo %s creado exitosamente.\n", fileName)
}

/*
	- Bienvenido Usuario, por favor ingrese sus credenciales:
	: Username
	: Pwd
	 -> ERR Credenciales invalidas
	 |
	 Validación del rol del usuario - Admin o Usr
	 |
	 ___________________________________________________
	 SI USUARIO
	 -> Print -> ERES UN USUARIO
	 -> Ver la cantidad de letras presionadas
	 ____________________________________________________
	 SI ADMIN
	 ->
*/
