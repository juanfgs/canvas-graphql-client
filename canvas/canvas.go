package canvas

import (
	"fmt"
	"github.com/juanfgs/canvas-client/shapes"
)

const (
	sizeX = 15
	sizeY = 10
)

type Canvas struct {
	ID       string
	Name     string
	Contents []shapes.Rectangle
}

func (c Canvas) Render() {
	for y := 0; y < sizeY; y++ {
		fmt.Println("")
		for x := 0; x < sizeX; x++ {
			var point string = " "
			for _, r := range c.Contents {
				if point == " " {
					point = r.GetPoint(x, y)
				}
			}
			fmt.Print(point)
		}
	}

	fmt.Println()
	fmt.Println()
}
