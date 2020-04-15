package main

import (
	"fmt"
	"math"
	"os"
	"log"
	"bytes"
	"io"
	"sync"

	"testing"
	"github.com/stretchr/testify/assert"
)

// Test case when wrong number of arguments are provided
func TestCheckArgsCountWrongNum(t *testing.T) {
	testArgsWrongNum := [][]string{
		{
			"Test wrong number of arguments 1",
			cmdInit,
			"test-args",
			"help",
		},
		{
			"Test wrong number of arguments 2",
			cmdInit,
			"build-word-chain",
			startFl,
			endFl,
			"end-word-test",
		},
		{
			"Test wrong number of arguments 3",
			cmdInit,
			cmdString,
			startFl,
			"start-word-test",
			endFl,
			"end-word-test",
			"end-word-test",
		},
	}

	expectedOutput := func(args []string) {
		fmt.Printf(ErrWrongArgsNum(len(args)-1))
		fmt.Printf(AvailCmdsOutput())
		fmt.Printf(RunHelpCmdOutput(args[0]))
	}

	for _, test := range testArgsWrongNum {
		expOutput := captureOutput(func() {
			expectedOutput(test[1:])
		})

		expErrorText := ErrWrongArgsNum(len(test)-2)

		var errOutput error
		resOuput := captureOutput(func() {
			errOutput = checkArgsCount(test[1:])
		})

		assert.Equal(t, expOutput, resOuput, fmt.Sprintf("Failed at: %s", test[0]))
		assert.Equal(t, expErrorText, errOutput.Error(), fmt.Sprintf("Failed at: %s", test[0]))
	}
}

// Test case when -help flag should be used
func TestCheckArgsCountHelpFlag(t *testing.T) {
	testArgsHelpFlag := [][]string{
		{
			"Test arguments with help flag 1",
			cmdInit,
			"test-args",
			"-help",
		},
		{
			"Test arguments with help flag 2",
			cmdInit,
			"test-args",
			"-help",
		},
	}

	for _, test := range testArgsHelpFlag {
		var errOutput error
		resOuput := captureOutput(func() {
			errOutput = checkArgsCount(test[1:])
		})

		assert.Equal(t, "", resOuput, fmt.Sprintf("Failed at: %s", test[0]))
		assert.Equal(t, nil, errOutput, fmt.Sprintf("Failed at: %s", test[0]))
	}
}

// Test case when correct number of arguments is provided
func TestCheckArgsCountCorrectNum(t *testing.T) {
	testArgs := [][]string{
		{
			"Test arguments 1",
			cmdInit,
			cmdString,
			startFl,
			"start-word-test",
			endFl,
			"end-word-test",
		},
		{
			"Test arguments 2",
			cmdInit,
			"test-command",
			"test-flag-start",
			"test-start-word",
			"test-flag-end",
			"test-end-flag",
		},
	}

	for _, test := range testArgs {
		var errOutput error
		resOuput := captureOutput(func() {
			errOutput = checkArgsCount(test[1:])
		})

		assert.Equal(t, "", resOuput, fmt.Sprintf("Failed at: %s", test[0]))
		assert.Equal(t, nil, errOutput, fmt.Sprintf("Failed at: %s", test[0]))
	}
}

// Test cases for characters that need to be cleared
func TestClearNonAlphCharsClear(t *testing.T) {
	testInput1 := []string{
		fmt.Sprintf("a!@#$^&*()_$&346%d", math.MaxInt64), 
		fmt.Sprintf("a!@#$^&*()_$&346%d", math.MaxInt8),
	}
	testInput2 := []string{
		"boundary-very-long-string-input58934^&*%$75389457", 
		"boundary-very-long-string-input58934^&*%$75389457",
	}
	testInput3 := []string{
		fmt.Sprintf("!@#$^&boundary-test-input*()_$&346%d", math.MaxInt8),
		fmt.Sprintf("a!@#$^&*()_$&346boundary-test-input%d", math.MaxInt8),
	}
	testInput4 := []string{
		"--boundary-test-input",
		"boundary-test-input--",
	}
	testInput5 := []string{
		"AA-boundary-test-input--1287324",
		fmt.Sprintf("a%dAAAA-boundary-test-input", math.MaxInt32),
	}

	resInput, err := clearNonAlphChars(testInput1)
	expectedTestInput1 := []string{"a", "a"}
	assert.Equal(t, expectedTestInput1, resInput, fmt.Sprint("Failed at testIput 1"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testIput 1"))

	resInput, err = clearNonAlphChars(testInput2)
	expectedTestInput2 := []string{"boundary-very-long-string-input", "boundary-very-long-string-input"}
	assert.Equal(t, expectedTestInput2, resInput, fmt.Sprint("Failed at testIput 2"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testIput 2"))

	resInput, err = clearNonAlphChars(testInput3)
	expectedTestInput3 := []string{"boundary-test-input", "aboundary-test-input"}
	assert.Equal(t, expectedTestInput3, resInput, fmt.Sprint("Failed at testIput 3"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testIput 3"))

	resInput, err = clearNonAlphChars(testInput4)
	expectedTestInput4 := []string{"boundary-test-input", "boundary-test-input"}
	assert.Equal(t, expectedTestInput4, resInput, fmt.Sprint("Failed at testIput 4"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testIput 4"))

	resInput, err = clearNonAlphChars(testInput5)
	expectedTestInput5 := []string{"AA-boundary-test-input", "aAAAA-boundary-test-input"}
	assert.Equal(t, expectedTestInput5, resInput, fmt.Sprint("Failed at testIput 5"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testIput 5"))
}

// Test cases for no characters to be cleared from input
func TestClearNonAlphCharsNotClear(t *testing.T) {
	testInput1 := []string{
		"a-boundary-string-input", 
		"a-boundary-string-input-a",
	}
	testInput2 := []string{
		"boundary-very-long-string-input", 
		"boundary-very-long-string-input",
	}
	
	resInput1, _ := clearNonAlphChars(testInput1)
	expectedTestInput1 := testInput1
	assert.Equal(t, expectedTestInput1, resInput1)

	resInput2, _ := clearNonAlphChars(testInput2)
	expectedTestInput2 := testInput2
	assert.Equal(t, expectedTestInput2, resInput2)
}

// Captute output from os.Stderr that needs to be checked across the package functions
// for testing
func captureOutput(f func()) string {
	// Initialize os.Pipeline - creates a pipe btewee nreader and writer with *os.File type
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// Replace strandad os output and err with pipeline
	// and set the output for a log package
	// log.Setouput changes the address of the output object
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer 
	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// We create a new goroutine because read and write cannot stay in one
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()

	writer.Close()
	return <- out
}
