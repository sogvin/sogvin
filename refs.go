package sogvin

import (
	"path"

	. "github.com/gregoryv/web"
)

func github(pth, label string) *Element {
	return A(
		Href(path.Join("https://github.com", pth)),
		label,
	)
}
