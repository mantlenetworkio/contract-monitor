package eth_adapter

import (
	"github.com/493labs/contract-monitor/adapters"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthAdapter struct {
	*ethclient.Client
}

func (l *Listener) MonitorContract(contract string) (adapters.Events, error) {
	panic("implement me")
}

//func (l *Listener) SubscribeBlock() (adapters.ReceiveChan, error) {
//	panic("implement me")
//}

func NewListener(url string) (*Listener, error) {
	var client *ethclient.Client
	var err error
	if client, err = ethclient.Dial(url); err != nil {
		return nil, err
	}
	return &Listener{
		Client: client,
	}, nil
}
