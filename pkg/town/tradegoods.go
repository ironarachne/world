package town

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// TradeGood is a trade good entry
type TradeGood struct {
	Name    string
	Quality string
	Amount  int
}

func (town Town) generateExportTradeGoods(min int, max int) []TradeGood {
	var good TradeGood
	var quality string

	goods := []TradeGood{}

	for _, p := range town.NotableProducers {
		quality = qualityFromSkillLevel(p.SkillLevel)
		for _, m := range p.GoodsMade {
			good = TradeGood{
				Name:    m,
				Quality: quality,
				Amount:  rand.Intn(3) + 1,
			}
			goods = append(goods, good)
		}
	}

	possibleGoods := getFarmGoods(town.Climate.Resources)

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

func (town Town) generateImportTradeGoods(min int, max int, resources []climate.Resource) []TradeGood {
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
