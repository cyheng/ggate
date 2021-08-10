package route

import "ggate/internal/context"

type Filter interface {
	Name() string
	Request(context *context.Context) error
	Response(context *context.Context) error
}

func GetDefaultFilters() ([]Filter, error) {
	return nil, nil
}

//TODO :根据StripPrefix=1 转换成对应的Filter
func createFilter(text string) (Filter, error) {
	return nil, nil
}
