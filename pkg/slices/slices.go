/*
Package slices provides convenience methods for dealing with slices
*/
package slices

import (
	"context"

	"github.com/ironarachne/world/pkg/random"
)

// ShuffleStringSlice randomly reorders a slice of strings
func ShuffleStringSlice(ctx context.Context, vals []string) []string {
	ret := make([]string, len(vals))
	perm := random.Perm(ctx, len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

// StringIn checks to see if a string is in an array of strings
func StringIn(item string, collection []string) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// StringSliceWithin checks to see if a slice of strings is wholly within another slice of strings
func StringSliceWithin(slice1 []string, slice2 []string) bool {
	if len(slice1) == 0 {
		return true
	}

	for _, i := range slice1 {
		if !StringIn(i, slice2) {
			return false
		}
	}
	return true
}

// StringSlicePartlyWithin checks to see if at least one element of a slice is in the second slice
func StringSlicePartlyWithin(slice1 []string, slice2 []string) bool {
	if len(slice1) == 0 {
		return true
	}

	for _, i := range slice1 {
		if StringIn(i, slice2) {
			return true
		}
	}
	return false
}
