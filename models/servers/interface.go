package server

// Result Serviceの結果を詰める
type Result struct {
	Data interface{}
}

// Server Serveの実行
type Server interface {
	// Serve
	Serve(i interface{}) error
}
