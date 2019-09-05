package buildings

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getRandomWindowDescriptor() (string, error) {
	descriptors := []string{
		"and narrow",
		"and ornate",
		"and oval-shaped",
		"and rectangular",
		"and round",
		"and square",
		"with stained glass",
		"with thick-paned glass",
		"with thin-paned glass",
		"and triangular",
		"and wide",
	}

	descriptor, err := random.String(descriptors)
	if err != nil {
		err = fmt.Errorf("Could not get window descriptor: %w", err)
		return "", err
	}

	return descriptor, nil
}

func getRandomWindowSize() (string, error) {
	sizes := []string{
		"large",
		"medium",
		"small",
	}

	size, err := random.String(sizes)

	if err != nil {
		err = fmt.Errorf("Could not get window size: %w", err)
		return "", err
	}

	return size, nil
}

func getRandomWindowStyle() (string, error) {
	description, err := getRandomWindowSize()
	if err != nil {
		err = fmt.Errorf("Could not generate window style: %w", err)
		return "", err
	}
	descriptor, err := getRandomWindowDescriptor()
	if err != nil {
		err = fmt.Errorf("Could not generate window style: %w", err)
		return "", err
	}

	description += " " + descriptor

	return description, nil
}
