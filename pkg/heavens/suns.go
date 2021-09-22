package heavens

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

// Sun is a sun
type Sun struct {
	Name       string
	Brightness int
	Color      string
	Size       int
}

func getRandomSunBrightness(ctx context.Context) int {
	return random.Intn(ctx, 5) + random.Intn(ctx, 5)
}

func getRandomSunColor(ctx context.Context) (string, error) {
	colors := map[string]int{
		"white":  1,
		"yellow": 35,
		"blue":   1,
		"red":    2,
		"orange": 2,
		"green":  1,
	}

	color, err := random.StringFromThresholdMap(ctx, colors)
	if err != nil {
		err = fmt.Errorf("Failed to generate random sun color: %w", err)
		return "", err
	}

	return color, nil
}

func getRandomSunSize(ctx context.Context) int {
	return random.Intn(ctx, 5) + random.Intn(ctx, 5)
}

func getRandomSun(ctx context.Context) (Sun, error) {
	sun := Sun{}
	sun.Name = "sun"
	color, err := getRandomSunColor(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate random sun: %w", err)
		return Sun{}, err
	}
	sun.Color = color
	sun.Brightness = getRandomSunBrightness(ctx)
	sun.Size = getRandomSunSize(ctx)

	return sun, nil
}

func getRandomSuns(ctx context.Context) ([]Sun, error) {
	suns := []Sun{}
	numberOfSuns := 1
	shallWeHaveMoreThanOneSun := random.Intn(ctx, 10) + 1
	if shallWeHaveMoreThanOneSun > 9 {
		numberOfSuns += random.Intn(ctx, 2)
	}

	for i := 0; i < numberOfSuns; i++ {
		sun, err := getRandomSun(ctx)
		if err != nil {
			err = fmt.Errorf("Failed to generate random suns: %w", err)
			return []Sun{}, err
		}
		suns = append(suns, sun)
	}

	return suns, nil
}
