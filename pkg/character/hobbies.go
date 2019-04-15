package character

import "math/rand"

// Hobby is a pasttime
type Hobby struct {
	Name           string
	RequiresOthers bool
	IsAdultOnly    bool
	IsChildOnly    bool
}

func getAllHobbies() []Hobby {
	hobbies := []Hobby{
		Hobby{
			Name:           "whittling",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		Hobby{
			Name:           "carving",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		Hobby{
			Name:           "darts",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		Hobby{
			Name:           "tournaments",
			RequiresOthers: false,
			IsAdultOnly:    true,
			IsChildOnly:    false,
		},
		Hobby{
			Name:           "fishing",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		Hobby{
			Name:           "hunting",
			RequiresOthers: false,
			IsAdultOnly:    true,
			IsChildOnly:    false,
		},
		Hobby{
			Name:           "painting",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		Hobby{
			Name:           "drinking",
			RequiresOthers: false,
			IsAdultOnly:    true,
			IsChildOnly:    false,
		},
	}

	return hobbies
}

func getHobbiesForAgeCategory(category string) []Hobby {
	potentialHobbies := getAllHobbies()
	hobbies := []Hobby{}

	if category == "infant" {
		hobbies = []Hobby{}
	} else if category == "child" {
		for _, h := range potentialHobbies {
			if !h.IsAdultOnly {
				hobbies = append(hobbies, h)
			}
		}
	} else {
		for _, h := range potentialHobbies {
			if !h.IsChildOnly {
				hobbies = append(hobbies, h)
			}
		}
	}

	return hobbies
}

func (character Character) getRandomHobby() Hobby {
	hobbies := getHobbiesForAgeCategory(character.AgeCategory)

	if len(hobbies) > 0 {
		return hobbies[rand.Intn(len(hobbies)-1)]
	}

	return Hobby{}
}
