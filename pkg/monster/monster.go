package monster

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/size"
)

// Monster is a monstrous creature or other intelligent threat
type Monster struct {
	Name                 string
	PluralName           string
	Description          string
	SizeCategory         size.Category
	Type                 string
	IdealTemperature     int
	TemperatureTolerance int
	NumAppearing         func() int
}

func getAllMonsters() []Monster {
	monsters := []Monster{}

	draconids := getAllDraconids()
	giants := getAllGiants()
	humanoids := getAllHumanoids()

	monsters = append(monsters, draconids...)
	monsters = append(monsters, giants...)
	monsters = append(monsters, humanoids...)

	return monsters
}

// Random returns a random monster
func Random() Monster {
	monsters := getAllMonsters()

	monster := monsters[rand.Intn(len(monsters))]

	return monster
}
