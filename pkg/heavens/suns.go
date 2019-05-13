package heavens

import (
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

func getRandomSunColor() string {
	colors := map[string]int{
		"white":  1,
		"yellow": 35,
		"blue":   1,
		"red":    2,
		"orange": 2,
		"green":  1,
	}

	return random.StringFromThresholdMap(colors)
}

func getRandomSunSize() int {
	return rand.Intn(5) + rand.Intn(5)
}

func getRandomSun() Sun {
	sun := Sun{}
	sun.Name = "sun"
	sun.Color = getRandomSunColor()
	sun.Brightness = getRandomSunBrightness()
	sun.Size = getRandomSunSize()

	return sun
}

func getRandomSuns() []Sun {
	suns := []Sun{}
	numberOfSuns := 1
	shallWeHaveMoreThanOneSun := rand.Intn(10) + 1
	if shallWeHaveMoreThanOneSun > 9 {
		numberOfSuns += rand.Intn(2)
	}

	for i := 0; i < numberOfSuns; i++ {
		suns = append(suns, getRandomSun())
	}

	return suns
}
