// +build !production,!staging,!test,!local

package env

import (
	"goa-study/valueobjects/config"
	"time"
)

func initialize() {
	OnDevelopment = true
	configfilename = "dev"
	TimeConfig = config.TimeConfig{
		TimeZone:       "Asia/Tokyo",
		TimeZoneOffset: 9 * 60 * 60,
		DateTimeLayout: "2006-01-02 15:04:05",
	}
	DataBaseTimeZone = time.FixedZone(TimeConfig.TimeZone, TimeConfig.TimeZoneOffset)
}
