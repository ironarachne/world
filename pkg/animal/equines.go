package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getEquines() []Animal {
	animals := []Animal{
		Animal{
			Name:       "horse",
			PluralName: "horses",
		},
		Animal{
			Name:       "donkey",
			PluralName: "donkeys",
		},
		Animal{
			Name:       "zebra",
			PluralName: "zebras",
		},
	}

	for _, a := range animals {
		a.AnimalType = "mammal"
		a.IsMount = true
		a.IsPackAnimal = true
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
				Commonality: 3,
			},
			resource.Resource{
				Name:        a.Name + " milk",
				Origin:      a.Name,
				Type:        "milk",
				Commonality: 3,
			},
			resource.Resource{
				Name:        a.Name,
				Origin:      a.Name,
				Type:        "meat",
				Commonality: 3,
			},
			resource.Resource{
				Name:        a.Name,
				Origin:      a.Name,
				Type:        "mount",
				Commonality: 7,
			},
			resource.Resource{
				Name:        a.Name,
				Origin:      a.Name,
				Type:        "pack animal",
				Commonality: 5,
			},
		}
	}

	return animals
}
