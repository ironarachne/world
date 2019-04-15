package culture

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/random"
)

// ClothingItem is a type of clothing item
type ClothingItem struct {
	Name           string
	Type           string
	IdealHeatLevel string
	Layer          int
}

// ClothingStyle describes what kind of clothing the culture wears
type ClothingStyle struct {
	CommonFemaleItems []ClothingItem
	CommonMaleItems   []ClothingItem
	CommonJewelry     []string
	CommonColors      []string
	CommonMaterials   []string
	DecorativeStyle   string
}

func (culture Culture) generateClothingStyle() ClothingStyle {
	style := ClothingStyle{}

	style.CommonMaterials = culture.getClothingFabrics()
	style.CommonFemaleItems = culture.getClothingItems()
	style.CommonMaleItems = culture.getClothingItems()

	style.CommonJewelry = culture.generateJewelry()

	style.DecorativeStyle = culture.randomDecorativeStyle()
	style.CommonColors = randomColorSet()

	return style
}

func (culture Culture) generateJewelry() []string {
	var jewelryItem string
	var gemProbability int

	jewelry := []string{}

	descriptors := []string{
		"brilliant",
		"gaudy",
		"lustrous",
		"ornate",
		"simple",
	}

	foundations := []string{
		"anklets",
		"bracelets",
		"chokers",
		"necklaces",
		"rings",
	}

	settings := []string{
		"adorned with",
		"decorated with",
		"set with",
	}

	materials := []string{}
	gems := []string{}

	for _, r := range culture.HomeClimate.Resources {
		if r.Type == "metal ingot" {
			materials = append(materials, strings.TrimSuffix(r.Name, " ingot"))
		} else if r.Type == "metal bar" {
			materials = append(materials, strings.TrimSuffix(r.Name, " bar"))
		} else if r.Type == "gem" {
			gems = append(gems, r.Name)
		}
	}

	numberOfJewelryPieces := rand.Intn(4) + 1

	for i := 0; i < numberOfJewelryPieces; i++ {
		jewelryItem = random.Item(descriptors) + " " + random.Item(materials) + " " + random.Item(foundations)
		if len(gems) > 0 {
			gemProbability = rand.Intn(10) + 1
			if gemProbability > 5 {
				jewelryItem += " " + random.Item(settings) + " " + random.Item(gems)
			}
		}

		jewelry = append(jewelry, jewelryItem)
	}

	return jewelry
}

func (culture Culture) getClothingItems() []ClothingItem {
	filteredItems := []ClothingItem{}

	items := getAllClothingItems()
	heatLevel := "warm"

	if culture.HomeClimate.Temperature < 5 {
		heatLevel = "cold"
	}

	items = getClothingItemsForHeat(heatLevel, items)

	fullOrIndividual := rand.Intn(10)
	if fullOrIndividual > 6 {
		full := getClothingItemsForType("full", items)
		if len(full) > 0 {
			filteredItems = append(filteredItems, full[rand.Intn(len(full))])
		}
	} else {
		bottoms := getClothingItemsForType("bottom", items)
		tops := getClothingItemsForType("top", items)
		if len(bottoms) > 0 {
			filteredItems = append(filteredItems, bottoms[rand.Intn(len(bottoms))])
		}
		if len(tops) > 0 {
			filteredItems = append(filteredItems, tops[rand.Intn(len(tops))])
		}
	}

	footwear := getClothingItemsForType("footwear", items)
	hats := getClothingItemsForType("headwear", items)
	overwear := getClothingItemsForType("overwear", items)
	underwear := getClothingItemsForType("underwear", items)
	waist := getClothingItemsForType("waist", items)

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

func getClothingItemsForHeat(heat string, source []ClothingItem) []ClothingItem {
	items := []ClothingItem{}

	for _, i := range source {
		if i.IdealHeatLevel == heat || i.IdealHeatLevel == "any" {
			items = append(items, i)
		}
	}

	return items
}

func getClothingItemsForLayer(layer int, source []ClothingItem) []ClothingItem {
	items := []ClothingItem{}

	for _, i := range source {
		if i.Layer == layer {
			items = append(items, i)
		}
	}

	return items
}

func getClothingItemsForType(clothingType string, source []ClothingItem) []ClothingItem {
	items := []ClothingItem{}

	for _, i := range source {
		if i.Type == clothingType {
			items = append(items, i)
		}
	}

	return items
}

func getAllClothingItems() []ClothingItem {
	items := []ClothingItem{
		ClothingItem{
			Name:           "short tunic",
			Type:           "top",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "knee-length tunic",
			Type:           "top",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "ankle-length tunic",
			Type:           "top",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		ClothingItem{
			Name:           "pantaloons",
			Type:           "bottom",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "sleeveless shirt",
			Type:           "top",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		ClothingItem{
			Name:           "vest",
			Type:           "top",
			IdealHeatLevel: "any",
			Layer:          2,
		},
		ClothingItem{
			Name:           "loincloth",
			Type:           "undergarment",
			IdealHeatLevel: "any",
			Layer:          0,
		},
		ClothingItem{
			Name:           "breeches",
			Type:           "bottom",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "long pants",
			Type:           "bottom",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "sandals",
			Type:           "footwear",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		ClothingItem{
			Name:           "knee-high boots",
			Type:           "footwear",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		ClothingItem{
			Name:           "short boots",
			Type:           "footwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "strapped boots",
			Type:           "footwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "turnshoes",
			Type:           "footwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "robe",
			Type:           "full",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		ClothingItem{
			Name:           "dress",
			Type:           "full",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "belt",
			Type:           "waist",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		ClothingItem{
			Name:           "sash",
			Type:           "waist",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		ClothingItem{
			Name:           "skirt",
			Type:           "bottom",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		ClothingItem{
			Name:           "kilt",
			Type:           "bottom",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		ClothingItem{
			Name:           "short gloves",
			Type:           "handwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "long gloves",
			Type:           "handwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "mittens",
			Type:           "handwear",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "short coat",
			Type:           "overwear",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		ClothingItem{
			Name:           "cloak",
			Type:           "overwear",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		ClothingItem{
			Name:           "long coat",
			Type:           "overwear",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		ClothingItem{
			Name:           "hooded cloak",
			Type:           "overwear",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		ClothingItem{
			Name:           "wide-brim hat",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "fez",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "conical hat",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "skullcap",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "kufi",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "conical cap",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		ClothingItem{
			Name:           "turban",
			Type:           "hat",
			IdealHeatLevel: "any",
			Layer:          1,
		},
	}

	return items
}

func (culture Culture) getClothingFabrics() []string {
	fabrics := []string{}

	for _, i := range culture.HomeClimate.Plants {
		if i.Name == "flax" {
			fabrics = append(fabrics, "linen")
		} else if i.IsFiber {
			fabrics = append(fabrics, i.Name)
		}
	}

	for _, i := range culture.HomeClimate.Animals {
		if i.GivesWool {
			fabrics = append(fabrics, i.Name+" wool")
		} else if i.GivesFur {
			fabrics = append(fabrics, i.Name+" fur")
		} else if i.Name == "cow" {
			fabrics = append(fabrics, "leather")
		} else if i.AnimalType == "reptile" {
			fabrics = append(fabrics, i.Name+" skin")
		} else if i.GivesHide {
			fabrics = append(fabrics, i.Name+" hide")
		}
	}

	return fabrics
}

func (culture Culture) randomDecorativeStyle() string {
	styles := []string{
		"beads",
		"complex embroidery",
		"geometric shapes",
		"knotwork",
		"plain",
		"simple embroidery",
		"stripes",
		"stylized animals",
		"tassels",
	}

	if climate.IsTypeInResources("ivory", culture.HomeClimate.Resources) {
		styles = append(styles, "ivory decorations")
	}

	if climate.IsTypeInResources("feathers", culture.HomeClimate.Resources) {
		styles = append(styles, "feather decorations")
	}

	if culture.HomeClimate.Temperature < 6 {
		styles = append(styles, "many layers")
	}

	return random.Item(styles)
}

func randomColorSet() []string {
	var newColor string
	var colorSet []string
	var saturation string

	colors := []string{
		"red",
		"blue",
		"green",
		"black",
		"white",
		"yellow",
		"purple",
		"orange",
		"grey",
	}

	for i := 0; i < 3; i++ {
		newColor = random.Item(colors)

		if !inSlice(newColor, []string{"white", "black"}) {
			saturation = randomSaturation()
			newColor = saturation + " " + newColor
		}

		if !inSlice(newColor, colorSet) {
			colorSet = append(colorSet, newColor)
		}
	}

	return colorSet
}

func randomSaturation() string {
	saturations := []string{
		"bright",
		"dark",
		"light",
		"moderate",
		"pastel",
		"subdued",
	}

	return random.Item(saturations)
}
