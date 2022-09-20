package canvas

import (
	"fmt"
	"github.com/juanfgs/canvas-client/shapes"
	"strings"
)

const (
	sizeX = 24
	sizeY = 11
)

type Canvas struct {
	ID       string
	Name     string
	Contents []shapes.Rectangle
}

func (c Canvas) GetContents() string {
	var sb strings.Builder // string builder should be a bit more efficient than concat
	for y := 0; y < sizeY; y++ {
		sb.WriteString("\n")
		for x := 0; x < sizeX; x++ {
			var point string = " "
			for _, r := range c.Contents {
				if point == " " {
					point = r.GetPoint(x, y)
				}
			}
			sb.WriteString(point)
		}
	}
	fmt.Println(sb.String())
	sb.WriteString("\n")
	return sb.String()
}
