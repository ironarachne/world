package heavens

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Sun is a sun
type Sun struct {
	Name       string
	Brightness int
	Color      string
	Size       int
}

func getRandomSunBrightness() int {
	return rand.Intn(5) + rand.Intn(5)
}

func getRandomSunColor() (string, error) {
	colors := map[string]int{
		"white":  1,
		"yellow": 35,
		"blue":   1,
		"red":    2,
		"orange": 2,
		"green":  1,
	}

	color, err := random.StringFromThresholdMap(colors)
	if err != nil {
		err = fmt.Errorf("Failed to generate random sun color: %w", err)
		return "", err
	}

	return color, nil
}

func getRandomSunSize() int {
	return rand.Intn(5) + rand.Intn(5)
}

func getRandomSun() (Sun, error) {
	sun := Sun{}
	sun.Name = "sun"
	color, err := getRandomSunColor()
	if err != nil {
		err = fmt.Errorf("Failed to generate random sun: %w", err)
		return Sun{}, err
	}
	sun.Color = color
	sun.Brightness = getRandomSunBrightness()
	sun.Size = getRandomSunSize()

	return sun, nil
}

func getRandomSuns() ([]Sun, error) {
	suns := []Sun{}
	numberOfSuns := 1
	shallWeHaveMoreThanOneSun := rand.Intn(10) + 1
	if shallWeHaveMoreThanOneSun > 9 {
		numberOfSuns += rand.Intn(2)
	}

	for i := 0; i < numberOfSuns; i++ {
		sun, err := getRandomSun()
		if err != nil {
			err = fmt.Errorf("Failed to generate random suns: %w", err)
			return []Sun{}, err
		}
		suns = append(suns, sun)
	}

	return suns, nil
}
