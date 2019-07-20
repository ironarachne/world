package goods

import (
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
	"math/rand"
)

// TradeGood is a trade good entry
type TradeGood struct {
	Name    string
	Quality string
	Amount  int
}

// GenerateExportTradeGoods produces a list of trade goods based on local production
func GenerateExportTradeGoods(min int, max int, professions []profession.Profession, resources []resource.Resource) []TradeGood {
	var filledPattern resource.Pattern
	var good TradeGood
	var description string
	var patterns []resource.Pattern
	var possiblePatterns []resource.Pattern
	var quality string
	var skillLevel int
	var resourcesForSlot []resource.Resource
	var resourceForSlot resource.Resource

	goods := []TradeGood{}
	possibleGoods := []TradeGood{}
	tradeGoodNames := []string{}
	amount := 0
	allPatterns := resource.AllPatterns()

	for _, p := range professions {
		possiblePatterns = []resource.Pattern{}
		skillLevel = rand.Intn(5)
		quality = qualityFromSkillLevel(skillLevel)
		patterns = resource.FindPatternsForProfession(p, allPatterns)
		for _, n := range patterns {
			if n.CanMake(resources) {
				possiblePatterns = append(possiblePatterns, n)
			}
		}
		if len(possiblePatterns) > 0 {
			for _, n := range possiblePatterns {
				filledPattern = resource.Pattern{
					Name: n.Name,
					Description: n.Description,
					Profession: n.Profession,
					Slots: []resource.Slot{},
				}
				for _, s := range n.Slots {
					resourcesForSlot = resource.ListOfType(s.RequiredType, resources)
					resourceForSlot = resource.Random(resourcesForSlot)
					s.Resource = resourceForSlot
					filledPattern.Slots = append(filledPattern.Slots, s)
				}
				description = filledPattern.Render()
				amount = rand.Intn(3)+1
				good = TradeGood{
					Name: description,
					Quality: quality,
					Amount: amount,
				}
				possibleGoods = append(possibleGoods, good)
			}
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
		goods = append(goods, good)
	}

	return goods
}

// GetAllTradeGoods converts a list of resources into a list of trade goods
func GetAllTradeGoods(resources []resource.Resource) []string {
	goods := []string{}

	for _, r := range resources {
		goods = append(goods, r.Name)
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
