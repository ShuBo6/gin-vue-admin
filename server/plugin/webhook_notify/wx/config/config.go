package config

type WX struct {
	Url string `mapstructure:"url" json:"url" yaml:"url"` // Url
	Key string `mapstructure:"key" json:"key" yaml:"key"` // key
}
