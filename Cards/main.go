package main

import "fmt"

func main() {
	myDeck := newDeck()
	myDeck.print()
	fmt.Println()
	remainingDeck, hand := deal(myDeck, 7)
	remainingDeck.print()
	fmt.Println()
	hand.print()
}
