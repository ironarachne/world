package heavens

import (
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

func getRandomStarColor() string {
	colors := map[string]int{
		"blue":   1,
		"green":  1,
		"orange": 1,
		"red":    2,
		"white":  10,
		"yellow": 3,
	}

	color := random.StringFromThresholdMap(colors)

	return color
}

func getRandomStar() Star {
	star := Star{}
	star.Name = "star"
	star.Color = getRandomStarColor()
	star.Brightness = getRandomStarBrightness()

	return star
}

func getRandomStars() []Star {
	stars := []Star{}
	numberOfStars := rand.Intn(16)

	for i := 0; i < numberOfStars; i++ {
		stars = append(stars, getRandomStar())
	}

	return stars
}
