package character

import "math/rand"

// Gender is a gender and its accompanying descriptors
type Gender struct {
	Name                 string
	Noun                 string
	PluralNoun           string
	AdolescentNoun       string
	PluralAdolescentNoun string
	ObjectPronoun        string
	PossessivePronoun    string
	ReflexivePronoun     string
	SubjectPronoun       string
}

func getAllGenders() []Gender {
	genders := []Gender{}
	genders = append(genders, getFemale())
	genders = append(genders, getMale())

	return genders
}

func getFemale() Gender {
	gender := Gender{
		Name:                 "female",
		Noun:                 "woman",
		PluralNoun:           "women",
		AdolescentNoun:       "girl",
		PluralAdolescentNoun: "girls",
		ObjectPronoun:        "her",
		PossessivePronoun:    "her",
		ReflexivePronoun:     "herself",
		SubjectPronoun:       "she",
	}

	return gender
}

func getMale() Gender {
	gender := Gender{
		Name:                 "male",
		Noun:                 "man",
		PluralNoun:           "men",
		AdolescentNoun:       "boy",
		PluralAdolescentNoun: "boys",
		ObjectPronoun:        "her",
		PossessivePronoun:    "his",
		ReflexivePronoun:     "himself",
		SubjectPronoun:       "he",
	}

	return gender
}

func getOppositeGender(gender Gender) Gender {
	if gender.Name == "male" {
		return getFemale()
	}

	return getMale()
}

func getRandomGender() Gender {
	genders := getAllGenders()

	return genders[rand.Intn(len(genders)-1)]
}
