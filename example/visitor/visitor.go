/*
Visitor pattern using geometric shapes.

Use this pattern when you need to allow to extend type behavior
without changing each type. The geometric shapes square, circle and
rectangle in this example all accept behaviors area and middle.

The behaviors can be used directly onto any shape but the
perceptibility is not as good as providing a method, compare;

	var area Area
	circle = Circle{radius: 10}

	// using directly and less perceptible
	circle.Accept(area) // ?? what does this mean
	circleArea := float64(area) // conversion required

With provided func which returns the correct type directly

	circle := Circle{radius: 12}
	area = CalcArea(circle)
*/
package visitor

import (
	"fmt"
	"math"
)

// Shape is a type whose behavior is extendable from external packages.
// The visited type.
type Shape interface {
	Accept(ShapeBehavior)
}

// ShapeBehavior apply to all shapes in this system, the visitor
type ShapeBehavior interface {
	Square(*Square)
	Circle(*Circle)
	Rectangle(*Rectangle)
}

func (me *Square) Accept(b ShapeBehavior)    { b.Square(me) }
func (me *Circle) Accept(b ShapeBehavior)    { b.Circle(me) }
func (me *Rectangle) Accept(b ShapeBehavior) { b.Rectangle(me) }

// ----------------------------------------
// Shapes

type Square struct {
	side int
}

type Circle struct {
	radius int
}

type Rectangle struct {
	width, height int
}

// ----------------------------------------
// Behaviors

// CalcArea calculates the area of any shape
func CalcArea(s Shape) float64 {
	var a Area
	s.Accept(&a)
	return float64(a)
}

// Area is a shape behavior
type Area float64

func (me *Area) Square(s *Square)       { *me = Area(s.side * s.side) }
func (me *Area) Circle(c *Circle)       { *me = Area(2.0 * math.Pi * float64(c.radius)) }
func (me *Area) Rectangle(r *Rectangle) { *me = Area(r.width * r.height) }

// FindMiddle returns x,y of a shape
func FindMiddle(s Shape) string {
	var mid Middle
	s.Accept(&mid)
	return fmt.Sprintf("%v,%v", mid.x, mid.y)
}

type Middle struct {
	x int
	y int
}

func (me *Middle) Square(s *Square)       { me.x = s.side / 2; me.y = me.x }
func (me *Middle) Circle(c *Circle)       { me.x = c.radius; me.y = c.radius }
func (me *Middle) Rectangle(r *Rectangle) { me.x = r.width / 2; me.y = r.height / 2 }
