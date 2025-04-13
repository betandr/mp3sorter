package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dhowden/tag"
)

func main() {
	f, err := os.Open("chokehold.mp3")
	if err != nil {
		fmt.Printf("error loading file: %v", err)
		return
	}
	defer f.Close()

	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Format())
	log.Print(m.Artist())
	log.Print(m.Title())
	log.Print(m.Genre())

}
