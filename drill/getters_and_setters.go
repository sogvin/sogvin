// Go idiomic getters and setters
//
// Exclude the Get prefix
package drill

import "fmt"

func init() {
	h := NewHotel()
	fmt.Print(h.String())
}

func NewHotel() *Building {
	var b Building
	b.SetArea(100)
	b.SetEntrances(2)

	for i := 1; i <= 20; i++ {
		var f Floor
		f.SetNumber(i)
		f.SetRoomCount(8)
		b.floors = append(b.floors, f)
	}
	return &b
}

type Building struct {
	area      int
	entrances int

	// complex relation
	floors []Floor
}

// group by behavior

// settings

func (me *Building) SetArea(v int)      { me.area = v }
func (me *Building) SetEntrances(v int) { me.entrances = v }

// attributes

func (me *Building) Area() int       { return me.area }
func (me *Building) Entrances() int  { return me.entrances }
func (me *Building) Floors() []Floor { return me.floors }

// operations

func (me *Building) String() string {
	return fmt.Sprintf("House with %v floors and %v rooms", len(me.floors), me.RoomCount())
}

func (me *Building) RoomCount() int {
	var sum int
	for _, f := range me.floors {
		sum += f.RoomCount()
	}
	return sum
}

// ----------------------------------------

type Floor struct {
	number    int
	roomCount int
}

// settings

func (me *Floor) SetNumber(v int)    { me.number = v }
func (me *Floor) SetRoomCount(v int) { me.roomCount = v }

// attributes

func (me *Floor) Number() int    { return me.number }
func (me *Floor) RoomCount() int { return me.roomCount }
