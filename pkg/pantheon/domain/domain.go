package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ironarachne/world/pkg/random"
)

// Data is a collection of domains
type Data struct {
	Domains []Domain `json:"domains"`
}

// Domain is an area of control
type Domain struct {
	Name              string   `json:"name"`
	AppearanceTraits  []string `json:"appearance_traits"`
	PersonalityTraits []string `json:"personality_traits"`
	HolyItems         []string `json:"holy_items"`
	HolySymbols       []string `json:"holy_symbols"`
}

// All returns all pre-defined domains
func All() ([]Domain, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/domains.json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Domain{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &d)

	all := d.Domains

	if len(all) == 0 {
		err = fmt.Errorf("no domains returned from database: domains.json")
		return []Domain{}, err
	}

	return all, nil
}

// AllAppearancesForDomains returns a string slice of all the appearances from a set of domains
func AllAppearancesForDomains(domains []Domain) []string {
	var appearances []string

	for _, d := range domains {
		appearances = append(appearances, d.AppearanceTraits...)
	}

	return appearances
}

// AllPersonalitiesForDomains returns a string slice of all the personalities from a set of domains
func AllPersonalitiesForDomains(domains []Domain) []string {
	var personalities []string

	for _, d := range domains {
		personalities = append(personalities, d.PersonalityTraits...)
	}

	return personalities
}

// RandomAppearanceFromDomains returns a random appearance given a set of domains
func RandomAppearanceFromDomains(ctx context.Context, domains []Domain) (string, error) {
	var possibleAppearances []string

	for _, d := range domains {
		possibleAppearances = append(possibleAppearances, d.AppearanceTraits...)
	}

	if len(possibleAppearances) == 0 {
		err := fmt.Errorf("Set of domain appearances is empty")
		return "", err
	}

	if len(possibleAppearances) == 1 {
		return possibleAppearances[0], nil
	}

	appearance := possibleAppearances[random.Intn(ctx, len(possibleAppearances))]

	return appearance, nil
}

// RandomPersonalityFromDomains returns a random personality given a set of domains
func RandomPersonalityFromDomains(ctx context.Context, domains []Domain) (string, error) {
	var possiblePersonalities []string

	for _, d := range domains {
		possiblePersonalities = append(possiblePersonalities, d.PersonalityTraits...)
	}

	if len(possiblePersonalities) == 0 {
		err := fmt.Errorf("Set of domain personalities is empty")
		return "", err
	}

	if len(possiblePersonalities) == 1 {
		return possiblePersonalities[0], nil
	}

	personality := possiblePersonalities[random.Intn(ctx, len(possiblePersonalities))]

	return personality, nil
}

// ByName returns a specific domain by name
func ByName(name string) (Domain, error) {
	domains, err := All()
	if err != nil {
		err = fmt.Errorf("failed to find domain with name: %w", err)
		return Domain{}, err
	}

	for _, d := range domains {
		if d.Name == name {
			return d, nil
		}
	}

	err = fmt.Errorf("failed to find domain with name " + name)
	return Domain{}, err
}

// Random returns a random domain from a slice of domains
func Random(ctx context.Context, domains []Domain) (Domain, error) {
	if len(domains) == 0 {
		err := fmt.Errorf("tried to get a random domain from an empty slice")
		return Domain{}, err
	}

	if len(domains) == 1 {
		return domains[0], nil
	}

	domain := domains[random.Intn(ctx, len(domains))]

	return domain, nil
}

// Exclude removes a slice of domains from the given slice of domains
func Exclude(domainsToRemove []Domain, superset []Domain) []Domain {
	result := []Domain{}

	for _, d := range superset {
		if !InSlice(d, domainsToRemove) {
			result = append(result, d)
		}
	}

	return result
}

// InSlice returns true if a given domain is in a given slice of domains
func InSlice(domain Domain, domains []Domain) bool {
	for _, d := range domains {
		if domain.Name == d.Name {
			return true
		}
	}
	return false
}
