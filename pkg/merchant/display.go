package merchant

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/goods"
)

// SimplifiedMerchant is a simplified version of merchant for display
type SimplifiedMerchant struct {
	Character character.SimplifiedCharacter `json:"character"`
	Goods     []goods.TradeGood             `json:"goods"`
}

// Simplify simplifies a merchant
func (merchant Merchant) Simplify(ctx context.Context) (SimplifiedMerchant, error) {
	sc, err := merchant.Character.Simplify(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify merchant: %w", err)
		return SimplifiedMerchant{}, err
	}
	goods := merchant.Goods

	return SimplifiedMerchant{
		Character: sc,
		Goods:     goods,
	}, nil
}

// RandomSimplified returns a random simplified merchant
func RandomSimplified(ctx context.Context) (SimplifiedMerchant, error) {
	merchant, err := Random(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify random merchant: %w", err)
		return SimplifiedMerchant{}, err
	}

	sm, err := merchant.Simplify(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify random merchant: %w", err)
		return SimplifiedMerchant{}, err
	}

	return sm, nil
}
