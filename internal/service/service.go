package service

import "ggate/pkg/config"

type Service struct {
}
type ServiceDefinition struct {
	Id  string `mapstructure:"id"`
	Uri string `mapstructure:"uri"`
	Qps int    `mapstructure:"qps"`

	Backends []Backend `mapstructure:"backends"`
}
type Backend struct {
	Id  string `mapstructure:"id"`
	Uri string `mapstructure:"uri"`
	Qps int    `mapstructure:"qps"`

	Weight int `mapstructure:"weight"`
}

type yamlServiceDefinitions struct {
	Services []ServiceDefinition `mapstructure:"services"`
}

func GetAllServicesBy(ServicesFile string) (map[string]Service, error) {
	result := make(map[string]Service, 0)
	yml := &yamlServiceDefinitions{}
	err := config.LoadYaml(ServicesFile, &yml)
	if err != nil {
		return nil, err
	}
	//TODO: tml转换成result
	return result, nil
}
