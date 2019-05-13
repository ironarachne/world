package heavens

// Heavens is the composition of the sky
type Heavens struct {
	Suns  []Sun
	Moons []Moon
	Stars []Star
}

// Generate procedurally generates a Heavens
func Generate() Heavens {
	heavens := Heavens{}

	heavens.Suns = getRandomSuns()
	heavens.Moons = getRandomMoons()
	heavens.Stars = getRandomStars()

	return heavens
}
