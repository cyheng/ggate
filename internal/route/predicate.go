package route

import (
	"errors"
	"fmt"
	"strings"
)

type PredicateDefinition struct {
	Name string
	Args map[string]string
}

func NewPredicateDefinition(text string) (*PredicateDefinition, error) {

	eqIdx := strings.IndexRune(text, '=')
	if eqIdx <= 0 {
		return nil, errors.New("Unable to parse PredicateDefinition text '" +
			text + "'" + ", must be of the form name=value")
	}
	result := &PredicateDefinition{
		Name: text[0:eqIdx],
		Args: make(map[string]string),
	}
	args := strings.Split(text[eqIdx+1:], ",")
	for i, v := range args {
		result.Args[fmt.Sprint("_genkey_", i)] = v
	}
	return result, nil
}
