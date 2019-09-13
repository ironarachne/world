package buildings

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

// BuildingStyle is a style of building
type BuildingStyle struct {
	Decorations string
	DoorStyle   string
	MaxStories  int
	RoofStyle   string
	StreetStyle string
	WallStyle   string
	WindowStyle string
}

// GenerateStyle generates a random building style
func GenerateStyle() (BuildingStyle, error) {
	decorations, err := getRandomDecorations()
	if err != nil {
		err = fmt.Errorf("Could not generate building style: %w", err)
		return BuildingStyle{}, err
	}
	doorStyle, err := getRandomDoorStyle()
	if err != nil {
		err = fmt.Errorf("Could not generate building style: %w", err)
		return BuildingStyle{}, err
	}
	streetStyle, err := getRandomStreetStyle()
	if err != nil {
		err = fmt.Errorf("Could not generate building style: %w", err)
		return BuildingStyle{}, err
	}
	wallStyle, err := getRandomWallStyle()
	if err != nil {
		err = fmt.Errorf("Could not generate building style: %w", err)
		return BuildingStyle{}, err
	}
	windowStyle, err := getRandomWindowStyle()
	if err != nil {
		err = fmt.Errorf("Could not generate building style: %w", err)
		return BuildingStyle{}, err
	}
	roofStyle, err := getRandomRoofStyle()
	if err != nil {
		err = fmt.Errorf("Could not generate building style: %w", err)
		return BuildingStyle{}, err
	}
	style := BuildingStyle{
		Decorations: decorations,
		DoorStyle:   doorStyle,
		MaxStories:  getRandomMaxStories(),
		RoofStyle:   roofStyle,
		StreetStyle: streetStyle,
		WallStyle:   wallStyle,
		WindowStyle: windowStyle,
	}

	return style, nil
}

func getRandomMaxStories() int {
	stories := map[int]int{
		1: 20,
		2: 5,
		3: 1,
	}

	storyCount := random.IntFromThresholdMap(stories)

	return storyCount
}
