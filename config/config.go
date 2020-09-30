package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

// Config 解析后的配置对象
type Config struct {
	// 监听端口
	Listen string `toml:"listen"`
	// 数据库
	DB   string `toml:"database"`
	Mode string `toml:"mode"`
}

var (
	cfg  *Config
	once sync.Once
)

var configFile = "./config.toml"

// NewConfig 解析路径指定的 toml 并返回
func NewConfig(path string) *Config {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}

// SetConfigFile 设置配置文件路径，在第一次调用 DefaultConfig 之前有效
func SetConfigFile(filePath string) {
	configFile = filePath
}

// DefaultConfig 获取到默认配置对象
func DefaultConfig() *Config {
	once.Do(func() {
		cfg = NewConfig(configFile)
	})
	return cfg
}
