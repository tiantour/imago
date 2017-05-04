package imago

import "unicode"

// Text text
var Text = &iText{}

type iText struct{}

// IsHan
func (t iText) IsHan(str string) bool {
	for _, v := range str {
		if !unicode.Is(unicode.Scripts["Han"], v) {
			return false
		}
	}
	return true
}

// IsDigit
func (t iText) IsDigit(str string) bool {
	for _, v := range str {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}
