package providers

import "mempool-monitor/types"

type ITransactionObserver interface {
	Subscribe() error
	GetObserver() <-chan *types.Transaction
	Close()
}
