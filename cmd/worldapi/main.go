package main

import (
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
)

// ErrorMessage is an HTTP error message
type ErrorMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
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
	// Sentry.io configuration
	sentryDSN := os.Getenv("SENTRY_DSN")

	sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	})

	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	// Seed randomly based on the current time
	rand.Seed(time.Now().UnixNano())

	// Set up router
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	r.Get("/buildingstyle", sentryHandler.HandleFunc(getBuildingStyleRandom))
	r.Get("/buildingstyle/{id}", sentryHandler.HandleFunc(getBuildingStyle))

	r.Get("/character", sentryHandler.HandleFunc(getCharacterRandom))
	r.Get("/character/{id}", sentryHandler.HandleFunc(getCharacter))

	r.Get("/clothingstyle", sentryHandler.HandleFunc(getClothingStyleRandom))
	r.Get("/clothingstyle/{id}", sentryHandler.HandleFunc(getClothingStyle))

	r.Get("/country", sentryHandler.HandleFunc(getCountryRandom))
	r.Get("/country/{id}", sentryHandler.HandleFunc(getCountry))

	r.Get("/culture", sentryHandler.HandleFunc(getCultureRandom))
	r.Get("/culture/{id}", sentryHandler.HandleFunc(getCulture))
	r.Post("/culture", sentryHandler.HandleFunc(getCultureFromArea))

	r.Get("/data/animals", sentryHandler.HandleFunc(dataAnimals))
	r.Get("/data/domains", sentryHandler.HandleFunc(dataDomains))
	r.Get("/data/fish", sentryHandler.HandleFunc(dataFish))
	r.Get("/data/heraldry/charges", sentryHandler.HandleFunc(dataCharges))
	r.Get("/data/heraldry/charges/tags", sentryHandler.HandleFunc(dataChargeTags))
	r.Get("/data/insects", sentryHandler.HandleFunc(dataInsects))
	r.Get("/data/minerals", sentryHandler.HandleFunc(dataMinerals))
	r.Get("/data/monsters", sentryHandler.HandleFunc(dataMonsters))
	r.Get("/data/patterns", sentryHandler.HandleFunc(dataPatterns))
	r.Get("/data/plants", sentryHandler.HandleFunc(dataPlants))
	r.Get("/data/professions", sentryHandler.HandleFunc(dataProfessions))
	r.Get("/data/races", sentryHandler.HandleFunc(dataRaces))
	r.Get("/data/soils", sentryHandler.HandleFunc(dataSoils))
	r.Get("/data/trees", sentryHandler.HandleFunc(dataTrees))

	r.Get("/foodstyle", sentryHandler.HandleFunc(getFoodStyleRandom))
	r.Get("/foodstyle/{id}", sentryHandler.HandleFunc(getFoodStyle))

	r.Get("/geography", sentryHandler.HandleFunc(getGeographicAreaRandom))
	r.Get("/geography/{id}", sentryHandler.HandleFunc(getGeographicArea))

	r.Get("/heavens", sentryHandler.HandleFunc(getHeavensRandom))
	r.Get("/heavens/{id}", sentryHandler.HandleFunc(getHeavens))

	r.Get("/heraldry", sentryHandler.HandleFunc(getHeraldryRandom))
	r.Get("/heraldry/{id}", sentryHandler.HandleFunc(getHeraldry))

	r.Get("/language", sentryHandler.HandleFunc(getLanguageRandom))
	r.Get("/language/{id}", sentryHandler.HandleFunc(getLanguage))

	r.Get("/merchant", sentryHandler.HandleFunc(getMerchantRandom))
	r.Get("/merchant/{id}", sentryHandler.HandleFunc(getMerchant))

	r.Get("/organization", sentryHandler.HandleFunc(getOrganizationRandom))
	r.Get("/organization/{id}", sentryHandler.HandleFunc(getOrganization))

	r.Get("/pantheon", sentryHandler.HandleFunc(getPantheonRandom))
	r.Get("/pantheon/{id}", sentryHandler.HandleFunc(getPantheon))

	r.Get("/race", sentryHandler.HandleFunc(getRaceRandom))
	r.Get("/race/{id}", sentryHandler.HandleFunc(getRace))

	r.Get("/region", sentryHandler.HandleFunc(getRegionRandom))
	r.Get("/region/{id}", sentryHandler.HandleFunc(getRegion))
	r.Post("/region", sentryHandler.HandleFunc(getRegionFromCulture))

	r.Get("/religion", sentryHandler.HandleFunc(getReligionRandom))
	r.Get("/religion/{id}", sentryHandler.HandleFunc(getReligion))

	r.Get("/town", sentryHandler.HandleFunc(getTownRandom))
	r.Get("/town/{id}", sentryHandler.HandleFunc(getTown))

	r.Get("/worldmap", sentryHandler.HandleFunc(getWorldMapRandom))
	r.Get("/worldmap/{id}", sentryHandler.HandleFunc(getWorldMap))

	r.Get("/", sentryHandler.HandleFunc(getRoot))

	domainName := os.Getenv("WORLDAPI_WEB_DOMAIN")
	dataDir := os.Getenv("WORLDAPI_DATA_PATH")
	saveDir := os.Getenv("WORLDAPI_SAVE_DIRECTORY")
	saveTarget := os.Getenv("WORLDAPI_SAVE_TARGET")

	// Initialize and start
	port := 7531
	fmt.Printf("World API is running on http://localhost:%d.\n", port)
	fmt.Printf("\nWorld API configuration:\n")
	fmt.Printf("Domain Name: %s\n", domainName)
	fmt.Printf("Data directory: %s\n", dataDir)
	fmt.Printf("Save directory: %s\n", saveDir)
	fmt.Printf("Save target: %s\n", saveTarget)
	fmt.Printf("\nHappy generating!\n\n")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
