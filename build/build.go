package build

// BuildGraph calls the internal build package functionality for 
// building shortest path and returns the bfspath result
// used for external calls from cmd package
func BuildGraph(initBound, endBound string) ([]string, error) {
	if err := calcShortestPath(initBound, endBound); err != nil {
		return nil, err
	}

	return bfspath, nil
}
