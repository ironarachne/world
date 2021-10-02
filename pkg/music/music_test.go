package music

import (
	"context"
	"testing"

	"github.com/ironarachne/world/pkg/random"
)

func TestGenerate(t *testing.T) {
	seed := "testing"

	style, err := Generate(random.WithSeed(context.Background(), seed))

	if err != nil {
		t.Errorf("Failed to generate music style: %e", err)
	}

	expected := "This style of music has a cross-rhythm with a moderate beat. It is very quiet, with two harmonies. It has simple melodies with a low pitch in a major key. Usually, it has a booming timbre."

	if style.Description != expected {
		t.Errorf("failed to get expected output: got '%s', expected '%s'", style.Description, expected)
	}
}
