package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/conlang"
	"github.com/ironarachne/world/pkg/country"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/geography"
	"github.com/ironarachne/world/pkg/heavens"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/merchant"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/religion"
	"github.com/ironarachne/world/pkg/town"
)

const contentType = "Content-Type"

func getBuildingStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	buildingStyle, err := buildings.GenerateStyle(random.WithSeed(r.Context(), id))
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := buildingStyle.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getBuildingStyleRandom(w http.ResponseWriter, r *http.Request) {
	buildingStyle, err := buildings.GenerateStyle(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := buildingStyle.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	o, err := character.RandomSimplified(random.WithSeed(r.Context(), id))
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
	o, err := character.RandomSimplified(r.Context())
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

func getClothingStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	o, err := clothing.Random(random.WithSeed(r.Context(), id))
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
	o, err := clothing.Random(r.Context())
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
	o, err := country.Generate(random.WithSeed(r.Context(), id))
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
	o, err := country.Generate(r.Context())
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
	randomCulture, err := culture.Random(random.WithSeed(r.Context(), id))
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

func getCultureFromArea(w http.ResponseWriter, r *http.Request) {
	var area geography.Area

	err := json.NewDecoder(r.Body).Decode(&area)
	if err != nil {
		handleError(w, r, err)
		return
	}

	cul, err := culture.Generate(r.Context(), area)
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
	randomCulture, err := culture.Random(r.Context())
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
	o, err := food.Random(random.WithSeed(r.Context(), id))
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
	o, err := food.Random(r.Context())
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

func getGeographicArea(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	o, err := geography.Generate(random.WithSeed(r.Context(), id))
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

func getGeographicAreaRandom(w http.ResponseWriter, r *http.Request) {
	o, err := geography.Generate(r.Context())
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
	o, err := heavens.Generate(random.WithSeed(r.Context(), id))
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
	o, err := heavens.Generate(r.Context())
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
	var chargeTag string
	var fieldType string

	id := chi.URLParam(r, "id")

	fieldTypes, ok := r.URL.Query()["shape"]

	if !ok || len(fieldTypes[0]) < 1 {
		fieldType = ""
	} else {
		fieldType = fieldTypes[0]
	}

	chargeTags, ok := r.URL.Query()["tag"]

	if !ok || len(chargeTags[0]) < 1 {
		chargeTag = ""
	} else {
		chargeTag = chargeTags[0]
	}
	o, err := heraldry.GenerateByParameters(random.WithSeed(r.Context(), id), fieldType, chargeTag)
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
	var chargeTag string
	var fieldType string

	fieldTypes, ok := r.URL.Query()["shape"]

	if !ok || len(fieldTypes[0]) < 1 {
		fieldType = ""
	} else {
		fieldType = fieldTypes[0]
	}

	chargeTags, ok := r.URL.Query()["tag"]

	if !ok || len(chargeTags[0]) < 1 {
		chargeTag = ""
	} else {
		chargeTag = chargeTags[0]
	}
	o, err := heraldry.GenerateByParameters(r.Context(), fieldType, chargeTag)
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
	randomLanguage, _, err := conlang.Generate(random.WithSeed(r.Context(), id))
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
	randomLanguage, _, err := conlang.Generate(r.Context())
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
	o, err := merchant.RandomSimplified(random.WithSeed(r.Context(), id))
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
	o, err := merchant.RandomSimplified(r.Context())
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
	o, err := organization.RandomSimplified(random.WithSeed(r.Context(), id))
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
	o, err := organization.RandomSimplified(r.Context())
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
	ctx := random.WithSeed(r.Context(), id)
	l, _, err := conlang.Generate(ctx)
	if err != nil {
		handleError(w, r, err)
		return
	}
	p, err := pantheon.Generate(ctx, 6, 15, l)
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := p.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getPantheonRandom(w http.ResponseWriter, r *http.Request) {
	l, _, err := conlang.Generate(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}
	p, err := pantheon.Generate(r.Context(), 6, 15, l)
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := p.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getRace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	o, err := race.RandomSimplified(random.WithSeed(r.Context(), id))
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

	o, err := race.RandomSimplified(r.Context())
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
	ctx := random.WithSeed(r.Context(), id)
	o, err := region.Random(ctx)
	if err != nil {
		handleError(w, r, err)
		return
	}

	so, err := o.Simplify(ctx)
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

	area, err := geography.Generate(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}

	reg, err := region.Generate(r.Context(), area, cul)
	if err != nil {
		handleError(w, r, err)
		return
	}

	sr, err := reg.Simplify(r.Context())
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
	o, err := region.Random(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}

	so, err := o.Simplify(r.Context())
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
	rel, err := religion.Random(random.WithSeed(r.Context(), id))
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := rel.Simplify()

	err = json.NewEncoder(w).Encode(o)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getReligionRandom(w http.ResponseWriter, r *http.Request) {
	rel, err := religion.Random(r.Context())
	if err != nil {
		handleError(w, r, err)
		return
	}
	o := rel.Simplify()

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
	_, err := w.Write([]byte(str.String()))
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func getTown(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	o, err := town.RandomSimplified(random.WithSeed(r.Context(), id))
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
	o, err := town.RandomSimplified(r.Context())
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
