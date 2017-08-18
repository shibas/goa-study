package config

import (
	"fmt"
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

// DatabaseEnv データベース接続設定を保持するstruct
type DatabaseEnv struct {
	Server   string
	Port     int
	User     string
	Pass     string
	Database string
}

// open DBコネクションを作成する
// エラーが発生した場合はnilを返す
func (d *DatabaseEnv) Conn() sqlbuilder.Database {
	host := fmt.Sprintf("%s:%d", d.Server, d.Port)
	setting := mysql.ConnectionURL{
		Database: d.Database,
		Host:     host,
		User:     d.User,
		Password: d.Pass,
		Options: map[string]string{
			"parseTime": "true",
			"loc":       "Asia/Tokyo",
		},
	}
	log.Printf("Open database connection Engine: mysql server: %s port: %d database: %s", d.Server, d.Port, d.Database)
	session, err := mysql.Open(setting)
	if err != nil {
		log.Println(err)
		return nil
	}
	return session
}
