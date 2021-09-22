package character

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getAllMotivations() []string {
	motivations := []string{
		"a desire to better oneself",
		"achieving destiny",
		"ambition",
		"balance in all things",
		"conquest",
		"conspiracy",
		"control",
		"creativity",
		"debauchery",
		"desperation",
		"distinguishing oneself",
		"duty",
		"escaping destiny",
		"fairness",
		"faith",
		"fame",
		"family",
		"fear",
		"freedom",
		"friends",
		"gaining acceptance",
		"greatness",
		"greed",
		"happiness",
		"hate",
		"honor",
		"justice",
		"knowledge",
		"laziness",
		"love",
		"loyalty",
		"money",
		"pleasure",
		"power",
		"religious devotion",
		"revenge",
		"rivalry",
		"romance",
		"safety",
		"stability",
		"survival",
		"the pursuit of perfection",
		"truth",
	}

	return motivations
}

func getRandomMotivation(ctx context.Context) (string, error) {
	motivations := getAllMotivations()

	motivation, err := random.String(ctx, motivations)
	if err != nil {
		err = fmt.Errorf("Could not generate face shape: %w", err)
		return "", err
	}
	return motivation, nil
}
