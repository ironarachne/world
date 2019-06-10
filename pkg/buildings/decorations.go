package buildings

import (
	"github.com/ironarachne/world/pkg/random"
)

func getRandomDecorations() string {
	decorations := []string{
		"statues stand outside doorways",
		"signs bearing the occupants' names stand out front",
		"geometric shapes are carved into the trim",
	}

	return random.String(decorations)
}
