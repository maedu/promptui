// Package promptui provides ui elements for the command line prompt.
package promptui

import "errors"

// ErrEOF is returned from prompts when EOF is encountered.
var ErrEOF = errors.New("^D")

// ErrInterrupt is returned from prompts when an interrupt (ctrl-c) is
// encountered.
var ErrInterrupt = errors.New("^C")

// ErrAbort is returned when confirm prompts are supplied "n"
var ErrAbort = errors.New("")

// ValidateFunc validates the given input. It should return a ValidationError
// if the input is not valid.
type ValidateFunc func(string) error

// SuccessfulValue returns a value as if it were entered via prompt, valid
func SuccessfulValue(label, value string) string {
	return IconGood + " " + label + ": " + faint(value)
}

// FailedValue returns a value as if it were entered via prompt, invalid
func FailedValue(label, value string) string {
	return IconBad + " " + label + ": " + faint(value)
}
