package food

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

func generateBread(ctx context.Context, resources []resource.Resource) (string, error) {
	var grain resource.Resource
	breadTypes := []string{
		"brick-like",
		"flat",
		"long",
		"round",
		"spongy",
		"thin",
	}

	flavors := []string{
		"bitter",
		"grainy",
		"hearty",
		"nutty",
		"savory",
		"sweet",
	}
	grains := resource.ByTag("flour", resources)

	if len(grains) == 0 {
		err := fmt.Errorf("Could not generate bread: no grains available")
		return "", err
	}

	if len(grains) == 1 {
		grain = grains[0]
	} else {
		grain = grains[random.Intn(ctx, len(grains))]
	}

	flavor, err := random.String(ctx, flavors)
	if err != nil {
		err = fmt.Errorf("Could not generate bread: %w", err)
		return "", err
	}
	breadType, err := random.String(ctx, breadTypes)
	if err != nil {
		err = fmt.Errorf("Could not generate bread: %w", err)
		return "", err
	}

	bread := flavor + " " + breadType + " " + grain.Name + " bread"

	return bread, nil
}

func randomBreads(ctx context.Context, resources []resource.Resource) ([]string, error) {
	var bread string
	var breads []string
	var err error

	grains := resource.ByTag("flour", resources)
	if len(grains) > 0 {
		numberOfBreads := random.Intn(ctx, 3) + 1
		for i := 0; i < numberOfBreads; i++ {
			bread, err = generateBread(ctx, resources)
			if err != nil {
				err = fmt.Errorf("Could not generate breads: %w", err)
				return []string{}, err
			}
			if !slices.StringIn(bread, breads) {
				breads = append(breads, bread)
			}
		}
	}

	return breads, nil
}
