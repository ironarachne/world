package merchant

import (
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/goods"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/town"
)

// Merchant is a merchant
type Merchant struct {
	Character character.Character
	Goods     []goods.TradeGood
}

// Generate returns a random merchant
func Generate(originTown town.Town) Merchant {
	character := character.Random()
	character.Profession = profession.ByName("merchant")

	goods := goods.GenerateMerchantGoods(10, 30, originTown.Resources)

	merchant := Merchant{
		Character: character,
		Goods:     goods,
	}

	return merchant
}

// Random returns a complete random merchant
func Random() Merchant {
	originCulture := culture.Random()
	originTown := town.Generate("metropolis", "random", originCulture)

	merchant := Generate(originTown)

	return merchant
}
