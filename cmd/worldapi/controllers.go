package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/conlang"
	"github.com/ironarachne/world/pkg/country"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/fish"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/heavens"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/heraldry/charge"
	"github.com/ironarachne/world/pkg/insect"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/merchant"
	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/monster"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/pantheon/domain"
	"github.com/ironarachne/world/pkg/plant"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/religion"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/town"
	"github.com/ironarachne/world/pkg/tree"
	"github.com/ironarachne/world/pkg/world"
)

const contentType = "Content-Type"

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
	patterns, err := resource.All()
	if err != nil {
		handleError(w, r, err)
		return
	}

	d := resource.Data{
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

func getBuildingStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o buildings.SimplifiedBuildingStyle

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	buildingStyle, err := buildings.GenerateStyle()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = buildingStyle.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getBuildingStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o buildings.SimplifiedBuildingStyle

	buildingStyle, err := buildings.GenerateStyle()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = buildingStyle.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o character.SimplifiedCharacter

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = character.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCharacterRandom(w http.ResponseWriter, r *http.Request) {
	var o character.SimplifiedCharacter

	o, err := character.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getClimate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	randomClimate, err := climate.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(randomClimate)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getClimateRandom(w http.ResponseWriter, r *http.Request) {
	randomClimate, err := climate.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(randomClimate)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getClothingStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o clothing.Style

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = clothing.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getClothingStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o clothing.Style

	o, err := clothing.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o country.Country

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = country.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCountryRandom(w http.ResponseWriter, r *http.Request) {
	var o country.Country

	o, err := country.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCulture(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	randomCulture, err := culture.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(randomCulture)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCultureFromClimate(w http.ResponseWriter, r *http.Request) {
	var clm climate.Climate

	err := json.NewDecoder(r.Body).Decode(&clm)
	if err != nil {
		handleError(w, r, err)
		return
	}

	cul, err := culture.Generate(clm)
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(cul)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCultureRandom(w http.ResponseWriter, r *http.Request) {
	randomCulture, err := culture.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(randomCulture)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getFoodStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o food.Style

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = food.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getFoodStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o food.Style

	o, err := food.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getHeavens(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o heavens.Heavens

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = heavens.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getHeavensRandom(w http.ResponseWriter, r *http.Request) {
	var o heavens.Heavens

	o, err := heavens.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getHeraldry(w http.ResponseWriter, r *http.Request) {
	var fieldType string
	id := chi.URLParam(r, "id")

	fieldTypes, ok := r.URL.Query()["shape"]

	if !ok || len(fieldTypes[0]) < 1 {
		fieldType = ""
	} else {
		fieldType = fieldTypes[0]
	}

	var o heraldry.Device

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if fieldType == "" {
		o, err = heraldry.Generate()
	} else {
		o, err = heraldry.GenerateByFieldName(fieldType)
	}
	if err != nil {
		handleError(w, r, err)
		return
	}

	sd := o.Simplify()

	err = json.NewEncoder(w).Encode(sd)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getHeraldryRandom(w http.ResponseWriter, r *http.Request) {
	var fieldType string
	fieldTypes, ok := r.URL.Query()["shape"]

	if !ok || len(fieldTypes[0]) < 1 {
		fieldType = ""
	} else {
		fieldType = fieldTypes[0]
	}

	var o heraldry.Device
	var err error

	if fieldType == "" {
		o, err = heraldry.Generate()
	} else {
		o, err = heraldry.GenerateByFieldName(fieldType)
	}
	if err != nil {
		handleError(w, r, err)
		return
	}

	sd := o.Simplify()

	err = json.NewEncoder(w).Encode(sd)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getLanguage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	randomLanguage, _, err := conlang.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(randomLanguage)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getLanguageRandom(w http.ResponseWriter, r *http.Request) {
	randomLanguage, _, err := conlang.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(randomLanguage)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o merchant.SimplifiedMerchant

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = merchant.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getMerchantRandom(w http.ResponseWriter, r *http.Request) {
	var o merchant.SimplifiedMerchant

	o, err := merchant.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o organization.SimplifiedOrganization

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = organization.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getOrganizationRandom(w http.ResponseWriter, r *http.Request) {
	var o organization.SimplifiedOrganization

	o, err := organization.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getPantheon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o pantheon.SimplifiedPantheon
	var l language.Language

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	l, _, err = conlang.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}
	p, err := pantheon.Generate(6, 15, l)
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = p.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getPantheonRandom(w http.ResponseWriter, r *http.Request) {
	var o pantheon.SimplifiedPantheon
	var l language.Language

	l, _, err := conlang.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}
	p, err := pantheon.Generate(6, 15, l)
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = p.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getRace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err := race.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getRaceRandom(w http.ResponseWriter, r *http.Request) {

	o, err := race.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getRegion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err := region.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	so, err := o.Simplify()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(so)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getRegionFromCulture(w http.ResponseWriter, r *http.Request) {
	var cul culture.Culture

	err := json.NewDecoder(r.Body).Decode(&cul)
	if err != nil {
		handleError(w, r, err)
		return
	}

	reg, err := region.Generate(cul.HomeClimate, cul)
	if err != nil {
		handleError(w, r, err)
		return
	}

	sr, err := reg.Simplify()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(sr)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getRegionRandom(w http.ResponseWriter, r *http.Request) {
	o, err := region.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	so, err := o.Simplify()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(so)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getReligion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o religion.SimplifiedReligion

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	rel, err := religion.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = rel.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getReligionRandom(w http.ResponseWriter, r *http.Request) {
	var o religion.SimplifiedReligion

	rel, err := religion.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = rel.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"buildingstyle",
		"character",
		"climate",
		"clothingstyle",
		"country",
		"culture",
		"data/animals",
		"data/domains",
		"data/fish",
		"data/heraldry/charges",
		"data/insects",
		"data/minerals",
		"data/monsters",
		"data/patterns",
		"data/plants",
		"data/professions",
		"data/races",
		"data/soils",
		"data/trees",
		"foodstyle",
		"heavens",
		"heraldry",
		"language",
		"merchant",
		"monster",
		"organization",
		"pantheon",
		"race",
		"region",
		"religion",
		"town",
		"world",
		"worldmap",
	}
	var str strings.Builder
	str.WriteString("<p>This is the World Generation API.</p>")
	str.WriteString("<ul>")
	for _, path := range paths {
		str.WriteString(fmt.Sprintf("<li><a href=\"/%s\">/%s</a></li>", path, path))
	}
	str.WriteString("</ul>")

	w.Header().Set(contentType, "text/html")
	w.Write([]byte(str.String()))
}

func getTown(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o town.SimplifiedTown

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = town.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getTownRandom(w http.ResponseWriter, r *http.Request) {
	var o town.SimplifiedTown

	o, err := town.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getWorld(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o world.World

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getWorldRandom(w http.ResponseWriter, r *http.Request) {
	var o world.World

	o, err := world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getWorldMap(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var l world.World

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	l, err = world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(l.WorldMap)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getWorldMapSVGImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var l world.World

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	l, err = world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := l.WorldMap.RenderAsSVG()

	w.Header().Set(contentType, "image/svg+xml")
	w.Write([]byte(o))
}

func getWorldMapTextImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var l world.World

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	l, err = world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := l.WorldMap.RenderAsText()

	w.Header().Set(contentType, "image/svg+xml")
	w.Write([]byte(o))
}

func getWorldMapRandom(w http.ResponseWriter, r *http.Request) {
	var l world.World

	l, err := world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(l.WorldMap)
	if err != nil {
		handleError(w, r, err)
		return
	}
}
