package heavens

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

const moonError = "failed to generate random moon: %w"

// Moon is a moon
type Moon struct {
	Name  string
	Color string
	Shape string
	Size  int
}

func getRandomMoonColor(ctx context.Context) (string, error) {
	colors := map[string]int{
		"white": 10,
		"grey":  5,
		"blue":  1,
		"red":   1,
	}

	color, err := random.StringFromThresholdMap(ctx, colors)
	if err != nil {
		err = fmt.Errorf("failed to generate random moon color: %w", err)
		return "", err
	}

	return color, nil
}

func getRandomMoonShape(ctx context.Context) (string, error) {
	shapes := map[string]int{
		"round":  10,
		"oblong": 2,
	}

	shape, err := random.StringFromThresholdMap(ctx, shapes)
	if err != nil {
		err = fmt.Errorf("failed to generate random moon shape: %w", err)
		return "", err
	}

	return shape, nil
}

func getRandomMoonSize(ctx context.Context) int {
	return random.Intn(ctx, 5) + random.Intn(ctx, 3)
}

func getRandomMoon(ctx context.Context) (Moon, error) {
	moon := Moon{}
	moon.Name = "moon"
	color, err := getRandomMoonColor(ctx)
	if err != nil {
		err = fmt.Errorf(moonError, err)
		return Moon{}, err
	}
	moon.Color = color
	shape, err := getRandomMoonShape(ctx)
	if err != nil {
		err = fmt.Errorf(moonError, err)
		return Moon{}, err
	}
	moon.Shape = shape
	moon.Size = getRandomMoonSize(ctx)

	return moon, nil
}

func getRandomMoons(ctx context.Context) ([]Moon, error) {
	moons := []Moon{}
	numberOfMoons := 1
	shallWeHaveMoreThanOneMoon := random.Intn(ctx, 10) + 1
	if shallWeHaveMoreThanOneMoon > 8 {
		numberOfMoons += random.Intn(ctx, 5)
	}

	for i := 0; i < numberOfMoons; i++ {
		moon, err := getRandomMoon(ctx)
		if err != nil {
			err = fmt.Errorf(moonError, err)
			return []Moon{}, err
		}
		moons = append(moons, moon)
	}

	return moons, nil
}
