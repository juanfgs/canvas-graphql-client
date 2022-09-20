package shapes

type Rectangle struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Fill    string `json:"fill"`
	Outline string `json:"outline"`
}

func (r Rectangle) GetPoint(x int, y int) string {
	if r.isInBounds(x, y) {
		return r.Fill
	}
	return " "
}

func (r Rectangle) isInBounds(x int, y int) bool {
	if x >= r.X && y >= r.Y {
		if x < r.X+r.Width && y < r.Y+r.Height {
			return true
		}
	}
	return false
}
