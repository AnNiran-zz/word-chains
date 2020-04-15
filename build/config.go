package build

import (
	"fmt"
)

// dictionary holds the source file location, where the words range is collected from
var dictionary = "dictionary/wordlist.txt"

var reverse bool

// Error outputs
var (
	ErrBoundNotFound = func(bound string) error {
		return fmt.Errorf("Provided word is not found in dictionary: %s", bound)
	}
)

// Types and variables set for heap memory allocations that are 
// extensively used across the functions

type universe []string

// wordsRange holds the list of words from dictionary
// that will be used for building the graph for finding shortest path, a.k.a. the unverse
var wordsRange universe

type edges []string

// edges is the list of neigboring words that a word-vertice can reach according
// the difference step of 1
var vReach edges

type queue []string

// vQueue is used for first-in-first-out frontier for running 
// the algorithm
var vQueue queue

// vpmap maps a verice to its predecessor
// vertice : predecessor
// used to remember visited vertices (words) as well
var vpmap = map[string]string{}

// bfspath holds the words comprising the shortest path of a breadth-first search algorithm
var bfspath []string
