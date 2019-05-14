package worldmap

// RenderAsText renders a world map to text
func (worldMap WorldMap) RenderAsText() string {
	output := ""

	characters := map[string]string{
		"coniferous forest": "!",
		"deciduous forest":  "*",
		"desert":            "_",
		"grassland":         "-",
		"marshland":         ";",
		"tropical":          "@",
		"mountain":          "^",
		"rainforest":        "?",
		"savanna":           "%",
		"steppe":            "s",
		"taiga":             "t",
		"tundra":            "u",
	}

	for _, r := range worldMap.Tiles {
		for _, t := range r {
			if t.TileType == "ocean" {
				output += "~"
			} else {
				output += characters[t.TileType]
			}
		}
		output += "\n"
	}

	return output
}
