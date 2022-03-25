package config

type Service struct {
	GRPCServer HTTPServer `yaml:"grpcServer"`
	HTTPServer HTTPServer `yaml:"httpServer"`
}

type HTTPServer struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
