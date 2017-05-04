package imago

import (
	"math"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// Gen gen
var Gen = &iGenerate{}

type iGenerate struct{}

// Number generate number
func (g *iGenerate) Number(length float64) int {
	i := int(math.Pow(10, length-1))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9*i) + i
}

// String generate string
func (g *iGenerate) String(length int) string {
	b := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	d := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		d = append(d, b[r.Intn(62)])
	}
	return string(d)
}

// Ulid generate ulid
func (g *iGenerate) ULID() string {
	t := time.Now()
	entrop := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entrop).String()
}
