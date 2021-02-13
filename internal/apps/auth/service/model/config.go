package model

type Database struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}
