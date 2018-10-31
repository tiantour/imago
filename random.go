package imago

import (
	"math"
	"math/rand"
	"time"
)

// Random random
type Random struct{}

// NewRandom new random
func NewRandom() *Random {
	return &Random{}
}

// Number random number
func (rd *Random) Number(length int) int {
	i := int(math.Pow(10, float64(length-1)))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9*i) + i
}

// Text random text
func (rd *Random) Text(length int) string {
	b := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	d := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		d = append(d, b[r.Intn(62)])
	}
	return string(d)
}
