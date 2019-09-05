package random

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"

	"github.com/ironarachne/world/pkg/slices"
)

// String returns a random string from a slice of strings
func String(items []string) (string, error) {
	if len(items) == 0 {
		err := fmt.Errorf("Tried to get a random value from an empty slice of strings")
		return "", err
	}
	if len(items) == 1 {
		return items[0], nil
	}
	item := items[rand.Intn(len(items))]

	return item, nil
}

// StringSubset returns a random slice of strings from a slice of strings
func StringSubset(items []string, maxItems int) ([]string, error) {
	result := []string{}
	newItem := ""

	if len(items) == 0 {
		err := fmt.Errorf("Tried to get a random value from an empty slice of strings")
		return []string{}, err
	}

	if maxItems > len(items) {
		maxItems = len(items)
	}

	for i := 0; i < maxItems; i++ {
		newItem = items[rand.Intn(len(items))]
		if !slices.StringIn(newItem, result) {
			result = append(result, newItem)
		}
	}

	return result, nil
}

// StringFromThresholdMap returns a random weighted string
func StringFromThresholdMap(items map[string]int) string {
	result := ""
	ceiling := 0
	start := 0
	var thresholds = make(map[string]int)

	for item, weight := range items {
		ceiling += weight
		thresholds[item] = start
		start += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, threshold := range thresholds {
		if threshold <= randomValue {
			result = item
		}
	}

	return result
}

// SeedFromString uses a string to seed the random number generator
func SeedFromString(source string) {
	h := md5.New()
	io.WriteString(h, source)
	seed := binary.BigEndian.Uint64(h.Sum(nil))
	rand.Seed(int64(seed))
}
