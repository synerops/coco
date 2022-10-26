package coco

import "errors"

type Level string

var ErrLevelAlreadyDefined = errors.New("the level name indicated is in use, please choose another one")

const (
	Error   Level = "error"
	Success Level = "success"
	Warning Level = "warning"
	Info    Level = "info"
)

func (l Level) String() string {
	return string(l)
}
