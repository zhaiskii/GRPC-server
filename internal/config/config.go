package config

type ConfigDatabase struct {
	Port_grpc string `yaml:"port_grpc" env:"GRPC_PORT" env-default:"50051"`
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port_http string `yaml:"port_http" env:"HTTP_PORT" env-default:"8082"`
}

func New() *ConfigDatabase{
	return &ConfigDatabase{}
}