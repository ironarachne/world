/*
Package heavens provides structures and tools for procedurally generating
suns, moons, and stars.
*/
package heavens

import (
	"context"
	"fmt"
)

const heavensError = "failed to generate random heavens: %w"

// Heavens is the composition of the sky
type Heavens struct {
	Suns  []Sun
	Moons []Moon
	Stars []Star
}

// Generate procedurally generates a Heavens
func Generate(ctx context.Context) (Heavens, error) {
	heavens := Heavens{}

	suns, err := getRandomSuns(ctx)
	if err != nil {
		err = fmt.Errorf(heavensError, err)
		return Heavens{}, err
	}
	heavens.Suns = suns
	moons, err := getRandomMoons(ctx)
	if err != nil {
		err = fmt.Errorf(heavensError, err)
		return Heavens{}, err
	}
	heavens.Moons = moons
	stars, err := getRandomStars(ctx)
	if err != nil {
		err = fmt.Errorf(heavensError, err)
		return Heavens{}, err
	}
	heavens.Stars = stars

	return heavens, nil
}
