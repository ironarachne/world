/*
Package character provides fictional character generation tools.
*/
package character

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/race"

	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/trait"

	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/gender"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

const characterError = "failed to generate character: %w"
const coupleError = "failed to generate couple: %w"
const cultureNameError = "failed to generate appropriate name for culture: %w"

// Character is a character
type Character struct {
	FirstName      string                    `json:"first_name"`
	LastName       string                    `json:"last_name"`
	Title          string                    `json:"title"`  // Title is the character's primary title
	Titles         []string                  `json:"titles"` // Titles is a list of all of the character's titles
	Heraldry       heraldry.SimplifiedDevice `json:"heraldry"`
	Gender         gender.Gender             `json:"gender"`
	Age            int                       `json:"age"`
	AgeCategory    age.Category              `json:"age_category"`
	Orientation    string                    `json:"orientation"`
	Height         int                       `json:"height"`
	Weight         int                       `json:"weight"`
	Profession     profession.Profession     `json:"profession"`
	Hobby          Hobby                     `json:"hobby"`
	NegativeTraits []string                  `json:"negative_traits"`
	PositiveTraits []string                  `json:"positive_traits"`
	Motivation     string                    `json:"motivation"`
	PhysicalTraits []trait.Trait             `json:"physical_traits"`
	Culture        culture.Culture           `json:"culture"`
	Race           species.Species           `json:"race"`
}

// Couple is a pair of partners
type Couple struct {
	Partner1        Character
	Partner2        Character
	CanHaveChildren bool
}

// Family is a family of characters
type Family struct {
	FamilyName string
	Parents    Couple
	Children   []Character
}

func getAppropriateName(ctx context.Context, gender string, culture culture.Culture) (string, string, error) {
	firstName, err := culture.Language.RandomMaleFirstName(ctx)
	if err != nil {
		err = fmt.Errorf(cultureNameError, err)
		return "", "", err
	}

	if gender == "female" {
		firstName, err = culture.Language.RandomFemaleFirstName(ctx)
		if err != nil {
			err = fmt.Errorf(cultureNameError, err)
			return "", "", err
		}
	}

	lastName, err := culture.Language.RandomFamilyName(ctx)
	if err != nil {
		err = fmt.Errorf(cultureNameError, err)
		return "", "", err
	}

	return firstName, lastName, nil
}

func randomOrientation(ctx context.Context) (string, error) {
	orientations := map[string]int{
		"straight": 100,
		"gay":      10,
		"bi":       15,
	}

	orientation, err := random.StringFromThresholdMap(ctx, orientations)
	if err != nil {
		err = fmt.Errorf("failed to generate random sexual orientation: %w", err)
		return "", err
	}

	return orientation, nil
}

// Generate generates a random character
func Generate(ctx context.Context, originCulture culture.Culture) (Character, error) {
	var t trait.Trait
	char := Character{}

	char.Gender = gender.Random(ctx)
	char.Culture = originCulture
	char.Race = char.Culture.PrimaryRace

	firstName, lastName, err := getAppropriateName(ctx, char.Gender.Name, char.Culture)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	char.FirstName = firstName
	char.LastName = lastName

	ageCategory, err := age.GetWeightedAgeCategory(ctx, char.Race.AgeCategories)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	char.AgeCategory = ageCategory
	char.Age = age.GetRandomAge(ctx, char.AgeCategory)

	orientation, err := randomOrientation(ctx)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	char.Orientation = orientation
	char.Profession, _ = profession.Random(ctx)
	char.Hobby = char.getRandomHobby(ctx)
	motivation, err := getRandomMotivation(ctx)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	char.Motivation = motivation

	negativeTraits, err := getRandomNegativeTraits(ctx, 2)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	char.NegativeTraits = negativeTraits
	positiveTraits, err := getRandomPositiveTraits(ctx, 3)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	char.PositiveTraits = positiveTraits

	raceTraits := []trait.Trait{}

	for _, i := range char.Race.CommonTraits {
		t, err = i.ToTrait(ctx)
		if err != nil {
			err = fmt.Errorf(characterError, err)
			return Character{}, err
		}
		raceTraits = append(raceTraits, t)
	}

	possibleTraits := char.Race.PossibleTraits
	possibleTraits = append(possibleTraits, race.GetCommonPossibleTraits()...)
	uniqueTraitTemplate, err := trait.RandomTemplate(ctx, possibleTraits)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	uniqueTrait, err := uniqueTraitTemplate.ToTrait(ctx)
	if err != nil {
		err = fmt.Errorf(characterError, err)
		return Character{}, err
	}
	raceTraits = append(raceTraits, uniqueTrait)
	char.PhysicalTraits = raceTraits

	char.Height = age.GetRandomHeight(ctx, char.Gender.Name, char.AgeCategory)
	char.Weight = age.GetRandomWeight(ctx, char.Gender.Name, char.AgeCategory)

	return char, nil
}

// GenerateCouple generates a couple
func GenerateCouple(ctx context.Context) (Couple, error) {
	char1, err := Random(ctx)
	if err != nil {
		err = fmt.Errorf(coupleError, err)
		return Couple{}, err
	}
	char2, err := Generate(ctx, char1.Culture)
	if err != nil {
		err = fmt.Errorf(coupleError, err)
		return Couple{}, err
	}
	canHaveChildren := false

	if char1.AgeCategory.Name == "child" {
		char1.AgeCategory = age.GetCategoryByName("adult", char1.Race.AgeCategories)
		char1.Age = age.GetRandomAge(ctx, char1.AgeCategory)
	}

	if char2.AgeCategory.Name == "child" {
		char2.AgeCategory = age.GetCategoryByName("adult", char2.Race.AgeCategories)
		char2.Age = age.GetRandomAge(ctx, char2.AgeCategory)
	}

	orientations := []string{char1.Orientation, char2.Orientation}

	if (char1.Orientation == "gay" && char2.Orientation == "straight") || (char1.Orientation == "straight" && char2.Orientation == "gay") {
		char2.Orientation = char1.Orientation
		orientations = []string{char1.Orientation}
	}

	if char1.Gender.Name == char2.Gender.Name && slices.StringIn("straight", orientations) {
		char2.Gender = char1.Gender.Opposite()
	} else if char1.Gender.Name != char2.Gender.Name && slices.StringIn("gay", orientations) {
		char2.Gender = char1.Gender
	}

	partnerFirstName, _, err := getAppropriateName(ctx, char2.Gender.Name, char2.Culture)
	if err != nil {
		err = fmt.Errorf(coupleError, err)
		return Couple{}, err
	}
	char2.FirstName = partnerFirstName

	if char1.Gender != char2.Gender {
		canHaveChildren = true
	}

	couple := Couple{char1, char2, canHaveChildren}

	return couple, nil
}

// GenerateAdultDescendent generates an adult character based on a couple
func GenerateAdultDescendent(ctx context.Context, couple Couple) (Character, error) {
	descendent, err := Generate(ctx, couple.Partner1.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate descendent: %w", err)
		return Character{}, err
	}

	descendent.LastName = couple.Partner1.LastName

	ac := age.GetCategoryByName("adult", couple.Partner1.Race.AgeCategories)

	descendent.Age = age.GetRandomAge(ctx, ac)
	descendent.AgeCategory = age.GetCategoryFromAge(descendent.Age, couple.Partner1.Race.AgeCategories)

	return descendent, nil
}

// GenerateChild generates a child character for a couple
func GenerateChild(ctx context.Context, couple Couple) (Character, error) {
	child, err := Generate(ctx, couple.Partner1.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate child: %w", err)
		return Character{}, err
	}

	child.LastName = couple.Partner1.LastName
	child.Age, child.AgeCategory = getAgeFromParents(ctx, couple)

	if child.AgeCategory.Name == "child" {
		child.Profession, _ = profession.ByName("none")
	}

	return child, nil
}

// GenerateCompatibleMate generates a character appropriate as a mate for another
func GenerateCompatibleMate(ctx context.Context, char Character) (Character, error) {
	mate, err := Generate(ctx, char.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate mate: %w", err)
		return Character{}, err
	}

	mate.Race = char.Race

	mate.Age = age.GetRandomAge(ctx, char.AgeCategory)
	mate.AgeCategory = age.GetCategoryFromAge(mate.Age, mate.Race.AgeCategories)

	if char.Orientation == "straight" {
		mate.Gender = char.Gender.Opposite()
		mate.Orientation = "straight"
	} else {
		mate.Gender = char.Gender
		mate.Orientation = "gay"
	}

	return mate, nil
}

// GenerateFamily generates a random family
func GenerateFamily(ctx context.Context) (Family, error) {
	child := Character{}
	children := []Character{}

	parents, err := GenerateCouple(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate family: %w", err)
		return Family{}, err
	}

	familyName := parents.Partner1.LastName

	parents.Partner2.LastName = familyName

	if parents.CanHaveChildren {
		for i := 0; i < random.Intn(ctx, 6); i++ {
			child, err = GenerateChild(ctx, parents)
			if err != nil {
				err = fmt.Errorf("Could not generate family: %w", err)
				return Family{}, err
			}
			child.LastName = familyName
			children = append(children, child)
		}
	}

	return Family{familyName, parents, children}, nil
}

// MarryCouple returns a couple from two characters
func MarryCouple(partner1 Character, partner2 Character) Couple {
	canHaveChildren := false

	if partner1.Gender.Name != partner2.Gender.Name {
		canHaveChildren = true
	}

	return Couple{partner1, partner2, canHaveChildren}
}

// Random generates a completely random character
func Random(ctx context.Context) (Character, error) {
	randomCulture, err := culture.Random(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate random character: %w", err)
		return Character{}, err
	}

	character, err := Generate(ctx, randomCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate random character: %w", err)
		return Character{}, err
	}

	return character, nil
}
