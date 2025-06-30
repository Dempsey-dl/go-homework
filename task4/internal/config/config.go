package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
		Mode    string `yaml:"mode"`
	} `yaml:"server"`
	Database struct {
		DSN     string `yaml:"dsn"`
		Dialect string `yaml:"dialect"`
	} `yaml:"database"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func Load(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err) // 增加错误上下文
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("解析 YAML 失败: %w", err) // 增加错误上下文
	}

	return &cfg, nil
}
