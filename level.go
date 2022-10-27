package coco

import "errors"

// Level is a custom type created and exposed with the purpose to bring the user a clear way to create their own levels
type Level string

// ErrLevelAlreadyDefined is an error message displayed when the user wants to create a new Level but the name selected
// already exists in the map
var ErrLevelAlreadyDefined = errors.New("the level name indicated is in use, please choose another one")

const (
	// Error commonly used to display messages that cannot be addressed by the application
	Error Level = "error"

	// Success commonly used to inform the user that the process has been finalized successfully
	Success Level = "success"

	// Warning commonly used to inform the user about an error but the application can continue regardless
	Warning Level = "warning"

	// Info commonly used for debugging purposes and hidden to the final user
	Info Level = "info"
)

// String converts the custom type Level to its real type
func (l Level) String() string {
	return string(l)
}
