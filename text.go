package imago

import "unicode"

// Text text
type Text struct{}

// NewText new text
func NewText() *Text {
	return &Text{}
}

// IsHan text is han
func (t Text) IsHan(str string) bool {
	for _, v := range str {
		if !unicode.Is(unicode.Scripts["Han"], v) {
			return false
		}
	}
	return true
}

// IsDigit text is digit
func (t Text) IsDigit(str string) bool {
	for _, v := range str {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}
