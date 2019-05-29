package character

import (
	"math"
	"math/rand"
	"strconv"

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
	Heraldry       heraldry.Heraldry
	Gender         gender.Gender
	Age            int
	AgeCategory    AgeCategory
	Orientation    string
	Height         int
	Weight         int
	Profession     string
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

func getAppropriateName(gender string, culture culture.Culture) (string, string) {
	firstName := culture.Language.RandomGenderedName(gender)
	lastName := culture.Language.RandomName()

	return firstName, lastName
}

func randomOrientation() string {
	orientations := map[string]int{
		"straight": 10,
		"gay":      1,
		"bi":       1,
	}

	return random.StringFromThresholdMap(orientations)
}

func (character Character) randomHeight() int {
	minHeight := character.Culture.Appearance.MinFemaleHeight
	maxHeight := character.Culture.Appearance.MaxFemaleHeight

	if character.Gender.Name == "male" {
		minHeight = character.Culture.Appearance.MinMaleHeight
		maxHeight = character.Culture.Appearance.MaxMaleHeight
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
	minWeight := character.Culture.Appearance.MinFemaleWeight
	maxWeight := character.Culture.Appearance.MaxFemaleWeight

	if character.Gender.Name == "male" {
		minWeight = character.Culture.Appearance.MinMaleWeight
		maxWeight = character.Culture.Appearance.MaxMaleWeight
	}

	weight := rand.Intn(maxWeight-minWeight) + minWeight

	if character.AgeCategory.Name == "child" {
		weight = int(float64(weight) * 0.4)
	} else if character.AgeCategory.Name == "infant" {
		weight = int(float64(weight) * 0.1)
	}

	return weight
}

func (character Character) randomFacialHair() string {
	if character.Gender.Name == "female" {
		return "none"
	}

	hairChance := rand.Intn(15)
	if hairChance > 8 {
		return random.String(character.Culture.Appearance.FacialHairStyles)
	}

	return "none"
}

// Generate generates a random character
func Generate(originCulture culture.Culture) Character {
	char := Character{}

	char.Gender = gender.Random()
	char.Culture = originCulture

	char.FirstName, char.LastName = getAppropriateName(char.Gender.Name, char.Culture)

	char.AgeCategory = getWeightedAgeCategory()
	char.Age = getRandomAge(char.AgeCategory)

	char.HairColor = random.String(char.Culture.Appearance.HairColors)
	if char.Gender.Name == "male" {
		char.HairStyle = random.String(char.Culture.Appearance.MaleHairStyles)
	} else {
		char.HairStyle = random.String(char.Culture.Appearance.FemaleHairStyles)
	}
	char.FacialHair = char.randomFacialHair()

	char.EyeColor = random.String(char.Culture.Appearance.EyeColors)
	char.EyeShape = char.Culture.Appearance.EyeShape
	char.FaceShape = char.Culture.Appearance.FaceShape
	char.MouthShape = char.Culture.Appearance.MouthShape
	char.NoseShape = char.Culture.Appearance.NoseShape
	char.SkinColor = random.String(char.Culture.Appearance.SkinColors)

	char.Orientation = randomOrientation()
	char.Profession = getRandomProfession()
	char.Hobby = char.getRandomHobby()
	char.Motivation = getRandomMotivation()

	char.NegativeTraits = getRandomNegativeTraits(2)
	char.PositiveTraits = getRandomPositiveTraits(3)

	char.Height = char.randomHeight()
	char.Weight = char.randomWeight()

	return char
}

// GenerateCouple generates a couple
func GenerateCouple() Couple {
	char1 := Random()
	char2 := Generate(char1.Culture)
	canHaveChildren := false

	if char1.AgeCategory.Name == "child" {
		char1.AgeCategory = getAgeCategoryByName("adult")
		char1.Age = getRandomAge(char1.AgeCategory)
	}

	if char2.AgeCategory.Name == "child" {
		char2.AgeCategory = getAgeCategoryByName("adult")
		char2.Age = getRandomAge(char2.AgeCategory)
	}

	orientations := []string{char1.Orientation, char2.Orientation}

	if (char1.Orientation == "gay" && char2.Orientation == "straight") || (char1.Orientation == "straight" && char2.Orientation == "gay") {
		char2.Orientation = char1.Orientation
		orientations = []string{char1.Orientation}
	}

	if char1.Gender.Name == char2.Gender.Name && slices.StringIn("straight", orientations) {
		char2.Gender = char1.Gender.Opposite()
		char2.FirstName, _ = getAppropriateName(char2.Gender.Name, char2.Culture)
	} else if char1.Gender.Name != char2.Gender.Name && slices.StringIn("gay", orientations) {
		char2.Gender = char1.Gender
		char2.FirstName, _ = getAppropriateName(char2.Gender.Name, char2.Culture)
	}

	if char1.Gender != char2.Gender {
		canHaveChildren = true
	}

	couple := Couple{char1, char2, canHaveChildren}

	return couple
}

// GenerateAdultDescendent generates an adult character based on a couple
func GenerateAdultDescendent(couple Couple) Character {
	descendent := Generate(couple.Partner1.Culture)

	descendent.LastName = couple.Partner1.LastName

	ac := getAgeCategoryByName("adult")

	descendent.Age = getRandomAge(ac)
	descendent.AgeCategory = getAgeCategoryFromAge(descendent.Age)

	return descendent
}

// GenerateChild generates a child character for a couple
func GenerateChild(couple Couple) Character {
	child := Generate(couple.Partner1.Culture)

	child.LastName = couple.Partner1.LastName
	child.Age, child.AgeCategory = getAgeFromParents(couple)

	if child.AgeCategory.Name == "child" {
		child.Profession = "none"
	}

	return child
}

// GenerateCompatibleMate generates a character appropriate as a mate for another
func GenerateCompatibleMate(char Character) Character {
	mate := Generate(char.Culture)

	mate.Age = getRandomAge(char.AgeCategory)
	mate.AgeCategory = getAgeCategoryFromAge(mate.Age)

	if char.Orientation == "straight" {
		mate.Gender = char.Gender.Opposite()
		mate.Orientation = "straight"
	} else {
		mate.Gender = char.Gender
		mate.Orientation = "gay"
	}

	return mate
}

// GenerateFamily generates a random family
func GenerateFamily() Family {
	child := Character{}
	children := []Character{}

	parents := GenerateCouple()

	familyName := parents.Partner1.LastName

	parents.Partner2.LastName = familyName

	if parents.CanHaveChildren {
		for i := 0; i < rand.Intn(6); i++ {
			child = GenerateChild(parents)
			child.LastName = familyName
			children = append(children, child)
		}
	}

	return Family{familyName, parents, children}
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
func Random() Character {
	randomCulture := culture.Generate()

	character := Generate(randomCulture)

	return character
}
