package clothing

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
)

// Item is a type of clothing item
type Item struct {
	Name           string
	Type           string
	IdealHeatLevel string
	Layer          int
}

func getItems(originClimate climate.Climate) []Item {
	filteredItems := []Item{}

	items := getAllItems()
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

func getAllItems() []Item {
	items := []Item{
		Item{
			Name:           "short tunic",
			Type:           "top",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "knee-length tunic",
			Type:           "top",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "ankle-length tunic",
			Type:           "top",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "pantaloons",
			Type:           "bottom",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "sleeveless shirt",
			Type:           "top",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "vest",
			Type:           "top",
			IdealHeatLevel: "any",
			Layer:          2,
		},
		Item{
			Name:           "loincloth",
			Type:           "undergarment",
			IdealHeatLevel: "any",
			Layer:          0,
		},
		Item{
			Name:           "breeches",
			Type:           "bottom",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "long pants",
			Type:           "bottom",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "sandals",
			Type:           "footwear",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "knee-high boots",
			Type:           "footwear",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "short boots",
			Type:           "footwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "strapped boots",
			Type:           "footwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "turnshoes",
			Type:           "footwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "robe",
			Type:           "full",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "dress",
			Type:           "full",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "belt",
			Type:           "waist",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "sash",
			Type:           "waist",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "skirt",
			Type:           "bottom",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "kilt",
			Type:           "bottom",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "short gloves",
			Type:           "handwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "long gloves",
			Type:           "handwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "mittens",
			Type:           "handwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "short coat",
			Type:           "overwear",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		Item{
			Name:           "cloak",
			Type:           "overwear",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "long coat",
			Type:           "overwear",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		Item{
			Name:           "hooded cloak",
			Type:           "overwear",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "wide-brim hat",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "fez",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "conical hat",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "skullcap",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "kufi",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "conical cap",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "turban",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
	}

	return items
}
