package animal

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
)

// Fish is a fish
type Fish struct {
	Name             string
	PluralName       string
	IdealTemperature int
	LivesInLakes     bool
	LivesInOceans    bool
	LivesInRivers    bool
	Resources        []resource.Resource
}

// AllFish returns all predefined fish
func AllFish() []Fish {
	fish := []Fish{}

	fish = append(fish, getLakeFish()...)
	fish = append(fish, getOceanFish()...)
	fish = append(fish, getRiverFish()...)

	for _, f := range fish {
		f.Resources = []resource.Resource{
			{
				Name:        f.Name,
				Origin:      f.Name,
				Type:        "meat",
				Commonality: 4,
			},
		}
	}

	return fish
}

// InSlice checks to see if a fish is in the given slice
func (fish Fish) InSlice(fishes []Fish) bool {
	isIt := false
	for _, a := range fishes {
		if a.Name == fish.Name {
			isIt = true
		}
	}

	return isIt
}

// RandomFish returns a random fish from a slice
func RandomFish(amount int, from []Fish) []Fish {
	var fish Fish

	randomFish := []Fish{}

	for i := 0; i < amount; i++ {
		fish = from[rand.Intn(len(from)-1)]
		if !fish.InSlice(randomFish) {
			randomFish = append(randomFish, fish)
		} else {
			i--
		}
	}

	return randomFish
}

func getLakeFish() []Fish {
	fish := []Fish{
		{
			Name:             "largemouth bass",
			PluralName:       "largemouth bass",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "channel catfish",
			PluralName:       "channel catfish",
			IdealTemperature: 7,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "bluegill",
			PluralName:       "bluegills",
			IdealTemperature: 7,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "coppernose bluegill",
			PluralName:       "coppernose bluegills",
			IdealTemperature: 7,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "sunfish",
			PluralName:       "sunfish",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "gizzard shad",
			PluralName:       "gizzard shads",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "golden shiner",
			PluralName:       "golden shiners",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "fathead minnow",
			PluralName:       "fathead minnows",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "mosquitofish",
			PluralName:       "mosquitofish",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "striped bass",
			PluralName:       "striped bass",
			IdealTemperature: 7,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "threadfin shad",
			PluralName:       "threadfin shads",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "tilapia",
			PluralName:       "tilapia",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "grass carp",
			PluralName:       "grass carp",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "koi",
			PluralName:       "koi",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "tiger muskie",
			PluralName:       "tiger muskies",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "muskellunge",
			PluralName:       "muskellunges",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "yellow perch",
			PluralName:       "yellow perches",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "black crappie",
			PluralName:       "black crappie",
			IdealTemperature: 7,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "white sucker",
			PluralName:       "white suckers",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "rainbow trout",
			PluralName:       "rainbow trouts",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "golden rainbow trout",
			PluralName:       "golden rainbow trouts",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "brook trout",
			PluralName:       "brook trouts",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "brown trout",
			PluralName:       "brown trouts",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
		{
			Name:             "tiger trout",
			PluralName:       "tiger trout",
			IdealTemperature: 5,
			LivesInLakes:     true,
			LivesInOceans:    false,
			LivesInRivers:    false,
		},
	}

	return fish
}

func getOceanFish() []Fish {
	fish := []Fish{
		{
			Name:             "bluefin tuna",
			PluralName:       "bluefin tuna",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "cod",
			PluralName:       "cod",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "goliath grouper",
			PluralName:       "goliath groupers",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "salmon",
			PluralName:       "salmon",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "trumpetfish",
			PluralName:       "trumpetfish",
			IdealTemperature: 7,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "wolffish",
			PluralName:       "wolffish",
			IdealTemperature: 7,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "butterflyfish",
			PluralName:       "butterflyfish",
			IdealTemperature: 7,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "beluga sturgeon",
			PluralName:       "beluga sturgeons",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "blue marlin",
			PluralName:       "blue marlins",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "blue tang",
			PluralName:       "blue tangs",
			IdealTemperature: 7,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "mackerel",
			PluralName:       "mackerels",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "clown triggerfish",
			PluralName:       "clown triggerfish",
			IdealTemperature: 8,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "fangtooth",
			PluralName:       "fangtooth",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "clownfish",
			PluralName:       "clownfish",
			IdealTemperature: 8,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "pufferfish",
			PluralName:       "pufferfish",
			IdealTemperature: 7,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "red snapper",
			PluralName:       "red snappers",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "ocean sunfish",
			PluralName:       "ocean sunfish",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "halibut",
			PluralName:       "halibut",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "sardine",
			PluralName:       "sardines",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "sailfish",
			PluralName:       "sailfish",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "moray eel",
			PluralName:       "moray eels",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "swordfish",
			PluralName:       "swordfish",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "summer flounder",
			PluralName:       "summer flounders",
			IdealTemperature: 7,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "flying fish",
			PluralName:       "flying fish",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
		{
			Name:             "yellowfin tuna",
			PluralName:       "yellowfin tuna",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    true,
			LivesInRivers:    false,
		},
	}

	return fish
}

func getRiverFish() []Fish {
	fish := []Fish{
		{
			Name:             "smallmouth bass",
			PluralName:       "smallmouth bass",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "white bass",
			PluralName:       "white bass",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "silver carp",
			PluralName:       "silver carp",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "bighead carp",
			PluralName:       "bighead carp",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "black carp",
			PluralName:       "black carp",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "carp",
			PluralName:       "carp",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "flathead catfish",
			PluralName:       "flathead catfish",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "white crappie",
			PluralName:       "white crappies",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "eel",
			PluralName:       "eels",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "long-nosed gar",
			PluralName:       "long-nosed gars",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "skipjack herring",
			PluralName:       "skipjack herrings",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "paddlefish",
			PluralName:       "paddlefish",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "northern pike",
			PluralName:       "northern pike",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "pumkinseed",
			PluralName:       "pumkinseeds",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "river redhorse",
			PluralName:       "river redhorse",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "sauger",
			PluralName:       "saugers",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "shovelnose sturgeon",
			PluralName:       "shovelnose sturgeons",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
		{
			Name:             "walleye",
			PluralName:       "walleye",
			IdealTemperature: 5,
			LivesInLakes:     false,
			LivesInOceans:    false,
			LivesInRivers:    true,
		},
	}

	return fish
}
