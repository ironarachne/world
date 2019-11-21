/*
Package heavens provides structures and tools for procedurally generating
suns, moons, and stars.
*/
package heavens

import "fmt"

// Heavens is the composition of the sky
type Heavens struct {
	Suns  []Sun
	Moons []Moon
	Stars []Star
}

// Generate procedurally generates a Heavens
func Generate() (Heavens, error) {
	heavens := Heavens{}

	suns, err := getRandomSuns()
	if err != nil {
		err = fmt.Errorf("Failed to generate random heavens: %w", err)
		return Heavens{}, err
	}
	heavens.Suns = suns
	moons, err := getRandomMoons()
	if err != nil {
		err = fmt.Errorf("Failed to generate random heavens: %w", err)
		return Heavens{}, err
	}
	heavens.Moons = moons
	stars, err := getRandomStars()
	if err != nil {
		err = fmt.Errorf("Failed to generate random heavens: %w", err)
		return Heavens{}, err
	}
	heavens.Stars = stars

	return heavens, nil
}
