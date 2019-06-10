package buildings

import (
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
func GenerateStyle() BuildingStyle {
	style := BuildingStyle{
		Decorations: getRandomDecorations(),
		DoorStyle:   getRandomDoorStyle(),
		MaxStories:  getRandomMaxStories(),
		RoofStyle:   getRandomRoofStyle(),
		StreetStyle: getRandomStreetStyle(),
		WallStyle:   getRandomWallStyle(),
		WindowStyle: getRandomWindowStyle(),
	}

	return style
}

func getRandomMaxStories() int {
	stories := map[int]int{
		1: 20,
		2: 5,
		3: 1,
	}

	return random.IntFromThresholdMap(stories)
}
