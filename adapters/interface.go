package adapters

type Adapter interface {
	// MonitorContract monitor a contract
	MonitorContract(contract string) (Events, error)
}


type Events interface {

}

type ReceiveChan chan interface{}