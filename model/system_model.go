package model

// 系统配置
type SystemConfig struct {
	Aiconfig        []*AiConfig          `json:"ai_config"`
	ActiveAiConfigs map[string]*AiConfig `json:"active_ai_configs"` // 根据不同的使用场景可能会选择不同的ai配置
}

// AiConfig AI配置
type AiConfig struct {
	ID          string   `json:"id"` // 配置id 用于精准匹配数据
	ApiKey      string   `json:"api_key"`
	BaseURL     string   `json:"base_url"`
	Models      []string `json:"models"`       // 模型列表
	ActiveModel string   `json:"active_model"` // 当前使用的模型
}
