package shapes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInBounds(t *testing.T) {
	var r = &Rectangle{
		X:       9,
		Y:       12,
		Width:   6,
		Height:  6,
		Fill:    "x",
		Outline: "o",
	}
	var inBoundsPoint = [2]int{10, 15}
	var outOfBoundsPoint = [2]int{1, 2}
	assert.True(t, r.isInBounds(inBoundsPoint[0], inBoundsPoint[1]), "Should be true")
	assert.False(t, r.isInBounds(outOfBoundsPoint[0], outOfBoundsPoint[1]), "Should be false")
}

func TestisOutline(t *testing.T) {
	var r = &Rectangle{
		X:       9,
		Y:       12,
		Width:   6,
		Height:  6,
		Fill:    "x",
		Outline: "o",
	}
	var isOutlinePoint = [2]int{9, 16}
	var isNotOutlinePoint = [2]int{10, 15}
	assert.True(t, r.isOutline(isOutlinePoint[0], isOutlinePoint[1]), "Should be true")
	assert.False(t, r.isOutline(isNotOutlinePoint[0], isNotOutlinePoint[1]), "Should be false")
}

func TestGetPoint(t *testing.T) {
	var r = &Rectangle{
		X:       9,
		Y:       12,
		Width:   6,
		Height:  6,
		Fill:    "x",
		Outline: "o",
	}
	var OutlinePoint = [2]int{9, 16}
	var InsidePoint = [2]int{10, 16}
	var OutsidePoint = [2]int{1, 1}
	assert.Equal(t, "o", r.GetPoint(OutlinePoint[0], OutlinePoint[1]))
	assert.Equal(t, "x", r.GetPoint(InsidePoint[0], InsidePoint[1]))
	assert.Equal(t, " ", r.GetPoint(OutsidePoint[0], OutsidePoint[1]))
}
