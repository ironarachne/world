/*
Package buildings provides structures and algorithms for producing
random architectural styles and individual building appearances.
*/
package buildings

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

const buildingError = "could not generate building style: %w"

// BuildingStyle is a style of building
type BuildingStyle struct {
	Description string `json:"description"`
	Decorations string `json:"decorations"`
	DoorStyle   string `json:"door_style"`
	MaxStories  int    `json:"max_stories"`
	RoofStyle   string `json:"roof_style"`
	StreetStyle string `json:"street_style"`
	WallStyle   string `json:"wall_style"`
	WindowStyle string `json:"window_style"`
}

// GenerateStyle generates a random building style
func GenerateStyle(ctx context.Context) (BuildingStyle, error) {
	decorations, err := getRandomDecorations(ctx)
	if err != nil {
		err = fmt.Errorf(buildingError, err)
		return BuildingStyle{}, err
	}
	doorStyle, err := getRandomDoorStyle(ctx)
	if err != nil {
		err = fmt.Errorf(buildingError, err)
		return BuildingStyle{}, err
	}
	maxStories, err := getRandomMaxStories(ctx)
	if err != nil {
		err = fmt.Errorf(buildingError, err)
		return BuildingStyle{}, err
	}
	streetStyle, err := getRandomStreetStyle(ctx)
	if err != nil {
		err = fmt.Errorf(buildingError, err)
		return BuildingStyle{}, err
	}
	wallStyle, err := getRandomWallStyle(ctx)
	if err != nil {
		err = fmt.Errorf(buildingError, err)
		return BuildingStyle{}, err
	}
	windowStyle, err := getRandomWindowStyle(ctx)
	if err != nil {
		err = fmt.Errorf(buildingError, err)
		return BuildingStyle{}, err
	}
	roofStyle, err := getRandomRoofStyle(ctx)
	if err != nil {
		err = fmt.Errorf(buildingError, err)
		return BuildingStyle{}, err
	}
	style := BuildingStyle{
		Decorations: decorations,
		DoorStyle:   doorStyle,
		MaxStories:  maxStories,
		RoofStyle:   roofStyle,
		StreetStyle: streetStyle,
		WallStyle:   wallStyle,
		WindowStyle: windowStyle,
	}
	style.Description = style.Describe()

	return style, nil
}

func getRandomMaxStories(ctx context.Context) (int, error) {
	stories := map[int]int{
		1: 20,
		2: 5,
		3: 1,
	}

	storyCount, err := random.IntFromThresholdMap(ctx, stories)
	if err != nil {
		err = fmt.Errorf("Failed to get random max stories: %w", err)
		return 0, err
	}

	return storyCount, nil
}
