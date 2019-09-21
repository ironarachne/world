package species

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/trait"
)

// Species is a species of living thing
type Species struct {
	Name           string
	PluralName     string
	Adjective      string
	Commonality    int
	PossibleTraits []trait.Template // Traits that individuals of this species *might* have
	CommonTraits   []trait.Template // Traits that all members of this species share
	AgeCategories  []age.Category
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Resources      []resource.Resource // These are resources that can be derived from this species
	Tags           []string
}
