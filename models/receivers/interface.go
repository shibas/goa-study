package receiver

// Receiver 結果を受け取る
type Receiver interface {
	Receive(i interface{}) error
}
