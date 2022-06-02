package rpc

import (
	"github.com/NethermindEth/juno/internal/db/block"
	"github.com/NethermindEth/juno/internal/db/transaction"
	"github.com/NethermindEth/juno/pkg/common"
)

func (x *BlockResponse) fromDatabaseBlock(y *block.Block) {
	x.BlockHash = common.BytesToFelt(y.Hash)
	x.ParentHash = common.BytesToFelt(y.ParentBlockHash)
	x.BlockNumber = y.BlockNumber
	x.Status = BlockStatus(y.Status)
	x.Sequencer = common.BytesToFelt(y.SequencerAddress)
	x.NewRoot = common.BytesToFelt(y.GlobalStateRoot)
	x.OldRoot = common.BytesToFelt(y.OldRoot)
	x.AcceptedTime = uint64(y.AcceptedTime)
}

func (x *Txn) fromDatabase(y *transaction.Transaction) {
	x.TxnHash = common.BytesToFelt(y.Hash)

	if tx := y.GetInvoke(); tx != nil {
		x.CallData = make([]common.Felt, len(tx.CallData))
		for i, data := range tx.CallData {
			x.CallData[i] = common.BytesToFelt(data)
		}
		x.EntryPointSelector = common.BytesToFelt(tx.EntryPointSelector)
		x.ContractAddress = common.BytesToFelt(tx.ContractAddress)
		x.MaxFee = common.BytesToFelt(tx.MaxFee)
		return
	}
}

func (x *TxnReceipt) fromDatabase(y *transaction.TransactionReceipt) {
	x.TxnHash = common.BytesToFelt(y.TxHash)
	x.Status = y.Status.String()
	x.StatusData = y.StatusData
	x.MessagesSent = make([]MsgToL1, len(y.MessagesSent))
	for i, msg := range y.MessagesSent {
		x.MessagesSent[i].fromDatabase(msg)
	}
	x.L1OriginMessage.fromDatabase(y.L1OriginMessage)
	x.Events = make([]Event, len(y.Events))
	for i, e := range y.Events {
		x.Events[i].fromDatabase(e)
	}
}

func (x *MsgToL1) fromDatabase(y *transaction.MessageToL1) {
	x.ToAddress = common.BytesToFelt(y.ToAddress)
	x.Payload = make([]common.Felt, len(y.Payload))
	for i, p := range y.Payload {
		x.Payload[i] = common.BytesToFelt(p)
	}
}

func (x *MsgToL2) fromDatabase(y *transaction.MessageToL2) {
	x.FromAddress = EthAddress(y.FromAddress)
	x.Payload = make([]common.Felt, len(y.Payload))
	for i, p := range y.Payload {
		x.Payload[i] = common.BytesToFelt(p)
	}
}

func (x *Event) fromDatabase(y *transaction.Event) {
	x.FromAddress = common.BytesToFelt(y.FromAddress)
	x.Keys = make([]common.Felt, len(y.Keys))
	for i, k := range y.Keys {
		x.Keys[i] = common.BytesToFelt(k)
	}
	x.Data = make([]common.Felt, len(y.Data))
	for i, d := range y.Data {
		x.Data[i] = common.BytesToFelt(d)
	}
}
