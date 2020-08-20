package main

import "sort"

type match struct {
	card            string
	item            string
	stack           int
	cardPrice       float32
	cardTotalPrice  float32
	itemPrice       float32
	benefit         float32
	relativeBenefit float32
}

func doMatches(cards []item, items []item, minChaos float32, maxChaos float32) []match {
	matchList := make([]match, 0)

	for _, card := range cards {
		for _, item := range items {
			matches := false

			for _, result := range card.Results {
				if result.Optional {
					continue
				}

				if result.Text == item.Name {
					matches = true
					break
				}
			}

			if !matches {
				continue
			}

			cardsTotalPrice := card.ChaosValue * float32(card.StackSize)
			benefit := item.ChaosValue - cardsTotalPrice
			relativeBenefit := benefit / cardsTotalPrice
			result := match{
				card:            card.Name,
				item:            item.Name,
				stack:           card.StackSize,
				cardPrice:       card.ChaosValue,
				cardTotalPrice:  cardsTotalPrice,
				itemPrice:       item.ChaosValue,
				benefit:         benefit,
				relativeBenefit: relativeBenefit,
			}

			matchList = append(matchList, result)
		}
	}

	matchListFiltered := make([]match, 0)
	for _, match := range matchList {
		if match.cardTotalPrice > minChaos && match.cardTotalPrice < maxChaos {
			matchListFiltered = append(matchListFiltered, match)
		}
	}

	sort.Slice(matchListFiltered, func(i, j int) bool {
		return matchListFiltered[i].benefit > matchListFiltered[j].benefit
	})

	return matchListFiltered
}
