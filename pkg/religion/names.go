package religion

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func randomReligionName() (string, error) {
	possibleNames := []string{
		"The Way",
		"The Path",
		"The Truth",
	}

	name, err := random.String(possibleNames)
	if err != nil {
		err = fmt.Errorf("could not generate religion name: %w", err)
		return "", err
	}

	return name, nil
}
