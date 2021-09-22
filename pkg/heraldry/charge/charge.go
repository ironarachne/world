package charge

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"sort"

	"github.com/fogleman/gg"

	"github.com/ironarachne/world/pkg/graphics"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"github.com/ironarachne/world/pkg/math"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

const chargeGroupError = "failed to get random charge group: %w"

// Charge is an interface for charges
type Charge interface {
	Blazon(count int, tincture string) (string, error)
	GetTags() []string
	Render(bodyTincture tincture.Tincture) image.Image
}

// Group is a group of charges with a common tincture
type Group struct {
	Blazon  string
	Charges []Charge
	tincture.Tincture
	Position string
}

// All returns all charges
func All() ([]Charge, error) {
	raster, err := AllRaster()
	if err != nil {
		err = fmt.Errorf("failed to load charges: %w", err)
		return nil, err
	}
	ordinaries := getOrdinaryCharges()

	length := len(raster) + len(ordinaries)

	charges := make([]Charge, length)
	l := 0
	for _, v := range raster {
		charges[l] = Charge(v)
		l++
	}
	for _, v := range ordinaries {
		charges[l] = Charge(v)
		l++
	}

	return charges, nil
}

// AllTags returns all unique charge tags
func AllTags() ([]string, error) {
	tags := []string{}
	ts := []string{}

	charges, err := All()
	if err != nil {
		err = fmt.Errorf("failed to get charge tags: %w", err)
		return []string{}, err
	}

	for _, c := range charges {
		ts = c.GetTags()
		for _, t := range ts {
			if !slices.StringIn(t, tags) {
				tags = append(tags, t)
			}
		}
	}

	sort.Strings(tags)

	return tags, nil
}

func blazonForCharge(count int, singular string, plural string, descriptor string, tincture string) (string, error) {
	blazon := ""

	if count == 1 {
		pronoun := words.Pronoun(singular)
		if descriptor != "" {
			blazon += pronoun + " " + singular + " " + descriptor + " " + tincture
		} else {
			blazon += pronoun + " " + singular + " " + tincture
		}

	} else if count > 1 {
		number := words.NumberToWord(count)
		if descriptor != "" {
			blazon += number + " " + plural + " " + descriptor + " " + tincture
		} else {
			blazon += number + " " + plural + " " + tincture
		}
	}

	return blazon, nil
}

func randomNumberOfCharges(ctx context.Context) (int, error) {
	possibleValues := map[int]int{
		1: 10,
		2: 5,
		3: 3,
	}

	value, err := random.IntFromThresholdMap(ctx, possibleValues)
	if err != nil {
		err = fmt.Errorf("failed to get random number of charges: %w", err)
		return 0, err
	}

	return value, nil
}

// MatchingTag returns all charges that match a tag
func MatchingTag(tag string) ([]Charge, error) {
	matching := []Charge{}

	charges, err := All()
	if err != nil {
		err = fmt.Errorf("failed to find charges matching tag: %w", err)
		return nil, err
	}

	for _, c := range charges {
		if slices.StringIn(tag, c.GetTags()) {
			matching = append(matching, c)
		}
	}

	return matching, nil
}

// Random returns a random charge
func Random(ctx context.Context) (Charge, error) {
	charges, err := All()
	if err != nil {
		err = fmt.Errorf("failed to get random charge: %w", err)
		return nil, err
	}

	charge := charges[random.Intn(ctx, len(charges))]

	return charge, nil
}

// RandomGroup returns a random group of charges
func RandomGroup(ctx context.Context, fieldTincture tincture.Tincture) (Group, error) {
	group, err := RandomGroupByParameters(ctx, fieldTincture, "", 0)
	if err != nil {
		err = fmt.Errorf("failed to get completely random charge group: %w", err)
		return Group{}, err
	}

	return group, nil
}

// RandomGroupByParameters returns a charge group matching the given parameters. Blank or zero values will return random results for the given parameter.
func RandomGroupByParameters(ctx context.Context, fieldTincture tincture.Tincture, tag string, numberOfCharges int) (Group, error) {
	var chargeObject Charge
	var err error
	group := Group{}

	var charges []Charge
	if tag == "" {
		chargeObject, err = Random(ctx)
	} else {
		chargeObject, err = RandomMatchingTag(ctx, tag)
	}
	if err != nil {
		err = fmt.Errorf(chargeGroupError, err)
		return Group{}, err
	}

	if numberOfCharges == 0 {
		numberOfCharges, err = randomNumberOfCharges(ctx)
	}
	if err != nil {
		err = fmt.Errorf(chargeGroupError, err)
		return Group{}, err
	}

	if slices.StringIn("full size", chargeObject.GetTags()) {
		numberOfCharges = 1
	}

	for i := 0; i < numberOfCharges; i++ {
		charges = append(charges, chargeObject)
	}

	group.Charges = charges
	t, err := tincture.RandomContrasting(ctx, fieldTincture, false)
	if err != nil {
		err = fmt.Errorf(chargeGroupError, err)
		return Group{}, err
	}

	group.Tincture = t
	p, err := randomChargePosition(ctx)
	if err != nil {
		err = fmt.Errorf(chargeGroupError, err)
		return Group{}, err
	}
	group.Position = p

	return group, nil
}

// RandomMatchingTag returns a random charge that matches a tag
func RandomMatchingTag(ctx context.Context, tag string) (Charge, error) {
	charges, err := MatchingTag(tag)
	if err != nil {
		err = fmt.Errorf("failed to find random charge matching tag: %w", err)
		return nil, err
	}

	if len(charges) == 0 {
		err := fmt.Errorf("failed to find random charge matching tag " + tag)
		return nil, err
	}

	if len(charges) == 1 {
		return charges[0], nil
	}

	charge := charges[random.Intn(ctx, len(charges))]
	return charge, nil
}

func randomChargePosition(ctx context.Context) (string, error) {
	positions := map[string]int{
		"in fess":  30,
		"in chief": 2,
		"in base":  1,
	}

	position, err := random.StringFromThresholdMap(ctx, positions)
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
func (group Group) RenderPNG(ctx context.Context, width int, height int) image.Image {
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
			alignment := random.Intn(ctx, 100)
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
