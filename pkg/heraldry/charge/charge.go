package charge

import (
	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/graphics"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"github.com/ironarachne/world/pkg/math"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
	"image"
	"image/color"
	"math/rand"
)

// Charge is a charge
type Charge struct {
	Identifier string
	Name       string
	Noun       string
	NounPlural string
	Descriptor string
	SingleOnly bool
	Tags       []string
	Render func(bodyTincture tincture.Tincture) image.Image
}

// Group is a group of charges with a common tincture
type Group struct {
	Blazon string
	Charges []Charge
	tincture.Tincture
	Position string
}

func all() []Charge {
	charges := []Charge{}

	animals := getAnimalCharges()
	heavens := getHeavensCharges()
	objects := getObjectCharges()
	ordinaries := getOrdinaryCharges()
	people := getPeopleCharges()
	plants := getPlantCharges()
	charges = append(charges, animals...)
	charges = append(charges, heavens...)
	charges = append(charges, objects...)
	charges = append(charges, ordinaries...)
	charges = append(charges, people...)
	charges = append(charges, plants...)

	for _, c := range charges {
		c.Tags = append(c.Tags, c.Identifier)
	}

	return charges
}

func randomNumberOfCharges() int {
	possibleValues := map[int]int{
		1: 10,
		2: 5,
		3: 3,
	}

	value := random.IntFromThresholdMap(possibleValues)

	return value
}

// MatchingTag returns all charges that match a tag
func MatchingTag(tag string) []Charge {
	matching := []Charge{}

	charges := all()

	for _, c := range charges {
		if slices.StringIn(tag, c.Tags) {
			matching = append(matching, c)
		}
	}

	return matching
}

// Random returns a random charge
func Random() Charge {
	charges := all()

	return charges[rand.Intn(len(charges))]
}

// RandomGroup returns a random group of charges
func RandomGroup(fieldTincture tincture.Tincture) Group {
	group := Group{}

	charge := Random()

	numberOfCharges := randomNumberOfCharges()
	if slices.StringIn("full size", charge.Tags) {
		numberOfCharges = 1
	}
	for i:=0;i<numberOfCharges;i++ {
		group.Charges = append(group.Charges, charge)
	}
	group.Tincture = tincture.RandomContrasting(fieldTincture, false)
	group.Position = randomChargePosition()

	return group
}

// SpecificGroup returns a charge group with a given tag and number of elements
func SpecificGroup(fieldTincture tincture.Tincture, tag string, numberOfCharges int) Group {
	group := Group{}

	var charges []Charge
	chargeObject := RandomMatchingTag(tag)

	if slices.StringIn("full size", chargeObject.Tags) {
		numberOfCharges = 1
	}
	for i:=0;i<numberOfCharges;i++ {
		charges = append(charges, chargeObject)
	}
	group.Charges = charges
	group.Tincture = tincture.RandomContrasting(fieldTincture, false)
	group.Position = randomChargePosition()

	return group
}

// RandomMatchingTag returns a random charge that matches a tag
func RandomMatchingTag(tag string) Charge {
	charges := MatchingTag(tag)

	if len(charges) == 0 {
		panic("No matching charges!") // TODO: Replace with error
	}

	if len(charges) == 1 {
		return charges[0]
	}

	return charges[rand.Intn(len(charges))]
}

func randomChargePosition() string {
	positions := map[string]int{
		"in fess": 30,
		"in chief": 2,
		"in base": 1,
	}

	position := random.StringFromThresholdMap(positions)

	return position
}

// RenderChargeFromFile renders a charge from a file using the given tincture and returns it as an Image
func RenderChargeFromFile(name string, t tincture.Tincture) image.Image {
	pattern := graphics.Pattern{
		Color: t.Color,
		PatternFileName: t.PatternFileName,
		Type: "solid",
	}

	linesPattern := graphics.Pattern{
		Type:            "solid",
		Color:           color.RGBA{R: 0, G: 0, B: 0, A: 255},
		PatternFileName: "",
	}

	if t.Name == "sable" {
		linesPattern.Color = color.RGBA{R:255, G:255, B: 255, A: 255}
	}

	if t.Type == "fur" {
		pattern.Type = "image"
	}

	path := "images/charges/" + name + ".png"
	linesPath := "images/charges/" + name + "-lines.png"

	body := graphics.RenderFilledImage(path, pattern)
	lines := graphics.RenderFilledImage(linesPath, linesPattern)

	dc := gg.NewContextForImage(body)
	dc.DrawImage(lines, 0,0)

	img := dc.Image()

	return img
}

// RenderPNG renders a charge group
func (group Group) RenderPNG(width int, height int) image.Image {
	var scaleFactor float64
	var scaleHeight int
	var scaleWidth int
	dc := gg.NewContext(width, height)
	c := group.Charges[0].Render(group.Tincture)
	chargeWidth := c.Bounds().Max.X
	chargeHeight := c.Bounds().Max.Y
	chargeMax := float64(math.Max(chargeHeight, chargeWidth))
	fieldDimension := float64(height)
	if chargeWidth > chargeHeight {
		fieldDimension = float64(width)
	}

	numberOfCharges := len(group.Charges)

	if slices.StringIn("full size", group.Charges[0].Tags) {
		sh := float64(height) / float64(chargeHeight)
		sw := float64(width) / float64(chargeWidth)
		dc.Scale(sw, sh)
		dc.DrawImage(c, 0, 0)
	} else {
		if numberOfCharges == 1 {
			scaleFactor = float64((fieldDimension * 0.7) / chargeMax)
			scaleHeight = int(float64(height) / scaleFactor)
			scaleWidth = int(float64(width) / scaleFactor)
			dc.Scale(scaleFactor, scaleFactor)
			dc.DrawImageAnchored(c, scaleWidth/2, scaleHeight/2, 0.5, 0.5)
		} else if numberOfCharges == 2 {
			scaleFactor = float64((fieldDimension * 0.45) / chargeMax)
			scaleHeight = int(float64(height) / scaleFactor)
			scaleWidth = int(float64(width) / scaleFactor)
			dc.Scale(scaleFactor, scaleFactor)
			alignment := rand.Intn(100)
			if alignment > 80 {
				dc.DrawImageAnchored(c, scaleWidth/4, scaleHeight/2, 0.5, 0.5)
				dc.DrawImageAnchored(c, (scaleWidth*3/4), scaleHeight/2, 0.5, 0.5)
			} else {
				dc.DrawImageAnchored(c, scaleWidth/2, (scaleHeight*9/32), 0.5, 0.5)
				dc.DrawImageAnchored(c, scaleWidth/2, (scaleHeight*7/10), 0.5, 0.5)
			}
		} else if numberOfCharges == 3 {
			scaleFactor = float64((fieldDimension * 0.45) / chargeMax)
			scaleHeight = int(float64(height) / scaleFactor)
			scaleWidth = int(float64(width) / scaleFactor)
			dc.Scale(scaleFactor, scaleFactor)
			dc.DrawImageAnchored(c, scaleWidth/4, (scaleHeight*9/32), 0.5, 0.5)
			dc.DrawImageAnchored(c, (scaleWidth*3/4), (scaleHeight*9/32), 0.5, 0.5)
			dc.DrawImageAnchored(c, scaleWidth/2, (scaleHeight*7/10), 0.5, 0.5)
		}
	}

	return dc.Image()
}

// RenderBlazon returns a string description of the blazon of a charge group
func (group Group) RenderBlazon() string {
	blazon := ""

	count := len(group.Charges)

	if count == 0 {
		blazon = ""
	} else if count == 1 {
		blazon += words.Pronoun(group.Charges[0].Noun) + " " + group.Charges[0].Name + " " + group.Tincture.Name
	} else if count == 2 {
		blazon += "two " + group.Charges[0].NounPlural + " " + group.Charges[0].Descriptor + " " + group.Tincture.Name
	} else if count == 3 {
		blazon += "three " + group.Charges[0].NounPlural + " " + group.Charges[0].Descriptor + " " + group.Tincture.Name
	}

	return blazon
}