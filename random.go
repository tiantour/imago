package imago

import (
	"math"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// random random
type random struct{}

// Random new random
var Random = new(random)

// Number random number
// date 2016-12-31
// author andy.jiang
func (g random) Number(length float64) int {
	i := int(math.Pow(10, length-1))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9*i) + i
}

// String random string
// date 2016-12-31
// author andy.jiang
func (g random) String(length int) string {
	b := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	d := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		d = append(d, b[r.Intn(62)])
	}
	return string(d)
}

// Ulid random ulid
// date 2016-12-31
// author andy.jiang
func (g random) ULID() string {
	t := time.Now()
	entrop := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entrop).String()
}
