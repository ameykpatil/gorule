package gorule

// ContainsString check if given element to search exists in the given string array
func ContainsString(array []string, searchElem string) bool {
	for _, elem := range array {
		if elem == searchElem {
			return true
		}
	}
	return false
}

// ContainsInt check if given element to search exists in the given int array
func ContainsInt(array []int, searchElem int) bool {
	for _, elem := range array {
		if elem == searchElem {
			return true
		}
	}
	return false
}

// ContainsAllString checks if an array contains all of the strings in another array
func ContainsAllString(array []string, searchElems []string) bool {
	for _, searchElem := range searchElems {
		if !ContainsString(array, searchElem) {
			return false
		}
	}
	return true
}

// ContainsAllInt checks if an array contains all of the ints in another array
func ContainsAllInt(array []int, searchElems []int) bool {
	for _, searchElem := range searchElems {
		if !ContainsInt(array, searchElem) {
			return false
		}
	}
	return true
}

// ContainsAnyString checks if an array contains any of the strings in another array
func ContainsAnyString(array []string, searchElems []string) bool {
	for _, searchElem := range searchElems {
		if ContainsString(array, searchElem) {
			return true
		}
	}
	return false
}

// ContainsAnyInt checks if an array contains any of the ints in another array
func ContainsAnyInt(array []int, searchElems []int) bool {
	for _, searchElem := range searchElems {
		if ContainsInt(array, searchElem) {
			return true
		}
	}
	return false
}
