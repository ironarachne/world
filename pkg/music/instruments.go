package music

import (
	"bytes"
	"math/rand"
	"strings"
	"text/template"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Instrument is a musical instrument
type Instrument struct {
	Name                   string
	Description            string
	Type                   string
	BaseMaterialOptions    []string
	SupportMaterialOptions []string
	BaseMaterial           string
	SupportMaterial        string
	DescriptionTemplate    string
}

// GenerateInstruments generates a set of musical instruments for a climate
func GenerateInstruments(originClimate climate.Climate) []Instrument {
	var instrument Instrument
	var materialType string
	var availableBaseMaterials []string
	var availableSupportMaterials []string
	var woodName string

	availableHides := []string{}
	availableMetals := []string{}
	availableWoods := []string{}
	availableMaterials := []string{}

	allInstruments := getAllInstruments()
	availableInstruments := []Instrument{}
	instruments := []Instrument{}

	for _, i := range originClimate.CommonMetals {
		availableMetals = append(availableMetals, i.Name)
	}

	for _, i := range originClimate.PreciousMetals {
		availableMetals = append(availableMetals, i.Name)
	}

	for _, i := range originClimate.Trees {
		woodName = i.Name
		if !strings.HasSuffix(i.Name, "wood") {
			woodName += "-wood"
		}
		availableWoods = append(availableWoods, woodName)
	}

	for _, i := range originClimate.Animals {
		if i.GivesHide {
			availableHides = append(availableHides, i.Name)
		}
	}

	if len(availableHides) > 0 {
		availableMaterials = append(availableMaterials, "hide")
	}
	if len(availableMetals) > 0 {
		availableMaterials = append(availableMaterials, "metal")
	}
	if len(availableWoods) > 0 {
		availableMaterials = append(availableMaterials, "wood")
	}

	for _, i := range allInstruments {
		if slices.StringSlicePartlyWithin(i.BaseMaterialOptions, availableMaterials) {
			if slices.StringSlicePartlyWithin(i.SupportMaterialOptions, availableMaterials) {
				availableInstruments = append(availableInstruments, i)
			}
		}
	}

	numberOfInstruments := rand.Intn(3) + 1

	for i := 0; i < numberOfInstruments; i++ {
		instrument = availableInstruments[rand.Intn(len(availableInstruments)-1)]
		availableBaseMaterials = []string{}
		availableSupportMaterials = []string{}

		for _, m := range instrument.BaseMaterialOptions {
			if slices.StringIn(m, availableMaterials) {
				availableBaseMaterials = append(availableBaseMaterials, m)
			}
		}

		for _, m := range instrument.SupportMaterialOptions {
			if slices.StringIn(m, availableMaterials) {
				availableSupportMaterials = append(availableSupportMaterials, m)
			}
		}

		materialType = random.String(availableBaseMaterials)
		if materialType == "hide" {
			instrument.BaseMaterial = random.String(availableHides)
		} else if materialType == "metal" {
			instrument.BaseMaterial = random.String(availableMetals)
		} else if materialType == "wood" {
			instrument.BaseMaterial = random.String(availableWoods)
		}

		materialType = random.String(availableSupportMaterials)
		if materialType == "hide" {
			instrument.SupportMaterial = random.String(availableHides)
		} else if materialType == "metal" {
			instrument.SupportMaterial = random.String(availableMetals)
		} else if materialType == "wood" {
			instrument.SupportMaterial = random.String(availableWoods)
		}

		instrument.Description = instrument.getDescription()

		instruments = append(instruments, instrument)
	}

	return instruments
}

func (instrument Instrument) getDescription() string {
	t := template.New("instrument description")

	var err error
	t, err = t.Parse(instrument.DescriptionTemplate)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, instrument); err != nil {
		panic(err)
	}

	result := tpl.String()

	return result
}

func getAllInstruments() []Instrument {
	instruments := []Instrument{
		Instrument{
			Name:                   "short flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		Instrument{
			Name:                   "long flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		Instrument{
			Name:                   "twin flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		Instrument{
			Name:                   "short harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "long harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "full harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "lyre",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "lijerica",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "long-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "pierced lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "short-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "single-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drone",
		},
		Instrument{
			Name:                   "multiple-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drones",
		},
		Instrument{
			Name:                   "hand drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		Instrument{
			Name:                   "short drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		Instrument{
			Name:                   "walking drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		Instrument{
			Name:                   "heavy drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
	}

	return instruments
}
