package route

import (
	"ggate/pkg/config"

	"github.com/sirupsen/logrus"
)

type Route struct {
	Id         string
	Uri        string
	Filters    []Filter    //过滤器
	Predicates []Predicate //断言，当满足这种条件后才会被转发
}
type RouteDefinition struct {
	Id         string   `mapstructure:"id"`
	Uri        string   `mapstructure:"uri"`
	Filters    []string `mapstructure:"filters"`
	Predicates []string `mapstructure:"predicates"`
}
type RouteConfig struct {
	Routes []RouteDefinition `mapstructure:"routes"`
}

func GetRoutesBy(RoutesFile string) ([]*Route, error) {
	result := make([]*Route, 0)
	routeConfig := &RouteConfig{}
	err := config.LoadYaml(RoutesFile, &routeConfig)
	if err != nil {
		logrus.Panic("Load RoutesFile fail:", err)
		return nil, err
	}

	for _, routeDefn := range routeConfig.Routes {
		route := &Route{
			Id:  routeDefn.Id,
			Uri: routeDefn.Uri,
		}
		for _, v := range routeDefn.Predicates {
			predi, err := createPredi(v)
			if err != nil {
				logrus.Error(err.Error())
				continue
			}
			route.Predicates = append(route.Predicates, predi)
		}
		for _, v := range routeDefn.Filters {
			filter, err := createFilter(v)
			if err != nil {
				logrus.Error(err.Error())

				continue
			}
			route.Filters = append(route.Filters, filter)
		}
		result = append(result, route)
	}
	return result, nil
}
