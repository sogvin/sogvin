package visitor

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_visitor_pattern(t *testing.T) {
	var (
		Square    = &Square{side: 2}
		Circle    = &Circle{radius: 3}
		Rectangle = &Rectangle{width: 3, height: 2}
		equals    = asserter.Wrap(t).Equals
	)

	equals(CalcArea(Square), 4.0)
	equals(CalcArea(Circle), 18.84955592153876)
	equals(CalcArea(Rectangle), 6.0)

	equals(FindMiddle(Square), "1,1")
	equals(FindMiddle(Circle), "3,3")
	equals(FindMiddle(Rectangle), "1,1")
}
