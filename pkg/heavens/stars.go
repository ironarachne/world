package heavens

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

// Star is a star in the sky
type Star struct {
	Name       string
	Color      string
	Brightness int
}

func getRandomStarBrightness(ctx context.Context) int {
	brightness := random.Intn(ctx, 5) + random.Intn(ctx, 5)

	return brightness
}

func getRandomStarColor(ctx context.Context) (string, error) {
	colors := map[string]int{
		"blue":   1,
		"green":  1,
		"orange": 1,
		"red":    2,
		"white":  10,
		"yellow": 3,
	}

	color, err := random.StringFromThresholdMap(ctx, colors)
	if err != nil {
		err = fmt.Errorf("Failed to generate random star color: %w", err)
		return "", err
	}

	return color, nil
}

func getRandomStar(ctx context.Context) (Star, error) {
	star := Star{}
	star.Name = "star"
	color, err := getRandomStarColor(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate random star: %w", err)
		return Star{}, err
	}
	star.Color = color
	star.Brightness = getRandomStarBrightness(ctx)

	return star, nil
}

func getRandomStars(ctx context.Context) ([]Star, error) {
	stars := []Star{}
	numberOfStars := random.Intn(ctx, 16)

	for i := 0; i < numberOfStars; i++ {
		star, err := getRandomStar(ctx)
		if err != nil {
			err = fmt.Errorf("Failed to generate random stars: %w", err)
			return []Star{}, err
		}
		stars = append(stars, star)
	}

	return stars, nil
}
