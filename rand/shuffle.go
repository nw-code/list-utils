package rand

import (
	"math/rand"
	"time"
)

func Shuffle(s []int) []int {
	rand.Seed(time.Now().UnixNano())
	shuffledSlice := []int{}
	// make a copy of the slice parameter
	sliceCopy := make([]int, len(s))
	copy(sliceCopy, s)

	for len(sliceCopy) > 0 {
		index := rand.Intn(len(sliceCopy))
		shuffledSlice = append(shuffledSlice, sliceCopy[index])
		// remove element
		sliceCopy = append(sliceCopy[:index], sliceCopy[index+1:]...)
	}
	return shuffledSlice
}
