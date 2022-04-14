// Go idiomic getters and setters
//
// Exclude the Get prefix and group methods by behavior. Ie.
// Set methods are settings and getters are attribute readers.
// And separate operations.
package drill

import "fmt"

func init() {
	spec := Specification{
		NumFloors: 20,
		Entrances: 2,
		Length:    30,
		Width:     45,
	}
	hotel := Build(spec)
	fmt.Print(hotel.String())
}

// Carrier struct with no logic, just use public fields
type Specification struct {
	Length    int
	Width     int
	NumFloors int
	Entrances int
}

func Build(s Specification) *Building {
	var b Building
	b.SetArea(s.Width * s.Length * s.NumFloors)
	b.SetEntrances(s.Entrances)

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
	floors    []Floor
}

func (me *Building) String() string { // operations
	return fmt.Sprintf(
		"%v floors, %v rooms, %v m^2",
		len(me.Floors()), me.RoomCount(), me.Area(),
	)
}

func (me *Building) RoomCount() int {
	var sum int
	for _, f := range me.floors {
		sum += f.RoomCount()
	}
	return sum
}

func (me *Building) SetArea(v int)      { me.area = v } // settings
func (me *Building) SetEntrances(v int) { me.entrances = v }

func (me *Building) Area() int       { return me.area } // attributes
func (me *Building) Entrances() int  { return me.entrances }
func (me *Building) Floors() []Floor { return me.floors }

type Floor struct {
	number    int
	roomCount int
}

func (me *Floor) SetNumber(v int)    { me.number = v }
func (me *Floor) SetRoomCount(v int) { me.roomCount = v }

func (me *Floor) Number() int    { return me.number }
func (me *Floor) RoomCount() int { return me.roomCount }
