package config

import (
	"github.com/spf13/viper"
	"log"
)

var CFG = &Config{}

// Config 配置文件结构体
type Config struct {
	Stage     Stage
	Cache     Cache
	Sqlmap    Sqlmap
	Burpsuite Burpsuite
}
type Stage struct {
	PacketSize int
	PythonPath string
}
type Sqlmap struct {
	Enable  bool
	Thread  int
	Path    string
	ShowLog bool
}

type Burpsuite struct {
	Enable          bool
	SingleScanCount int
	ExeCmd          string
	ShowLog         bool
}

func LoadConfig(path string) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	err = viper.Unmarshal(&CFG)
	return err
}
