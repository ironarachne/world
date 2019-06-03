package language

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

// WritingSystem is a system of writing
type WritingSystem struct {
	Name           string
	Classification string
	StrokeType     string
}

func randomStrokeType() string {
	var newStroke string
	strokes := []string{}

	strokeTypes := []string{
		"angular lines",
		"arcs",
		"boxes",
		"circles",
		"dots",
		"flowing lines",
		"half-circles",
		"half-loops",
		"loops",
		"sharp angles",
		"slashes",
		"swirls",
		"triangles",
	}

	for i := 0; i < 3; i++ {
		newStroke = strokeTypes[rand.Intn(len(strokeTypes))]
		if !slices.StringIn(newStroke, strokes) {
			strokes = append(strokes, newStroke)
		}
	}

	return words.CombinePhrases(strokes)
}

func randomWritingSystem() WritingSystem {
	var writingSystem WritingSystem

	classifications := []string{
		"abjad",
		"abugida",
		"alphabet",
		"ideographic",
		"pictographic",
		"semanto-phonetic",
		"syllabary",
	}

	writingSystem.Classification = random.String(classifications)
	writingSystem.StrokeType = randomStrokeType()

	return writingSystem
}
