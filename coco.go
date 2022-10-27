// Package coco is a combination of utilities that make it easy to print logs in a readable way
// for those applications that use the command line. Its goal is to be simple and easy to understand,
// without extra complexities, in simple words, you get what you set
package coco

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Log structs encapsulates the list of levels defined by default and new ones that might be created by the user
type Log struct {
	opts *Option
}

// New returns a newly created log object
func New(opts *Option) *Log {
	return &Log{
		opts: opts,
	}
}

// Error is an alias to the main Log function, which gives the user a quick way to use this Level
func (l *Log) Error(err error) {
	level := l.opts.levels[Error]
	label := strings.ToUpper(Error.String())

	fmt.Fprintf(level.Writer, level.Format, color.New(level.Color).Sprint(label), err)
}

// Success is an alias to the main Log function, which gives the user a quick way to use this Level
func (l *Log) Success(msg string) {
	l.Log(Success, msg)
}

// Warning is an alias to the main Log function, which gives the user a quick way to use this Level
func (l *Log) Warning(msg string) {
	l.Log(Warning, msg)
}

// Info is an alias to the main Log function, which gives the user a quick way to use this Level
func (l *Log) Info(msg string) {
	l.Log(Info, msg)
}

// Log is the function that combines the settings defined in Options and the levels defined in Level in order to use
// a variety of alternatives for the particular uses of any application
func (l *Log) Log(lev Level, msg string) {
	level := l.opts.levels[lev]
	label := strings.ToUpper(lev.String())

	fmt.Fprintf(level.Writer, level.Format, color.New(level.Color).Sprint(label), msg)
}
