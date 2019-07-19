package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getBigCats() []Animal {
	animals := []Animal{
		Animal{
			Name:       "bobcat",
			PluralName: "bobcats",
		},
		Animal{
			Name:       "cheetah",
			PluralName: "cheetahs",
		},
		Animal{
			Name:       "cougar",
			PluralName: "cougars",
		},
		Animal{
			Name:       "jaguar",
			PluralName: "jaguars",
		},
		Animal{
			Name:       "leopard",
			PluralName: "leopards",
		},
		Animal{
			Name:       "lion",
			PluralName: "lions",
		},
		Animal{
			Name:       "tiger",
			PluralName: "tigers",
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
		a.MinTemperature = 6
		a.MaxTemperature = 10
		a.Size = size.GetCategoryByName("medium")
		a.Resources = []resource.Resource{
			resource.Resource{
				Name:        a.Name + " hide",
				Origin:      a.Name,
				Type:        "hide",
				Commonality: 5,
			},
			resource.Resource{
				Name:        a.Name + " fangs",
				Origin:      a.Name,
				Type:        "teeth",
				Commonality: 3,
			},
			resource.Resource{
				Name:        a.Name,
				Origin:      a.Name,
				Type:        "meat",
				Commonality: 5,
			},
		}
	}

	return animals
}
