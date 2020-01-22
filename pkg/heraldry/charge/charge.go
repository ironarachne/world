package charge

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/graphics"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"github.com/ironarachne/world/pkg/math"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

// Charge is an interface for charges
type Charge interface {
	Blazon(count int, tincture string) (string, error)
	GetTags() []string
	Render(bodyTincture tincture.Tincture) image.Image
}

// Raster is a charge that is rendered from a raster file
type Raster struct {
	Identifier string
	Name       string
	Noun       string
	NounPlural string
	Descriptor string
	SingleOnly bool
	Tags       []string
}

// Group is a group of charges with a common tincture
type Group struct {
	Blazon  string
	Charges []Charge
	tincture.Tincture
	Position string
}

// Blazon returns the blazon for a raster charge
func (r Raster) Blazon(count int, tincture string) (string, error) {
	blazon, err := blazonForCharge(count, r.Noun, r.NounPlural, r.Descriptor, tincture)
	if err != nil {
		err = fmt.Errorf("failed to generate blazon for raster charge: %w", err)
		return "", err
	}

	return blazon, nil
}

func (r Raster) GetTags() []string {
	return r.Tags
}

// Render renders a raster charge from a file
func (rc Raster) Render(bodyTincture tincture.Tincture) image.Image {
	img := RenderChargeFromFile(rc.Identifier, bodyTincture)

	return img
}

func all() []Charge {
	animals := getAnimalCharges()
	heavens := getHeavensCharges()
	objects := getObjectCharges()
	ordinaries := getOrdinaryCharges()
	people := getPeopleCharges()
	plants := getPlantCharges()

	length := len(animals) + len(heavens) + len(objects) + len(ordinaries) + len(people) + len(plants)

	charges := make([]Charge, length)
	l := 0
	for _, v := range animals {
		charges[l] = Charge(v)
		l++
	}
	for _, v := range heavens {
		charges[l] = Charge(v)
		l++
	}
	for _, v := range objects {
		charges[l] = Charge(v)
		l++
	}
	for _, v := range ordinaries {
		charges[l] = Charge(v)
		l++
	}
	for _, v := range people {
		charges[l] = Charge(v)
		l++
	}
	for _, v := range plants {
		charges[l] = Charge(v)
		l++
	}

	return charges
}

func blazonForCharge(count int, singular string, plural string, descriptor string, tincture string) (string, error) {
	blazon := ""

	if count == 1 {
		pronoun := words.Pronoun(singular)
		blazon += pronoun + " " + singular + " " + tincture
	} else if count > 1 {
		number, err := words.NumberToWord(count)
		if err != nil {
			err = fmt.Errorf("failed to generate blazon: %w", err)
			return "", err
		}
		if descriptor != "" {
			blazon += number + " " + plural + " " + descriptor + " " + tincture
		} else {
			blazon += number + " " + plural + " " + tincture
		}
	}

	return blazon, nil
}

func randomNumberOfCharges() (int, error) {
	possibleValues := map[int]int{
		1: 10,
		2: 5,
		3: 3,
	}

	value, err := random.IntFromThresholdMap(possibleValues)
	if err != nil {
		err = fmt.Errorf("failed to get random number of charges: %w", err)
		return 0, err
	}

	return value, nil
}

// MatchingTag returns all charges that match a tag
func MatchingTag(tag string) []Charge {
	matching := []Charge{}

	charges := all()

	for _, c := range charges {
		if slices.StringIn(tag, c.GetTags()) {
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
func RandomGroup(fieldTincture tincture.Tincture) (Group, error) {
	group := Group{}

	charge := Random()

	numberOfCharges, err := randomNumberOfCharges()
	if err != nil {
		err = fmt.Errorf("failed to get random charge group: %w", err)
		return Group{}, err
	}
	if slices.StringIn("full size", charge.GetTags()) {
		numberOfCharges = 1
	}
	for i := 0; i < numberOfCharges; i++ {
		group.Charges = append(group.Charges, charge)
	}
	t, err := tincture.RandomContrasting(fieldTincture, false)
	if err != nil {
		err = fmt.Errorf("failed to get random charge group: %w", err)
		return Group{}, err
	}
	group.Tincture = t
	p, err := randomChargePosition()
	if err != nil {
		err = fmt.Errorf("failed to get random charge group: %w", err)
		return Group{}, err
	}
	group.Position = p

	return group, nil
}

// SpecificGroup returns a charge group with a given tag and number of elements
func SpecificGroup(fieldTincture tincture.Tincture, tag string, numberOfCharges int) (Group, error) {
	group := Group{}

	var charges []Charge
	chargeObject, err := RandomMatchingTag(tag)
	if err != nil {
		err = fmt.Errorf("failed to get specific charge group: %w", err)
		return Group{}, err
	}

	if slices.StringIn("full size", chargeObject.GetTags()) {
		numberOfCharges = 1
	}
	for i := 0; i < numberOfCharges; i++ {
		charges = append(charges, chargeObject)
	}
	group.Charges = charges
	t, err := tincture.RandomContrasting(fieldTincture, false)
	if err != nil {
		err = fmt.Errorf("failed to get specific charge group: %w", err)
		return Group{}, err
	}
	group.Tincture = t
	p, err := randomChargePosition()
	if err != nil {
		err = fmt.Errorf("failed to get specific charge group: %w", err)
		return Group{}, err
	}
	group.Position = p

	return group, nil
}

// RandomMatchingTag returns a random charge that matches a tag
func RandomMatchingTag(tag string) (Charge, error) {
	charges := MatchingTag(tag)

	if len(charges) == 0 {
		err := fmt.Errorf("failed to find charge matching tag " + tag)
		return nil, err
	}

	if len(charges) == 1 {
		return charges[0], nil
	}

	charge := charges[rand.Intn(len(charges))]
	return charge, nil
}

func randomChargePosition() (string, error) {
	positions := map[string]int{
		"in fess":  30,
		"in chief": 2,
		"in base":  1,
	}

	position, err := random.StringFromThresholdMap(positions)
	if err != nil {
		err = fmt.Errorf("failed to get random charge position: %w", err)
		return "", err
	}

	return position, nil
}

// RenderChargeFromFile renders a charge from a file using the given tincture and returns it as an Image
func RenderChargeFromFile(name string, t tincture.Tincture) image.Image {
	pattern := graphics.Pattern{
		Color:           t.Color,
		PatternFileName: t.PatternFileName,
		Type:            "solid",
	}

	linesPattern := graphics.Pattern{
		Type:            "solid",
		Color:           color.RGBA{R: 0, G: 0, B: 0, A: 255},
		PatternFileName: "",
	}

	if t.Name == "sable" {
		linesPattern.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	}

	if t.Type == "fur" {
		pattern.Type = "image"
	}

	path := "images/charges/" + name + ".png"
	linesPath := "images/charges/" + name + "-lines.png"

	body := graphics.RenderFilledImage(path, pattern)
	lines := graphics.RenderFilledImage(linesPath, linesPattern)

	dc := gg.NewContextForImage(body)
	dc.DrawImage(lines, 0, 0)

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

	if slices.StringIn("full size", group.Charges[0].GetTags()) {
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
				dc.DrawImageAnchored(c, (scaleWidth * 3 / 4), scaleHeight/2, 0.5, 0.5)
			} else {
				dc.DrawImageAnchored(c, scaleWidth/2, (scaleHeight * 9 / 32), 0.5, 0.5)
				dc.DrawImageAnchored(c, scaleWidth/2, (scaleHeight * 7 / 10), 0.5, 0.5)
			}
		} else if numberOfCharges == 3 {
			scaleFactor = float64((fieldDimension * 0.45) / chargeMax)
			scaleHeight = int(float64(height) / scaleFactor)
			scaleWidth = int(float64(width) / scaleFactor)
			dc.Scale(scaleFactor, scaleFactor)
			dc.DrawImageAnchored(c, scaleWidth/4, (scaleHeight * 9 / 32), 0.5, 0.5)
			dc.DrawImageAnchored(c, (scaleWidth * 3 / 4), (scaleHeight * 9 / 32), 0.5, 0.5)
			dc.DrawImageAnchored(c, scaleWidth/2, (scaleHeight * 7 / 10), 0.5, 0.5)
		}
	}

	return dc.Image()
}

// RenderBlazon returns a string description of the blazon of a charge group
func (group Group) RenderBlazon() (string, error) {
	count := len(group.Charges)

	blazon, err := group.Charges[0].Blazon(count, group.Tincture.Name)
	if err != nil {
		err = fmt.Errorf("failed to render blazon for charge group: %w", err)
		return "", err
	}

	return blazon, nil
}
