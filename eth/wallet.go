package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewWallet creates are wallet for the given params.
func NewWallet(rpc string, chainId *big.Int, privateKey string) (*Wallet, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	key, from, err := ParseHexKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &Wallet{client, chainId, key, from}, nil
}

//  ParseHexKey converts a hex private key to an ecdsa private key.
func ParseHexKey(hexkey string) (*ecdsa.PrivateKey, common.Address, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(hexkey, "0x"))
	if err != nil {
		return nil, common.Address{}, errors.New("failed to parse private key")
	}
	pubkey, ok := (key.Public()).(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, errors.New("failed to private key to public key")
	}
	return key, crypto.PubkeyToAddress(*pubkey), nil
}

// Wallet is wrappers ethclient and private key.
type Wallet struct {
	Client     *ethclient.Client
	ChainId    *big.Int
	PrivateKey *ecdsa.PrivateKey
	From       common.Address
}

// Returns the current nonce value.
func (w *Wallet) GetNonce(ctx context.Context) (uint64, error) {
	nonce, err := w.Client.PendingNonceAt(ctx, w.From)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// Returns transaction authorization data.
func (w *Wallet) GetTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	signer := types.LatestSignerForChainID(w.ChainId)
	return &bind.TransactOpts{
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, signer, w.PrivateKey)
		},
		From:    w.From,
		Context: ctx,
		// GasLimit: uint64(1_000_000),
		// GasPrice: big.NewInt(1_000_000_000),,
	}, nil
}

// Returns options to fine tune a contract call request.
func (w *Wallet) GetCallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{From: w.From, Context: ctx}
}

// Wait for transaction to be completed.
// Note: Will wait forever, should cancel.
func (w *Wallet) WaitForTransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := w.Client.TransactionReceipt(ctx, hash)
		if err != nil {
			if err.Error() != "not found" {
				return nil, err
			}
			time.Sleep(1 * time.Second)
			continue
		}
		if receipt.Status != 0 {
			return receipt, nil
		}
		time.Sleep(1 * time.Second)
	}
}
