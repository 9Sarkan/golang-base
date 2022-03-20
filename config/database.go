package config

type Database struct {
	MongoDB Mongo `yaml:"mongoDB"`
}

type Mongo struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}
