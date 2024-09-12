package main

import (
	"fmt"

	"static-site-generator/pkg/textnodes"
)

func main() {
	textNode := textnodes.NewTextNode("This is a text node", "bold", "https://www.boot.dev")
	fmt.Println(textNode)
}
