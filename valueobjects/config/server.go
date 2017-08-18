package config

import "fmt"

// ServerConfig サーバー接続情報を定義
type ServerConfig struct {
	URIScheme   string
	HostName    string
	PortNum     int
	DocHostName string
	DocPort     int
	ImageDir    string
}

// APIHostString APIホスト文字列（ドメイン・ポート）を返却
func (s *ServerConfig) APIHostString() string {
	if s.PortNum <= 0 {
		return s.HostName
	}
	return fmt.Sprintf("%s:%d", s.HostName, s.PortNum)
}

// APIHostStringOnDoc APIホスト文字列（ドメイン・ポート）を返却
func (s *ServerConfig) APIHostStringOnDoc() string {
	if s.DocPort <= 0 {
		return s.DocHostName
	}
	return fmt.Sprintf("%s:%d", s.DocHostName, s.DocPort)
}
