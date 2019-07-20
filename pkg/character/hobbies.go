package character

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/race"
)

// Hobby is a pasttime
type Hobby struct {
	Name           string
	RequiresOthers bool
	IsAdultOnly    bool
	IsChildOnly    bool
}

func getAllHobbies() []Hobby {
	hobbies := []Hobby{
		{
			Name:           "archery",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "carving",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "dancing",
			RequiresOthers: true,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "darts",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "drinking",
			RequiresOthers: false,
			IsAdultOnly:    true,
			IsChildOnly:    false,
		},
		{
			Name:           "feasting",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "fishing",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "gambling",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "games of strategy",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "hawking",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "hunting",
			RequiresOthers: false,
			IsAdultOnly:    true,
			IsChildOnly:    false,
		},
		{
			Name:           "painting",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "participating in tournaments",
			RequiresOthers: false,
			IsAdultOnly:    true,
			IsChildOnly:    false,
		},
		{
			Name:           "watching tournaments",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
		{
			Name:           "whittling",
			RequiresOthers: false,
			IsAdultOnly:    false,
			IsChildOnly:    false,
		},
	}

	return hobbies
}

func getHobbiesForAgeCategory(category race.AgeCategory) []Hobby {
	potentialHobbies := getAllHobbies()
	hobbies := []Hobby{}

	if category.Name == "infant" {
		hobbies = []Hobby{}
	} else if category.Name == "child" {
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
