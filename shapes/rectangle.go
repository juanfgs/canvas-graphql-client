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
	if r.isOutline(x, y) {
		return r.Outline
	}
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

func (r Rectangle) isOutline(x int, y int) bool {
	if x == r.X && y >= r.Y && y < r.Y+r.Height {
		return true
	}
	if y == r.Y && x >= r.X && x < r.X+r.Width {
		return true
	}
	if x == r.X+r.Width && y >= r.Y && y < r.Y+r.Height {
		return true
	}
	if y == r.Y+r.Height && x >= r.X && x <= r.X+r.Width {
		return true
	}

	return false
}
