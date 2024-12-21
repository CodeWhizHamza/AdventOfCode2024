package main

import (
	"fmt"
)

func main() {
	lines := getLines()
	charactersMatrix := getCharactersMatrix(lines)
	fmt.Println(charactersMatrix)
}
