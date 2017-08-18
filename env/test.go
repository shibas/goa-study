// +build test

package env

import (
	"time"
	"goa-study/valueobjects/config"
)

func initialize() {
	OnDevelopment = true
	configfilename = "test"
	TimeConfig = config.TimeConfig{
		TimeZone:       "Asia/Tokyo",
		TimeZoneOffset: 9 * 60 * 60,
		DateTimeLayout: "2006-01-02 15:04:05",
	}
	DataBaseTimeZone = time.FixedZone(TimeConfig.TimeZone, TimeConfig.TimeZoneOffset)
}
