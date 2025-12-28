package main

import (
	"flag"
	"fmt"
)

func main(){
	length:= flag.Int("length",12,"Password length")
	flag.Parse();

	fmt.Println("Password length:", *length)
}