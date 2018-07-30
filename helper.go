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

// ConvertToStringSlice converts interface array to string array
func ConvertToStringSlice(slice []interface{}) []string {
	stringSlice := make([]string, len(slice))
	for i, item := range slice {
		itemString, ok := item.(string)
		if !ok {
			err := errors.New("found elements of non string type")
			panic(err)
		}
		stringSlice[i] = itemString
	}
	return stringSlice
}

// ConvertToIntSlice converts interface slice to int slice
func ConvertToIntSlice(slice []interface{}) []int {
	intSlice := make([]int, len(slice))
	for i, item := range slice {
		var itemInt int
		switch item.(type) {
		case int:
			itemInt = item.(int)
		case int8:
			itemInt = int(item.(int8))
		case int16:
			itemInt = int(item.(int16))
		case int32:
			itemInt = int(item.(int32))
		case int64:
			itemInt = int(item.(int64))
		case float32:
			itemInt = int(item.(float32))
		case float64:
			itemInt = int(item.(float64))
		default:
			err := errors.New("found elements of non int type")
			panic(err)
		}
		intSlice[i] = itemInt
	}
	return intSlice
}
