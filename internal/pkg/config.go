package pkg

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

var (
	config Config
)

// Config config
type Config struct {
	Websites Websites   `toml:"websites"`
	Mail     MailConfig `toml:"mail"`
}

// LoadFile 加载文件
func LoadFile(filename string) error {
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetConfig 获取配置
func GetConfig() *Config {
	return &config
}
