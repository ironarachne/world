package clothing

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
)

// Item is a type of clothing item
type Item struct {
	Name           string
	Type           string
	MaterialType   string
	IdealHeatLevel string
	Layer          int
}

func addMaterials(items []Item, hides []string, fabrics []string) []Item {
	var newItem Item
	var result []Item

	for _, i := range items {
		newItem = i
		if i.MaterialType == "fabric" {
			newItem.Name = random.String(fabrics) + " " + newItem.Name
			result = append(result, newItem)
		} else if i.MaterialType == "hide " {
			newItem.Name = random.String(hides) + " " + newItem.Name
			result = append(result, newItem)
		}
	}

	return result
}

func getItems(originClimate climate.Climate) []Item {
	filteredItems := []Item{}

	items := All()
	heatLevel := "warm"

	if originClimate.Temperature < 5 {
		heatLevel = "cold"
	}

	items = getItemsForHeat(heatLevel, items)

	fullOrIndividual := rand.Intn(10)
	if fullOrIndividual > 6 {
		full := getItemsForType("full", items)
		if len(full) > 0 {
			filteredItems = append(filteredItems, full[rand.Intn(len(full))])
		}
	} else {
		bottoms := getItemsForType("bottom", items)
		tops := getItemsForType("top", items)
		if len(bottoms) > 0 {
			filteredItems = append(filteredItems, bottoms[rand.Intn(len(bottoms))])
		}
		if len(tops) > 0 {
			filteredItems = append(filteredItems, tops[rand.Intn(len(tops))])
		}
	}

	footwear := getItemsForType("footwear", items)
	hats := getItemsForType("headwear", items)
	overwear := getItemsForType("overwear", items)
	underwear := getItemsForType("underwear", items)
	waist := getItemsForType("waist", items)

	if len(footwear) > 0 {
		filteredItems = append(filteredItems, footwear[rand.Intn(len(footwear))])
	}
	if len(hats) > 0 {
		filteredItems = append(filteredItems, hats[rand.Intn(len(hats))])
	}
	if len(overwear) > 0 {
		filteredItems = append(filteredItems, overwear[rand.Intn(len(overwear))])
	}
	if len(underwear) > 0 {
		filteredItems = append(filteredItems, underwear[rand.Intn(len(underwear))])
	}
	if len(waist) > 0 {
		filteredItems = append(filteredItems, waist[rand.Intn(len(waist))])
	}

	filteredItems = addMaterials(filteredItems, getHides(originClimate), getFabrics(originClimate))

	return filteredItems
}

func getItemsForHeat(heat string, source []Item) []Item {
	items := []Item{}

	for _, i := range source {
		if i.IdealHeatLevel == heat || i.IdealHeatLevel == "any" {
			items = append(items, i)
		}
	}

	return items
}

func getItemsForLayer(layer int, source []Item) []Item {
	items := []Item{}

	for _, i := range source {
		if i.Layer == layer {
			items = append(items, i)
		}
	}

	return items
}

func getItemsForType(clothingType string, source []Item) []Item {
	items := []Item{}

	for _, i := range source {
		if i.Type == clothingType {
			items = append(items, i)
		}
	}

	return items
}

// All returns all clothing items
func All() []Item {
	items := []Item{}

	items = append(items, getHandwear()...)
	items = append(items, getFull()...)
	items = append(items, getWaist()...)
	items = append(items, getBottoms()...)
	items = append(items, getHats()...)
	items = append(items, getTops()...)
	items = append(items, getFootwear()...)
	items = append(items, getOverwear()...)

	return items
}
