package config

// HTTPClientConfig http.Clientの接続に関するConfig
type HTTPClientConfig struct {
	DefaultTimeout    int
	DefaultRetryLimit int
}
