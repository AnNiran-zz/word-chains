package build

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

// 
var testRange1 = []string{"aunt", "bunt", "bent", "rent"}
var testRange2 = []string{"bat", "bet", "pet", "pea", "tea"}
var testRange3 = []string{"bat", "bay", "say", "spy"}

func TestCalcShortestPathRubyCode(t *testing.T) {
	dictionary = "testdictionary/wordlist.txt"
	testInput := []struct{
		input    []string
		expected []string
		name     string
	}{
		{
			input:    []string{"ruby", "code"},
			expected: []string{"ruby", "rubs", "robs", "robe", "rode", "code"},
			name:     "test ruby-code chain",
		},
		{			
			input:    []string{"code", "ruby"},
			expected: []string{"code", "rode", "robe", "robs", "rubs", "ruby"},
			name:     "test code-ruby chain",
		},
	}

	for _, test := range testInput {
		err := calcShortestPath(test.input[0], test.input[1])
		assert.Equal(t, nil, err, fmt.Sprint("Test failed for nil error"))
		assert.Equal(t, test.expected, bfspath, fmt.Sprintf("Test failed at: %s", test.name))
	
		wordsRange = []string{}
		vpmap = map[string]string{}
		bfspath = []string{}
		vReach = []string{}
		vQueue = []string{}
	}
}

func TestCalcShortestPathLeadGold(t *testing.T) {
	dictionary = "testdictionary/wordlist.txt"
	testInput := []struct{
		input    []string
		expected []string
		name     string
	}{
		{
			input:    []string{"lead", "gold"},
			expected: []string{"lead", "head", "held", "hold", "gold"},
			name:     "test lead-gold chain",
		},
		{
			input:    []string{"gold", "lead"},
			expected: []string{"gold", "hold", "held", "head", "lead"},
			name:     "test gold-lead chain",
		},
	}

	for _, test := range testInput {
		err := calcShortestPath(test.input[0], test.input[1])

		assert.Equal(t, nil, err, fmt.Sprint("Test failed for nil error"))
		assert.Equal(t, test.expected, bfspath, fmt.Sprintf("Test failed at: %s", test.name))
	
		wordsRange = []string{}
		vpmap = map[string]string{}
		bfspath = []string{}
		vReach = []string{}
		vQueue = []string{}
	}
}

func TestCalcShortestPathCatDog(t *testing.T) {
	dictionary = "testdictionary/wordlist.txt"
	testInput := []struct{
		input    []string
		expected []string
		name     string
	}{
		{
			input:    []string{"cat", "dog"},
			expected: []string{"cat", "cot", "cog", "dog"},
			name:     "test cat-dog chain",
		},
		{
			input:    []string{"dog", "cat"},
			expected: []string{"dog", "cog", "cot", "cat"},
			name:     "test dog-cat chain",
		},
	}

	for _, test := range testInput {
		err := calcShortestPath(test.input[0], test.input[1])
		assert.Equal(t, nil, err, fmt.Sprint("Test failed for nil error"))
		assert.Equal(t, test.expected, bfspath, fmt.Sprintf("Test failed at: %s", test.name))
	
		wordsRange = []string{}
		vpmap = map[string]string{}
		bfspath = []string{}
		vReach = []string{}
		vQueue = []string{}
	}
}

func TestCalcShortestPathBuildGuilt(t *testing.T) {
	dictionary = "testdictionary/wordlist.txt"
	testInput := []struct{
		input    []string
		expected []string
		name     string
	}{
		{
			input:    []string{"build","guilt"},
			expected: []string{"build", "built", "guilt"},
			name:     "test build-guilt chain",
		},
		{
			input:    []string{"guilt","build"},
			expected: []string{"guilt", "built", "build"},
			name:     "test guilt-build chain",
		},
	}

	for _, test := range testInput {
		err := calcShortestPath(test.input[0], test.input[1])
		assert.Equal(t, nil, err, fmt.Sprint("Test failed for nil error"))
		assert.Equal(t, test.expected, bfspath, fmt.Sprintf("Test failed at: %s", test.name))
	
		wordsRange = []string{}
		vpmap = map[string]string{}
		bfspath = []string{}
		vReach = []string{}
		vQueue = []string{}
	}
}

func TestCalcShortestPathCaseTest(t *testing.T) {
	dictionary = "testdictionary/wordlist.txt"
	testInput := []struct{
		input    []string
		expected []string
		name     string
	}{
		{
			input:    []string{"case", "test"},
			expected: []string{"case", "cast", "last", "lest", "test"},
			name:     "test case-test chain",
		},
		{
			input:    []string{"test", "case"},
			expected: []string{"test", "lest", "last", "cast", "case"},
			name:     "test test-case chain",
		},
	}

	for _, test := range testInput {
		err := calcShortestPath(test.input[0], test.input[1])
		assert.Equal(t, nil, err, fmt.Sprint("Test failed for nil error"))
		assert.Equal(t, test.expected, bfspath, fmt.Sprintf("Test failed at: %s", test.name))
	
		wordsRange = []string{}
		vpmap = map[string]string{}
		bfspath = []string{}
		vReach = []string{}
		vQueue = []string{}
	}
}

func TestCalcShortestPathEmpty(t *testing.T) {
	dictionary = "testdictionary/wordlist.txt"
	testInput := []struct{
		input    []string
		expected []string
		name     string
	}{
		{
			input:    []string{"guilt","extraordinary"},
			expected: []string{},
			name:     "test build-guilt chain",
		},
		{
			input:    []string{"case", "blue"},
			expected: []string{},
			name:     "test case-blue chain empty",
		},
	}

	for _, test := range testInput {
		err := calcShortestPath(test.input[0], test.input[1])
		assert.Equal(t, nil, err, fmt.Sprint("Test failed for nil error"))
		assert.Equal(t, test.expected, bfspath, fmt.Sprintf("Test failed at: %s", test.name))
	
		wordsRange = []string{}
		vpmap = map[string]string{}
		bfspath = []string{}
		vReach = []string{}
		vQueue = []string{}
	}
}

func TestCalcShortestPathCasePink(t *testing.T) {
	dictionary = "testdictionary/wordlist.txt"
	testInput := []struct{
		input    []string
		expected []string
		name     string
	}{
		{
			input:    []string{"case", "pink"},
			expected: []string{"case", "cave", "pave", "pane", "pine", "pink"},
			name:     "test case-pink chain",
		},
		{
			input:    []string{"pink", "case"},
			expected: []string{"pink", "pine", "pane", "pave", "cave", "case"},
			name:     "test pink-case chain",
		},
	}

	for _, test := range testInput {
		err := calcShortestPath(test.input[0], test.input[1])
		assert.Equal(t, nil, err, fmt.Sprint("Test failed for nil error"))
		assert.Equal(t, test.expected, bfspath, fmt.Sprintf("Test failed at: %s", test.name))
	
		wordsRange = []string{}
		vpmap = map[string]string{}
		bfspath = []string{}
		vReach = []string{}
		vQueue = []string{}
	}
}

// Test case for no neighbors existing
func TestNeighborsZero(t *testing.T) {
	var result edges

	wordsRange = testRange1
	neighbors("wept")
	assert.Equal(t, vReach, result)

	wordsRange = testRange3
	neighbors("see")
	assert.Equal(t, vReach, result)

	wordsRange = testRange2
	neighbors("spy")
	assert.Equal(t, vReach, result)

	wordsRange = testRange3
	neighbors("tea")
	assert.Equal(t, vReach, result)
}

// Test case for one neighbor existing
func TestNeighborsOne(t *testing.T) {
	var result edges

	wordsRange = testRange1
	neighbors("bend")
	assert.Equal(t, len(vReach), 1)

	wordsRange = testRange2
	neighbors("bee")
	assert.Equal(t, len(vReach), 1)

	wordsRange = testRange3
	neighbors("tea")
	assert.Equal(t, vReach, result)

	// Remove vReach content because we clean it in the beginnig of calling neighbors
	vReach = nil
}

// Test case for multiple neighbors existing
func TestNeighborsMultiple(t *testing.T) {
	wordsRange = testRange1
	neighbors("sent")
	assert.Equal(t, len(vReach), 2)

	wordsRange = testRange2
	neighbors("pee")
	assert.Equal(t, len(vReach), 2)

	wordsRange = testRange3
	neighbors("lay")
	assert.Equal(t, len(vReach), 2)

	// Remove vReach content because we clean it in the beginnig of calling neighbors
	vReach = nil
}

// Test neighboring property according to required step of 1
// character difference between words
func TestIsNeighborFalse(t *testing.T) {
	var testInputNotNeighbors = []struct{
		testDesc  string
		reference string
		word      string
	}{
		{
			testDesc:  "Test not neighbors 1",
			reference: "abases",
			word:      "abandon",
		},
		{
			testDesc:  "Test not neighbors 1",
			reference: "aback",
			word:      "abate",
		},
		{
			testDesc:  "Test not neighbors 3",
			reference: "abundant",
			word:      "embodied",
		},
		{
			testDesc:  "Test not neighbors 4",
			reference: "emerald",
			word:      "nightly",
		},
	}

	for _, args := range testInputNotNeighbors {
		result := isNeighbor(args.reference, args.word)
		assert.Equal(t, result, false, args.testDesc)
	}
}

// Test neighboring property according to required step of 1
// character difference between words
func TestIsNeighborTrue(t *testing.T) {
	var testInputNeighbors = []struct{
		testDesc  string
		reference string
		word      string
	}{
		{
			testDesc:  "Test neighbors 1",
			reference: "nicked",
			word:      "nickel",
		},
		{
			testDesc:  "Test neighbors 1",
			reference: "quaked",
			word:      "quakes",
		},
		{
			testDesc:  "Test neighbors 3",
			reference: "built",
			word:      "quilt",
		},
		{
			testDesc:  "Test neighbors 4",
			reference: "quo",
			word:      "duo",
		},
	}

	for _, args := range testInputNeighbors {
		result := isNeighbor(args.reference, args.word)
		assert.Equal(t, result, true, args.testDesc)
	}
}

