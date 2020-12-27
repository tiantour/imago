package imago

import (
	"regexp"
	"strconv"
	"strings"
)

// Veriry verify
type Veriry struct{}

// NewVerify new verify
func NewVerify() *Veriry {
	return &Veriry{}
}

// Card verify card
func (v *Veriry) Card(s string) bool {
	var x string
	var y int
	var z = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	for i := 0; i < 18; i++ {
		if i == 17 {
			x = strings.ToUpper(s[i : i+1])
			continue
		}
		t, _ := strconv.Atoi(s[i : i+1])
		y += z[i] * t
	}

	y = y % 11
	y = (12 - y) % 11

	ok := x == strconv.Itoa(y)
	if y == 10 {
		ok = x != "X"
	}
	return ok

}

// Email verify email
func (v *Veriry) Email(s string) bool {
	pattern := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(s)
}

// Phone verify phone
func (v *Veriry) Phone(s string) bool {
	pattern := `^(13[0-9]|14[0-9]|15[0-9]|16[0-9]|17[0-9]|18[0-9]|19[0-9])\d{8}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(s)
}

// Telephone verify telephone
func (v *Veriry) Telephone(s string) bool {
	pattern := `^(\d{3}-\d{8}|\d{4}-\d{7})$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(s)
}
