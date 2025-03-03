package config

type RegistryAuthConfig struct {
	APISecretKey string `env:"API_SECRET_KEY,required"`
}
