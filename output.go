package coco

import (
	"io"

	"github.com/fatih/color"
)

// Output is the exported struct that brings the user the possibility to manipulate existing or custom Level
type Output struct {
	Format string
	Writer io.Writer
	Color  color.Attribute
}
