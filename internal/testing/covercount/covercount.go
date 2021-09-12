package covercount

import (
	"fmt"
	"io"
)

const (
	BackInBlack = iota
	Summertime
	RockAroundTheClock
)

func sing(w io.Writer, song int) error {
	switch song {
	case BackInBlack:
		fmt.Fprint(w, "Back in black...")
	case Summertime:
		fmt.Fprint(w, "Its summertime...")
	case RockAroundTheClock:
		fmt.Fprint(w, "Rock, rock, rock around...")
	default:
		return fmt.Errorf("I don't know that one")
	}
	return nil
}
