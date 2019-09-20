package character

import (
	"fmt"
	"math/rand"

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

// Character is a character
type Character struct {
	FirstName      string
	LastName       string
	Title          string
	Heraldry       heraldry.Device
	Gender         gender.Gender
	Age            int
	AgeCategory    age.Category
	Orientation    string
	Height         string
	Weight         string
	Profession     profession.Profession
	Hobby          Hobby
	NegativeTraits []string
	PositiveTraits []string
	Motivation     string
	PhysicalTraits []trait.Trait
	Culture        culture.Culture
	Race           species.Species
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

func getAppropriateName(gender string, culture culture.Culture) (string, string, error) {
	firstName, err := culture.Language.RandomGenderedName(gender)
	if err != nil {
		err = fmt.Errorf("Could not generate appropriate name for culture: %w", err)
		return "", "", err
	}
	lastName, err := culture.Language.RandomName()
	if err != nil {
		err = fmt.Errorf("Could not generate appropriate name for culture: %w", err)
		return "", "", err
	}

	return firstName, lastName, nil
}

func randomOrientation() (string, error) {
	orientations := map[string]int{
		"straight": 100,
		"gay":      10,
		"bi":       15,
	}

	orientation, err := random.StringFromThresholdMap(orientations)
	if err != nil {
		err = fmt.Errorf("Could not generate random sexual orientation: %w", err)
		return "", err
	}

	return orientation, nil
}

func (character Character) randomHeight() string {
	height := species.RandomHeight(character.Gender.Name, character.AgeCategory, character.AgeCategory.SizeCategory)

	return height
}

func (character Character) randomWeight() string {
	weight := species.RandomWeight(character.Gender.Name, character.AgeCategory, character.AgeCategory.SizeCategory)

	return weight
}

// Generate generates a random character
func Generate(originCulture culture.Culture) (Character, error) {
	var t trait.Trait
	char := Character{}

	char.Gender = gender.Random()
	char.Culture = originCulture
	char.Race = char.Culture.PrimaryRace

	firstName, lastName, err := getAppropriateName(char.Gender.Name, char.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	char.FirstName = firstName
	char.LastName = lastName

	ageCategory, err := age.GetWeightedAgeCategory(char.Race.AgeCategories)
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	char.AgeCategory = ageCategory
	char.Age = age.GetRandomAge(char.AgeCategory)

	orientation, err := randomOrientation()
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	char.Orientation = orientation
	char.Profession = profession.Random()
	char.Hobby = char.getRandomHobby()
	motivation, err := getRandomMotivation()
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	char.Motivation = motivation

	negativeTraits, err := getRandomNegativeTraits(2)
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	char.NegativeTraits = negativeTraits
	positiveTraits, err := getRandomPositiveTraits(3)
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	char.PositiveTraits = positiveTraits

	raceTraits := []trait.Trait{}

	for _, i := range char.Race.CommonTraits {
		t, err = i.ToTrait()
		if err != nil {
			err = fmt.Errorf("Could not generate character: %w", err)
			return Character{}, err
		}
		raceTraits = append(raceTraits, t)
	}

	uniqueTraitTemplate, err := trait.RandomTemplate(char.Race.PossibleTraits)
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	uniqueTrait, err := uniqueTraitTemplate.ToTrait()
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	raceTraits = append(raceTraits, uniqueTrait)
	char.PhysicalTraits = raceTraits

	char.Height = char.randomHeight()
	char.Weight = char.randomWeight()

	return char, nil
}

// GenerateCouple generates a couple
func GenerateCouple() (Couple, error) {
	char1, err := Random()
	if err != nil {
		err = fmt.Errorf("Could not generate couple: %w", err)
		return Couple{}, err
	}
	char2, err := Generate(char1.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate couple: %w", err)
		return Couple{}, err
	}
	canHaveChildren := false

	if char1.AgeCategory.Name == "child" {
		char1.AgeCategory = age.GetCategoryByName("adult", char1.Race.AgeCategories)
		char1.Age = age.GetRandomAge(char1.AgeCategory)
	}

	if char2.AgeCategory.Name == "child" {
		char2.AgeCategory = age.GetCategoryByName("adult", char2.Race.AgeCategories)
		char2.Age = age.GetRandomAge(char2.AgeCategory)
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

	partnerFirstName, _, err := getAppropriateName(char2.Gender.Name, char2.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate couple: %w", err)
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
func GenerateAdultDescendent(couple Couple) (Character, error) {
	descendent, err := Generate(couple.Partner1.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate descendent: %w", err)
		return Character{}, err
	}

	descendent.LastName = couple.Partner1.LastName

	ac := age.GetCategoryByName("adult", couple.Partner1.Race.AgeCategories)

	descendent.Age = age.GetRandomAge(ac)
	descendent.AgeCategory = age.GetCategoryFromAge(descendent.Age, couple.Partner1.Race.AgeCategories)

	return descendent, nil
}

// GenerateChild generates a child character for a couple
func GenerateChild(couple Couple) (Character, error) {
	child, err := Generate(couple.Partner1.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate child: %w", err)
		return Character{}, err
	}

	child.LastName = couple.Partner1.LastName
	child.Age, child.AgeCategory = getAgeFromParents(couple)

	if child.AgeCategory.Name == "child" {
		child.Profession = profession.ByName("none")
	}

	return child, nil
}

// GenerateCompatibleMate generates a character appropriate as a mate for another
func GenerateCompatibleMate(char Character) (Character, error) {
	mate, err := Generate(char.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate mate: %w", err)
		return Character{}, err
	}

	mate.Race = char.Race

	mate.Age = age.GetRandomAge(char.AgeCategory)
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
func GenerateFamily() (Family, error) {
	child := Character{}
	children := []Character{}

	parents, err := GenerateCouple()
	if err != nil {
		err = fmt.Errorf("Could not generate face shape: %w", err)
		return Family{}, err
	}

	familyName := parents.Partner1.LastName

	parents.Partner2.LastName = familyName

	if parents.CanHaveChildren {
		for i := 0; i < rand.Intn(6); i++ {
			child, err = GenerateChild(parents)
			if err != nil {
				err = fmt.Errorf("Could not generate child for family: %w", err)
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
func Random() (Character, error) {
	randomCulture, err := culture.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate random character: %w", err)
		return Character{}, err
	}

	character, err := Generate(randomCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate random character: %w", err)
		return Character{}, err
	}

	return character, nil
}
