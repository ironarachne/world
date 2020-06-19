package religion

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func randomReligionName(ctx context.Context) (string, error) {
	possibleNames := []string{
		"The Way",
		"The Path",
		"The Truth",
	}

	name, err := random.String(ctx, possibleNames)
	if err != nil {
		err = fmt.Errorf("could not generate religion name: %w", err)
		return "", err
	}

	return name, nil
}
