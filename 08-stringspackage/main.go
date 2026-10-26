package main

import (
	"fmt"
	"08-stringspackage/piscine"
)

func main() {
	w, c, h := piscine.SentenceStats("Hello world")
	fmt.Println(w, c, h)

	w, c, h = piscine.SentenceStats("say HELLO to everyone")
	fmt.Println(w, c, h)

	w, c, h = piscine.SentenceStats("no greeting here")
	fmt.Println(w, c, h)

	w, c, h = piscine.SentenceStats("")
	fmt.Println(w, c, h)
}