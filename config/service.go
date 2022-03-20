package config

type Service struct {
	HTTP HTTPServer `yaml:"http"`
}

type HTTPServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
