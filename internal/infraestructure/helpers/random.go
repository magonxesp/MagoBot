package helpers

import (
	"math/rand"
	"time"
)

func RandomInt(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return min + rand.Intn(max-min)
}
