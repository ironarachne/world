package buildings

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getRandomStreetStyle(ctx context.Context) (string, error) {
	styles := []string{
		"narrow",
		"wide",
	}

	style, err := random.String(ctx, styles)
	if err != nil {
		err = fmt.Errorf("Could not find street style: %w", err)
		return "", err
	}

	return style, nil
}
