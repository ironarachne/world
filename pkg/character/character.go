package character

import (
	"math/rand"

	"github.com/ironarachne/utility"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Character is a character
type Character struct {
	FirstName           string
	LastName            string
	Gender              string
	Age                 int
	AgeCategory         string
	Orientation         string
	Height              int
	Weight              int
	Profession          string
	Hobby               Hobby
	PsychologicalTraits []string
	Motivation          string
	HairColor           string
	HairStyle           string
	FacialHair          string
	EyeColor            string
	FaceShape           string
	MouthShape          string
	NoseShape           string
	Culture             culture.Culture
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

func getOppositeGender(gender string) string {
	if gender == "male" {
		return "female"
	}

	return "male"
}

func getAppropriateName(gender string, culture culture.Culture) (string, string) {
	firstName := culture.Language.RandomGenderedName(gender)
	lastName := culture.Language.RandomName()

	return firstName, lastName
}

func getAgeCategoryRange(ageCategory string) (int, int) {
	minAge := 0
	maxAge := 1

	if ageCategory == "infant" {
		minAge = 0
		maxAge = 1
	} else if ageCategory == "child" {
		minAge = 2
		maxAge = 12
	} else if ageCategory == "teenager" {
		minAge = 13
		maxAge = 19
	} else if ageCategory == "young adult" {
		minAge = 20
		maxAge = 25
	} else if ageCategory == "adult" {
		minAge = 26
		maxAge = 69
	} else if ageCategory == "elderly" {
		minAge = 70
		maxAge = 110
	}

	return minAge, maxAge
}

func getRandomAge(ageCategory string) int {
	minAge, maxAge := getAgeCategoryRange(ageCategory)

	return rand.Intn(maxAge-minAge) + minAge
}

func getAgeCategoryFromAge(age int) string {
	ageCategory := ""
	minAge := 0
	maxAge := 0

	for category := range ageCategories {
		minAge, maxAge = getAgeCategoryRange(category)
		if age >= minAge && age <= maxAge {
			ageCategory = category
		}
	}

	return ageCategory
}

func getAgeFromParents(parents Couple) (int, string) {
	lowestAge := utility.Min(parents.Partner1.Age, parents.Partner2.Age)

	childAge := rand.Intn(lowestAge-18) + 1
	childAgeCategory := getAgeCategoryFromAge(childAge)

	return childAge, childAgeCategory
}

func randomOrientation() string {
	return random.StringFromThresholdMap(orientations)
}

func (character Character) randomHeight() int {
	minHeight := character.Culture.Appearance.MinFemaleHeight
	maxHeight := character.Culture.Appearance.MaxFemaleHeight

	if character.Gender == "male" {
		minHeight = character.Culture.Appearance.MinMaleHeight
		maxHeight = character.Culture.Appearance.MaxMaleHeight
	}

	height := rand.Intn(maxHeight-minHeight) + minHeight

	if character.AgeCategory == "child" {
		height = int(float64(height) * 0.6)
	} else if character.AgeCategory == "infant" {
		height = int(float64(height) * 0.1)
	}

	return height
}

func (character Character) randomWeight() int {
	minWeight := character.Culture.Appearance.MinFemaleWeight
	maxWeight := character.Culture.Appearance.MaxFemaleWeight

	if character.Gender == "male" {
		minWeight = character.Culture.Appearance.MinMaleWeight
		maxWeight = character.Culture.Appearance.MaxMaleWeight
	}

	weight := rand.Intn(maxWeight-minWeight) + minWeight

	if character.AgeCategory == "child" {
		weight = int(float64(weight) * 0.4)
	} else if character.AgeCategory == "infant" {
		weight = int(float64(weight) * 0.1)
	}

	return weight
}

func (character Character) randomFacialHair() string {
	if character.Gender == "female" {
		return "none"
	}

	hairChance := rand.Intn(10)
	if hairChance > 8 {
		return random.String(character.Culture.Appearance.FacialHairStyles)
	}

	return "none"
}

// Generate generates a random character
func Generate() Character {
	char := Character{}

	char.Gender = random.String(genders)
	char.Culture = culture.Generate()

	char.FirstName, char.LastName = getAppropriateName(char.Gender, char.Culture)

	char.AgeCategory = random.StringFromThresholdMap(ageCategories)
	char.Age = getRandomAge(char.AgeCategory)

	char.HairColor = random.String(char.Culture.Appearance.HairColors)
	if char.Gender == "male" {
		char.HairStyle = random.String(char.Culture.Appearance.MaleHairStyles)
	} else {
		char.HairStyle = random.String(char.Culture.Appearance.FemaleHairStyles)
	}
	char.FacialHair = char.randomFacialHair()

	char.EyeColor = random.String(char.Culture.Appearance.EyeColors)
	char.FaceShape = char.Culture.Appearance.FaceShape
	char.MouthShape = char.Culture.Appearance.MouthShape
	char.NoseShape = char.Culture.Appearance.NoseShape

	char.Orientation = randomOrientation()
	char.Profession = random.String(professions)
	char.Hobby = char.getRandomHobby()
	char.Motivation = random.String(motivations)

	for i := 0; i < 2; i++ {
		trait := random.String(traits)
		for slices.StringIn(trait, char.PsychologicalTraits) {
			trait = random.String(traits)
		}
		char.PsychologicalTraits = append(char.PsychologicalTraits, trait)
	}

	char.Height = char.randomHeight()
	char.Weight = char.randomWeight()

	return char
}

// GenerateCharacterOfCulture generates a random character with a given culture
func GenerateCharacterOfCulture(culture culture.Culture) Character {
	character := Generate()
	character = character.SetCulture(culture)

	return character
}

// GenerateCouple generates a couple
func GenerateCouple() Couple {
	char1 := Generate()
	char2 := Generate()
	canHaveChildren := false

	if char1.AgeCategory == "child" {
		char1.AgeCategory = "adult"
		char1.Age = getRandomAge(char1.AgeCategory)
	}

	if char2.AgeCategory == "child" {
		char2.AgeCategory = "adult"
		char2.Age = getRandomAge(char2.AgeCategory)
	}

	orientations := []string{char1.Orientation, char2.Orientation}

	if (char1.Orientation == "gay" && char2.Orientation == "straight") || (char1.Orientation == "straight" && char2.Orientation == "gay") {
		char2.Orientation = char1.Orientation
		orientations = []string{char1.Orientation}
	}

	if char1.Gender == char2.Gender && utility.ItemInCollection("straight", orientations) {
		char2.Gender = getOppositeGender(char1.Gender)
		char2.FirstName, _ = getAppropriateName(char2.Gender, char2.Culture)
	} else if char1.Gender != char2.Gender && utility.ItemInCollection("gay", orientations) {
		char2.Gender = char1.Gender
		char2.FirstName, _ = getAppropriateName(char2.Gender, char2.Culture)
	}

	if char1.Gender != char2.Gender {
		canHaveChildren = true
	}

	couple := Couple{char1, char2, canHaveChildren}

	return couple
}

// GenerateAdultDescendent generates an adult character based on a couple
func GenerateAdultDescendent(couple Couple) Character {
	descendent := Generate()

	descendent.LastName = couple.Partner1.LastName

	descendent.Age = getRandomAge("adult")
	descendent.AgeCategory = getAgeCategoryFromAge(descendent.Age)

	return descendent
}

// GenerateChild generates a child character for a couple
func GenerateChild(couple Couple) Character {
	child := Generate()

	child.LastName = couple.Partner1.LastName
	child.Age, child.AgeCategory = getAgeFromParents(couple)

	if child.AgeCategory == "child" {
		child.Profession = "none"
	}

	return child
}

// GenerateCompatibleMate generates a character appropriate as a mate for another
func GenerateCompatibleMate(char Character) Character {
	mate := Generate()

	mate.Age = getRandomAge(char.AgeCategory)
	mate.AgeCategory = getAgeCategoryFromAge(mate.Age)

	if char.Orientation == "straight" {
		mate.Gender = getOppositeGender(char.Gender)
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

// MarryCouple returns a couple from two characters
func MarryCouple(partner1 Character, partner2 Character) Couple {
	canHaveChildren := false

	if partner1.Gender != partner2.Gender {
		canHaveChildren = true
	}

	return Couple{partner1, partner2, canHaveChildren}
}

// SetCulture sets the culture of the character
func (character Character) SetCulture(culture culture.Culture) Character {
	newCharacter := character

	newCharacter.Culture = culture
	newCharacter.FirstName, newCharacter.LastName = getAppropriateName(newCharacter.Gender, newCharacter.Culture)

	return newCharacter
}
