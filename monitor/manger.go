package monitor

import (
	"github.com/493labs/contract-monitor/adapters"
	"github.com/493labs/contract-monitor/config"
	"github.com/493labs/contract-monitor/logger"
	"sync"
)

type Manager struct {
	*sync.Mutex
	Monitors []*config.ContractMonitor
	errC     chan struct{}
}

func NewManager() (*Manager, error) {
	return &Manager{
		Monitors: make([]*config.ContractMonitor, 0),
		errC:     make(chan struct{}, 0),
	}, nil
}

func (m *Manager) RegisterMonitor(cm *config.ContractMonitor) {
	m.Monitors = append(m.Monitors, cm)
}

func (m *Manager) Start() {
	m.Lock()
	defer m.Unlock()

	log := logger.GetLogger(logger.ModuleMonitor)
	for _, cfg := range m.Monitors {
		go func() {
			// create client
			adapters.ClientFactory(cfg)
			monitorConf.Network

			//
		}()
	}

}

func (m *Manager) Stop() {
	m.Lock()
	defer m.Unlock()

}
