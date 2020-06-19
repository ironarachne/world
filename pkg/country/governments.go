package country

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/heraldry"
)

// Government is a political structure
type Government struct {
	Type    string
	Leaders []character.Character
}

func (country Country) getNewMonarchy(ctx context.Context) (Government, error) {
	government := Government{}
	government.Type = "monarchy"

	monarch, err := character.Generate(ctx, country.DominantCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate monarch: %w", err)
		return Government{}, err
	}
	monarch.ChangeAge(ctx, random.Intn(ctx, 30)+20)

	if monarch.Gender.Name == "male" {
		monarch.Title = "King"
		monarch.Profession, _ = profession.ByName("noble")
	} else {
		monarch.Title = "Queen"
		monarch.Profession, _ = profession.ByName("noble")
	}

	device, err := heraldry.Generate(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate monarch: %w", err)
		return Government{}, err
	}
	monarch.Heraldry = device.Simplify()

	government.Leaders = append(government.Leaders, monarch)

	return government, nil
}
