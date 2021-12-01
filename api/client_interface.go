package api

type ConnectClickHouse interface {
	Ping() (ok bool, err error)
	CreateTables() (err error)
	AddMetrics(request CollectRequest) (err error)
}
