package charge

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/ironarachne/world/pkg/heraldry/tincture"
)

// Data is a struct containing a slice of Charges
type Data struct {
	Charges []Raster `json:"charges"`
}

// Raster is a charge that is rendered from a raster file
type Raster struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	Noun       string   `json:"noun"`
	NounPlural string   `json:"noun_plural"`
	Descriptor string   `json:"descriptor"`
	SingleOnly bool     `json:"single_only"`
	Tags       []string `json:"tags"`
}

// AllRaster returns all raster charges
func AllRaster() ([]Raster, error) {
	raster, err := Load()
	if err != nil {
		err = fmt.Errorf("failed to load raster charges: %w", err)
		return []Raster{}, err
	}

	return raster, nil
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

// Load returns all predefined raster charges from a JSON file on disk
func Load() ([]Raster, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/charges.json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Raster{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &d)

	all := d.Charges

	if len(all) == 0 {
		err = fmt.Errorf("no charges returned from database: charges.json")
		return []Raster{}, err
	}

	return all, nil
}
