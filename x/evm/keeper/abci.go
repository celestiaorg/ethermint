package keeper

import (
	"encoding/json"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

// BeginBlock sets the sdk Context and EIP155 chain id to the Keeper.
func (k *Keeper) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	k.WithChainID(ctx)
}

// EndBlock also retrieves the bloom filter value from the transient store and commits it to the
// KVStore. The EVM end block logic doesn't update the validator set, thus it returns
// an empty slice.
func (k *Keeper) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	// Gas costs are handled within msg handler so costs should be ignored
	infCtx := ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	bloom := ethtypes.BytesToBloom(k.GetBlockBloomTransient(infCtx).Bytes())
	k.EmitBlockBloomEvent(infCtx, bloom)

	ethTxs := k.GetTxs(ctx)
	JSONtxs, err := json.MarshalIndent(ethTxs, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("End Block Txs: %s\n", string(JSONtxs))
	var txRoot common.Hash
	if len(ethTxs) == 0 {
		txRoot = ethtypes.EmptyRootHash
	} else {
		hasher := trie.NewStackTrie(nil)
		txRoot = ethtypes.DeriveSha(ethtypes.Transactions(ethTxs), hasher)
	}
	fmt.Printf("End Block TxRoot: %s\n", txRoot)
	k.EmitTxRootEvent(ctx, txRoot)

	return []abci.ValidatorUpdate{}
}
