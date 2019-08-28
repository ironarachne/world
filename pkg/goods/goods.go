package goods

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

// TradeGood is a trade good entry
type TradeGood struct {
	Name       string `json:"name"`
	Quality    string `json:"quality"`
	Amount     int    `json:"amount"`
	PriceEach  int    `json:"price_each"`
	PriceTotal int    `json:"price_total"`
}

// InSlice checks to see if a good is in a list of goods
func (good TradeGood) InSlice(goods []TradeGood) bool {
	for _, g := range goods {
		if g.Name == good.Name {
			return true
		}
	}

	return false
}

// GenerateMerchantGoods creates a list of specific trade goods for use in shop lists
func GenerateMerchantGoods(min int, max int, resources []resource.Resource) []TradeGood {
	var good TradeGood
	var quality string
	var skillLevel int

	goods := []TradeGood{}
	possibleGoods := []TradeGood{}
	tradeGoodNames := []string{}
	amount := 0

	for _, r := range resources {
		if len(r.Tags) > 0 {
			amount = rand.Intn(3) + 1
			skillLevel = rand.Intn(5)
			quality = qualityFromSkillLevel(skillLevel)
			good = TradeGood{
				Name:      r.Name,
				Quality:   quality,
				Amount:    amount,
				PriceEach: price(skillLevel, r),
			}
			good.PriceTotal = good.PriceEach * good.Amount
			possibleGoods = append(possibleGoods, good)
		}
	}

	numberOfGoods := rand.Intn(max+1-min) + min

	for i := 0; i < numberOfGoods; i++ {
		good = possibleGoods[rand.Intn(len(possibleGoods))]
		if !slices.StringIn(good.Name, tradeGoodNames) {
			goods = append(goods, good)
			tradeGoodNames = append(tradeGoodNames, good.Name)
		}
	}

	return goods
}

// GenerateExportTradeGoods produces a list of trade goods based on given resources
func GenerateExportTradeGoods(min int, max int, resources []resource.Resource) []TradeGood {
	var good TradeGood
	var quality string
	var skillLevel int

	goods := []TradeGood{}
	possibleGoods := []TradeGood{}
	tradeGoodNames := []string{}
	amount := 0

	for _, r := range resources {
		if len(r.Tags) > 0 {
			amount = rand.Intn(3) + 1
			skillLevel = rand.Intn(5)
			quality = qualityFromSkillLevel(skillLevel)
			good = TradeGood{
				Name:      r.Tags[0],
				Quality:   quality,
				Amount:    amount,
				PriceEach: price(skillLevel, r),
			}
			good.PriceTotal = good.PriceEach * good.Amount
			possibleGoods = append(possibleGoods, good)
		}
	}

	numberOfGoods := rand.Intn(max+1-min) + min

	for i := 0; i < numberOfGoods; i++ {
		good = possibleGoods[rand.Intn(len(possibleGoods))]
		if !slices.StringIn(good.Name, tradeGoodNames) {
			goods = append(goods, good)
			tradeGoodNames = append(tradeGoodNames, good.Name)
		}
	}

	return goods
}

// GenerateImportTradeGoods produces a list of trade goods based on externally-available resources
func GenerateImportTradeGoods(min int, max int, resources []resource.Resource) []TradeGood {
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
		if !good.InSlice(goods) {
			goods = append(goods, good)
		}
	}

	return goods
}

// GetAllTradeGoods converts a list of resources into a list of trade goods
func GetAllTradeGoods(resources []resource.Resource) []string {
	goods := []string{}

	for _, r := range resources {
		if len(r.Tags) > 0 && !slices.StringIn(r.Tags[0], goods) {
			goods = append(goods, r.Tags[0])
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
