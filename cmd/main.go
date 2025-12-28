package main

import (
	"flag"
	"fmt"
	"os"

	"passgen/internal/generator"
)

func main() {
	length := flag.Int("length", 12, "Password length")
	flag.Parse()

	password, err := generator.Generate(*length)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Generated Password:")
	fmt.Println(password)
}
