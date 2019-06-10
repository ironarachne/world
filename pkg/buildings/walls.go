package buildings

import (
	"github.com/ironarachne/world/pkg/random"
)

func getRandomWallStyle() string {
	styles := []string{
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

	return random.String(styles)
}
