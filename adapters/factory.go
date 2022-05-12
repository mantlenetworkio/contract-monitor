package adapters

import (
	"fmt"
	"github.com/493labs/contract-monitor/adapters/eth_adapter"
	"github.com/493labs/contract-monitor/config"
)

func ClientFactory(cfg *config.ContractMonitor) (BLockChainClient, error) {
	switch cfg.Network {
	case string(ETH):
		return eth_adapter.NewListener(cfg.Methods)
	//case BSC:
	//	return nil, fmt.Errorf("not implied yet")
	default:
		return nil, fmt.Errorf("not implied yet")
	}
}
