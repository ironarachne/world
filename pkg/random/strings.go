package random

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math/rand"
)

// String returns a random string from an array of strings
func String(items []string) string {
	if len(items) == 0 {
		panic("Tried to get a random value from an empty slice of strings")
	}
	return items[rand.Intn(len(items))]
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
