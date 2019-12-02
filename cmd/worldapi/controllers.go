package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/conlang"
	"github.com/ironarachne/world/pkg/country"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/heavens"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/merchant"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/religion"
	"github.com/ironarachne/world/pkg/town"
	"github.com/ironarachne/world/pkg/world"
)

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

	json.NewEncoder(w).Encode(o)
}

func getBuildingStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o buildings.SimplifiedBuildingStyle

	buildingStyle, err := buildings.GenerateStyle()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = buildingStyle.Simplify()

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getCharacterRandom(w http.ResponseWriter, r *http.Request) {
	var o character.SimplifiedCharacter

	o, err := character.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(randomClimate)
}

func getClimateRandom(w http.ResponseWriter, r *http.Request) {
	randomClimate, err := climate.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(randomClimate)
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

	json.NewEncoder(w).Encode(o)
}

func getClothingStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o clothing.Style

	o, err := clothing.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getCountryRandom(w http.ResponseWriter, r *http.Request) {
	var o country.Country

	o, err := country.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
}

func getCulture(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o culture.SimplifiedCulture

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
	o = randomCulture.Simplify()

	json.NewEncoder(w).Encode(o)
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

	simplifiedCulture := cul.Simplify()

	json.NewEncoder(w).Encode(simplifiedCulture)
}

func getCultureRandom(w http.ResponseWriter, r *http.Request) {
	var o culture.SimplifiedCulture

	randomCulture, err := culture.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = randomCulture.Simplify()

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getFoodStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o food.Style

	o, err := food.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getHeavensRandom(w http.ResponseWriter, r *http.Request) {
	var o heavens.Heavens

	o, err := heavens.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(sd)
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

	json.NewEncoder(w).Encode(sd)
}

func getLanguage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o language.SimplifiedLanguage

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
	o = randomLanguage.Simplify()

	json.NewEncoder(w).Encode(o)
}

func getLanguageRandom(w http.ResponseWriter, r *http.Request) {
	var o language.SimplifiedLanguage

	randomLanguage, _, err := conlang.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = randomLanguage.Simplify()

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getMerchantRandom(w http.ResponseWriter, r *http.Request) {
	var o merchant.SimplifiedMerchant

	o, err := merchant.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getOrganizationRandom(w http.ResponseWriter, r *http.Request) {
	var o organization.SimplifiedOrganization

	o, err := organization.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getRaceRandom(w http.ResponseWriter, r *http.Request) {

	o, err := race.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
}

func getRegion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o region.SimplifiedRegion

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o, err = region.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
}

func getRegionRandom(w http.ResponseWriter, r *http.Request) {
	var o region.SimplifiedRegion

	o, err := region.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getReligionRandom(w http.ResponseWriter, r *http.Request) {
	var o religion.SimplifiedReligion

	rel, err := religion.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = rel.Simplify()

	json.NewEncoder(w).Encode(o)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"buildingstyle",
		"character",
		"climate",
		"clothingstyle",
		"country",
		"culture",
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

	w.Header().Set("Content-Type", "text/html")
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

	json.NewEncoder(w).Encode(o)
}

func getTownRandom(w http.ResponseWriter, r *http.Request) {
	var o town.SimplifiedTown

	o, err := town.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(o)
}

func getWorldRandom(w http.ResponseWriter, r *http.Request) {
	var o world.World

	o, err := world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
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

	json.NewEncoder(w).Encode(l.WorldMap)
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

	w.Header().Set("Content-Type", "image/svg+xml")
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

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(o))
}

func getWorldMapRandom(w http.ResponseWriter, r *http.Request) {
	var l world.World

	l, err := world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(l.WorldMap)
}
