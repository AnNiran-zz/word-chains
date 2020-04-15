package build

// calcShortestPath creates the graph dynamically -
// when a vertice is visited - its neighboring words list is generated
// 
func calcShortestPath(initBound, endBound string) error {
	var err error

	// Find words range between to the two bounds - comprise the universe
	// the universe is the set of all unvisited nodes - unvisited set
	wordsRange, err = findRange(initBound, endBound)
	if err != nil {
		return err
	}

	// vQueue is the queue functioning at FIFO principle
	// we add the inital boundary vertice (start word) to the queue - 
	// initial boundary vertice is with zero tentative value
	// 
	// rest of the unvisited vertices are theoretically with infinite tentative values
	// until they are visited
	vQueue = enqueue(vQueue, initBound)

	//
	for i := 0; i < len(vQueue); i++ {
		// For the current vertice - get its neighbors with distance 1
		// save them inside vReach (vertice reachable) variable
		neighbors(vQueue[i])

		// Loop over the received values in vReach 
		// check if each new vertice is present in the map, if not insert it
		// with predecessor value of the current queue element
		for _, edge := range vReach {
			if _, ok := vpmap[edge]; !ok {
				vpmap[edge] = vQueue[i]
				vQueue = enqueue(vQueue, edge)
			}
		}

	}

	getPath(initBound, endBound)
	return nil
}

// getPath constructs a path using bfspath variable
// the path is built in reverse starting from the ending vertice and looking at its parents
func getPath(initBound, endBound string) {
	if _, ok := vpmap[endBound]; !ok {
		return
	}

	vertice := endBound
	// loop in reverse direction by looking at parents at each step
	// at each step check if local variable word equals the initial vertice - if not 
	// add it to the path at the beginning of the bfspath
	// 
	for vertice != initBound {
		// add the vertice at current step to the beginning of the bfspath
		bfspath = append([]string{vertice}, bfspath...)

		// Change the vertice value to its parent
		vertice = vpmap[vertice]
	}

	// Add the initial vertice last because we are moving in reverse order
	bfspath = append([]string{initBound}, bfspath...)
}

// neighbors returns the edges-words that can be reached 
// from a given reference word- vertice according to the set diference step of 1
func neighbors(ref string) {
	vReach = nil

	for _, word := range wordsRange {
		if isNeighbor(ref, word) {
			vReach = append(vReach, word)
		}
	}	
}

// isNeighbor calculates the distance between a reference word 
// and another word of the same length and returns true
// if the difference is exactly one character
func isNeighbor(ref, word string) bool {
	diff := 0
	for i := 0; i < len(ref); i++ {
		if ref[i] != word[i] {
			diff++
		}
	}
	return diff == 1
}

// enqueue pushes an element to the end of the vertice queue
func enqueue(q []string, element string) []string {
	q = append(q, element)
	return q
}

// dequeue removes the first element from the queue
// and move one step backward all other elements
func dequeue(q []string) []string {
	return q[1:]
}
