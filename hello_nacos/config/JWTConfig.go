package config

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
