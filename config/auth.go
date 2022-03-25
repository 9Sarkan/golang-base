package config

type AuthCreditional struct {
	HMACSecret  string `yaml:"hmacSecret"`
	RSASecret   string `yaml:"rsaSecret"`
	ECDSASecret string `yaml:"ecdsaSecret"`
}
