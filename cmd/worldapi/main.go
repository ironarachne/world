package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/country"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/heavens"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/merchant"
	"github.com/ironarachne/world/pkg/monster"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/religion"
	"github.com/ironarachne/world/pkg/town"
	"github.com/ironarachne/world/pkg/world"
)

// ErrorMessage is an HTTP error message
type ErrorMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
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

	json.NewEncoder(w).Encode(o)
}

func getBuildingStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o buildings.SimplifiedBuildingStyle

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

	o, err := character.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
}

func getClimate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o climate.SimplifiedClimate

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	randomClimate, err := climate.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = randomClimate.Simplify()

	json.NewEncoder(w).Encode(o)
}

func getClimateRandom(w http.ResponseWriter, r *http.Request) {
	var o climate.SimplifiedClimate

	rand.Seed(time.Now().UnixNano())

	randomClimate, err := climate.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = randomClimate.Simplify()

	json.NewEncoder(w).Encode(o)
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

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

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

func getCultureRandom(w http.ResponseWriter, r *http.Request) {
	var o culture.SimplifiedCulture

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

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

	randomLanguage, err := language.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = randomLanguage.Simplify()

	json.NewEncoder(w).Encode(o)
}

func getLanguageRandom(w http.ResponseWriter, r *http.Request) {
	var o language.SimplifiedLanguage

	rand.Seed(time.Now().UnixNano())

	randomLanguage, err := language.Generate()
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

	rand.Seed(time.Now().UnixNano())

	o, err := merchant.RandomSimplified()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(o)
}

func getMonster(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o monster.SimplifiedMonster

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o = monster.Random().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getMonsterRandom(w http.ResponseWriter, r *http.Request) {
	var o monster.SimplifiedMonster

	rand.Seed(time.Now().UnixNano())

	o = monster.Random().Simplify()

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

	rand.Seed(time.Now().UnixNano())

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

	l, err = language.Generate()
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

	rand.Seed(time.Now().UnixNano())

	l, err := language.Generate()
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

	var o race.SimplifiedRace

	err := random.SeedFromString(id)
	if err != nil {
		handleError(w, r, err)
		return
	}

	o = race.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getRaceRandom(w http.ResponseWriter, r *http.Request) {
	var o race.SimplifiedRace

	rand.Seed(time.Now().UnixNano())

	o = race.RandomSimplified()

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

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

	rel, err := religion.Random()
	if err != nil {
		handleError(w, r, err)
		return
	}
	o = rel.Simplify()

	json.NewEncoder(w).Encode(o)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	o := "This is the World Generation API."

	json.NewEncoder(w).Encode(o)
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

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

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

	rand.Seed(time.Now().UnixNano())

	l, err := world.Generate()
	if err != nil {
		handleError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(l.WorldMap)
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	sentry.CaptureException(err)
	log.Println(err)
	http.Error(w, "There was a problem processing your request.", 500)
}

func errorTracking(h http.Handler) http.Handler {
	sentryDSN := os.Getenv("SENTRY_DSN")

	if sentryDSN != "" {
		sentry.Init(sentry.ClientOptions{
			Dsn: sentryDSN,
		})

		sentryHandler := sentryhttp.New(sentryhttp.Options{})

		fn := func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}

		return sentryHandler.HandleFunc(fn)
	}

	return h
}

func main() {
	sentryDSN := os.Getenv("SENTRY_DSN")

	sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	})

	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/buildingstyle", sentryHandler.HandleFunc(getBuildingStyleRandom))
	r.Get("/buildingstyle/{id}", sentryHandler.HandleFunc(getBuildingStyle))

	r.Get("/character", sentryHandler.HandleFunc(getCharacterRandom))
	r.Get("/character/{id}", sentryHandler.HandleFunc(getCharacter))

	r.Get("/climate", sentryHandler.HandleFunc(getClimateRandom))
	r.Get("/climate/{id}", sentryHandler.HandleFunc(getClimate))

	r.Get("/clothingstyle", sentryHandler.HandleFunc(getClothingStyleRandom))
	r.Get("/clothingstyle/{id}", sentryHandler.HandleFunc(getClothingStyle))

	r.Get("/country", sentryHandler.HandleFunc(getCountryRandom))
	r.Get("/country/{id}", sentryHandler.HandleFunc(getCountry))

	r.Get("/culture", sentryHandler.HandleFunc(getCultureRandom))
	r.Get("/culture/{id}", sentryHandler.HandleFunc(getCulture))

	r.Get("/foodstyle", sentryHandler.HandleFunc(getFoodStyleRandom))
	r.Get("/foodstyle/{id}", sentryHandler.HandleFunc(getFoodStyle))

	r.Get("/heavens", sentryHandler.HandleFunc(getHeavensRandom))
	r.Get("/heavens/{id}", sentryHandler.HandleFunc(getHeavens))

	r.Get("/heraldry", sentryHandler.HandleFunc(getHeraldryRandom))
	r.Get("/heraldry/{id}", sentryHandler.HandleFunc(getHeraldry))

	r.Get("/language", sentryHandler.HandleFunc(getLanguageRandom))
	r.Get("/language/{id}", sentryHandler.HandleFunc(getLanguage))

	r.Get("/merchant", sentryHandler.HandleFunc(getMerchantRandom))
	r.Get("/merchant/{id}", sentryHandler.HandleFunc(getMerchant))

	r.Get("/monster", sentryHandler.HandleFunc(getMonsterRandom))
	r.Get("/monster/{id}", sentryHandler.HandleFunc(getMonster))

	r.Get("/organization", sentryHandler.HandleFunc(getOrganizationRandom))
	r.Get("/organization/{id}", sentryHandler.HandleFunc(getOrganization))

	r.Get("/pantheon", sentryHandler.HandleFunc(getPantheonRandom))
	r.Get("/pantheon/{id}", sentryHandler.HandleFunc(getPantheon))

	r.Get("/race", sentryHandler.HandleFunc(getRaceRandom))
	r.Get("/race/{id}", sentryHandler.HandleFunc(getRace))

	r.Get("/region", sentryHandler.HandleFunc(getRegionRandom))
	r.Get("/region/{id}", sentryHandler.HandleFunc(getRegion))

	r.Get("/religion", sentryHandler.HandleFunc(getReligionRandom))
	r.Get("/religion/{id}", sentryHandler.HandleFunc(getReligion))

	r.Get("/town", sentryHandler.HandleFunc(getTownRandom))
	r.Get("/town/{id}", sentryHandler.HandleFunc(getTown))

	r.Get("/world", sentryHandler.HandleFunc(getWorldRandom))
	r.Get("/world/{id}", sentryHandler.HandleFunc(getWorld))
	r.Get("/world/{id}/map", sentryHandler.HandleFunc(getWorldMapSVGImage))

	r.Get("/worldmap", sentryHandler.HandleFunc(getWorldMapRandom))
	r.Get("/worldmap/{id}", sentryHandler.HandleFunc(getWorldMap))
	r.Get("/worldmap/{id}/image", sentryHandler.HandleFunc(getWorldMapSVGImage))
	r.Get("/worldmap/{id}/textimage", sentryHandler.HandleFunc(getWorldMapTextImage))

	r.Get("/", sentryHandler.HandleFunc(getRoot))

	fmt.Println("World Generator API is online.")
	log.Fatal(http.ListenAndServe(":7531", r))
}
