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
	var activeTile worldmap.Tile
	var activeTileCoordinates grid.Coordinate
	var homeTile worldmap.Tile
	var homeTileCoordinates grid.Coordinate
	var newCountry country.Country
	var newRegion region.Region
	var newRegions []region.Region
	var numRegions int
	var regionCoordinates []grid.Coordinate
	var contiguousAvailableCoords []grid.Coordinate
	var contiguousAvailableRegionCoords []grid.Coordinate
	var adjacentTiles []worldmap.Tile
	var adjacentRegionTiles []worldmap.Tile
	var pattern string
	var style string

	world := World{}

	world.Name = "WORLD"
	world.WorldMap = worldmap.Generate(50, 80)

	numCountries := rand.Intn(20) + 10
	baseRegions := 40 - numCountries
	availableCoords := world.WorldMap.AllLandCoordinates()

	for i := 0; i < numCountries; i++ {
		newCountry = country.Generate()
		numRegions = rand.Intn(12) + baseRegions
		newRegions = []region.Region{}
		regionCoordinates = []grid.Coordinate{}

		homeTileCoordinates = availableCoords[rand.Intn(len(availableCoords)-1)]
		homeTile = world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X]
		homeTile.IsInhabited = true
		world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X] = homeTile
		newRegion = region.Generate(homeTile.TileType)
		newCountry.DominantCulture = newRegion.Culture
		regionCoordinates = append(regionCoordinates, homeTileCoordinates)
		activeTile = homeTile
		for j := 1; j < newRegion.Class.NumberOfTiles; j++ {
			adjacentTiles = world.WorldMap.GetAdjacentTiles(activeTile)
			contiguousAvailableCoords = []grid.Coordinate{}
			for _, t := range adjacentTiles {
				if !t.IsOcean && !t.IsInhabited {
					contiguousAvailableCoords = append(contiguousAvailableCoords, t.Coordinate)
				}
			}
			if len(contiguousAvailableCoords) > 0 {
				if len(contiguousAvailableCoords) == 1 {
					activeTileCoordinates = contiguousAvailableCoords[0]
				} else {
					activeTileCoordinates = contiguousAvailableCoords[rand.Intn(len(contiguousAvailableCoords)-1)]
				}
				activeTile = world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X]
				activeTile.IsInhabited = true
				world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X] = activeTile
				regionCoordinates = append(regionCoordinates, activeTileCoordinates)
			}
		}
		newRegion = newRegion.AssignTiles(regionCoordinates)
		world.OccupiedCoordinates = append(world.OccupiedCoordinates, regionCoordinates...)
		availableCoords = grid.RemoveCoordinatesFromSlice(world.OccupiedCoordinates, availableCoords)
		newRegions = append(newRegions, newRegion)

		for n := 0; n < numRegions; n++ {
			regionCoordinates = []grid.Coordinate{}
			contiguousAvailableCoords = []grid.Coordinate{}
			adjacentTiles = world.WorldMap.GetAdjacentTiles(homeTile)
			for _, t := range adjacentTiles {
				if !t.IsOcean && !t.IsInhabited {
					contiguousAvailableCoords = append(contiguousAvailableCoords, t.Coordinate)
				}
			}
			if len(contiguousAvailableCoords) > 0 {
				if len(contiguousAvailableCoords) == 1 {
					homeTileCoordinates = contiguousAvailableCoords[0]
				} else {
					homeTileCoordinates = contiguousAvailableCoords[rand.Intn(len(contiguousAvailableCoords)-1)]
				}
				homeTile = world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X]
				homeTile.IsInhabited = true
				world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X] = homeTile
				newRegion = region.Generate(homeTile.TileType)
				newRegion = newRegion.SetCulture(newCountry.DominantCulture)
				regionCoordinates = append(regionCoordinates, homeTileCoordinates)
				activeTile = homeTile
				for j := 1; j < newRegion.Class.NumberOfTiles; j++ {
					adjacentRegionTiles = world.WorldMap.GetAdjacentTiles(activeTile)
					contiguousAvailableRegionCoords = []grid.Coordinate{}
					for _, t := range adjacentRegionTiles {
						if !t.IsOcean && !t.IsInhabited {
							contiguousAvailableRegionCoords = append(contiguousAvailableRegionCoords, t.Coordinate)
						}
					}
					if len(contiguousAvailableRegionCoords) > 0 {
						if len(contiguousAvailableRegionCoords) == 1 {
							activeTileCoordinates = contiguousAvailableRegionCoords[0]
						} else {
							activeTileCoordinates = contiguousAvailableRegionCoords[rand.Intn(len(contiguousAvailableRegionCoords)-1)]
						}
						activeTile = world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X]
						activeTile.IsInhabited = true
						world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X] = activeTile
						regionCoordinates = append(regionCoordinates, activeTileCoordinates)
					}
				}
				newRegion = newRegion.AssignTiles(regionCoordinates)
				world.OccupiedCoordinates = append(world.OccupiedCoordinates, regionCoordinates...)
				availableCoords = grid.RemoveCoordinatesFromSlice(world.OccupiedCoordinates, availableCoords)
				newRegions = append(newRegions, newRegion)
			}
		}
		newCountry.Regions = newRegions
		world.Countries = append(world.Countries, newCountry)
	}

	for _, c := range world.Countries {
		pattern = ""
		style = "fill:" + c.Heraldry.Tinctures[0] + ";fill-opacity:0.5"
		if len(c.Heraldry.Patterns) > 0 {
			pattern = c.Heraldry.Patterns[0]
		}
		boundary = worldmap.CreateBoundary(c.Name, style, pattern, c.GetAllTileCoordinates())
		boundaries = append(boundaries, boundary)
	}
	world.WorldMap.Boundaries = boundaries
	world.WorldMap.SVG = world.WorldMap.RenderAsSVG()

	return world
}
