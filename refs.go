package sogvin

import (
	. "github.com/gregoryv/web"
)

func github(pth, label string) *Element {
	return A(
		Href("https://github.com/"+pth),
		label,
	)
}
