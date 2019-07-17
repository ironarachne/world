package monster

import "math/rand"

// Monster is a monstrous creature or other intelligent threat
type Monster struct {
	Name                 string
	PluralName           string
	Description          string
	SizeCategory         string
	Type                 string
	IdealTemperature     int
	TemperatureTolerance int
	NumAppearing         func() int
}

func getAllMonsters() []Monster {
	monsters := []Monster{}

	humanoids := getAllHumanoids()

	monsters = append(monsters, humanoids...)

	return monsters
}

// Random returns a random monster
func Random() Monster {
	monsters := getAllMonsters()

	monster := monsters[rand.Intn(len(monsters))]

	return monster
}
