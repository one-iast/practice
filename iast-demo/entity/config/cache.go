package config

import "fmt"

type Cache struct {
	Redis Redis `json:"redis"`
}
type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbNum    int    `json:"dbNum"`
}

func (r *Redis) DSN() string {
	return fmt.Sprintf("redis://%s:%s@%s:%s/%d", r.Username, r.Password, r.Host, r.Port, r.DbNum)
}
