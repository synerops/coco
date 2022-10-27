package coco

import (
	"io"
	"os"

	"github.com/fatih/color"
)

// Option is a public struct that brings the user the possibility to modify default configurations on each level.
// In addition, Option brings the possibility to create new levels
type Option struct {
	levels map[Level]Output
}

// GlobalFormat used by all the levels and indicates how to display the level and its corresponding message
//
//	[SUCCESS]: This is an example of a success message following the GlobalFormat
const GlobalFormat = "[%s]: %s\n"

// Default is the proposal created by Coco to configure the default levels, based on regular considerations but open
// to any change the user wants to make
func Default() *Option {
	return &Option{
		levels: map[Level]Output{
			Error: {
				Format: GlobalFormat,
				Writer: os.Stderr,
				Color:  color.FgRed,
			},
			Success: {
				Format: GlobalFormat,
				Writer: os.Stdout,
				Color:  color.FgGreen,
			},
			Warning: {
				Format: GlobalFormat,
				Writer: os.Stdout,
				Color:  color.FgYellow,
			},
			Info: {
				Format: GlobalFormat,
				Writer: io.Discard,
				Color:  color.FgHiCyan,
			},
		},
	}
}

// SetFormat brings the user the possibility to change the GlobalFormat of a specific Level
func (o *Option) SetFormat(l Level, f string) {
	if level, ok := o.levels[l]; ok {
		level.Format = f

		o.levels[l] = level
	}
}

// SetWriter brings the user the possibility to change the default Writer (io.Writer) of a specific Level
func (o *Option) SetWriter(l Level, w io.Writer) {
	if level, ok := o.levels[l]; ok {
		level.Writer = w

		o.levels[l] = level
	}
}

// SetColor brings the user the possibility to change the color (color.Attribute) of a specific Level
func (o *Option) SetColor(l Level, c color.Attribute) {
	if level, ok := o.levels[l]; ok {
		level.Color = c

		o.levels[l] = level
	}
}

// NewLevel brings the user the possibility to create custom Levels
func (o *Option) NewLevel(name Level, output Output) error {
	if _, ok := o.levels[name]; ok {
		return ErrLevelAlreadyDefined
	}

	o.levels[name] = output

	return nil
}
