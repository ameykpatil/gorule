package gorule

import "strings"

//FlattenJSON method converts the nested json to flattened map
func FlattenJSON(property map[string]interface{}) map[string]interface{} {
	//create asimple flattened json
	flattenedProfile := make(map[string]interface{})

	//loop for creating flattened property
	for key, value := range property {
		switch child := value.(type) {
		//only in case of nested json we recursively call flattenProfile
		case map[string]interface{}:
			nestedJSON := FlattenJSON(child)
			for nestedKey, nestedValue := range nestedJSON {
				flattenedProfile[key+"."+nestedKey] = nestedValue
			}
		default:
			flattenedProfile[key] = value
		}
	}

	return flattenedProfile
}

//Unflatten function unflattens a map into nested json
func Unflatten(m map[string]interface{}) map[string]interface{} {
	var tree = make(map[string]interface{})
	for key, value := range m {
		ks := strings.Split(key, ".")
		tr := tree
		for _, tk := range ks[:len(ks)-1] {
			trnew, ok := tr[tk]
			if !ok {
				trnew = make(map[string]interface{})
				tr[tk] = trnew
			}
			tr = trnew.(map[string]interface{})
		}
		tr[ks[len(ks)-1]] = value
	}
	return tree
}
