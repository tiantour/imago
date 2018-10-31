package imago

import (
	"fmt"
)

// Query query
type Query struct{}

// NewQuery new query
func NewQuery() *Query {
	return &Query{}
}

// Judge query judge
func (q *Query) Judge(number int32, name string) string {
	judge := "<> ?"
	if number != 0 {
		judge = "= ?"
	}
	return fmt.Sprintf("%s %s", name, judge)
}
