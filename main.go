package main

import (
	"fmt"
)

const league = "Harvest"

func main() {
	cards := getItems("DivinationCard", league)
	items := getResultItems(league)

	matches := doMatches(cards, items, 0, 9999999)

	printMatches(matches)
}

func printMatches(matches []match) {
	for i, match := range matches {
		if i > 20 {
			break
		}

		fmt.Printf("============================\n")
		fmt.Printf("Card:              %s\n", match.card)
		fmt.Printf("Item:              %s\n", match.item)
		fmt.Printf("Card price:        %.2f\n", match.cardPrice)
		fmt.Printf("Card stack price:  %.2f\n", match.cardTotalPrice)
		fmt.Printf("Item price:        %.2f\n", match.itemPrice)
		fmt.Printf("Profit:            %.2f\n", match.benefit)
		fmt.Printf("Profit:            %.2f%%\n", 100.0*match.relativeBenefit)
	}
}
