package biome

import "testing"

func TestScore(t *testing.T) {
	b := Biome{
		Type:             "terrestrial",
		PrecipitationMax: 70,
		PrecipitationMin: 30,
		TemperatureMin:   40,
		TemperatureMax:   70,
	}

	score := b.Score("terrestrial", 50, 50)

	if score == 0 {
		t.Error("failed to identify biome is in acceptable ranges")
	}

	if score != 95 {
		t.Error("failed to correctly score biome: got ", score, " instead of expected 95")
	}

	score = b.Score("marine", 50, 50)

	if score != 0 {
		t.Error("failed to zero score incorrect biome type")
	}

	score = b.Score("terrestrial", 10, 50)

	if score != 0 {
		t.Error("failed to zero score precipitation below acceptable range")
	}

	score = b.Score("terrestrial", 90, 50)

	if score != 0 {
		t.Error("failed to zero score precipitation above acceptable range")
	}

	score = b.Score("terrestrial", 50, 10)

	if score != 0 {
		t.Error("failed to zero score temperature below acceptable range")
	}

	score = b.Score("terrestrial", 50, 80)

	if score != 0 {
		t.Error("failed to zero score temperature above acceptable range")
	}
}
