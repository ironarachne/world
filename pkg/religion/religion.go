/*
Package religion implements fantasy religions
*/
package religion

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/conlang"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/random"
)

// Religion is a religion
type Religion struct {
	Name           string            `json:"name"`
	CommonName     string            `json:"common_name"`
	Class          Class             `json:"class"`
	Pantheon       pantheon.Pantheon `json:"pantheon"`
	GatheringPlace string            `json:"gathering_place"`
	Virtues        []string          `json:"virtues"`
}

// Generate procedurally generates a religion
func Generate(ctx context.Context, originLanguage language.Language) (Religion, error) {
	religion := Religion{}

	class, err := getWeightedClass(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}
	religion.Class = class
	gatheringPlace, err := religion.randomGatheringPlace(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}
	religion.GatheringPlace = gatheringPlace

	if religion.Class.PantheonMaxSize > 0 {
		newPantheon, err := pantheon.Generate(ctx, religion.Class.PantheonMinSize, religion.Class.PantheonMaxSize, originLanguage)
		if err != nil {
			err = fmt.Errorf("Could not generate religion: %w", err)
			return Religion{}, err
		}
		religion.Pantheon = newPantheon
	}

	religionName, err := randomReligionName(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate religion: %w", err)
		return Religion{}, err
	}
	name := originLanguage.TranslatePhrase(religionName)
	religion.CommonName = religionName
	religion.Name = name

	virtues, err := getRandomVirtues(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate religion: %w", err)
		return Religion{}, err
	}
	religion.Virtues = virtues

	return religion, nil
}

func (religion Religion) randomGatheringPlace(ctx context.Context) (string, error) {
	gatheringPlace, err := random.String(ctx, religion.Class.GatheringPlaces)
	if err != nil {
		err = fmt.Errorf("Could not get random gathering place: %w", err)
		return "", err
	}
	return gatheringPlace, nil
}

// Random generates a completely random religion
func Random(ctx context.Context) (Religion, error) {
	originLanguage, _, err := conlang.Generate(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}
	religion, err := Generate(ctx, originLanguage)
	if err != nil {
		err = fmt.Errorf("Could not generate random religion: %w", err)
		return Religion{}, err
	}

	return religion, nil
}
