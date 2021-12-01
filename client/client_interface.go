package client

import "analytics/api"

type ConnectClickHouse interface {
	Ping() (ok bool, err error)
	CreateTables() (err error)
	AddMetrics(request api.CollectRequest) (err error)
}
