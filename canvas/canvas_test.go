package canvas

import (
	"github.com/juanfgs/canvas-client/shapes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	fixture1, err := os.ReadFile("../fixtures/1.txt")
	fixture2, err := os.ReadFile("../fixtures/2.txt")
	if err != nil {
		panic(err)
	}

	var c = &Canvas{
		ID:   "1234-5432-5432-6543",
		Name: "test Canvas",
		Contents: []shapes.Rectangle{{
			X:       3,
			Y:       2,
			Width:   5,
			Height:  3,
			Fill:    "X",
			Outline: "@",
		}, {
			X:       10,
			Y:       3,
			Width:   14,
			Height:  6,
			Fill:    "O",
			Outline: "X",
		},
		},
	}

	assert.Equal(t, string(fixture1), c.GetContents())
	var c2 = &Canvas{
		ID:   "1234-5432-5432-6543",
		Name: "test Canvas",
		Contents: []shapes.Rectangle{{
			X:       14,
			Y:       0,
			Width:   7,
			Height:  6,
			Fill:    ".",
			Outline: "",
		}, {
			X:       5,
			Y:       5,
			Width:   5,
			Height:  3,
			Fill:    "X",
			Outline: "X",
		}, {
			X:       0,
			Y:       3,
			Width:   8,
			Height:  4,
			Fill:    "",
			Outline: "O",
		}},
	}

	assert.Equal(t, string(fixture2), c2.GetContents())
}
