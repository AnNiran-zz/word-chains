package build

import (
	"os"
	//"fmt"
	"bufio"
	"sort"
)

// openDictionary opens the dictionary file for reading
func openDictionary() (*os.File, error) {
	return os.Open(dictionary)
}

// findRange returns list of words upon which search graph is going to be built
//
// Finding start bound:
// 1. Comparing first byte -> skip all words that do not start with the specific byte
// 2. Compare length -> when words start with the same byte - check lengths before proceeding to cheging each byte
// 3. When words start with the specific byte and are the same lengths - check each byte; at first mismatch - skip
//
// Finding end bound:
// 1. Compare first byte -> add to range all words that do not start with the specific byte
// 2. Compare length -> when words start with the same byte but are with different length - add them to range
// 3. Compare each byte - when words start with specific byte, are the same length, proceed with checking each byte
func findRange(strtBound, endBound string) ([]string, error) {
	// Read dictionary file
	file, err := openDictionary()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Scan each line from the dictionary
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	// Check order of start and end words and reverse their positions if reversed
	if !orderInput(strtBound, endBound) {
		strtBound, endBound = endBound, strtBound
	}

	strtBoundByts, endBoundByts := []byte(strtBound), []byte(endBound)

	var words []string
	var startIsSet, endIsSet bool
	
	// Read words from dictionary line by line
	for scanner.Scan() {
		// Set up initial word bound
		if startIsSet {
			goto FindEndBound
		}

		// If the first byte of both words is not the same
		// proceed with the next word from the dictionary
		if strtBoundByts[0] != scanner.Bytes()[0] {
			// If the search moved to the next letter -> no match if found for the provided
			// starting word
			if strtBoundByts[0] < scanner.Bytes()[0] {
				// 
				return nil, ErrBoundNotFound(strtBound)
			}

			goto End
		}

		// First bytes are the same:
		// Compare lengths - go to next loop at first mismatch
		if len(strtBoundByts) != len(scanner.Bytes()) {
			goto End
		}

		// Compare bytes from second position
		// Go to next loop at first mismatch
		for i := 1; i < len(strtBoundByts); i++ {
			if strtBoundByts[i] != scanner.Bytes()[i] {
				goto End
			}
		}

		startIsSet = true
		goto AddToRange
	
// Set up the ending word bound
// Mirror the steps
FindEndBound:

		// If lengths of bound and scanned word are not the same, move to the next round
		// Here we check lengths first because words with non-matching first bytes are added
		// to the range
		if len(endBoundByts) != len(scanner.Bytes()) {
			goto End
		}

		// Compare first bytes
		// If mismatch - add to the range
		if endBoundByts[0] != scanner.Bytes()[0] {
			// If the search moved to the next letter -> no match if found for the provided
			// ending word
			if endBoundByts[0] < scanner.Bytes()[0] {
				// 
				return nil, ErrBoundNotFound(endBound)
			}
			goto AddToRange
		}

		// Loop over the bytes of endBound in order
		// stop at the first mismatch and add the current record to the range
		for i := 1; i < len(endBoundByts); i++ {
			
			// If no match - continue adding to the range
			if endBoundByts[i] != scanner.Bytes()[i] {
				goto AddToRange
			}
		}

		// Set endIsSet to true as anchor for breaking out of the loop
		endIsSet = true

AddToRange:
		words = append(words, scanner.Text())

End:
		// If end bound is set we break out of the loop
		if endIsSet {
			break
		}
	}

	return words, nil
}

// Check order of two bounds
// return true if they are in order
// return false if they are in reverse order
func orderInput(bound1, bound2 string) bool {
	order := []string{bound1, bound2}
	sort.Strings(order)

	return order[0] == bound1
}