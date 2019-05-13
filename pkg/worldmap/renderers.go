package worldmap

// RenderAsText renders a world map to text
func (worldMap WorldMap) RenderAsText() string {
	output := ""

	for _, r := range worldMap.Tiles {
		for _, t := range r {
			if t.IsOcean {
				output += " "
			} else {
				output += "█"
			}
		}
		output += "\n"
	}

	return output
}
