package town

import (
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/resource"
)

func getProducers(size int, resources []resource.Resource) []profession.Profession {
	var producer []profession.Profession
	var producers []profession.Profession

	possibleProducers := resource.GetPossibleProfessions(resources)

	numberOfProducers := 0

	if size < 20 {
		numberOfProducers = 10
	} else if size < 50 {
		numberOfProducers = 15
	} else if size < 100 {
		numberOfProducers = 20
	} else if size < 500 {
		numberOfProducers = 25
	} else if size < 1000 {
		numberOfProducers = 40
	} else if size < 5000 {
		numberOfProducers = 45
	} else if size < 10000 {
		numberOfProducers = 50
	} else {
		numberOfProducers = 100
	}

	for i := 0; i < numberOfProducers; i++ {
		producer = profession.RandomSet(1, possibleProducers)
		producers = append(producers, producer[0])
	}

	return producers
}
