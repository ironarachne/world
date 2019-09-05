package race

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/ironarachne/world/pkg/random"
)

// Appearance is a set of physical characteristics
type Appearance struct {
	MaxFemaleHeight   int
	MinFemaleHeight   int
	MaxMaleHeight     int
	MinMaleHeight     int
	MaxFemaleWeight   int
	MinFemaleWeight   int
	MaxMaleWeight     int
	MinMaleWeight     int
	FaceShapes        []string
	EarShapes         []string
	EyeShapes         []string
	NoseShapes        []string
	MouthShapes       []string
	FacialHairStyles  []string
	HairColors        []string
	HairConsistencies []string
	FemaleHairStyles  []string
	MaleHairStyles    []string
	SkinColors        []string
	EyeColors         []string
	UniqueTraits      []string
}

func (race Race) generateRandomSubraceAppearance() (Appearance, error) {
	appearance := Appearance{}

	skinColors, err := random.StringSubset(race.Appearance.SkinColors, 2)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.SkinColors = skinColors
	hairColors, err := random.StringSubset(race.Appearance.HairColors, 2)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.HairColors = hairColors
	eyeColors, err := random.StringSubset(race.Appearance.EyeColors, 2)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.EyeColors = eyeColors
	hairConsistencies, err := random.StringSubset(race.Appearance.HairConsistencies, 1)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.HairConsistencies = hairConsistencies
	femaleHairStyles, err := random.StringSubset(race.Appearance.FemaleHairStyles, 2)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.FemaleHairStyles = femaleHairStyles
	maleHairStyles, err := random.StringSubset(race.Appearance.MaleHairStyles, 2)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.MaleHairStyles = maleHairStyles
	facialHairStyles, err := random.StringSubset(race.Appearance.FacialHairStyles, 3)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.FacialHairStyles = facialHairStyles
	earShapes, err := random.StringSubset(race.Appearance.EarShapes, 1)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.EarShapes = earShapes
	eyeShapes, err := random.StringSubset(race.Appearance.EyeShapes, 1)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.EyeShapes = eyeShapes
	faceShapes, err := random.StringSubset(race.Appearance.FaceShapes, 1)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.FaceShapes = faceShapes
	mouthShapes, err := random.StringSubset(race.Appearance.MouthShapes, 1)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.MouthShapes = mouthShapes
	noseShapes, err := random.StringSubset(race.Appearance.NoseShapes, 1)
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
		return Appearance{}, err
	}
	appearance.NoseShapes = noseShapes

	appearance.MinFemaleHeight = rand.Intn(24) + 52
	appearance.MaxFemaleHeight = appearance.MinFemaleHeight + rand.Intn(12) + 1
	appearance.MinMaleHeight = rand.Intn(24) + 52
	appearance.MaxMaleHeight = appearance.MinMaleHeight + rand.Intn(12) + 1
	appearance.MinFemaleWeight = rand.Intn(30) + 90
	appearance.MaxFemaleWeight = appearance.MinFemaleWeight + rand.Intn(80) + 1
	appearance.MinMaleWeight = rand.Intn(30) + 140
	appearance.MaxMaleWeight = appearance.MinMaleWeight + rand.Intn(100) + 1

	if len(race.Appearance.UniqueTraits) > 0 {
		uniqueTraits, err := random.StringSubset(race.Appearance.UniqueTraits, 1)
		if err != nil {
			err = fmt.Errorf("Could not generate random subrace appearance: %w", err)
			return Appearance{}, err
		}
		appearance.UniqueTraits = uniqueTraits
	}

	return appearance, nil
}

func heightToString(heightInInches int) string {
	feet := strconv.Itoa(int(math.Floor(float64(heightInInches) / 12.0)))
	inches := strconv.Itoa(int(math.Mod(float64(heightInInches), 12.0)))

	display := feet + "'" + inches + "\""

	return display
}

func (race Race) getAverageHeight(gender string) int {
	min := math.Min(float64(race.Appearance.MinFemaleHeight), float64(race.Appearance.MinMaleHeight))
	max := math.Max(float64(race.Appearance.MaxFemaleHeight), float64(race.Appearance.MaxMaleHeight))

	if gender == "male" {
		min = float64(race.Appearance.MinMaleHeight)
		max = float64(race.Appearance.MaxMaleHeight)
	} else if gender == "female" {
		min = float64(race.Appearance.MinFemaleHeight)
		max = float64(race.Appearance.MaxFemaleHeight)
	}

	return int((min + max) / 2)
}

func (race Race) getAverageWeight(gender string) int {
	min := math.Min(float64(race.Appearance.MinFemaleWeight), float64(race.Appearance.MinMaleWeight))
	max := math.Max(float64(race.Appearance.MaxFemaleWeight), float64(race.Appearance.MaxMaleWeight))

	if gender == "male" {
		min = float64(race.Appearance.MinMaleWeight)
		max = float64(race.Appearance.MaxMaleWeight)
	} else if gender == "female" {
		min = float64(race.Appearance.MinFemaleWeight)
		max = float64(race.Appearance.MaxFemaleWeight)
	}

	return int((min + max) / 2)
}
