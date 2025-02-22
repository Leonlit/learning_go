package main

import (
	"fmt"
	"nmapParser/generator"
	"nmapParser/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, require a file path/name!")
		return
	}

	fileData, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Can't open file:", os.Args[1])
		panic(err)
	}

	doc, err := parser.ParseNmap(fileData)
	if err != nil {
		panic(err)
	}

	err = generator.HTMLGenerator(doc)
	if err != nil {
		panic(err)
	}

	// Test printing the extracted values
	fmt.Printf("Up: %d\n", doc.RunStats.Hosts.Up)
	fmt.Printf("Down: %d\n", doc.RunStats.Hosts.Down)
	fmt.Printf("Total: %d\n", doc.RunStats.Hosts.Total)
}
