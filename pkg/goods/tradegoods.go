package goods

import (
	"bytes"
	"html/template"
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// GenerateExportTradeGoods produces a list of trade goods based on local production
func GenerateExportTradeGoods(min int, max int, producers []Producer, resources []climate.Resource) []TradeGood {
	var good TradeGood
	var quality string

	goods := []TradeGood{}
	possibleGoods := []TradeGood{}
	tradeGoodNames := []string{}
	amount := 0

	farmGoods := getFarmGoods(resources)

	for _, f := range farmGoods {
		amount = rand.Intn(3) + 1
		good.Name = f
		good.Amount = amount
		good.Quality = randomQuality()
		possibleGoods = append(goods, good)
	}

	for _, p := range producers {
		quality = qualityFromSkillLevel(p.SkillLevel)
		for _, pattern := range p.Patterns {
			pattern.Material1 = randomMaterialFromType(pattern.Need1, resources)
			if pattern.Need2 != "" {
				pattern.Material2 = randomMaterialFromType(pattern.Need2, resources)
			}
			if pattern.Need3 != "" {
				pattern.Material3 = randomMaterialFromType(pattern.Need3, resources)
			}

			good = TradeGood{
				Name:    pattern.getName(),
				Quality: quality,
				Amount:  rand.Intn(3) + 1,
			}
			possibleGoods = append(possibleGoods, good)
		}
	}

	numberOfGoods := rand.Intn(max+1-min) + min

	for i := 0; i < numberOfGoods; i++ {
		good = possibleGoods[rand.Intn(len(possibleGoods)-1)]
		if !slices.StringIn(good.Name, tradeGoodNames) {
			goods = append(goods, good)
			tradeGoodNames = append(tradeGoodNames, good.Name)
		}
	}

	return goods
}

// GenerateImportTradeGoods produces a list of trade goods based on externally-available resources
func GenerateImportTradeGoods(min int, max int, resources []climate.Resource) []TradeGood {
	var good TradeGood

	goods := []TradeGood{}

	possibleGoods := GetAllTradeGoods(resources)

	numberOfGoods := rand.Intn(max+1-min) + min
	amount := 0
	newItem := ""

	for i := 0; i < numberOfGoods; i++ {
		good = TradeGood{}
		newItem = random.String(possibleGoods)
		amount = rand.Intn(3) + 1
		good.Name = newItem
		good.Amount = amount
		good.Quality = randomQuality()
		goods = append(goods, good)
	}

	return goods
}

// GetAllTradeGoods converts a list of resources into a list of trade goods
func GetAllTradeGoods(resources []climate.Resource) []string {
	goods := []string{}

	for _, resource := range resources {
		goods = append(goods, resource.Name)
	}

	return goods
}

func getFarmGoods(resources []climate.Resource) []string {
	goods := []string{}

	goodTypes := []string{
		"eggs",
		"fruit",
		"grain",
		"meat",
		"milk",
		"vegetable",
	}

	for _, resource := range resources {
		if slices.StringIn(resource.Type, goodTypes) {
			goods = append(goods, resource.Name)
		}
	}

	return goods
}

func (pattern Pattern) getName() string {
	var tplOutput bytes.Buffer

	tmpl, err := template.New(pattern.Name).Parse(pattern.NameTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tplOutput, pattern)
	if err != nil {
		panic(err)
	}
	name := tplOutput.String()

	return name
}

func randomMaterialFromType(goodType string, resources []climate.Resource) string {
	var possibles []string

	for _, r := range resources {
		if r.Type == goodType {
			possibles = append(possibles, r.Name)
		}
	}

	return random.String(possibles)
}

func randomQuality() string {
	qualities := map[string]int{
		"exceptional":  1,
		"fine":         2,
		"":             11,
		"questionable": 2,
		"pathetic":     1,
	}

	return random.StringFromThresholdMap(qualities)
}
