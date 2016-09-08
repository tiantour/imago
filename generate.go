package imago

import (
	"math"
	"math/rand"
	"time"
)

// Number 生成随机数字
func (g *igenerate) Number(length float64) (n int) {
	i := int(math.Pow(10, length-1))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n = r.Intn(9*i) + i
	return
}

// String 生成随机字符串
func (g *igenerate) String(length int) (s string) {
	b := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	d := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		d = append(d, b[r.Intn(62)])
	}
	s = string(d)
	return
}
