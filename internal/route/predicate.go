package route

import (
	"errors"
	"fmt"
	"ggate/internal/context"
	"strings"
)

type Predicate interface {
	Name() string
	Apply(ctx *context.Context) bool
}
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

//TODO :根据Path=/registry/** 转换成对应的Predicate
func createPredi(text string) (Predicate, error) {
	return nil, nil
}
