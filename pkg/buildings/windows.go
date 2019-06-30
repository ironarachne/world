package buildings

import (
	"github.com/ironarachne/world/pkg/random"
)

func getRandomWindowDescriptor() string {
	descriptors := []string{
		"and narrow",
		"and ornate",
		"and oval-shaped",
		"and rectangular",
		"and round",
		"and square",
		"with stained glass",
		"with thick-paned glass",
		"with thin-paned glass",
		"and triangular",
		"and wide",
	}

	return random.String(descriptors)
}

func getRandomWindowSize() string {
	sizes := []string{
		"large",
		"medium",
		"small",
	}

	return random.String(sizes)
}

func getRandomWindowStyle() string {
	description := getRandomWindowSize()

	description += " " + getRandomWindowDescriptor()

	return description
}
