/*
Nap implements real sleep (offering sub-second busy waits).
The first real value argument (base) in seconds determines the base duration.
When given a real number as second argument (variation)
the duration will vary uniform randomly plus minus the 2nd argument.

Default mode just sleeps for a second.
This behavior is identical to when given a single argument of 1.

Usage:

	nap [base [variation]]

The flags are:

	-h
	    display help message

	-v
	    display the version string

Examples:

❯ nap 0.75 0.1  # sleeps for a randomly selected duration in [0.65, 0.85] seconds

Caveat emptor: Maybe this command is not useful for your use cases.
*/
package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULT_SECS = 1.0 // Base value in seconds When no arguments are given
	HELP         = `Nap implements real sleep (offering sub-second busy waits). The first real
value argument (base) in seconds determines the base duration. When given a real
number as second argument (variation) the duration will vary uniform randomly
plus minus the 2nd argument.

Default mode just sleeps for a second. This behavior is identical to
when given a single argument of 1.

Usage:

    nap [base [variation]]

The flags are:

    -h
        display help message

    -v
        display the version string

Examples:

❯ nap 0.75 0.1 # sleeps for a randomly selected duration in [0.65, 0.85] seconds

Caveat emptor: Maybe this command is not useful for your use cases.` // update after go doc nap shows different help
	MILLIS  = 1000 // Rescaling to adapt to the time.Duration domain
	USAGE   = "usage: nap [base [variation]]"
	VERSION = "v2023.9.20"
)

// ParseFloat parses the given text as positive float and returns code, message, and parsed value as tuple.
// The scope value enables the caller to amend the message with context information.
// If the conversion fails, the value is set to Not-a-Number (NaN).
func ParseFloat(text string, scope string) (int, string, float64) {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 2, fmt.Sprintf(USAGE + scope), math.NaN()
	}
	return 0, "", math.Max(0.0, value)
}

// ParseBase attempts to parse the base duration.
// If succeeding, returns the value and nil, else returns NaN and the error.
func ParseBase(text string) (float64, error) {
	code, message, base := ParseFloat(text, " - failed to parse the base duration")
	if code != 0 {
		return math.NaN(), errors.New(message)
	}
	return base, nil
}

// ParseVariation attempts to parse the variation in the context of base.
// If succeeding, ensures the variation is less than or equal to base and returns the value and nil,
// else returns NaN and the error.
func ParseVariation(text string, base float64) (float64, error) {
	code, message, variation := ParseFloat(text, " - failed to parse the variation of duration")
	if code != 0 {
		return math.NaN(), errors.New(message)
	}
	variation = math.Min(base, variation)
	return variation, nil
}

// RandomDuration derives a random duration in seconds given a base and the variation.
func RandomDuration(base float64, variation float64) float64 {
	span := variation + variation
	center := base - variation
	return rand.Float64()*span + center
}

// HelpRequested searches for typical help requests in the command arguments and returns the result.
func HelpRequested(args []string) bool {
	for _, cmd := range args {
		cmdLower := strings.ToLower(cmd)
		if strings.Contains(cmdLower, "h") || strings.Contains(cmdLower, "?") {
			return true
		}
	}
	return false
}

// VersionRequested searches for typical version requests in the command arguments and returns the result.
func VersionRequested(args []string) bool {
	for _, cmd := range args {
		cmdLower := strings.ToLower(cmd)
		if strings.Contains(cmdLower, "v") {
			return true
		}
	}
	return false
}

// HandleAnyErrors prints out the error and exits with code if err is not nil else does nothing.
func HandleAnyErrors(w io.Writer, err error, code int) {
	if err != nil {
		fmt.Fprintln(w, err)
		os.Exit(code)
	}
}

// Execute implements a testable main function.
func Execute(w io.Writer, seed int64, args []string) int {
	if len(args) > 0 && HelpRequested(args) {
		fmt.Fprintln(w, HELP)
		return 0
	}
	if len(args) > 0 && VersionRequested(args) {
		fmt.Fprintln(w, VERSION)
		return 0
	}

	if seed != 0 {
		Seed(seed)
	}
	duration := DEFAULT_SECS
	if len(args) > 0 {
		base, err := ParseBase(args[0])
		HandleAnyErrors(w, err, 2)
		duration = base // duration if no variation requested
		if len(args) > 1 {
			variation, err := ParseVariation(args[1], base)
			HandleAnyErrors(w, err, 2)
			duration = RandomDuration(base, variation) // duration adapted to variation
		}

	}

	time.Sleep(time.Duration(MILLIS*duration) * time.Millisecond)
	return 0
}

// Seed the random number generator.
func Seed(seed int64) {
	rand.Seed(seed)
}

// main perceives the request and acts on it.
func main() {
	os.Exit(Execute(os.Stdout, time.Now().UnixNano(), os.Args[1:]))
}
