package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	fmt.Print(cards)
}