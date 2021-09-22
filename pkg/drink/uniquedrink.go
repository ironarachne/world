package drink

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pattern"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

// Method is a way of making an alcoholic beverage
type Method struct {
	Name            string `json:"name"`
	BaseResourceTag string `json:"base_resource_tag"`
	Producer        string `json:"producer"`
}

func generateUniqueDrinkPattern(ctx context.Context, lang language.Language, resources []resource.Resource) (pattern.Pattern, error) {
	name, err := lang.NewWord(ctx)
	if err != nil {
		err = fmt.Errorf("failed to generate unique drink pattern: %w", err)
		return pattern.Pattern{}, err
	}
	method, err := getRandomMethod(ctx, resources)
	if err != nil {
		err = fmt.Errorf("failed to generate unique drink pattern: %w", err)
		return pattern.Pattern{}, err
	}

	filteredResources := resource.ByTag(method.BaseResourceTag, resources)

	baseResource := resource.Random(ctx, filteredResources)

	producer, _ := profession.ByName(method.Producer)

	pattern := pattern.Pattern{
		Name:         name,
		NameTemplate: name,
		Description:  "a beverage called " + name + ", which is " + method.Name + " from " + baseResource.Name,
		Tags: []string{
			"alcohol",
			"beverage",
		},
		ProfessionName: producer.Name,
		Slots: []pattern.Slot{
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

func getRandomMethod(ctx context.Context, resources []resource.Resource) (Method, error) {
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
		err := fmt.Errorf("failed to find drink production method matching given resources")
		return Method{}, err
	}

	if len(drinkMethods) == 1 {
		return drinkMethods[0], nil
	}

	method := drinkMethods[random.Intn(ctx, len(drinkMethods))]

	return method, nil
}
