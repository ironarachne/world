/*
Package monster provides monster implementations of the species.Species struct
*/
package monster

import "github.com/ironarachne/world/pkg/species"

func getAllMonsters() []species.Species {
	monsters := []species.Species{}

	draconids := getAllDraconids()
	giants := getAllGiants()
	humanoids := getAllHumanoids()

	monsters = append(monsters, draconids...)
	monsters = append(monsters, giants...)
	monsters = append(monsters, humanoids...)

	return monsters
}
