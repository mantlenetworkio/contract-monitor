package monitor

import "github.com/493labs/contract-monitor/adapters"

type Monitor struct {
	ID           string
	ChainAdapter adapters.Adapter
	Handler
	errC chan struct{}
}
