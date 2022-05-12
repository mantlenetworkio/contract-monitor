package adapters

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TransactionFilter(items types.Transactions, fn func(*types.Transaction) bool) types.Transactions {
	var filtered types.Transactions

	for _, tx := range items {
		if fn(tx) {
			filtered = append(filtered, tx)
		}
	}

	return filtered
}

func IsContractCreation(tx *types.Transaction) bool {
	return tx.To() == nil
}

func GetTransactionFrom(tx *types.Transaction) (common.Address, error) {
	return types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
}
