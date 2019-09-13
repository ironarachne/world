package heraldry

import (
	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
	"image"
	"math"
	"math/rand"
)

// Variation is a variation of the field
type Variation struct {
	Name              string
	Blazon            string
	NumberOfTinctures int
	Tinctures         []tincture.Tincture
	Commonality       int
	Render            func(width int, height int, tinctures []tincture.Tincture) image.Image
}

func allVariations() []Variation {
	variations := []Variation{
		{
			Name:              "barry",
			NumberOfTinctures: 2,
			Commonality:       5,
			Render: func(width int, height int, tinctures []tincture.Tincture) image.Image {
				dc := gg.NewContext(width, height)
				barHeight := height/10
				for i:=0;i<10;i++ {
					dc.DrawRectangle(0,float64(i*barHeight), float64(width), float64((i+1)*barHeight))
					if math.Mod(float64(i), 2) == 0 {
						tinctures[0].Fill(dc)
					} else {
						tinctures[1].Fill(dc)
					}
				}
				return dc.Image()
			},
		},
		{
			Name:              "bendy",
			NumberOfTinctures: 2,
			Commonality:       5,
			Render: func(width int, height int, tinctures []tincture.Tincture) image.Image {
				dc := gg.NewContext(width, height)
				barWidth := height/6
				for i:=-10;i<20;i++ {
					dc.MoveTo(float64(i*barWidth), 0)
					dc.LineTo(float64((i+1)*barWidth), 0)
					dc.LineTo(float64((i+6)*barWidth), float64(height))
					dc.LineTo(float64((i+5)*barWidth), float64(height))
					dc.ClosePath()
					if math.Mod(float64(i), 2) == 0 {
						tinctures[0].Fill(dc)
					} else {
						tinctures[1].Fill(dc)
					}
				}
				return dc.Image()
			},
		},
		{
			Name:              "bendy sinister",
			NumberOfTinctures: 2,
			Commonality:       5,
			Render: func(width int, height int, tinctures []tincture.Tincture) image.Image {
				dc := gg.NewContext(width, height)
				barWidth := height/6
				for i:=0;i<20;i++ {
					dc.MoveTo(float64(i*barWidth), 0)
					dc.LineTo(float64((i+1)*barWidth), 0)
					dc.LineTo(float64((i-5)*barWidth), float64(height))
					dc.LineTo(float64((i-6)*barWidth), float64(height))
					dc.ClosePath()
					if math.Mod(float64(i), 2) == 0 {
						tinctures[0].Fill(dc)
					} else {
						tinctures[1].Fill(dc)
					}
				}
				return dc.Image()
			},
		},
		{
			Name:              "chequy",
			NumberOfTinctures: 2,
			Commonality:       10,
			Render: func(width int, height int, tinctures []tincture.Tincture) image.Image {
				dc := gg.NewContext(width, height)
				boxSize := width/10
				drawType := 0
				for y:=0;y<(height/boxSize)+1;y++ {
					if math.Mod(float64(y), 2) == 0 {
						drawType = 1
					}
					for x:=0;x<10;x++ {
						dc.DrawRectangle(float64(x*boxSize),float64(y*boxSize), float64((x+1)*boxSize), float64((y+1)*boxSize))
						if drawType == 0 {
							tinctures[0].Fill(dc)
							drawType = 1
						} else {
							tinctures[1].Fill(dc)
							drawType = 0
						}
					}
					drawType = 0
				}
				return dc.Image()
			},
		},
		{
			Name:              "paly",
			NumberOfTinctures: 2,
			Commonality:       5,
			Render: func(width int, height int, tinctures []tincture.Tincture) image.Image {
				dc := gg.NewContext(width, height)
				barWidth := width/10
				for i:=0;i<10;i+=2 {
					dc.DrawRectangle(float64(i*barWidth),0, float64((i+1)*barWidth), float64(height))
					tinctures[0].Fill(dc)
					dc.DrawRectangle(float64((i+1)*barWidth),0, float64((i+3)*barWidth), float64(height))
					tinctures[1].Fill(dc)
				}
				return dc.Image()
			},
		},
		{
			Name:              "",
			NumberOfTinctures: 1,
			Commonality:       300,
			Render: func(width int, height int, tinctures []tincture.Tincture) image.Image {
				dc := gg.NewContext(width, height)
				dc.DrawRectangle(0,0, float64(width), float64(height))
				tinctures[0].Fill(dc)
				return dc.Image()
			},
		},
	}

	return variations
}

func randomVariation() Variation {
	all := allVariations()
	return all[rand.Intn(len(all))]
}

func randomWeightedVariation() Variation {
	all := allVariations()

	weights := map[string]int{}

	for _, c := range all {
		weights[c.Name] = c.Commonality
	}

	name := random.StringFromThresholdMap(weights)

	for _, c := range all {
		if c.Name == name {
			return c
		}
	}

	return Variation{}
}

func generateVariation(initialTincture tincture.Tincture) Variation {
	var tinc tincture.Tincture
	var tinctureNames []string
	var possible []tincture.Tincture

	variation := randomWeightedVariation()

	lastTincture := initialTincture

	if variation.NumberOfTinctures > 1 {
		if initialTincture.Type == "fur" {
			possible = tincture.Complementary(initialTincture, false)
			initialTincture = tincture.RandomWeighted(possible)
		}
	}

	variation.Tinctures = append(variation.Tinctures, initialTincture)

	for i := 1; i < variation.NumberOfTinctures; i++ {
		possible = tincture.All()
		possible = tincture.Remove(lastTincture, possible)
		tinc = tincture.Random(possible)
		variation.Tinctures = append(variation.Tinctures, tinc)
		lastTincture = tinc
	}

	for _, t := range variation.Tinctures {
		tinctureNames = append(tinctureNames, t.Name)
	}

	variation.Blazon = variation.Name
	if variation.NumberOfTinctures > 1 {
		variation.Blazon += " " + words.CombinePhrases(tinctureNames)
	} else {
		variation.Blazon += variation.Tinctures[0].Name
	}

	return variation
}
