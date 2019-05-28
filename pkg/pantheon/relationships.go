package pantheon

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Relationship is a unidirectional relationship between deities
type Relationship struct {
	Target     string
	Descriptor string
	Type       string
}

// RelationshipType is a type of relationship
type RelationshipType struct {
	Name        string
	Descriptors []string
}

// GenerateRelationships generates relationships between deities
func (pantheon Pantheon) GenerateRelationships() map[string]Deity {
	var descriptor string
	var relationship Relationship
	var inverse Relationship
	var inverseDeity Deity
	var relationshipType RelationshipType
	var target string

	modifiedDeities := make(map[string]Deity)

	relationshipTypes := getAllRelationshipTypes()

	for _, deity := range pantheon.Deities {
		for i := 0; i < 3; i++ {
			target = randomDeityNameFromMap(pantheon.Deities)
			relationshipType = relationshipTypes[rand.Intn(len(relationshipTypes))]
			descriptor = random.String(relationshipType.Descriptors)

			if deity.Name != target {
				if relationshipType.Name == "parent" {
					if deity.Gender.Name == "male" {
						descriptor = "is the father of"
					} else if deity.Gender.Name == "female" {
						descriptor = "is the mother of"
					}
				} else if relationshipType.Name == "child" {
					if deity.Gender.Name == "male" {
						descriptor = "is the son of"
					} else if deity.Gender.Name == "female" {
						descriptor = "is the daughter of"
					}
				}
				relationship = Relationship{Target: target, Descriptor: descriptor, Type: relationshipType.Name}
				if !isRelationshipADuplicate(relationship, deity.Relationships) {
					deity.Relationships = append(deity.Relationships, relationship)
				}

				modifiedDeities[deity.Name] = deity
			}
		}
	}

	for _, deity := range modifiedDeities {
		for _, r := range deity.Relationships {
			inverse = deity.getInverseRelationship(r)
			if !isRelationshipADuplicate(inverse, modifiedDeities[r.Target].Relationships) {
				inverseDeity = modifiedDeities[r.Target]
				inverseDeity.Relationships = append(inverseDeity.Relationships, inverse)
				modifiedDeities[r.Target] = inverseDeity
			}
		}
	}

	return modifiedDeities
}

func getAllRelationshipTypes() []RelationshipType {
	types := []RelationshipType{
		RelationshipType{
			Name: "parent",
			Descriptors: []string{
				"created",
				"gave birth to",
				"is parent of",
			},
		},
		RelationshipType{
			Name: "child",
			Descriptors: []string{
				"was created by",
				"is child of",
			},
		},
		RelationshipType{
			Name: "friend",
			Descriptors: []string{
				"is a confidant of",
				"is a friend of",
				"is an ally of",
			},
		},
		RelationshipType{
			Name: "enemy",
			Descriptors: []string{
				"despises",
				"fears",
				"hates",
				"is the enemy of",
				"is the eternal rival of",
				"is the hated foe of",
			},
		},
		RelationshipType{
			Name: "lover",
			Descriptors: []string{
				"loves",
				"adores",
			},
		},
		RelationshipType{
			Name: "opinion",
			Descriptors: []string{
				"is amused by",
				"is bored by",
				"is chagrined by",
				"is suspicious of",
				"respects",
				"tends to avoid",
				"worries about",
			},
		},
	}

	return types
}

func getRelationshipType(name string) RelationshipType {
	allTypes := getAllRelationshipTypes()

	for _, t := range allTypes {
		if t.Name == name {
			return t
		}
	}

	return RelationshipType{}
}

func (deity Deity) getInverseRelationship(relationship Relationship) Relationship {
	inverseType := relationship.Type

	if relationship.Type == "parent" {
		inverseType = "child"
	} else if relationship.Type == "child" {
		inverseType = "parent"
	}

	inverse := getRelationshipType(inverseType)
	inverseDescriptor := random.String(inverse.Descriptors)

	return Relationship{Target: deity.Name, Descriptor: inverseDescriptor, Type: inverseType}
}

func isRelationshipADuplicate(relationship Relationship, relationships []Relationship) bool {
	for _, r := range relationships {
		if relationship.Type == r.Type && relationship.Target == r.Target {
			return true
		}
	}
	return false
}
