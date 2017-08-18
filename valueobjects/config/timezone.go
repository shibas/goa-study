package config

type TimeConfig struct {
	TimeZone       string `mapstructure:"timezone"`
	TimeZoneOffset int    `mapstructure:"timezoneoffset"`
	DateTimeLayout string `mapstructure:"datetimelayout"`
}
