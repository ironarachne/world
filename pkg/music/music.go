package music

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
)

// Style is a style of music
type Style struct {
	Description    string   `json:"description"`
	InstrumentTags []string `json:"instrument_tags"`
}

var (
	rhythms = map[string]int{
		"a single rhythm":    100,
		"a cross-rhythm":     10,
		"complex polyrhythm": 1,
	}

	beats = map[string]int{
		"very fast": 5,
		"fast":      5,
		"moderate":  10,
		"slow":      5,
		"very slow": 5,
	}

	dynamics = map[string]int{
		"very quiet": 5,
		"quiet":      15,
		"loud":       15,
		"very loud":  5,
	}

	harmonies = map[string]int{
		"simple harmony": 10,
		"two harmonies":  1,
		"no harmony":     5,
	}

	melodies = map[string]int{
		"simple":    10,
		"complex":   2,
		"wandering": 2,
		"chaotic":   1,
	}

	pitches = map[string]int{
		"low":    5,
		"medium": 5,
		"high":   5,
	}

	keys = map[string]int{
		"major": 10,
		"minor": 2,
	}

	timbres = []string{
		"booming",
		"bright",
		"brilliant",
		"dark",
		"emotional",
		"full",
		"mellow",
		"metallic",
		"nasal",
		"reedy",
		"resonant",
		"rough",
		"smooth",
	}
)

// Generate procedurally creates a style of music
func Generate(ctx context.Context) (Style, error) {
	description := "This style of music has "

	rhythm, err := random.StringFromThresholdMap(ctx, rhythms)
	if err != nil {
		err = fmt.Errorf("failed to get a rhythm: %w", err)
		return Style{}, err
	}

	description += rhythm + " with "

	beat, err := random.StringFromThresholdMap(ctx, beats)
	if err != nil {
		err = fmt.Errorf("failed to get a beat: %w", err)
		return Style{}, err
	}

	description += words.Pronoun(beat) + " " + beat + " beat. It is "

	dynamic, err := random.StringFromThresholdMap(ctx, dynamics)
	if err != nil {
		err = fmt.Errorf("failed to get a dynamic: %w", err)
		return Style{}, err
	}

	description += dynamic + ", with "

	harmony, err := random.StringFromThresholdMap(ctx, harmonies)
	if err != nil {
		err = fmt.Errorf("failed to get a harmony: %w", err)
		return Style{}, err
	}

	description += harmony + ". It has "

	melody, err := random.StringFromThresholdMap(ctx, melodies)
	if err != nil {
		err = fmt.Errorf("failed to get a melody: %w", err)
		return Style{}, err
	}

	if rhythm == "a single rhythm" {
		description += words.Pronoun(melody) + " "
	}

	description += melody + " "

	if rhythm == "a single rhythm" {
		description += "melody"
	} else {
		description += "melodies"
	}

	description += " with "

	pitch, err := random.StringFromThresholdMap(ctx, pitches)
	if err != nil {
		err = fmt.Errorf("failed to get a pitch: %w", err)
		return Style{}, err
	}

	description += words.Pronoun(pitch) + " " + pitch + " pitch in a "

	key, err := random.StringFromThresholdMap(ctx, keys)
	if err != nil {
		err = fmt.Errorf("failed to get a key: %w", err)
		return Style{}, err
	}

	description += key + " key. Usually, it has "

	timbre, err := random.String(ctx, timbres)
	if err != nil {
		err = fmt.Errorf("failed to get a timbre: %w", err)
		return Style{}, err
	}

	description += words.Pronoun(timbre) + " " + timbre + " timbre."

	tags := []string{
		timbre,
		pitch,
	}

	style := Style{
		Description:    description,
		InstrumentTags: tags,
	}

	return style, nil
}
