package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

var regex = regexp.MustCompile("(<.*>)|({)|(})|(\r)")

type cardResult struct {
	Text     string
	Optional bool
}

type item struct {
	ID         int
	Name       string
	StackSize  int
	Results    []cardResult `json:"explicitModifiers"`
	ChaosValue float32
	Count      int
	Links      int
}

type data struct {
	Lines []item
}

func getResultItems(league string) []item {
	items := append(getItems("Essence", league), getItems("UniqueArmour", league)...)
	items = append(items, getItems("UniqueFlask", league)...)
	items = append(items, getItems("UniqueWeapon", league)...)
	items = append(items, getItems("UniqueAccessory", league)...)
	items = append(items, getItems("UniqueJewel", league)...)
	items = append(items, getItems("Prophecy", league)...)

	return items
}

func getItems(class string, league string) []item {
	res, err := http.Get("https://poe.ninja/api/data/itemoverview?league=" + league + "&type=" + class)
	if err != nil {
		log.Panicln(err)
	}

	var result data
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Panicln(err)
	}

	return filterItems(result.Lines)
}

func (c *cardResult) UnmarshalJSON(data []byte) error {
	var a struct {
		Text     string
		Optional bool
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	c.Text = regex.ReplaceAllString(a.Text, "")
	c.Optional = a.Optional

	return nil
}

func filterItems(items []item) []item {
	result := make([]item, 0)
	groups := make(map[string]item)

	for _, item := range items {
		group, exists := groups[item.Name]

		if !exists {
			groups[item.Name] = item
			continue
		}

		if item.ChaosValue < group.ChaosValue {
			groups[item.Name] = item
		}
	}

	for _, item := range groups {
		if item.Count > 20 {
			result = append(result, item)
		}
	}

	return result
}
