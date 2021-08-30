// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package txpool

import (
	"context"
	"github.com/ledgerwatch/erigon-lib/kv"
	"sync"
)

// Ensure, that PoolMock does implement Pool.
// If this is not the case, regenerate this file with moq.
var _ Pool = &PoolMock{}

// PoolMock is a mock implementation of Pool.
//
// 	func TestSomethingThatUsesPool(t *testing.T) {
//
// 		// make and configure a mocked Pool
// 		mockedPool := &PoolMock{
// 			AddNewGoodPeerFunc: func(peerID PeerID)  {
// 				panic("mock out the AddNewGoodPeer method")
// 			},
// 			GetRlpFunc: func(tx kv.Tx, hash []byte) ([]byte, error) {
// 				panic("mock out the GetRlp method")
// 			},
// 			IdHashKnownFunc: func(tx kv.Tx, hash []byte) (bool, error) {
// 				panic("mock out the IdHashKnown method")
// 			},
// 			OnNewBlockFunc: func(stateChanges map[string]senderInfo, unwindTxs TxSlots, minedTxs TxSlots, protocolBaseFee uint64, pendingBaseFee uint64, blockHeight uint64, blockHash [32]byte, senders *SendersCache) error {
// 				panic("mock out the OnNewBlock method")
// 			},
// 			OnNewTxsFunc: func(ctx context.Context, db kv.RoDB, newTxs TxSlots, senders *SendersCache) error {
// 				panic("mock out the OnNewTxs method")
// 			},
// 			StartedFunc: func() bool {
// 				panic("mock out the Started method")
// 			},
// 		}
//
// 		// use mockedPool in code that requires Pool
// 		// and then make assertions.
//
// 	}
type PoolMock struct {
	// AddNewGoodPeerFunc mocks the AddNewGoodPeer method.
	AddNewGoodPeerFunc func(peerID PeerID)

	// GetRlpFunc mocks the GetRlp method.
	GetRlpFunc func(tx kv.Tx, hash []byte) ([]byte, error)

	// IdHashKnownFunc mocks the IdHashKnown method.
	IdHashKnownFunc func(tx kv.Tx, hash []byte) (bool, error)

	// OnNewBlockFunc mocks the OnNewBlock method.
	OnNewBlockFunc func(stateChanges map[string]senderInfo, unwindTxs TxSlots, minedTxs TxSlots, protocolBaseFee uint64, pendingBaseFee uint64, blockHeight uint64, blockHash [32]byte, senders *SendersCache) error

	// OnNewTxsFunc mocks the OnNewTxs method.
	OnNewTxsFunc func(ctx context.Context, db kv.RoDB, newTxs TxSlots, senders *SendersCache) error

	// StartedFunc mocks the Started method.
	StartedFunc func() bool

	// calls tracks calls to the methods.
	calls struct {
		// AddNewGoodPeer holds details about calls to the AddNewGoodPeer method.
		AddNewGoodPeer []struct {
			// PeerID is the peerID argument value.
			PeerID PeerID
		}
		// GetRlp holds details about calls to the GetRlp method.
		GetRlp []struct {
			// Tx is the tx argument value.
			Tx kv.Tx
			// Hash is the hash argument value.
			Hash []byte
		}
		// IdHashKnown holds details about calls to the IdHashKnown method.
		IdHashKnown []struct {
			// Tx is the tx argument value.
			Tx kv.Tx
			// Hash is the hash argument value.
			Hash []byte
		}
		// OnNewBlock holds details about calls to the OnNewBlock method.
		OnNewBlock []struct {
			// StateChanges is the stateChanges argument value.
			StateChanges map[string]senderInfo
			// UnwindTxs is the unwindTxs argument value.
			UnwindTxs TxSlots
			// MinedTxs is the minedTxs argument value.
			MinedTxs TxSlots
			// ProtocolBaseFee is the protocolBaseFee argument value.
			ProtocolBaseFee uint64
			// PendingBaseFee is the pendingBaseFee argument value.
			PendingBaseFee uint64
			// BlockHeight is the blockHeight argument value.
			BlockHeight uint64
			// BlockHash is the blockHash argument value.
			BlockHash [32]byte
			// Senders is the senders argument value.
			Senders *SendersCache
		}
		// OnNewTxs holds details about calls to the OnNewTxs method.
		OnNewTxs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db kv.RoDB
			// NewTxs is the newTxs argument value.
			NewTxs TxSlots
			// Senders is the senders argument value.
			Senders *SendersCache
		}
		// Started holds details about calls to the Started method.
		Started []struct {
		}
	}
	lockAddNewGoodPeer sync.RWMutex
	lockGetRlp         sync.RWMutex
	lockIdHashKnown    sync.RWMutex
	lockOnNewBlock     sync.RWMutex
	lockOnNewTxs       sync.RWMutex
	lockStarted        sync.RWMutex
}

// AddNewGoodPeer calls AddNewGoodPeerFunc.
func (mock *PoolMock) AddNewGoodPeer(peerID PeerID) {
	callInfo := struct {
		PeerID PeerID
	}{
		PeerID: peerID,
	}
	mock.lockAddNewGoodPeer.Lock()
	mock.calls.AddNewGoodPeer = append(mock.calls.AddNewGoodPeer, callInfo)
	mock.lockAddNewGoodPeer.Unlock()
	if mock.AddNewGoodPeerFunc == nil {
		return
	}
	mock.AddNewGoodPeerFunc(peerID)
}

// AddNewGoodPeerCalls gets all the calls that were made to AddNewGoodPeer.
// Check the length with:
//     len(mockedPool.AddNewGoodPeerCalls())
func (mock *PoolMock) AddNewGoodPeerCalls() []struct {
	PeerID PeerID
} {
	var calls []struct {
		PeerID PeerID
	}
	mock.lockAddNewGoodPeer.RLock()
	calls = mock.calls.AddNewGoodPeer
	mock.lockAddNewGoodPeer.RUnlock()
	return calls
}

// GetRlp calls GetRlpFunc.
func (mock *PoolMock) GetRlp(tx kv.Tx, hash []byte) ([]byte, error) {
	callInfo := struct {
		Tx   kv.Tx
		Hash []byte
	}{
		Tx:   tx,
		Hash: hash,
	}
	mock.lockGetRlp.Lock()
	mock.calls.GetRlp = append(mock.calls.GetRlp, callInfo)
	mock.lockGetRlp.Unlock()
	if mock.GetRlpFunc == nil {
		var (
			bytesOut []byte
			errOut   error
		)
		return bytesOut, errOut
	}
	return mock.GetRlpFunc(tx, hash)
}

// GetRlpCalls gets all the calls that were made to GetRlp.
// Check the length with:
//     len(mockedPool.GetRlpCalls())
func (mock *PoolMock) GetRlpCalls() []struct {
	Tx   kv.Tx
	Hash []byte
} {
	var calls []struct {
		Tx   kv.Tx
		Hash []byte
	}
	mock.lockGetRlp.RLock()
	calls = mock.calls.GetRlp
	mock.lockGetRlp.RUnlock()
	return calls
}

// IdHashKnown calls IdHashKnownFunc.
func (mock *PoolMock) IdHashKnown(tx kv.Tx, hash []byte) (bool, error) {
	callInfo := struct {
		Tx   kv.Tx
		Hash []byte
	}{
		Tx:   tx,
		Hash: hash,
	}
	mock.lockIdHashKnown.Lock()
	mock.calls.IdHashKnown = append(mock.calls.IdHashKnown, callInfo)
	mock.lockIdHashKnown.Unlock()
	if mock.IdHashKnownFunc == nil {
		var (
			bOut   bool
			errOut error
		)
		return bOut, errOut
	}
	return mock.IdHashKnownFunc(tx, hash)
}

// IdHashKnownCalls gets all the calls that were made to IdHashKnown.
// Check the length with:
//     len(mockedPool.IdHashKnownCalls())
func (mock *PoolMock) IdHashKnownCalls() []struct {
	Tx   kv.Tx
	Hash []byte
} {
	var calls []struct {
		Tx   kv.Tx
		Hash []byte
	}
	mock.lockIdHashKnown.RLock()
	calls = mock.calls.IdHashKnown
	mock.lockIdHashKnown.RUnlock()
	return calls
}

// OnNewBlock calls OnNewBlockFunc.
func (mock *PoolMock) OnNewBlock(stateChanges map[string]senderInfo, unwindTxs TxSlots, minedTxs TxSlots, protocolBaseFee uint64, pendingBaseFee uint64, blockHeight uint64, blockHash [32]byte, senders *SendersCache) error {
	callInfo := struct {
		StateChanges    map[string]senderInfo
		UnwindTxs       TxSlots
		MinedTxs        TxSlots
		ProtocolBaseFee uint64
		PendingBaseFee  uint64
		BlockHeight     uint64
		BlockHash       [32]byte
		Senders         *SendersCache
	}{
		StateChanges:    stateChanges,
		UnwindTxs:       unwindTxs,
		MinedTxs:        minedTxs,
		ProtocolBaseFee: protocolBaseFee,
		PendingBaseFee:  pendingBaseFee,
		BlockHeight:     blockHeight,
		BlockHash:       blockHash,
		Senders:         senders,
	}
	mock.lockOnNewBlock.Lock()
	mock.calls.OnNewBlock = append(mock.calls.OnNewBlock, callInfo)
	mock.lockOnNewBlock.Unlock()
	if mock.OnNewBlockFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.OnNewBlockFunc(stateChanges, unwindTxs, minedTxs, protocolBaseFee, pendingBaseFee, blockHeight, blockHash, senders)
}

// OnNewBlockCalls gets all the calls that were made to OnNewBlock.
// Check the length with:
//     len(mockedPool.OnNewBlockCalls())
func (mock *PoolMock) OnNewBlockCalls() []struct {
	StateChanges    map[string]senderInfo
	UnwindTxs       TxSlots
	MinedTxs        TxSlots
	ProtocolBaseFee uint64
	PendingBaseFee  uint64
	BlockHeight     uint64
	BlockHash       [32]byte
	Senders         *SendersCache
} {
	var calls []struct {
		StateChanges    map[string]senderInfo
		UnwindTxs       TxSlots
		MinedTxs        TxSlots
		ProtocolBaseFee uint64
		PendingBaseFee  uint64
		BlockHeight     uint64
		BlockHash       [32]byte
		Senders         *SendersCache
	}
	mock.lockOnNewBlock.RLock()
	calls = mock.calls.OnNewBlock
	mock.lockOnNewBlock.RUnlock()
	return calls
}

// OnNewTxs calls OnNewTxsFunc.
func (mock *PoolMock) OnNewTxs(ctx context.Context, db kv.RoDB, newTxs TxSlots) error {
	callInfo := struct {
		Ctx     context.Context
		Db      kv.RoDB
		NewTxs  TxSlots
		Senders *SendersCache
	}{
		Ctx:     ctx,
		Db:      db,
		NewTxs:  newTxs,
		Senders: senders,
	}
	mock.lockOnNewTxs.Lock()
	mock.calls.OnNewTxs = append(mock.calls.OnNewTxs, callInfo)
	mock.lockOnNewTxs.Unlock()
	if mock.OnNewTxsFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.OnNewTxsFunc(ctx, db, newTxs, senders)
}

// OnNewTxsCalls gets all the calls that were made to OnNewTxs.
// Check the length with:
//     len(mockedPool.OnNewTxsCalls())
func (mock *PoolMock) OnNewTxsCalls() []struct {
	Ctx     context.Context
	Db      kv.RoDB
	NewTxs  TxSlots
	Senders *SendersCache
} {
	var calls []struct {
		Ctx     context.Context
		Db      kv.RoDB
		NewTxs  TxSlots
		Senders *SendersCache
	}
	mock.lockOnNewTxs.RLock()
	calls = mock.calls.OnNewTxs
	mock.lockOnNewTxs.RUnlock()
	return calls
}

// Started calls StartedFunc.
func (mock *PoolMock) Started() bool {
	callInfo := struct {
	}{}
	mock.lockStarted.Lock()
	mock.calls.Started = append(mock.calls.Started, callInfo)
	mock.lockStarted.Unlock()
	if mock.StartedFunc == nil {
		var (
			bOut bool
		)
		return bOut
	}
	return mock.StartedFunc()
}

// StartedCalls gets all the calls that were made to Started.
// Check the length with:
//     len(mockedPool.StartedCalls())
func (mock *PoolMock) StartedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStarted.RLock()
	calls = mock.calls.Started
	mock.lockStarted.RUnlock()
	return calls
}
