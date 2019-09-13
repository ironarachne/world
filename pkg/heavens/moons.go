package heavens

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Moon is a moon
type Moon struct {
	Name  string
	Color string
	Shape string
	Size  int
}

func getRandomMoonColor() (string, error) {
	colors := map[string]int{
		"white": 10,
		"grey":  5,
		"blue":  1,
		"red":   1,
	}

	color, err := random.StringFromThresholdMap(colors)
	if err != nil {
		err = fmt.Errorf("Failed to generate random moon color: %w", err)
		return "", err
	}

	return color, nil
}

func getRandomMoonShape() (string, error) {
	shapes := map[string]int{
		"round":  10,
		"oblong": 2,
	}

	shape, err := random.StringFromThresholdMap(shapes)
	if err != nil {
		err = fmt.Errorf("Failed to generate random moon shape: %w", err)
		return "", err
	}

	return shape, nil
}

func getRandomMoonSize() int {
	return rand.Intn(5) + rand.Intn(3)
}

func getRandomMoon() (Moon, error) {
	moon := Moon{}
	moon.Name = "moon"
	color, err := getRandomMoonColor()
	if err != nil {
		err = fmt.Errorf("Failed to generate random moon: %w", err)
		return Moon{}, err
	}
	moon.Color = color
	shape, err := getRandomMoonShape()
	if err != nil {
		err = fmt.Errorf("Failed to generate random moon: %w", err)
		return Moon{}, err
	}
	moon.Shape = shape
	moon.Size = getRandomMoonSize()

	return moon, nil
}

func getRandomMoons() ([]Moon, error) {
	moons := []Moon{}
	numberOfMoons := 1
	shallWeHaveMoreThanOneMoon := rand.Intn(10) + 1
	if shallWeHaveMoreThanOneMoon > 8 {
		numberOfMoons += rand.Intn(5)
	}

	for i := 0; i < numberOfMoons; i++ {
		moon, err := getRandomMoon()
		if err != nil {
			err = fmt.Errorf("Failed to generate random moon: %w", err)
			return []Moon{}, err
		}
		moons = append(moons, moon)
	}

	return moons, nil
}
