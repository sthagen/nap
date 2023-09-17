package main

import (
	"bytes"
	"errors"
	"math"
	"os"
	"os/exec"
	"testing"
)

// TestParseBaseValid calls nap.ParseBase with a valid value.
func TestParseBaseValid(t *testing.T) {
	text := "3.1415"
	want := 3.1415
	base, err := ParseBase(text)
	if base != want || err != nil {
		t.Fatalf(`ParseBase("3.1415") = %v, %v, want match for %v, nil`, base, err, want)
	}
}

// TestParseBaseWrongType calls nap.ParseBase with a non-convertible value.
func TestParseBaseWrongType(t *testing.T) {
	text := "foo"
	want := math.NaN()
	msg := errors.New(USAGE + " - failed to parse the base duration")
	base, err := ParseBase(text)
	if !math.IsNaN(base) || err == nil {
		t.Fatalf(`ParseBase("foo") = %v, %v, want match for %v, %v`, base, err, want, msg)
	}
}

// TestParseVariationValid calls nap.ParseVariation with valid values.
func TestParseVariationValid(t *testing.T) {
	text := "1"
	want := 1.0
	base := 3.1415
	variation, err := ParseVariation(text, base)
	if variation != want || err != nil {
		t.Fatalf(`ParseVariation("1", 3.1415) = %v, %v, want match for %v, nil`, base, err, want)
	}
}

// TestParseVariationWrongType calls nap.ParseVariation with a non-convertible value for variation.
func TestParseVariationWrongType(t *testing.T) {
	text := "foo"
	want := math.NaN()
	base := 3.1415
	msg := errors.New(USAGE + " - failed to parse the variation of duration")
	variation, err := ParseVariation(text, base)
	if !math.IsNaN(variation) || err == nil {
		t.Fatalf(`ParseVariation("foo", 3.1415) = %v, %v, want match for %v, %v`, variation, err, want, msg)
	}
}

// TestRandomDurationZeroVariation calls nap.RandomDuration with valid values and corner case of variation zero (0).
func TestRandomDurationZeroVariation(t *testing.T) {
	base := 3.1415
	variation := 0.0
	want := base
	duration := RandomDuration(base, variation)
	if duration != want {
		t.Fatalf(`RandomDuration(3.1415, 0.0) = %v, want match for %v`, duration, want)
	}
}

// TestHelpRequested calls nap.HelpRequested with a match in the list
func TestHelpRequested(t *testing.T) {
	args := []string{"1", "help"}
	helpdDetected := HelpRequested(args)
	if !helpdDetected {
		t.Fatalf(`HelpRequested(["1", "help"]) = %v, %v, want match for %v`, helpdDetected, args, true)
	}
}

// TestHelpRequested calls nap.HelpRequested with no match in the list
func TestHelpRequestedNone(t *testing.T) {
	args := []string{"1", "hel", "/h", "--Help", "-H", "?"}
	helpdDetected := HelpRequested(args)
	if helpdDetected {
		t.Fatalf(`HelpRequested(["1", "hel", "/h", "--Help", "-H", "?"]) = %v, %v, want match for %v`, helpdDetected, args, false)
	}
}

// TestHandleAnyErrorsNone calls nap.HandleAnyErrors with no error so no output and simple return.
func TestHandleAnyErrorsNone(t *testing.T) {
	tt := []struct {
		description string
		error       error
		expect      string
		code        int
	}{
		{"empty", nil, "", 0},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			var output bytes.Buffer
			HandleAnyErrors(&output, tc.error, tc.code)
			if tc.expect != output.String() {
				t.Errorf("got %s but expected %s", output.String(), tc.expect)
			}
		})
	}
}

// TestHandleAnyErrorsSome calls nap.HandleAnyErrors with an error so expect output and process exit with code.
func TestHandleAnyErrorsSome(t *testing.T) {
	error := errors.New("An error.")
	want := "An error.\n"
	code := 2

	// Run the crashing code when FLAG is set
	if os.Getenv("FLAG") == "1" {
		var output bytes.Buffer
		HandleAnyErrors(&output, error, code)
		if want != output.String() {
			t.Errorf("got %s but expected %s", output.String(), want)
		}
		return
	}
	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestHandleAnyErrorsSome")
	cmd.Env = append(os.Environ(), "FLAG=1")
	err := cmd.Run()

	// Cast the error as *exec.ExitError and compare the result
	e, ok := err.(*exec.ExitError)
	expectedErrorString := "exit status 2"
	if expectedErrorString != e.Error() {
		t.Errorf("got %s but expected %s", e.Error(), expectedErrorString)
	}
	if !ok {
		t.Errorf("got %v but expected %v", ok, true)
	}
}

// TestExecute calls nap.Execute with valid numeric arguments.
func TestExecute(t *testing.T) {
	args := []string{"0.2", "0.1"}
	var output bytes.Buffer
	code := Execute(&output, 42, args)
	if code != 0 {
		t.Errorf(`Execute(["0.2", "0.1"]) = %v, want match for %v"`, code, 0)
	}
	if "" != output.String() {
		t.Errorf("got %s but expected ''", output.String())
	}
}

// TestExecuteHelp calls nap.Execute with a valid help argument.
func TestExecuteHelp(t *testing.T) {
	args := []string{"--help"}
	var output bytes.Buffer
	code := Execute(&output, 0, args)
	if code != 0 {
		t.Errorf(`Execute(["--help"]) = %v, want match for %v"`, code, 0)
	}
	if HELP+"\n" != output.String() {
		t.Errorf("got %s but expected 'the help text'", output.String())
	}
}

// TestSeed calls nap.Seed with a constant argument.
func TestSeed(t *testing.T) {
	Seed(42)
}
