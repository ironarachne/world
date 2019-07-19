package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getCanines() []Animal {
	animals := []Animal{
		Animal{
			Name:       "coyote",
			PluralName: "coyotes",
		},
		Animal{
			Name:       "fox",
			PluralName: "foxes",
		},
		Animal{
			Name:       "wolf",
			PluralName: "wolves",
		},
	}

	for _, a := range animals {
		a.AnimalType = "mammal"
		a.EatsAnimals = true
		a.EatsPlants = false
		a.IsMount = false
		a.IsPackAnimal = false
		a.IsScavenger = false
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 0
		a.MaxTemperature = 10
		a.Size = size.GetCategoryByName("medium")
		a.Resources = []resource.Resource{
			resource.Resource{
				Name:        a.Name + " hide",
				Origin:      a.Name,
				Type:        "hide",
				Commonality: 4,
			},
			resource.Resource{
				Name:        a.Name + " fangs",
				Origin:      a.Name,
				Type:        "teeth",
				Commonality: 4,
			},
			resource.Resource{
				Name:        a.Name,
				Origin:      a.Name,
				Type:        "meat",
				Commonality: 4,
			},
		}
	}

	return animals
}
