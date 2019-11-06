package religion

import (
	"fmt"

	"github.com/ironarachne/world/pkg/conlang"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/random"
)

// Religion is a religion
type Religion struct {
	Name           string
	CommonName     string
	Class          Class
	Pantheon       pantheon.Pantheon
	GatheringPlace string
	Virtues        []string
}

// Generate procedurally generates a religion
func Generate(originLanguage language.Language) (Religion, error) {
	religion := Religion{}

	class, err := getWeightedClass()
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}
	religion.Class = class
	gatheringPlace, err := religion.randomGatheringPlace()
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}
	religion.GatheringPlace = gatheringPlace

	if religion.Class.PantheonMaxSize > 0 {
		newPantheon, err := pantheon.Generate(religion.Class.PantheonMinSize, religion.Class.PantheonMaxSize, originLanguage)
		if err != nil {
			err = fmt.Errorf("Could not generate religion: %w", err)
			return Religion{}, err
		}
		religion.Pantheon = newPantheon
	}

	religionName, err := randomReligionName()
	if err != nil {
		err = fmt.Errorf("Could not generate religion: %w", err)
		return Religion{}, err
	}
	name := originLanguage.TranslatePhrase(religionName)
	religion.CommonName = religionName
	religion.Name = name

	virtues, err := getRandomVirtues()
	if err != nil {
		err = fmt.Errorf("Could not generate religion: %w", err)
		return Religion{}, err
	}
	religion.Virtues = virtues

	return religion, nil
}

func (religion Religion) randomGatheringPlace() (string, error) {
	gatheringPlace, err := random.String(religion.Class.GatheringPlaces)
	if err != nil {
		err = fmt.Errorf("Could not get random gathering place: %w", err)
		return "", err
	}
	return gatheringPlace, nil
}

// Random generates a completely random religion
func Random() (Religion, error) {
	originLanguage, _, err := conlang.Generate()
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}
	religion, err := Generate(originLanguage)
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}

	return religion, nil
}
