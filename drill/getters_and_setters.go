// Go idiomic getters and setters
//
// Exclude the Get prefix and group methods by behavior. Ie.
// Set methods are settings and getters are attribute readers.
// And separate operations.
package drill

import "fmt"

func init() {
	fmt.Print(NewHotel().String())
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
	floors    []Floor
}

// group by behavior
func (me *Building) String() string {
	return fmt.Sprintf(
		"House with %v floors and %v rooms",
		len(me.floors), me.RoomCount(),
	)
}

func (me *Building) RoomCount() int {
	var sum int
	for _, f := range me.floors {
		sum += f.RoomCount()
	}
	return sum
}

func (me *Building) SetArea(v int)      { me.area = v }
func (me *Building) SetEntrances(v int) { me.entrances = v }

func (me *Building) Area() int       { return me.area }
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
