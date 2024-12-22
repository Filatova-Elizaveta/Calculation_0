package main

import (
	"fmt"

	"https://github.com/Filatova-Elizaveta/Calculation_0/internal/application"
)

func main() {
	fmt.Printf("%s", "Server is starting...\nWarning: Do not close this console!\nIf don't understand how to work with server, please, read README.md\n")
	app := application.New()
	app.StartServer()
}
