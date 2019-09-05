package buildings

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getRandomStreetStyle() (string, error) {
	styles := []string{
		"narrow",
		"wide",
	}

	style, err := random.String(styles)
	if err != nil {
		err = fmt.Errorf("Could not find street style: %w", err)
		return "", err
	}

	return style, nil
}
