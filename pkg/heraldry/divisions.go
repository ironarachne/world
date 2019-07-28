package heraldry

import (
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"math/rand"
)

// Division is a division of the field
type Division struct {
	Name   string
	Blazon string
	tincture.Tincture
}

func allDivisions() []Division {
	divisions := []Division{
		{
			Name: "bend",
			Blazon: "Per bend ",
		},
		{
			Name: "bendsinister",
			Blazon: "Per bend sinister ",
		},
		{
			Name: "fess",
			Blazon: "Per fess ",
		},
		{
			Name: "pale",
			Blazon: "Per pale ",
		},
		{
			Name: "plain",
			Blazon: "",
		},
		{
			Name: "quarterly",
			Blazon: "Quarterly ",
		},
		{
			Name: "saltire",
			Blazon: "Per saltire ",
		},
		{
			Name: "chevron",
			Blazon: "Per chevron",
		},
	}

	return divisions
}

func randomDivision() Division {
	divisions := allDivisions()
	return divisions[rand.Intn(len(divisions))]
}
