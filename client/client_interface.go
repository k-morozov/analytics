package client

type ConnectClickHouse interface {
	Ping() (ok bool, err error)
	Send() bool
}
