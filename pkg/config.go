package pkg

import (
	"github.com/BurntSushi/toml"
	"github.com/capyflow/Allspark-go/ds"
	scfg "github.com/capyflow/mediaStorage/config"
)

// 配置文件
type Config struct {
	Port           *string      `json:"port" toml:"port"`         // 端口
	DatabaseConfig *ds.DsConfig `json:"database" toml:"database"` // 数据库配置
	S3             *scfg.S3     `json:"s3" toml:"s3"`             // S3 配置
	Jwt            *scfg.Jwt    `json:"jwt" toml:"jwt"`           // JWT 配置
}

// LoadConfig 从指定路径加载TOML配置文件
func LoadConfig(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		// 打印错误信息
		return nil, err
	}
	return &cfg, nil
}
