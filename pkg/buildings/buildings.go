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
	maxStories, err := getRandomMaxStories()
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
		MaxStories:  maxStories,
		RoofStyle:   roofStyle,
		StreetStyle: streetStyle,
		WallStyle:   wallStyle,
		WindowStyle: windowStyle,
	}

	return style, nil
}

func getRandomMaxStories() (int, error) {
	stories := map[int]int{
		1: 20,
		2: 5,
		3: 1,
	}

	storyCount, err := random.IntFromThresholdMap(stories)
	if err != nil {
		err = fmt.Errorf("Failed to get random max stories: %w", err)
		return 0, err
	}

	return storyCount, nil
}
