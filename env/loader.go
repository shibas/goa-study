package env

// 環境設定を保持するパッケージ

import (
	"bytes"
	"fmt"
	"log"
	"time"
	"goa-study/assets"
	"goa-study/valueobjects/config"

	"upper.io/db.v3/lib/sqlbuilder"

	"github.com/deadcheat/goacors"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// 環境設定として保持する情報の構造体
var (
	// .goで定義
	OnDevelopment    = false
	configfilename   string
	DataBaseTimeZone *time.Location
	TimeConfig       config.TimeConfig

	// tomlで設定
	Server           *config.ServerConfig
	HTTPClientConfig *config.HTTPClientConfig
	CorsConf         goacors.GoaCORSConfig
	Externals        *config.ExternalsConfig

	// DB
	writeDB            *config.DatabaseEnv
	readDB             *config.DatabaseEnv
	Connection         sqlbuilder.Database
	ReadOnlyConnection sqlbuilder.Database

	// RabbitMQ
	amqpEnv        *config.AMQPEnv
	AMQPConnection *amqp.Connection
)

// init パッケージの初期化処理
func init() {
	// 初期設定
	initialize()
	InitializeEnv()
}

const (
	defaultTimeout    = 1
	defaultRetryLimit = 3
)

// InitializeEnv パッケージの初期化処理
func InitializeEnv() {
	// 設定ファイル
	confDir := "config"
	fileName := configfilename
	ext := "toml"
	filePath := fmt.Sprintf("%s/%s.%s", confDir, fileName, ext)

	// Asset経由でファイルを取得しに行く
	viper.SetConfigType(ext)
	data, err := assets.Asset(filePath)
	if err != nil {
		log.Panic(err, filePath)
	}

	// go-bindataのAsset、およびconfigファイルのbindataが存在する場合
	viper.ReadConfig(bytes.NewBuffer(data))
	// 設定読み込み
	_ = viper.UnmarshalKey("server", &Server)
	_ = viper.UnmarshalKey("httpclient", &HTTPClientConfig)
	_ = viper.UnmarshalKey("cors", &CorsConf)
	_ = viper.UnmarshalKey("writedb", &writeDB)
	_ = viper.UnmarshalKey("readdb", &readDB)
	_ = viper.UnmarshalKey("externals", &Externals)
	_ = viper.UnmarshalKey("amqpEnv", &amqpEnv)

	if HTTPClientConfig == nil {
		HTTPClientConfig = &config.HTTPClientConfig{
			DefaultTimeout:    defaultTimeout,
			DefaultRetryLimit: defaultRetryLimit,
		}
	}
}

// IgniteDBConnection DBコネクションをスタートする
func IgniteDBConnection() {

	// DBをスタート
	if writeDB != nil {
		Connection = writeDB.Conn()
		Connection.SetLogging(OnDevelopment)
	}
	if readDB != nil {
		ReadOnlyConnection = readDB.Conn()
		ReadOnlyConnection.SetLogging(OnDevelopment)
	}
}

// IgniteAMQPConnection RabbitMQコネクションをスタートする
func IgniteAMQPConnection() {
	if amqpEnv != nil {
		AMQPConnection = amqpEnv.Conn()
	}
}
