package route

import (
	"ggate/pkg/config"

	"github.com/sirupsen/logrus"
)

type Route struct {
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

func NewRouteDefinition(text string) {

}

func GetRoutesBy(RoutesFile string) ([]Route, error) {
	result := make([]Route, 0)
	routeConfig := &RouteConfig{}
	err := config.LoadYaml(RoutesFile, &routeConfig)
	if err != nil {
		logrus.Panic("Load RoutesFile fail:", err)
		return nil, err
	}
	//TODO :根据routeConfig 转换成对应的[]Route
	return result, nil
}
