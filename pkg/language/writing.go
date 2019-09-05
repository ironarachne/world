package language

import (
	"fmt"
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

func randomWritingSystem() (WritingSystem, error) {
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

	classification, err := random.String(classifications)
	if err != nil {
		err = fmt.Errorf("Could not generate writing system: %w", err)
		return WritingSystem{}, err
	}
	writingSystem.Classification = classification
	writingSystem.StrokeType = randomStrokeType()

	return writingSystem, nil
}
