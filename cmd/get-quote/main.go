package main

import (
	"fmt"
	"os"

	"github.com/eubide/get-quote/pkg/randomline"
)

func main() {
	// Check if a file parameter was provided
	if len(os.Args) < 2 {
		fmt.Printf("Uso: %s <nombre_fichero>\n", os.Args[0])
		fmt.Println("Debe proporcionar un nombre de fichero")
		os.Exit(1)
	}

	// Get the file parameter
	fileName := os.Args[1]

	// Get a random line from the file
	configPath := ".get-quote.yaml"
	line, err := randomline.GetRandomLine(fileName, configPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Print the random line without a newline
	fmt.Print(line)
}
