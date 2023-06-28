package server

type LogConfig struct {
	Level     string `mapstructure:"level"`
	TimeZone  string `mapstructure:"timeZone"`
	LogName   string `mapstructure:"logName"`
	LogSuffix string `mapstructure:"logSuffix"`
	MaxBackup int    `mapstructure:"maxBackup"`
}
