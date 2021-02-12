package model

type IConfigReader interface {
	Read() (*Config, error)
}

type Config struct {
	Port     string   `json:"port"`
	Database database `json:"database"`
}

type database struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}
