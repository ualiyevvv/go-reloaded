package main

import (
	g "go-reloaded/pkg"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Println("ERROR: must need 2 files")
		return
	}
	text := string(g.ReadFile(args[0]))
	arrtext := g.Split(text)
	// fmt.Println(text)
	newtext := ""
	for _, text := range arrtext {
		text = g.NumFilter(text)
		text = g.ParticlesFilter(text)
		text = g.CasesFilter(text)
		text = g.PunctuationsFilter(text)
		newtext += text
	}

	g.WriteToFile(args[1], newtext)

}
