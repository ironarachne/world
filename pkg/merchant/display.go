package merchant

import (
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/goods"
)

// SimplifiedMerchant is a simplified version of merchant for display
type SimplifiedMerchant struct {
	Character character.SimplifiedCharacter `json:"character"`
	Goods     []goods.TradeGood             `json:"goods"`
}

// Simplify simplifies a merchant
func (merchant Merchant) Simplify() SimplifiedMerchant {
	sc := merchant.Character.Simplify()
	goods := merchant.Goods

	return SimplifiedMerchant{
		Character: sc,
		Goods:     goods,
	}
}

// RandomSimplified returns a random simplified merchant
func RandomSimplified() SimplifiedMerchant {
	merchant := Random()

	return merchant.Simplify()
}
