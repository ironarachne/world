package heavens

import (
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

func getRandomMoonColor() string {
	colors := map[string]int{
		"white": 10,
		"grey":  5,
		"blue":  1,
		"red":   1,
	}

	return random.StringFromThresholdMap(colors)
}

func getRandomMoonShape() string {
	shapes := map[string]int{
		"round":  10,
		"oblong": 2,
	}

	return random.StringFromThresholdMap(shapes)
}

func getRandomMoonSize() int {
	return rand.Intn(5) + rand.Intn(3)
}

func getRandomMoon() Moon {
	moon := Moon{}
	moon.Name = "moon"
	moon.Color = getRandomMoonColor()
	moon.Shape = getRandomMoonShape()
	moon.Size = getRandomMoonSize()

	return moon
}

func getRandomMoons() []Moon {
	moons := []Moon{}
	numberOfMoons := 1
	shallWeHaveMoreThanOneMoon := rand.Intn(10) + 1
	if shallWeHaveMoreThanOneMoon > 8 {
		numberOfMoons += rand.Intn(5)
	}

	for i := 0; i < numberOfMoons; i++ {
		moons = append(moons, getRandomMoon())
	}

	return moons
}
