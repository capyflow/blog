package pkg

import (
	"github.com/BurntSushi/toml"
	scfg "github.com/capyflow/mediaStorage/config"
)

type Config struct {
	BlogPort *string      `toml:"blog_port" json:"blog_port"` // 博客服务端口
	Group    *string      `toml:"group" json:"group"`
	Server   *scfg.Server `toml:"server" json:"server"`
	BlogUser struct {
		Email    *string `toml:"email" json:"email"`
		Password *string `toml:"password" json:"password"`
	} `toml:"blog_user" json:"blog_user"`
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
