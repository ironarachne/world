package character

import (
	"context"

	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Hobby is a pasttime
type Hobby struct {
	Name                  string
	RequiresOthers        bool
	PossibleAgeCategories []string
}

func getAllHobbies() []Hobby {
	hobbies := []Hobby{
		{
			Name:           "archery",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
		{
			Name:           "carving",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
		{
			Name:           "composing poetry",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "composing stories",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "dancing",
			RequiresOthers: true,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
		{
			Name:           "drinking",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "feasting",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "fishing",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
		{
			Name:           "gambling",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "hawking",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "hunting",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "painting",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
		{
			Name:           "participating in tournaments",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "playing darts",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
		{
			Name:           "playing games of strategy",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
			},
		},
		{
			Name:           "watching tournaments",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
		{
			Name:           "whittling",
			RequiresOthers: false,
			PossibleAgeCategories: []string{
				"adult",
				"elderly",
				"young adult",
				"child",
			},
		},
	}

	return hobbies
}

func getHobbiesForAgeCategory(category age.Category) []Hobby {
	potentialHobbies := getAllHobbies()
	hobbies := []Hobby{}

	for _, h := range potentialHobbies {
		if slices.StringIn(category.Name, h.PossibleAgeCategories) {
			hobbies = append(hobbies, h)
		}
	}

	return hobbies
}

func (character Character) getRandomHobby(ctx context.Context) Hobby {
	hobbies := getHobbiesForAgeCategory(character.AgeCategory)

	if len(hobbies) == 1 {
		return hobbies[0]
	}

	if len(hobbies) > 1 {
		return hobbies[random.Intn(ctx, len(hobbies))]
	}

	return Hobby{}
}
