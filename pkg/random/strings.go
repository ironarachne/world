package random

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"strconv"

	"github.com/ironarachne/world/pkg/slices"
)

// SeedString returns a random string that can be used to seed a generator
func SeedString() string {
	b := make([]byte, 13)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

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
		} else {
			i--
		}
	}

	return result, nil
}

// StringFromThresholdMap returns a random weighted string
func StringFromThresholdMap(items map[string]int) (string, error) {
	ceiling := 0

	for _, weight := range items {
		ceiling += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, weight := range items {
		randomValue -= weight
		if randomValue <= 0 {
			return item, nil
		}
	}

	err := fmt.Errorf("Could not find value for " + strconv.Itoa(randomValue))
	return "", err
}

// SeedFromString uses a string to seed the random number generator
func SeedFromString(source string) error {
	h := md5.New()
	_, err := io.WriteString(h, source)
	if err != nil {
		err = fmt.Errorf("Failed to seed random number generator: %w", err)
		return err
	}
	seed := binary.BigEndian.Uint64(h.Sum(nil))
	rand.Seed(int64(seed))
	return nil
}
