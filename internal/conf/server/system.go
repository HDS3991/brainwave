package server

type SystemConf struct {
	SSL        string `mapstructure:"ssl"`
	Mode       string `mapstructure:"mode"`
	Host       string `mapstructure:"host"`
	HttpPort   string `mapstructure:"httpPort"`
	Version    string `mapstructure:"version"`
	EncryptKey string `mapstructure:"encryptKey"`
	LogPath    string `mapstructure:"logPath"`
	DB         DBConf `mapstructure:"db"`
}

type DBConf struct {
	Path              string `mapstructure:"path"`
	FileName          string `mapstructure:"fileName"`
	MaxIdleConns      int    `mapstructure:"maxIdleConns"`
	MaxOpenConns      int    `mapstructure:"maxOpenConns"`
	MaxLifeTimeSecond int    `mapstructure:"maxLifeTimeSecond"`
}
