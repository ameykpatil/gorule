# gorule
Rule engine in golang

```
package main

import "github.com/ameykpatil/gorule"

type A struct {
	ID string
	B  B
}

type B struct {
	C string
}

func main() {
	rule := gorule.Rule{
		Operator: And,
		Rules: []Rule{
			{
				Path:       "B.C",
				Comparator: Eq,
				Value:      "abc",
			},
			{
				Path:       "ID",
				Comparator: Eq,
				Value:      "id",
			},
		},
	}

	obj := A{
		ID: "id",
		B: B{
			C: "abc",
		},
	}

	fmt.Println(gorule.Apply(obj, rule))
}
```
