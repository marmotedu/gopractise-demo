package main

import "github.com/ory/ladon"

// StringEqualCondition is an exemplary condition.
type StringEqualCondition struct {
	Equals string `json:"equals"`
}

// Fulfills returns true if the given value is a string and is the
// same as in StringEqualCondition.Equals
func (c *StringEqualCondition) Fulfills(value interface{}, _ *ladon.Request) bool {
	s, ok := value.(string)

	return ok && s == c.Equals
}

// GetName returns the condition's name.
func (c *StringEqualCondition) GetName() string {
	return "StringEqualCondition"
}

var pol = &ladon.DefaultPolicy{
	// ...
	Conditions: ladon.Conditions{
		"some-arbitrary-key": &StringEqualCondition{
			Equals: "the-value-should-be-this",
		},
	},
}

func main() {
}
