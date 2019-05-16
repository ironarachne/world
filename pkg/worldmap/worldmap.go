package worldmap

// WorldMap is a world map. It contains tiles.
type WorldMap struct {
	Height     int
	Width      int
	Tiles      [][]Tile
	Boundaries []Boundary
	SVG        string
}

// Generate creates a world map
func Generate(height int, width int) WorldMap {
	worldMap := WorldMap{}

	worldMap.Height = height
	worldMap.Width = width

	tiles := worldMap.initializeTiles()
	worldMap.Tiles = tiles
	worldMap.Tiles = worldMap.generateLandSIR()
	worldMap.Tiles = worldMap.removeArtifactOceanTiles()
	worldMap.Tiles = worldMap.setTileTemperatures()
	worldMap.Tiles = worldMap.setTileHumidities()
	worldMap.Tiles = worldMap.setTileTypes()

	return worldMap
}
