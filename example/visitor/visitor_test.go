package visitor

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_visitor_pattern(t *testing.T) {
	var (
		square    = &Square{side: 2}
		circle    = &Circle{radius: 3}
		rectangle = &Rectangle{width: 3, height: 2}
		equals    = asserter.Wrap(t).Equals
	)

	equals(CalcArea(square), 4.0)
	equals(CalcArea(circle), 18.84955592153876)
	equals(CalcArea(rectangle), 6.0)

	equals(FindMiddle(square), "1,1")
	equals(FindMiddle(circle), "3,3")
	equals(FindMiddle(rectangle), "1,1")
}
