package heavens

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Star is a star in the sky
type Star struct {
	Name       string
	Color      string
	Brightness int
}

func getRandomStarBrightness() int {
	brightness := rand.Intn(5) + rand.Intn(5)

	return brightness
}

func getRandomStarColor() (string, error) {
	colors := map[string]int{
		"blue":   1,
		"green":  1,
		"orange": 1,
		"red":    2,
		"white":  10,
		"yellow": 3,
	}

	color, err := random.StringFromThresholdMap(colors)
	if err != nil {
		err = fmt.Errorf("Failed to generate random star color: %w", err)
		return "", err
	}

	return color, nil
}

func getRandomStar() (Star, error) {
	star := Star{}
	star.Name = "star"
	color, err := getRandomStarColor()
	if err != nil {
		err = fmt.Errorf("Failed to generate random star: %w", err)
		return Star{}, err
	}
	star.Color = color
	star.Brightness = getRandomStarBrightness()

	return star, nil
}

func getRandomStars() ([]Star, error) {
	stars := []Star{}
	numberOfStars := rand.Intn(16)

	for i := 0; i < numberOfStars; i++ {
		star, err := getRandomStar()
		if err != nil {
			err = fmt.Errorf("Failed to generate random stars: %w", err)
			return []Star{}, err
		}
		stars = append(stars, star)
	}

	return stars, nil
}
