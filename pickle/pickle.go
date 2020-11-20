// Package pickle implements functions which include random picking and generating
package pickle

import (
	"time"
	"math/rand"
)

// StringSliceChoice picks an object or item from a string slice and returns a random value
func StringSliceChoice(slice []string) string{
	rand.Seed(time.Now().Unix())

	// select a random index
	index := rand.Intn(len(slice))

	// return the slice index value
	return slice[index]
}