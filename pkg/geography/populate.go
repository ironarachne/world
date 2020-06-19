package geography

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/fish"
	"github.com/ironarachne/world/pkg/insect"
	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/plant"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/tree"
)

func getAnimals(ctx context.Context, humidity int, temperature int, prevalence int, tags []string) ([]species.Species, error) {
	animals, err := getFilteredAnimals(humidity, temperature, tags)
	if err != nil {
		err = fmt.Errorf("failed to get animals for geographic area: %w", err)
		return []species.Species{}, err
	}
	fishes, err := getFilteredFish(humidity, temperature, tags)
	if err != nil {
		err = fmt.Errorf("failed to get animals for geographic area: %w", err)
		return []species.Species{}, err
	}
	insects, err := getFilteredInsects(humidity, temperature, tags)
	if err != nil {
		err = fmt.Errorf("failed to get animals for geographic area: %w", err)
		return []species.Species{}, err
	}

	numberOfAnimals := int(((float64(prevalence) / 100) * 20) + 4)
	numberOfFish := int(((float64(prevalence) / 100) * 20) + 3)
	numberOfInsects := int(((float64(prevalence) / 100) * 20) + 1)

	animals = species.Random(ctx, numberOfAnimals, animals)
	animals = append(animals, species.Random(ctx, numberOfFish, fishes)...)
	animals = append(animals, species.Random(ctx, numberOfInsects, insects)...)

	return animals, nil
}

func getFilteredAnimals(humidity int, temperature int, tags []string) ([]species.Species, error) {
	animals, err := animal.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered animals for geographic area: %w", err)
		return []species.Species{}, err
	}
	animals = species.FilterHumidity(humidity, animals)
	animals = species.FilterTemperature(temperature, animals)
	animals = species.ByTagIn(tags, animals)

	return animals, nil
}

func getFilteredFish(humidity int, temperature int, tags []string) ([]species.Species, error) {
	fishes, err := fish.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered fish for geographic area: %w", err)
		return []species.Species{}, err
	}
	fishes = species.FilterHumidity(humidity, fishes)
	fishes = species.FilterTemperature(temperature, fishes)
	fishes = species.ByTagIn(tags, fishes)

	return fishes, nil
}

func getFilteredInsects(humidity int, temperature int, tags []string) ([]species.Species, error) {
	insects, err := insect.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered insects for geographic area: %w", err)
		return []species.Species{}, err
	}
	insects = species.FilterHumidity(humidity, insects)
	insects = species.FilterTemperature(temperature, insects)
	insects = species.ByTagIn(tags, insects)

	return insects, nil
}

func ensurePlantsHaveGrains(ctx context.Context, plants []species.Species) ([]species.Species, error) {
	grains := species.ByTag("grain", plants)
	if len(grains) == 0 {
		allPlants, err := plant.All()
		if err != nil {
			err = fmt.Errorf("failed to get all plants for fetching grains: %w", err)
			return []species.Species{}, err
		}
		grain, err := species.RandomWithResourceTag(ctx, "grain", allPlants)
		if err != nil {
			err = fmt.Errorf("failed to get grain plant for fetching grains: %w", err)
			return []species.Species{}, err
		}
		plants = append(plants, grain)
	}

	return plants, nil
}

func getPlants(ctx context.Context, humidity int, temperature int, prevalence int, tags []string) ([]species.Species, error) {
	plants, err := getFilteredPlants(humidity, temperature, tags)
	if err != nil {
		err = fmt.Errorf("failed to get plants for geographic area: %w", err)
		return []species.Species{}, err
	}
	trees, err := getFilteredTrees(humidity, temperature, tags)
	if err != nil {
		err = fmt.Errorf("failed to get plants for geographic area: %w", err)
		return []species.Species{}, err
	}

	numberOfPlants := int(((float64(prevalence) / 100) * 30) + 4)
	numberOfTrees := int(((float64(prevalence) / 100) * 20) + 1)

	plants = species.Random(ctx, numberOfPlants, plants)
	plants, err = ensurePlantsHaveGrains(ctx, plants)
	if err != nil {
		err = fmt.Errorf("failed to get grain plants for area: %w", err)
		return []species.Species{}, err
	}
	plants = append(plants, species.Random(ctx, numberOfTrees, trees)...)

	return plants, nil
}

func getFilteredPlants(humidity int, temperature int, tags []string) ([]species.Species, error) {
	plants, err := plant.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered plants for geographic area: %w", err)
		return []species.Species{}, err
	}
	plants = species.FilterHumidity(humidity, plants)
	plants = species.FilterTemperature(temperature, plants)
	plants = species.ByTagIn(tags, plants)

	return plants, nil
}

func getFilteredTrees(humidity int, temperature int, tags []string) ([]species.Species, error) {
	trees, err := tree.All()
	if err != nil {
		err = fmt.Errorf("failed to get filtered trees for geographic area: %w", err)
		return []species.Species{}, err
	}
	trees = species.FilterHumidity(humidity, trees)
	trees = species.FilterTemperature(temperature, trees)
	trees = species.ByTagIn(tags, trees)

	return trees, nil
}

func getMinerals(ctx context.Context) ([]mineral.Mineral, error) {
	var minerals []mineral.Mineral

	all, err := mineral.All()
	if err != nil {
		err = fmt.Errorf("failed to get minerals: %w", err)
		return []mineral.Mineral{}, err
	}

	metals := mineral.ByTag("metal", all)

	filteredMetals, err := mineral.RandomWeightedSet(ctx, 5, metals)
	if err != nil {
		err = fmt.Errorf("failed to get minerals: %w", err)
		return []mineral.Mineral{}, err
	}

	gems := mineral.ByTag("gem", all)
	filteredGems := mineral.Random(ctx, 3, gems)

	stones := mineral.ByTag("stone", all)
	filteredStones := mineral.Random(ctx, 2, stones)
	salt := mineral.ByTag("salt", all)

	minerals = append(minerals, filteredMetals...)
	minerals = append(minerals, filteredGems...)
	minerals = append(minerals, filteredStones...)
	minerals = append(minerals, salt...)

	return minerals, nil
}

func getSoils(ctx context.Context, oceanDistance int, humidity int, temperature int) ([]soil.Soil, error) {
	var s []soil.Soil
	var soils []soil.Soil
	var filtered []soil.Soil

	tags := []string{
		"dirt",
	}

	all, err := soil.All()
	if err != nil {
		err = fmt.Errorf("failed to get soils for geographic area: %w", err)
		return []soil.Soil{}, err
	}

	if oceanDistance < 10 || humidity < 15 && temperature > 40 {
		tags = append(tags, "sand")
	}

	if humidity > 50 {
		tags = append(tags, "clay")
	}

	if temperature > 30 && humidity > 30 {
		tags = append(tags, "loam")
	}

	if temperature > 40 && humidity < 40 {
		tags = append(tags, "silt")
	}

	for _, t := range tags {
		soils, err = soil.ByTag(t, all)
		if err != nil {
			err = fmt.Errorf("failed to get soils for geographic area: %w", err)
			return []soil.Soil{}, err
		}
		s = soil.Random(ctx, 1, soils)
		filtered = append(filtered, s...)
	}

	return filtered, nil
}
