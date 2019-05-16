package world

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/country"
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/worldmap"
)

// World is a fantasy world
type World struct {
	Name                string
	WorldMap            worldmap.WorldMap
	Countries           []country.Country
	OccupiedCoordinates []grid.Coordinate
}

// Generate procedurally generates a world
func Generate() World {
	var boundary worldmap.Boundary
	var boundaries []worldmap.Boundary
	var homeTile worldmap.Tile
	var homeTileCoordinates grid.Coordinate
	var newCountry country.Country
	var newRegion region.Region
	var newRegions []region.Region
	var numRegions int
	var regionCoordinates []grid.Coordinate
	var contiguousAvailableCoords []grid.Coordinate
	var adjacentTiles []worldmap.Tile

	world := World{}

	world.Name = "WORLD"
	world.WorldMap = worldmap.Generate(50, 80)

	numCountries := rand.Intn(4) + 1
	availableCoords := world.WorldMap.AllLandCoordinates()

	for i := 0; i < numCountries; i++ {
		newCountry = country.Generate()
		numRegions = rand.Intn(12) + 1
		newRegions = []region.Region{}
		regionCoordinates = []grid.Coordinate{}

		homeTileCoordinates = availableCoords[rand.Intn(len(availableCoords)-1)]
		homeTile = world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X]
		homeTile.IsInhabited = true
		world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X] = homeTile
		newRegion = region.Generate(homeTile.TileType)
		newCountry.DominantCulture = newRegion.Culture
		regionCoordinates = append(regionCoordinates, homeTileCoordinates)
		newRegion = newRegion.AssignTiles(regionCoordinates)
		world.OccupiedCoordinates = append(world.OccupiedCoordinates, homeTileCoordinates)
		availableCoords = grid.RemoveCoordinatesFromSlice(world.OccupiedCoordinates, availableCoords)
		newRegions = append(newRegions, newRegion)

		for n := 0; n < numRegions; n++ {
			contiguousAvailableCoords = []grid.Coordinate{}
			adjacentTiles = world.WorldMap.GetAdjacentTiles(homeTile)
			for _, t := range adjacentTiles {
				if !t.IsOcean && !t.IsInhabited {
					contiguousAvailableCoords = append(contiguousAvailableCoords, t.Coordinate)
				}
			}
			if len(contiguousAvailableCoords) > 0 {
				homeTileCoordinates = contiguousAvailableCoords[rand.Intn(len(contiguousAvailableCoords)-1)]
				homeTile = world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X]
				homeTile.IsInhabited = true
				world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X] = homeTile
				newRegion = region.Generate(homeTile.TileType)
				newRegion = newRegion.SetCulture(newCountry.DominantCulture)
				regionCoordinates = append(regionCoordinates, homeTileCoordinates)
				newRegion = newRegion.AssignTiles(regionCoordinates)
				world.OccupiedCoordinates = append(world.OccupiedCoordinates, homeTileCoordinates)
				availableCoords = grid.RemoveCoordinatesFromSlice(world.OccupiedCoordinates, availableCoords)
				newRegions = append(newRegions, newRegion)
			}
		}
		newCountry.Regions = newRegions
		world.Countries = append(world.Countries, newCountry)
	}

	for _, c := range world.Countries {
		boundary = worldmap.CreateBoundary(c.Name, c.GetAllTileCoordinates())
		boundaries = append(boundaries, boundary)
	}
	world.WorldMap.Boundaries = boundaries
	world.WorldMap.SVG = world.WorldMap.RenderAsSVG()

	return world
}
