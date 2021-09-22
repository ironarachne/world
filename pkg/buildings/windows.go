package buildings

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getRandomWindowDescriptor(ctx context.Context) (string, error) {
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

	descriptor, err := random.String(ctx, descriptors)
	if err != nil {
		err = fmt.Errorf("Could not get window descriptor: %w", err)
		return "", err
	}

	return descriptor, nil
}

func getRandomWindowSize(ctx context.Context) (string, error) {
	sizes := []string{
		"large",
		"medium",
		"small",
	}

	size, err := random.String(ctx, sizes)

	if err != nil {
		err = fmt.Errorf("Could not get window size: %w", err)
		return "", err
	}

	return size, nil
}

func getRandomWindowStyle(ctx context.Context) (string, error) {
	description, err := getRandomWindowSize(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate window style: %w", err)
		return "", err
	}
	descriptor, err := getRandomWindowDescriptor(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate window style: %w", err)
		return "", err
	}

	description += " " + descriptor

	return description, nil
}
