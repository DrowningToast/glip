package config

type AuthConfig struct {
	JWTSecret string `env:"JWT_SECRET,required"`

	JWTExpirationTime int    `env:"JWT_EXPIRATION_TIME,required"`
	JWTAdminApiSecret string `env:"JWT_ADMIN_API_SECRET,required"`
}
