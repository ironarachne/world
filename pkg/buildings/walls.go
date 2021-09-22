package buildings

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getWallDescriptors() []string {
	return []string{
		"interspersed with narrow columns",
		"decorated with strips of metal",
		"painted with bright colors",
		"that's meticulously cleaned",
		"covered over with plaster",
		"with small decorative arches near the roof",
		"with large decorative arches",
	}
}

func getWallMaterials() []string {
	return []string{
		"carved stone",
		"heavy wooden boards",
		"insulated wood",
		"narrow bricks",
		"stone slabs",
		"thick bricks",
		"wattle and daub",
		"wide bricks",
		"wooden slats",
	}
}

func getRandomWallStyle(ctx context.Context) (string, error) {
	materials := getWallMaterials()
	descriptors := getWallDescriptors()

	style, err := random.String(ctx, materials)
	if err != nil {
		err = fmt.Errorf("Could not generate wall style: %w", err)
		return "", err
	}
	descriptor, err := random.String(ctx, descriptors)
	if err != nil {
		err = fmt.Errorf("Could not generate wall style: %w", err)
		return "", err
	}
	style += " " + descriptor

	return style, nil
}
