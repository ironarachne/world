package character

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/ironarachne/world/pkg/profession"

	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/gender"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Character is a character
type Character struct {
	FirstName      string
	LastName       string
	Title          string
	Heraldry       heraldry.Heraldry
	Gender         gender.Gender
	Age            int
	AgeCategory    race.AgeCategory
	Orientation    string
	Height         int
	Weight         int
	Profession     profession.Profession
	Hobby          Hobby
	NegativeTraits []string
	PositiveTraits []string
	Motivation     string
	HairColor      string
	HairStyle      string
	FacialHair     string
	EyeColor       string
	EyeShape       string
	FaceShape      string
	MouthShape     string
	NoseShape      string
	SkinColor      string
	Culture        culture.Culture
	Race           race.Race
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

func randomOrientation() string {
	orientations := map[string]int{
		"straight": 100,
		"gay":      10,
		"bi":       15,
	}

	orientation := random.StringFromThresholdMap(orientations)

	return orientation
}

func (character Character) randomHeight() int {
	minHeight := character.Race.Appearance.MinFemaleHeight
	maxHeight := character.Race.Appearance.MaxFemaleHeight

	if character.Gender.Name == "male" {
		minHeight = character.Race.Appearance.MinMaleHeight
		maxHeight = character.Race.Appearance.MaxMaleHeight
	}

	height := rand.Intn(maxHeight-minHeight) + minHeight

	if character.AgeCategory.Name == "child" {
		height = int(float64(height) * 0.6)
	} else if character.AgeCategory.Name == "infant" {
		height = int(float64(height) * 0.1)
	}

	return height
}

func (character Character) randomWeight() int {
	minWeight := character.Race.Appearance.MinFemaleWeight
	maxWeight := character.Race.Appearance.MaxFemaleWeight

	if character.Gender.Name == "male" {
		minWeight = character.Race.Appearance.MinMaleWeight
		maxWeight = character.Race.Appearance.MaxMaleWeight
	}

	weight := rand.Intn(maxWeight-minWeight) + minWeight

	if character.AgeCategory.Name == "child" {
		weight = int(float64(weight) * 0.4)
	} else if character.AgeCategory.Name == "infant" {
		weight = int(float64(weight) * 0.1)
	}

	return weight
}

func (character Character) randomFacialHair() (string, error) {
	if character.Gender.Name == "female" {
		return "none", nil
	}

	hairChance := rand.Intn(15)
	if hairChance > 8 {
		facialHair, err := random.String(character.Race.Appearance.FacialHairStyles)
		if err != nil {
			err = fmt.Errorf("Could not get facial hair style: %w", err)
			return "", err
		}
		return facialHair, nil
	}

	return "none", nil
}

// Generate generates a random character
func Generate(originCulture culture.Culture) (Character, error) {
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

	char.AgeCategory = char.Race.GetWeightedAgeCategory()
	char.Age = race.GetRandomAge(char.AgeCategory)

	hairColor, err := random.String(char.Race.Appearance.HairColors)
	if err != nil {
		err = fmt.Errorf("Could not generate hair color: %w", err)
		return Character{}, err
	}
	char.HairColor = hairColor

	hairStyle, err := random.String(char.Race.Appearance.FemaleHairStyles)

	if char.Gender.Name == "male" {
		hairStyle, err = random.String(char.Race.Appearance.MaleHairStyles)
	}
	if err != nil {
		err = fmt.Errorf("Could not generate hair style: %w", err)
		return Character{}, err
	}
	char.HairStyle = hairStyle
	facialHair, err := char.randomFacialHair()
	if err != nil {
		err = fmt.Errorf("Could not generate character: %w", err)
		return Character{}, err
	}
	char.FacialHair = facialHair

	eyeColor, err := random.String(char.Race.Appearance.EyeColors)
	if err != nil {
		err = fmt.Errorf("Could not generate eye color: %w", err)
		return Character{}, err
	}
	char.EyeColor = eyeColor
	eyeShape, err := random.String(char.Race.Appearance.EyeShapes)
	if err != nil {
		err = fmt.Errorf("Could not generate eye shape: %w", err)
		return Character{}, err
	}
	char.EyeShape = eyeShape
	faceShape, err := random.String(char.Race.Appearance.FaceShapes)
	if err != nil {
		err = fmt.Errorf("Could not generate face shape: %w", err)
		return Character{}, err
	}
	char.FaceShape = faceShape
	mouthShape, err := random.String(char.Race.Appearance.MouthShapes)
	if err != nil {
		err = fmt.Errorf("Could not generate mouth shape: %w", err)
		return Character{}, err
	}
	char.MouthShape = mouthShape
	noseShape, err := random.String(char.Race.Appearance.NoseShapes)
	if err != nil {
		err = fmt.Errorf("Could not generate nose shape: %w", err)
		return Character{}, err
	}
	char.NoseShape = noseShape
	skinColor, err := random.String(char.Race.Appearance.SkinColors)
	if err != nil {
		err = fmt.Errorf("Could not generate skin color: %w", err)
		return Character{}, err
	}
	char.SkinColor = skinColor

	char.Orientation = randomOrientation()
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
		char1.AgeCategory = char1.Race.GetAgeCategoryByName("adult")
		char1.Age = race.GetRandomAge(char1.AgeCategory)
	}

	if char2.AgeCategory.Name == "child" {
		char2.AgeCategory = char2.Race.GetAgeCategoryByName("adult")
		char2.Age = race.GetRandomAge(char2.AgeCategory)
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

	ac := couple.Partner1.Race.GetAgeCategoryByName("adult")

	descendent.Age = race.GetRandomAge(ac)
	descendent.AgeCategory = couple.Partner1.Race.GetAgeCategoryFromAge(descendent.Age)

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

	mate.Age = race.GetRandomAge(char.AgeCategory)
	mate.AgeCategory = mate.Race.GetAgeCategoryFromAge(mate.Age)

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

// HeightSimplified returns a string in the common format for height
func (character Character) HeightSimplified() string {
	feet := strconv.Itoa(int(math.Floor(float64(character.Height) / 12.0)))
	inches := strconv.Itoa(int(math.Mod(float64(character.Height), 12.0)))

	display := feet + "'" + inches + "\""

	return display
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
