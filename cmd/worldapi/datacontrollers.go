package main

import (
	"encoding/json"
	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/fish"
	"github.com/ironarachne/world/pkg/heraldry/charge"
	"github.com/ironarachne/world/pkg/insect"
	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/monster"
	"github.com/ironarachne/world/pkg/pantheon/domain"
	"github.com/ironarachne/world/pkg/pattern"
	"github.com/ironarachne/world/pkg/plant"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/tree"
	"net/http"
)

func dataAnimals(w http.ResponseWriter, r *http.Request) {
	animals, err := animal.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := species.Data{
		Species: animals,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataCharges(w http.ResponseWriter, r *http.Request) {
	charges, err := charge.AllRaster()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := charge.Data{
		Charges: charges,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataChargeTags(w http.ResponseWriter, r *http.Request) {
	tags, err := charge.AllTags()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(tags)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataDomains(w http.ResponseWriter, r *http.Request) {
	all, err := domain.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := domain.Data{
		Domains: all,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataFish(w http.ResponseWriter, r *http.Request) {
	all, err := fish.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := species.Data{
		Species: all,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataInsects(w http.ResponseWriter, r *http.Request) {
	all, err := insect.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := species.Data{
		Species: all,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataMinerals(w http.ResponseWriter, r *http.Request) {
	all, err := mineral.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := mineral.Data{
		Minerals: all,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataMonsters(w http.ResponseWriter, r *http.Request) {
	all, err := monster.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := species.Data{
		Species: all,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataPatterns(w http.ResponseWriter, r *http.Request) {
	patterns, err := pattern.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := pattern.Data{
		Patterns: patterns,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataProfessions(w http.ResponseWriter, r *http.Request) {
	professions, err := profession.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := profession.Data{
		Professions: professions,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataPlants(w http.ResponseWriter, r *http.Request) {
	plants, err := plant.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := species.Data{
		Species: plants,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataRaces(w http.ResponseWriter, r *http.Request) {
	all, err := race.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := species.Data{
		Species: all,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataSoils(w http.ResponseWriter, r *http.Request) {
	all, err := soil.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := soil.Data{
		Soils: all,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func dataTrees(w http.ResponseWriter, r *http.Request) {
	trees, err := tree.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := species.Data{
		Species: trees,
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		handleError(w, r, err)
		return
	}
}
