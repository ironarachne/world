package drink

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/resource"
)

// Method is a way of making an alcoholic beverage
type Method struct {
	Name            string `json:"name"`
	BaseResourceTag string `json:"base_resource_tag"`
	Producer        string `json:"producer"`
}

func generateUniqueDrinkPattern(lang language.Language, resources []resource.Resource) (resource.Pattern, error) {
	name, err := lang.NewWord()
	if err != nil {
		err = fmt.Errorf("failed to generate unique drink pattern: %w", err)
		return resource.Pattern{}, err
	}
	method := getRandomMethod(resources)

	filteredResources := resource.ByTag(method.BaseResourceTag, resources)

	baseResource := resource.Random(filteredResources)

	producer, _ := profession.ByName(method.Producer)

	pattern := resource.Pattern{
		Name:         name,
		NameTemplate: name,
		Description:  "a beverage called " + name + ", which is " + method.Name + " from " + baseResource.Name,
		Tags: []string{
			"alcohol",
			"beverage",
		},
		Profession: producer,
		Slots: []resource.Slot{
			{
				Name:                "body",
				RequiredTag:         baseResource.Name,
				DescriptionTemplate: name,
			},
		},
		Value: 1,
	}

	return pattern, nil
}

func getRandomMethod(resources []resource.Resource) Method {
	var drinkMethods []Method

	grainMethods := []Method{
		{
			Name:            "fermented",
			BaseResourceTag: "grain",
			Producer:        "brewer",
		},
		{
			Name:            "distilled",
			BaseResourceTag: "grain",
			Producer:        "distiller",
		},
	}

	fruitMethods := []Method{
		{
			Name:            "fermented",
			BaseResourceTag: "fruit",
			Producer:        "vintner",
		},
	}

	milkMethods := []Method{
		{
			Name:            "fermented",
			BaseResourceTag: "milk",
			Producer:        "distiller",
		},
	}

	grains := resource.ByTag("grain", resources)
	fruit := resource.ByTag("fruit", resources)
	milk := resource.ByTag("milk", resources)

	if len(grains) > 0 {
		drinkMethods = append(drinkMethods, grainMethods...)
	}

	if len(fruit) > 0 {
		drinkMethods = append(drinkMethods, fruitMethods...)
	}

	if len(milk) > 0 {
		drinkMethods = append(drinkMethods, milkMethods...)
	}

	if len(drinkMethods) == 0 {
		panic("no resources for making unique beverages!")
	}

	if len(drinkMethods) == 1 {
		return drinkMethods[0]
	}

	method := drinkMethods[rand.Intn(len(drinkMethods))]

	return method
}
