package music

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Style is a music style
type Style struct {
	Structure   int
	Vocals      int
	Beat        int
	Tonality    int
	Descriptors []string
	Instruments []Instrument
}

func getAllStyleDescriptors() []string {
	descriptors := []string{
		"airy",
		"bombastic",
		"booming",
		"breathy",
		"bright",
		"cheerful",
		"driving",
		"dynamic",
		"energetic",
		"ethereal",
		"euphonic",
		"fast",
		"full-toned",
		"haunting",
		"lilting",
		"lofty",
		"mellifluous",
		"mellow",
		"melodic",
		"moody",
		"operatic",
		"orotund",
		"percussive",
		"powerful",
		"primitive",
		"regimented",
		"resonant",
		"rigid",
		"savage",
		"somber",
		"structured",
		"tumid",
		"uplifting",
		"vibrant",
		"warm",
	}

	return descriptors
}

func getRandomStyleDescriptors(maxDescriptors int) []string {
	var randomDescriptor string
	descriptors := []string{}

	possibleDescriptors := getAllStyleDescriptors()

	numberOfDescriptors := rand.Intn(maxDescriptors) + 1

	for i := 0; i < numberOfDescriptors; i++ {
		randomDescriptor = random.String(possibleDescriptors)
		if !slices.StringIn(randomDescriptor, descriptors) {
			descriptors = append(descriptors, randomDescriptor)
		} else {
			i--
		}
	}

	return descriptors
}

// GenerateStyle generates a random musical style given a set of instruments
func GenerateStyle(instruments []Instrument) Style {
	style := Style{}

	style.Beat = rand.Intn(3)
	style.Structure = rand.Intn(3)
	style.Tonality = rand.Intn(3)
	style.Vocals = rand.Intn(3)

	style.Descriptors = getRandomStyleDescriptors(4)
	style.Instruments = instruments

	return style
}
