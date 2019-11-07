package fish

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined fish
func All() []species.Species {
	fish := []species.Species{}

	fish = append(fish, getLakeFish()...)
	fish = append(fish, getOceanFish()...)
	fish = append(fish, getRiverFish()...)

	for _, f := range fish {
		f.Resources = []resource.Resource{
			{
				Name:         f.Name,
				Origin:       f.Name,
				MainMaterial: f.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 4,
				Value:       1,
			},
		}
	}

	return fish
}
