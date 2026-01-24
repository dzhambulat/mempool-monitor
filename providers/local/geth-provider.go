package local

import (
	"context"
	core_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"mempool-monitor/types"
	"os"
)

type GethProvider struct {
	observer chan *types.Transaction
	rpcURL   string
	client   *ethclient.Client
}


func NewGethProvider(rpcURL string) *GethProvider {
	return &GethProvider{
		rpcURL: rpcURL,
	}
}

func (p *GethProvider) Subscribe() error {
    p.observer = make(chan *types.Transaction, 100)
    nodeUri := os.Getenv("NODE_URI")
    client, err := ethclient.Dial("wss://" + nodeUri)
    if err != nil {
        return err
    }
    p.client = client

    headers := make(chan *core_types.Header)
    sub, err := p.client.SubscribeNewHead(context.Background(), headers)
    if err != nil {
        return err
    }

    go func() {
        defer sub.Unsubscribe()
        defer close(p.observer)
        for {
            select {
            case err := <-sub.Err():
                log.Println(err)
                return
            case header := <-headers:
                block, err := p.client.BlockByHash(context.Background(), header.Hash())
                if err != nil {
                    log.Println(err)
                    continue
                }

                for _, tx := range block.Transactions() {
                    from, err := core_types.Sender(core_types.LatestSignerForChainID(tx.ChainId()), tx)
                    if err != nil {
                        log.Println(err)
                        continue
                    }
                    var to []byte
                    if tx.To() != nil {
                        to = tx.To().Bytes()
                    }

                    p.observer <- &types.Transaction{
                        Hash:     tx.Hash().Bytes(),
                        From:     from.Bytes(),
                        To:       to,
                        CallData: tx.Data(),
                    }
                }
            }
        }
    }()

    return nil
}


func (p *GethProvider) GetObserver() <-chan *types.Transaction {
	return p.observer
}

func (p *GethProvider) Close() {
	close(p.observer)
}
