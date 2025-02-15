package config

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	HttpPort       int    `json:"httpPort"`
	Secret         string `json:"secret"`
	AuthExpMin     uint   `json:"authExpMin"`
	AuthRefreshMin uint   `json:"authRefreshMin"`
}
