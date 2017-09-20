package imago

import (
	"unicode"
	"unicode/utf8"
)

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

// FilterEmoji filterEmoji
func (t Text) FilterEmoji(str string) string {
	data := ""
	for _, value := range str {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			data += string(value)
		}
	}
	return data
}
