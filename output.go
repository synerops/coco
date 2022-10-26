package coco

import (
	"io"

	"github.com/fatih/color"
)

type Output struct {
	Format string
	Writer io.Writer
	Color  color.Attribute
}
