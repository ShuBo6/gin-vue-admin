package config

type Feishu struct {
	Url   string `mapstructure:"url" json:"url" yaml:"url"`       // Url
	Token string `mapstructure:"token" json:"token" yaml:"token"` // token
}
