package diff

// EditOperation represents the type of edit operation (INSERT, DELETE, MATCH)
type EditOperation int

// Enumeration of the edit operation types
const (
	INSERT EditOperation = iota
	DELETE
	MATCH
)

// String method returns a string representation of the EditOperation
func (op EditOperation) String() string {
	switch op {
	case INSERT:
		return "INSERT"
	case DELETE:
		return "DELETE"
	case MATCH:
		return "MATCH"
	default:
		return "UNKNOWN"
	}
}

// MyersDiff computes the diff between two slices of strings (source and destination)
func MyersDiff(source, destination []string) []string {
	editScript := myersShortestEditSequence(source, destination)
	return generateDiffOutput(source, destination, editScript)
}

// generateDiffOutput generates the diff output from the editScript
func generateDiffOutput(source, destination []string, editScript []EditOperation) []string {
	var operations []string
	sourceIndex, destinationIndex := 0, 0
	for _, operation := range editScript {
		switch operation {
		case INSERT:
			operations = append(operations, "\n+ "+destination[destinationIndex])
			destinationIndex++
		case MATCH:
			operations = append(operations, "\n  "+source[sourceIndex])
			sourceIndex++
			destinationIndex++
		case DELETE:
			operations = append(operations, "\n- "+source[sourceIndex])
			sourceIndex++
		}
	}

	return operations
}

// myersShortestEditSequence computes the shortest edit sequence between source and destination using Myers' algorithm
func myersShortestEditSequence(source, destination []string) []EditOperation {
	sourceLength := len(source)
	destinationLength := len(destination)
	maxLength := sourceLength + destinationLength
	var trace []map[int]int
	var currentX, currentY int

loop:
	for editDistance := 0; editDistance <= maxLength; editDistance++ {
		currentVector := make(map[int]int, editDistance+2)
		trace = append(trace, currentVector)

		// Initialize with common prefix
		if editDistance == 0 {
			commonLength := 0
			for len(source) > commonLength && len(destination) > commonLength && source[commonLength] == destination[commonLength] {
				commonLength++
			}
			currentVector[0] = commonLength
			if commonLength == sourceLength && commonLength == destinationLength {
				break loop
			}
			continue
		}

		lastVector := trace[editDistance-1]

		// Iterate through the diagonals
		for diagonal := -editDistance; diagonal <= editDistance; diagonal += 2 {
			// Choose the direction to move (down or right)
			if diagonal == -editDistance || (diagonal != editDistance && lastVector[diagonal-1] < lastVector[diagonal+1]) {
				currentX = lastVector[diagonal+1]
			} else {
				currentX = lastVector[diagonal-1] + 1
			}

			currentY = currentX - diagonal

			// Move diagonally while elements match
			whileMatch(source, destination, &currentX, &currentY)

			currentVector[diagonal] = currentX

			// If the end of both source and destination is reached, break the loop
			if currentX == sourceLength && currentY == destinationLength {
				break loop
			}
		}
	}

	// Generate the edit script by backtracking through the trace
	return backtrack(trace, sourceLength, destinationLength)
}

// whileMatch moves diagonally (increments currentX and currentY) while elements at the current positions in source and destination match
func whileMatch(source, destination []string, currentX, currentY *int) {
	sourceLength := len(source)
	destinationLength := len(destination)

	for *currentX < sourceLength && *currentY < destinationLength && source[*currentX] == destination[*currentY] {
		*currentX = *currentX + 1
		*currentY = *currentY + 1
	}
}

// backtrack generates the edit script by backtracking through the trace
func backtrack(trace []map[int]int, sourceLength, destinationLength int) []EditOperation {
	var editScript []EditOperation

	currentX := sourceLength
	currentY := destinationLength
	var diagonal, prevDiagonal, prevX, prevY int

	for editDistance := len(trace) - 1; editDistance > 0; editDistance-- {
		diagonal = currentX - currentY
		lastVector := trace[editDistance-1]

		// Choose the direction to move (down or right)
		if diagonal == -editDistance || (diagonal != editDistance && lastVector[diagonal-1] < lastVector[diagonal+1]) {
			prevDiagonal = diagonal + 1
		} else {
			prevDiagonal = diagonal - 1
		}

		prevX = lastVector[prevDiagonal]
		prevY = prevX - prevDiagonal

		// Move diagonally while elements match
		for currentX > prevX && currentY > prevY {
			editScript = append(editScript, MATCH)
			currentX -= 1
			currentY -= 1
		}

		// Add the appropriate edit operation
		if currentX == prevX {
			editScript = append(editScript, INSERT)
		} else {
			editScript = append(editScript, DELETE)
		}

		currentX, currentY = prevX, prevY
	}

	// Add remaining moves
	if trace[0][0] != 0 {
		for i := 0; i < trace[0][0]; i++ {
			editScript = append(editScript, MATCH)
		}
	}

	// Reverse the edit script to get the correct order
	return reverseEditScript(editScript)
}

// reverseEditScript reverses the order of the edit operations in the editScript
func reverseEditScript(editScript []EditOperation) []EditOperation {
	reversedScript := make([]EditOperation, len(editScript))
	for i, operation := range editScript {
		reversedScript[len(editScript)-i-1] = operation
	}

	return reversedScript
}
