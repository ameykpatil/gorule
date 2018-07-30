package gorule

import (
	"encoding/json"
	"fmt"
	"strings"
	"errors"
)

//Operator is enum for boolean operator And, Or, Not
type Operator string

const (
	And Operator = "AND"
	Or  Operator = "OR"
	Not Operator = "NOT"
)

//Comparator is enum for supported comparisons
type Comparator string

const (
	Eq          Comparator = "="
	Neq         Comparator = "!="
	Lt          Comparator = "<"
	Lte         Comparator = "<="
	Gt          Comparator = ">"
	Gte         Comparator = ">="
	Contains    Comparator = "contains"
	ContainsAll Comparator = "contains all"
	ContainsAny Comparator = "contains any"
)

// Rule is a combination of conditions to satisfy
type Rule struct {
	// There can be either 'Path, Comparator & Value' OR 'Rules & Operator'
	// Path, Comparator, Value means leaf node
	Path       string      `json:"path,omitempty"`
	Comparator Comparator  `json:"comparator,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	// Rules & Operators means non-leaf node
	Rules    []Rule   `json:"rules,omitempty"`
	Operator Operator `json:"operator,omitempty"`
}

// Apply checks if rule applies to an object
func Apply(object interface{}, rule Rule) bool {
	var objMap map[string]interface{}
	objJSON, err := json.Marshal(object)
	if err == nil {
		err = json.Unmarshal(objJSON, &objMap)
	}
	if err != nil {
		return false
	}

	flatObjMap := FlattenJSON(objMap)
	return evaluate(flatObjMap, rule)
}

func evaluate(flatObjMap map[string]interface{}, rule Rule) bool {
	json, er := json.Marshal(rule)
	fmt.Println(string(json), er)

	if rule.Path != "" {
		return match(flatObjMap, rule.Path, rule.Comparator, rule.Value)
	}
	result := false
	switch rule.Operator {
	case And:
		result = true
		for _, singleRule := range rule.Rules {
			result = result && evaluate(flatObjMap, singleRule)
		}
	case Or:
		for _, singleRule := range rule.Rules {
			result = result || evaluate(flatObjMap, singleRule)
		}
	case Not:
		if len(rule.Rules) < 1 {
			singleRule := rule.Rules[0]
			result = !evaluate(flatObjMap, singleRule)
		} else {
			e := errors.New("Not operator can not be with multiple rules")
			panic(e)
		}
	}
	return result
}

func match(flatObjMap map[string]interface{}, path string, comparator Comparator, value interface{}) bool {
	for key, val := range flatObjMap {
		if typeOf(val) != typeOf(value) {
			err := errors.New("unequal types " + typeOf(val) + " & " + typeOf(value))
			panic(err)
		}
		if key == path {
			switch comparator {
			case Eq:
				return val == value
			case Neq:
				return val != value
			case Lt:
				return lessThan(val, value)
			case Lte:
				return lessThanEqualTo(val, value)
			case Gt:
				return greaterThan(val, value)
			case Gte:
				return greaterThanEqualTo(val, value)
			case Contains:
				return contains(val, value)
			case ContainsAll:
				return containsAll(val, value)
			case ContainsAny:
				return containsAny(val, value)
			}
		}
	}
	return false
}

func lessThan(firstValue interface{}, secondValue interface{}) bool {
	switch firstValue.(type) {
	case int, int8, int16, int32, int64:
		first, _ := firstValue.(int64)
		second, _ := secondValue.(int64)
		return first < second
	case float32, float64:
		first, _ := firstValue.(float64)
		second, _ := secondValue.(float64)
		return first < second
	default:
		err := errors.New("< operator not allowed on " + typeOf(firstValue) + " & " + typeOf(secondValue))
		panic(err)
	}
}

func lessThanEqualTo(firstValue interface{}, secondValue interface{}) bool {
	switch firstValue.(type) {
	case int, int8, int16, int32, int64:
		first, _ := firstValue.(int64)
		second, _ := secondValue.(int64)
		return first <= second
	case float32, float64:
		first, _ := firstValue.(float64)
		second, _ := secondValue.(float64)
		return first <= second
	default:
		err := errors.New("<= operator not allowed on " + typeOf(firstValue) + " & " + typeOf(secondValue))
		panic(err)
	}
}

func greaterThan(firstValue interface{}, secondValue interface{}) bool {
	switch firstValue.(type) {
	case int, int8, int16, int32, int64:
		first, _ := firstValue.(int64)
		second, _ := secondValue.(int64)
		return first > second
	case float32, float64:
		first, _ := firstValue.(float64)
		second, _ := secondValue.(float64)
		return first > second
	default:
		err := errors.New("> operator not allowed on " + typeOf(firstValue) + " & " + typeOf(secondValue))
		panic(err)
	}
}

func greaterThanEqualTo(firstValue interface{}, secondValue interface{}) bool {
	switch firstValue.(type) {
	case int, int8, int16, int32, int64:
		first, _ := firstValue.(int64)
		second, _ := secondValue.(int64)
		return first >= second
	case float32, float64:
		first, _ := firstValue.(float64)
		second, _ := secondValue.(float64)
		return first >= second
	default:
		err := errors.New(">= operator not allowed on " + typeOf(firstValue) + " & " + typeOf(secondValue))
		panic(err)
	}
}

func contains(firstValue interface{}, secondValue interface{}) bool {
	switch firstValue.(type) {
	case string:
		first, _ := firstValue.(string)
		second, _ := secondValue.(string)
		return strings.Contains(first, second)
	default:
		err := errors.New("contains operator not allowed on " + typeOf(firstValue) + " & " + typeOf(secondValue))
		panic(err)
	}
}

func containsAll(firstValue interface{}, secondValue interface{}) bool {
	switch firstValue.(type) {
	case []string:
		first, _ := firstValue.([]string)
		second, _ := secondValue.([]string)
		return ContainsAllString(first, second)
	case []int:
		first, _ := firstValue.([]int)
		second, _ := secondValue.([]int)
		return ContainsAllInt(first, second)
	case []interface{}:
		// when firstValue is of the type []interface
		// convert both the values into []interface
		first, _ := firstValue.([]interface{})
		second, _ := secondValue.([]interface{})
		if len(first) > 0 && len(second) > 0 {
			// check the type of an inner item of the slice
			// convert slice of interface to corresponding slice of the type
			switch first[0].(type) {
			case string:
				firstString := ConvertToStringSlice(first)
				secondString := ConvertToStringSlice(second)
				return ContainsAllString(firstString, secondString)
			case int, int8, int16, int32, int64, float32, float64:
				firstInt := ConvertToIntSlice(first)
				secondInt := ConvertToIntSlice(second)
				return ContainsAllInt(firstInt, secondInt)
			default:
				err := errors.New("containAny operator not allowed on types " + typeOf(first[0]) + " & " + typeOf(second[0]))
				panic(err)
			}
		}
		return false	
	default:
		err := errors.New("containAll operator not allowed on " + typeOf(firstValue) + " & " + typeOf(secondValue))
		panic(err)
	}
}

func containsAny(firstValue interface{}, secondValue interface{}) bool {
	switch firstValue.(type) {
	case []string:
		first, _ := firstValue.([]string)
		second, _ := secondValue.([]string)
		return ContainsAnyString(first, second)
	case []int:
		first, _ := firstValue.([]int)
		second, _ := secondValue.([]int)
		return ContainsAnyInt(first, second)
	case []interface{}:
		first, _ := firstValue.([]interface{})
		second, _ := secondValue.([]interface{})
		if len(first) > 0 && len(second) > 0 {
			switch first[0].(type) {
			case string:
				firstString := ConvertToStringSlice(first)
				secondString := ConvertToStringSlice(second)
				return ContainsAnyString(firstString, secondString)
			case int, int8, int16, int32, int64, float32, float64:
				firstInt := ConvertToIntSlice(first)
				secondInt := ConvertToIntSlice(second)
				return ContainsAnyInt(firstInt, secondInt)
			default:
				err := errors.New("containAny operator not allowed on types " + typeOf(first[0]) + " & " + typeOf(second[0]))
				panic(err)
			}
		}
		return false	
	default:
		err := errors.New("containAny operator not allowed on " + typeOf(firstValue) + " & " + typeOf(secondValue))
		panic(err)
	}
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
