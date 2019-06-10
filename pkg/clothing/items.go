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

func getAllItems() []Item {
	items := []Item{
		Item{
			Name:           "short tunic",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "knee-length tunic",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "ankle-length tunic",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "pantaloons",
			Type:           "bottom",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "sleeveless shirt",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "vest",
			Type:           "top",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          2,
		},
		Item{
			Name:           "loincloth",
			Type:           "undergarment",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          0,
		},
		Item{
			Name:           "breeches",
			Type:           "bottom",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "long pants",
			Type:           "bottom",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "sandals",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "knee-high boots",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "short boots",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "strapped boots",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "turnshoes",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "robe",
			Type:           "full",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "dress",
			Type:           "full",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "belt",
			Type:           "waist",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "sash",
			Type:           "waist",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "skirt",
			Type:           "bottom",
			MaterialType:   "fabric",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "kilt",
			Type:           "bottom",
			MaterialType:   "fabric",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "short gloves",
			Type:           "handwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "long gloves",
			Type:           "handwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "mittens",
			Type:           "handwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "short coat",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		Item{
			Name:           "cloak",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "long coat",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		Item{
			Name:           "hooded cloak",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "wide-brim hat",
			Type:           "hat",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "fez",
			Type:           "hat",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "conical hat",
			Type:           "hat",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "skullcap",
			Type:           "hat",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "kufi",
			Type:           "hat",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "conical cap",
			Type:           "hat",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "turban",
			Type:           "hat",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
	}

	return items
}
