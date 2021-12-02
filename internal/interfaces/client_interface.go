package interfaces

import (
	"analytics/internal/model"
)

type ConnectClickHouse interface {
	Ping() (ok bool, err error)
	CreateTables() (err error)
	AddMetrics(request model.CollectRequest) (err error)
}
