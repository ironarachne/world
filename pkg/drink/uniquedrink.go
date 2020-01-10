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
		err = fmt.Errorf("could not generate unique drink pattern: %w", err)
		return resource.Pattern{}, err
	}
	method := getRandomMethod()

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

func getRandomMethod() Method {
	drinkMethods := []Method{
		{
			Name:            "brewed",
			BaseResourceTag: "grain",
			Producer:        "brewer",
		},
		{
			Name:            "distilled",
			BaseResourceTag: "grain",
			Producer:        "distiller",
		},
		{
			Name:            "fermented",
			BaseResourceTag: "fruit",
			Producer:        "vintner",
		},
	}

	method := drinkMethods[rand.Intn(len(drinkMethods))]

	return method
}
