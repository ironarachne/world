package music

import (
	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedInstrument is a simplified version of an instrument for display
type SimplifiedInstrument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SimplifiedStyle is a simplified versino of a music style for display
type SimplifiedStyle struct {
	Description string                 `json:"description"`
	Instruments []SimplifiedInstrument `json:"instruments"`
}

// Simplify returns a simplified instrument
func (instrument Instrument) Simplify() SimplifiedInstrument {
	return SimplifiedInstrument{
		Name:        instrument.Name,
		Description: instrument.Description,
	}
}

// Simplify returns a simplified music style
func (style Style) Simplify() SimplifiedStyle {
	instruments := []SimplifiedInstrument{}

	for _, i := range style.Instruments {
		instruments = append(instruments, i.Simplify())
	}

	return SimplifiedStyle{
		Description: style.Describe(),
		Instruments: instruments,
	}
}

// Describe returns the description of a style
func (style Style) Describe() string {
	description := "This music is "
	description += words.CombinePhrases(style.Descriptors)
	description += "."

	return description
}
