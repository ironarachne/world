package merchant

import (
	"fmt"

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
func Generate(originTown town.Town) (Merchant, error) {
	character, err := character.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate merchant: %w", err)
		return Merchant{}, err
	}
	character.Profession = profession.ByName("merchant")

	goods := goods.GenerateMerchantGoods(10, 30, originTown.Resources)

	merchant := Merchant{
		Character: character,
		Goods:     goods,
	}

	return merchant, nil
}

// Random returns a complete random merchant
func Random() (Merchant, error) {
	originCulture, err := culture.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate random merchant: %w", err)
		return Merchant{}, err
	}
	originTown, err := town.Generate("metropolis", "random", originCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate random merchant: %w", err)
		return Merchant{}, err
	}

	merchant, err := Generate(originTown)
	if err != nil {
		err = fmt.Errorf("Could not generate random merchant: %w", err)
		return Merchant{}, err
	}

	return merchant, nil
}
