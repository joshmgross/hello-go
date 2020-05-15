package main

import (
    "fmt"
    "os"
)

func main() {
	name := os.Getenv("INPUT_NAME")
	phrase := os.Getenv("INPUT_PHRASE")
	fmt.Printf("%s says %s.\n", name, phrase)
	

	fmt.Printf("::set-output name=%s::%s\n", "name", "Go")
	fmt.Printf("::set-output name=%s::%s\n", "phrase", "I'm a gopher")
}
