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

func getBuildingStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o buildings.SimplifiedBuildingStyle

	random.SeedFromString(id)

	o = buildings.GenerateStyle().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getBuildingStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o buildings.SimplifiedBuildingStyle

	rand.Seed(time.Now().UnixNano())

	o = buildings.GenerateStyle().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o character.SimplifiedCharacter

	random.SeedFromString(id)

	o = character.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getCharacterRandom(w http.ResponseWriter, r *http.Request) {
	var o character.SimplifiedCharacter

	rand.Seed(time.Now().UnixNano())

	o = character.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getClimate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o climate.SimplifiedClimate

	random.SeedFromString(id)

	o = climate.Generate().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getClimateRandom(w http.ResponseWriter, r *http.Request) {
	var o climate.SimplifiedClimate

	rand.Seed(time.Now().UnixNano())

	o = climate.Generate().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getClothingStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o clothing.Style

	random.SeedFromString(id)

	o = clothing.Random()

	json.NewEncoder(w).Encode(o)
}

func getClothingStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o clothing.Style

	rand.Seed(time.Now().UnixNano())

	o = clothing.Random()

	json.NewEncoder(w).Encode(o)
}

func getCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o country.Country

	random.SeedFromString(id)

	o = country.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCountryRandom(w http.ResponseWriter, r *http.Request) {
	var o country.Country

	rand.Seed(time.Now().UnixNano())

	o = country.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCulture(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o culture.SimplifiedCulture

	random.SeedFromString(id)

	o = culture.Random().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getCultureRandom(w http.ResponseWriter, r *http.Request) {
	var o culture.SimplifiedCulture

	rand.Seed(time.Now().UnixNano())

	o = culture.Random().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getFoodStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o food.Style

	random.SeedFromString(id)

	o = food.Random()

	json.NewEncoder(w).Encode(o)
}

func getFoodStyleRandom(w http.ResponseWriter, r *http.Request) {
	var o food.Style

	rand.Seed(time.Now().UnixNano())

	o = food.Random()

	json.NewEncoder(w).Encode(o)
}

func getHeavens(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o heavens.Heavens

	random.SeedFromString(id)

	o = heavens.Generate()

	json.NewEncoder(w).Encode(o)
}

func getHeavensRandom(w http.ResponseWriter, r *http.Request) {
	var o heavens.Heavens

	rand.Seed(time.Now().UnixNano())

	o = heavens.Generate()

	json.NewEncoder(w).Encode(o)
}

func getHeraldry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o heraldry.Heraldry

	random.SeedFromString(id)

	o = heraldry.GenerateHeraldry()

	json.NewEncoder(w).Encode(o)
}

func getHeraldryRandom(w http.ResponseWriter, r *http.Request) {
	var o heraldry.Heraldry

	rand.Seed(time.Now().UnixNano())

	o = heraldry.GenerateHeraldry()

	json.NewEncoder(w).Encode(o)
}

func getLanguage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o language.SimplifiedLanguage

	random.SeedFromString(id)

	o = language.Generate().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getLanguageRandom(w http.ResponseWriter, r *http.Request) {
	var o language.SimplifiedLanguage

	rand.Seed(time.Now().UnixNano())

	o = language.Generate().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o merchant.SimplifiedMerchant

	random.SeedFromString(id)

	o = merchant.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getMerchantRandom(w http.ResponseWriter, r *http.Request) {
	var o merchant.SimplifiedMerchant

	rand.Seed(time.Now().UnixNano())

	o = merchant.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getMonster(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o monster.SimplifiedMonster

	random.SeedFromString(id)

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

	random.SeedFromString(id)

	o = organization.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getOrganizationRandom(w http.ResponseWriter, r *http.Request) {
	var o organization.SimplifiedOrganization

	rand.Seed(time.Now().UnixNano())

	o = organization.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getPantheon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o pantheon.SimplifiedPantheon
	var l language.Language

	random.SeedFromString(id)

	l = language.Generate()
	o = pantheon.Generate(6, 15, l).Simplify()

	json.NewEncoder(w).Encode(o)
}

func getPantheonRandom(w http.ResponseWriter, r *http.Request) {
	var o pantheon.SimplifiedPantheon
	var l language.Language

	rand.Seed(time.Now().UnixNano())

	l = language.Generate()
	o = pantheon.Generate(6, 15, l).Simplify()

	json.NewEncoder(w).Encode(o)
}

func getRace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o race.SimplifiedRace

	random.SeedFromString(id)

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

	random.SeedFromString(id)

	o = region.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getRegionRandom(w http.ResponseWriter, r *http.Request) {
	var o region.SimplifiedRegion

	rand.Seed(time.Now().UnixNano())

	o = region.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getReligion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o religion.SimplifiedReligion

	random.SeedFromString(id)

	o = religion.Random().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getReligionRandom(w http.ResponseWriter, r *http.Request) {
	var o religion.SimplifiedReligion

	rand.Seed(time.Now().UnixNano())

	o = religion.Random().Simplify()

	json.NewEncoder(w).Encode(o)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	o := "This is the World Generation API."

	json.NewEncoder(w).Encode(o)
}

func getTown(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o town.SimplifiedTown

	random.SeedFromString(id)

	o = town.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getTownRandom(w http.ResponseWriter, r *http.Request) {
	var o town.SimplifiedTown

	rand.Seed(time.Now().UnixNano())

	o = town.RandomSimplified()

	json.NewEncoder(w).Encode(o)
}

func getWorld(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o world.World

	random.SeedFromString(id)

	o = world.Generate()

	json.NewEncoder(w).Encode(o)
}

func getWorldRandom(w http.ResponseWriter, r *http.Request) {
	var o world.World

	rand.Seed(time.Now().UnixNano())

	o = world.Generate()

	json.NewEncoder(w).Encode(o)
}

func getWorldMap(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var l world.World

	random.SeedFromString(id)

	l = world.Generate()

	json.NewEncoder(w).Encode(l.WorldMap)
}

func getWorldMapSVGImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var l world.World

	random.SeedFromString(id)

	l = world.Generate()
	o := l.WorldMap.RenderAsSVG()

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(o))
}

func getWorldMapTextImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var l world.World

	random.SeedFromString(id)

	l = world.Generate()
	o := l.WorldMap.RenderAsText()

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(o))
}

func getWorldMapRandom(w http.ResponseWriter, r *http.Request) {
	var l world.World

	rand.Seed(time.Now().UnixNano())

	l = world.Generate()

	json.NewEncoder(w).Encode(l.WorldMap)
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

	if sentryDSN == "" {
		panic("Sentry not set up!")
	}

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
