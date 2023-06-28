package server

type SystemConf struct {
	SSL        string `mapstructure:"ssl"`
	Mode       string `mapstructure:"mode"`
	Host       string `mapstructure:"host"`
	HttpPort   string `mapstructure:"httpPort"`
	Version    string `mapstructure:"version"`
	EncryptKey string `mapstructure:"encryptKey"`
	LogPath    string `mapstructure:"logPath"`
}
