package climate

import (
	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/plant"
)

// Resource is a refined resource derived from raw materials
type Resource struct {
	Name   string
	Origin string
	Type   string
}

// IsTypeInResources checks to see if a given type is present in a collection of resources
func IsTypeInResources(typeName string, resources []Resource) bool {
	for _, i := range resources {
		if i.Type == typeName {
			return true
		}
	}

	return false
}

// ListResourcesOfType returns all resources in the given collection that match the type
func ListResourcesOfType(typeName string, resources []Resource) []Resource {
	filteredResources := []Resource{}

	for _, i := range resources {
		if i.Type == typeName {
			filteredResources = append(filteredResources, i)
		}
	}

	return filteredResources
}

func resourcesFromAnimal(source animal.Animal) []Resource {
	resources := []Resource{}

	if source.GivesBone {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "bone"})
	}
	if source.GivesEggs {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "eggs"})
	}
	if source.GivesHide {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "hide"})
	}
	if source.GivesHorn {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "horn"})
	}
	if source.GivesMeat {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "meat"})
	}
	if source.GivesMilk {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "milk"})
	}
	if source.IsMount {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "mount"})
	}
	if source.IsPackAnimal {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "pack animal"})
	}
	if source.IsVenomous {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "toxin"})
	}

	return resources
}

func resourcesFromFish(source animal.Fish) []Resource {
	resources := []Resource{}

	resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "meat"})

	return resources
}

func resourcesFromMineral(source Mineral) []Resource {
	resources := []Resource{}

	if source.HasOre {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "ore"})
	}
	if source.IsGem {
		resources = append(resources, Resource{Name: source.PluralName, Origin: source.Name, Type: "gem"})
	}
	if source.IsMetal {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "metal bar"})
	}
	if source.IsPrecious && source.IsMetal {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "metal ingot"})
	}
	if source.IsStone {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "stone block"})
	}
	if source.IsEdible {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "spice"})
	}

	return resources
}

func resourcesFromPlant(source plant.Plant) []Resource {
	resources := []Resource{}

	if source.IsFiber {
		if source.Name == "flax" {
			resources = append(resources, Resource{Name: "linen", Origin: source.Name, Type: "fabric"})
		} else {
			resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "fabric"})
		}
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "rope"})
	}
	if source.IsFruit {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "fruit"})
	}
	if source.IsGrain {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "grain"})
	}
	if source.IsHerb {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "herb"})
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "spice"})
	}
	if source.IsMedicine {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "medicine"})
	}
	if source.IsNut {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "nut"})
	}
	if source.IsToxic {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "toxin"})
	}
	if source.IsVegetable {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "vegetable"})
	}

	return resources
}

func resourcesFromSoil(source Soil) []Resource {
	resources := []Resource{}

	if source.UsedForAgriculture {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "soil"})
	}
	if source.UsedForPottery {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "clay"})
	}
	if source.IsSand {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "sand"})
	}

	return resources
}

func resourcesFromTree(source plant.Tree) []Resource {
	resources := []Resource{}

	resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "wood"})

	if source.GivesFruit {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "fruit"})
	}
	if source.GivesNuts {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "nut"})
	}
	if source.IsSpice {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "spice"})
	}

	return resources
}
