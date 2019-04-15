package climate

import (
	"math/rand"
	"time"
)

func inSlice(i string, check []string) bool {
	for _, j := range check {
		if i == j {
			return true
		}
	}
	return false
}

func shuffle(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func combinePhrases(phrases []string) string {
	combined := ""
	for index, i := range phrases {
		if index == len(phrases)-1 && len(phrases) > 2 {
			combined += "and " + i
		} else if index == len(phrases)-1 && len(phrases) == 2 {
			combined += " and " + i
		} else if index < len(phrases)-1 && len(phrases) > 2 {
			combined += i + ", "
		} else {
			combined += i
		}
	}

	return combined
}
