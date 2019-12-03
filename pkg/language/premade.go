package language

import (
	"github.com/ironarachne/namegen"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/writing"
)

// PremadeCommon returns the Common language
func PremadeCommon() (Language, error) {
	var name string
	var maleNames []string
	var femaleNames []string
	var familyNames []string

	wordList := map[string]string{}

	writingSystem := writing.System{
		Name:           "Common",
		Classification: "alphabet",
		StrokeType:     "lines and arcs",
	}

	femaleNameGenerator := namegen.NameGeneratorFromType("fantasy", "female")
	maleNameGenerator := namegen.NameGeneratorFromType("fantasy", "male")

	for i := 0; i < 100; i++ {
		name = femaleNameGenerator.FirstName("female")
		if !slices.StringIn(name, femaleNames) {
			femaleNames = append(femaleNames, name)
		} else {
			i--
		}
	}

	for i := 0; i < 100; i++ {
		name = maleNameGenerator.FirstName("male")
		if !slices.StringIn(name, maleNames) {
			maleNames = append(maleNames, name)
		} else {
			i--
		}
	}

	for i := 0; i < 100; i++ {
		name = femaleNameGenerator.LastName()
		if !slices.StringIn(name, familyNames) {
			familyNames = append(familyNames, name)
		} else {
			i--
		}
	}

	townNames := []string{
		"Aerilon",
		"Aquarin",
		"Aramoor",
		"Azmar",
		"Begger's Hole",
		"Black Hollow",
		"Blue Field",
		"Briar Glen",
		"Brickelwhyte",
		"Broken Shield",
		"Boatwright",
		"Bullmar",
		"Carran",
		"City of Fire",
		"Coalfell",
		"Cullfield",
		"Darkwell",
		"Deathfall",
		"Doonatel",
		"Dry Gulch",
		"Easthaven",
		"Ecrin",
		"Erast",
		"Far Water",
		"Firebend",
		"Fool's March",
		"Frostford",
		"Goldcrest",
		"Goldenleaf",
		"Greenflower",
		"Garen's Well",
		"Haran",
		"Hillfar",
		"Hogsfeet",
		"Hollyhead",
		"Hull",
		"Hwen",
		"Icemeet",
		"Ironforge",
		"Irragin",
		"Jarren's Outpost",
		"Jongvale",
		"Kara's Vale",
		"Knife's Edge",
		"Lakeshore",
		"Leeside",
		"Lullin",
		"Marren's Eve",
		"Millstone",
		"Moonbright",
		"Mountmend",
		"Nearon",
		"New Cresthill",
		"Northpass",
		"Nuxvar",
		"Oakheart",
		"Oar's Rest",
		"Old Ashton",
		"Orrinshire",
		"Ozryn",
		"Pavv",
		"Pella's Wish",
		"Pinnella Pass",
		"Pran",
		"Quan Ma",
		"Queenstown",
		"Ramshorn",
		"Red Hawk",
		"Rivermouth",
		"Saker Keep",
		"Seameet",
		"Ship's Haven",
		"Silverkeep",
		"South Warren",
		"Snake's Canyon",
		"Snowmelt",
		"Squall's End",
		"Swordbreak",
		"Tarrin",
		"Three Streams",
		"Trudid",
		"Ubbin Falls",
		"Ula'ree",
		"Veritas",
		"Violl's Garden",
		"Wavemeet",
		"Whiteridge",
		"Willowdale",
		"Windrip",
		"Wintervale",
		"Wellspring",
		"Westwend",
		"Wolfden",
		"Xan's Bequest",
		"Xynnar",
		"Yarrin",
		"Yellowseed",
		"Zanfeld",
		"Zeffari",
		"Zumka",
	}

	newWordPrefixes := []string{
		"bu",
		"ha",
		"hi",
		"ka",
		"re",
		"sa",
		"se",
		"tra",
		"tre",
	}

	newWordSuffixes := []string{
		"auch",
		"lly",
		"ouch",
		"sh",
		"shy",
		"usch",
	}

	lang := Language{
		Name:             "Common",
		Adjective:        "Common",
		Descriptors:      []string{"rough"},
		Description:      "a language spoken widely and used for trade",
		WritingSystem:    writingSystem,
		WordList:         wordList,
		FemaleFirstNames: femaleNames,
		MaleFirstNames:   maleNames,
		FamilyNames:      familyNames,
		TownNames:        townNames,
		IsTonal:          false,
		NewWordPrefixes:  newWordPrefixes,
		NewWordSuffixes:  newWordSuffixes,
	}

	return lang, nil
}
