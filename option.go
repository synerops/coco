package coco

import (
	"io"
	"os"

	"github.com/fatih/color"
)

type Option struct {
	levels map[Level]Output
}

const GlobalFormat = "[%s]: %s\n"

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

func (o *Option) SetFormat(l Level, f string) {
	if level, ok := o.levels[l]; ok {
		level.Format = f

		o.levels[l] = level
	}
}

func (o *Option) SetWriter(l Level, w io.Writer) {
	if level, ok := o.levels[l]; ok {
		level.Writer = w

		o.levels[l] = level
	}
}

func (o *Option) SetColor(l Level, c color.Attribute) {
	if level, ok := o.levels[l]; ok {
		level.Color = c

		o.levels[l] = level
	}
}

func (o *Option) NewLevel(name Level, output Output) error {
	if _, ok := o.levels[name]; ok {
		return ErrLevelAlreadyDefined
	}

	o.levels[name] = output

	return nil
}
