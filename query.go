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
func (q *Query) Judge(name string, number int32) string {
	judge := "<> ?"
	if number != 0 {
		judge = "= ?"
	}
	return fmt.Sprintf("%s %s", name, judge)
}
