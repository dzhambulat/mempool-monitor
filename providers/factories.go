package providers

import (
	"mempool-monitor/providers/local"
)

func GetLocalGethProvider(rpcURL string) ITransactionObserver {
	return local.NewGethProvider(rpcURL)
}
