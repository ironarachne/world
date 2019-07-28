package charge

import (
	svg "github.com/ajstarks/svgo"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"github.com/ironarachne/world/pkg/slices"
	"math/rand"
)

// Charge is a charge
type Charge struct {
	Identifier string
	Name       string
	Noun       string
	NounPlural string
	Descriptor string
	SingleOnly bool
	Tags       []string
	Render     func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int)
}

// Group is a group of charges with a common tincture
type Group struct {
	Charges []Charge
	tincture.Tincture
}

func all() []Charge {
	charges := []Charge{
		AntelopeRampant,
		BattleAxe,
		BatVolant,
		BearHeadCouped,
		BearRampant,
		BearStatant,
		BeeVolant,
		Bell,
		BoarHeadErased,
		BoarPassant,
		BoarRampant,
		BugleHorn,
		BullPassant,
		BullRampant,
		Castle,
		Cock,
		Cockatrice,
		Crown,
		DolphinHauriant,
		DoubleHeadedEagleDisplayed,
		DragonPassant,
		DragonRampant,
		EagleDisplayed,
		EaglesHeadErased,
		FoxPassant,
		FoxSejant,
		GryphonPassant,
		Hare,
		HareSalient,
		Heron,
		HorsePassant,
		HorseRampant,
		LeopardPassant,
		LionPassant,
		LionRampant,
		LionsHeadErased,
		OrdinaryBend,
		OrdinaryBordure,
		OrdinaryChevron,
		OrdinaryChief,
		OrdinaryCross,
		OrdinaryFess,
		OrdinaryLozenge,
		OrdinaryPale,
		OrdinaryPall,
		OrdinaryPile,
		OrdinaryRoundel,
		OrdinarySaltire,
		Owl,
		PegasusPassant,
		RamRampant,
		RamStatant,
		Rose,
		SeaHorse,
		Squirrel,
		StagLodged,
		StagStatant,
		SunInSplendor,
		TigerPassant,
		TigerRampant,
		Tower,
		TwoAxesInSaltire,
		TwoBonesInSaltire,
		UnicornStatant,
		WolfPassant,
		WolfRampant,
		Wyvern,
	}

	return charges
}

// MatchingTag returns all charges that match a tag
func MatchingTag(tag string) []Charge {
	matching := []Charge{}

	charges := all()

	for _, c := range charges {
		if slices.StringIn(tag, c.Tags) {
			matching = append(matching, c)
		}
	}

	return matching
}

// Random returns a random charge
func Random() Charge {
	charges := all()
	//BrokenCharges = [4]Charge{AntelopePassant, GryphonSegreant, UnicornRampant, SerpentNowed}
	return charges[rand.Intn(len(charges))]
}

// RandomMatchingTag returns a random charge that matches a tag
func RandomMatchingTag(tag string) Charge {
	charges := MatchingTag(tag)

	return charges[rand.Intn(len(charges))]
}
