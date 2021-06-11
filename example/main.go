package main

import (
	"fmt"

	markdown "github.com/skyaxl/go-term-markdown"
)

func main() {
	output := markdown.Render("# Ola eu sou um novo mund \n"+
		"```\neu sou um codigo\n```\n"+
		"## Negrito\n"+
		"**negrito** \n"+
		"## Italico\n"+
		"*italico* \n",
		100, 4)
	fmt.Println(string(output))
}
