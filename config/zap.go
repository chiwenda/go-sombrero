package config

type Zap struct {
	Level  string `mapstructure:"level" json:"level" yml:"level"`
	Prefix string `mapstructure:"prefix" json:"prefix" yml:"prefix"`
}
