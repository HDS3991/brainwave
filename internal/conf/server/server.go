package server

type Conf struct {
	System SystemConf `mapstructure:"system"`
	Log    LogConfig  `mapstructure:"log"`
}
