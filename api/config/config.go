package config

type ApiConfig struct {
	Name string   `mapstructure:"name"`
	Tags []string `mapstructure:"tags"`
}
type ServiceConfig struct {
	Edge string `mapstructure:"edge"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ServerConfig struct {
	Api        ApiConfig     `mapstructure:"api"`
	Service    ServiceConfig `mapstructure:"service"`
	MySQL      MySQLConfig   `mapstructure:"mysql"`
	Consul     ConsulConfig  `mapstructure:"consul"`
	CommonName string        `mapstructure:"common_name"`
	Debug      bool          `mapstructure:"debug"`
}
