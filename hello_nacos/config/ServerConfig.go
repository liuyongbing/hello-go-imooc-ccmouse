package config

// User API 服务配置
type ServerConfig struct {
	Host        string        `mapstructure:"host" json:"host"`
	Port        int           `mapstructure:"port" json:"port"`
	Name        string        `mapstructure:"name" json:"name"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}
