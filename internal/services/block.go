package services

import (
	"context"

	"github.com/NethermindEth/juno/internal/config"
	"github.com/NethermindEth/juno/internal/db"
	"github.com/NethermindEth/juno/internal/db/block"
	"github.com/NethermindEth/juno/internal/log"
)

// BlockService is a service to manage the block database. Before
// using the service, it must be configured with the Setup method;
// otherwise, the value will be the default. To stop the service, call the
// Close method.
var BlockService blockService

type blockService struct {
	service
	manager *block.Manager
}

// Setup is used to configure the service before it's started. The database
// param is the database where the transactions will be stored.
func (s *blockService) Setup(database db.Databaser) {
	if s.service.Running() {
		// notest
		s.logger.Panic("trying to Setup with service running")
	}
	s.manager = block.NewManager(database)
}

// Run starts the service. If the Setup method is not called before, the default
// values are used.
func (s *blockService) Run() error {
	if s.logger == nil {
		s.logger = log.Default.Named("Block Service")
	}

	if err := s.service.Run(); err != nil {
		// notest
		return err
	}

	s.setDefaults()
	return nil
}

func (s *blockService) setDefaults() {
	if s.manager == nil {
		// notest
		database := db.NewKeyValueDb(config.DataDir+"/block", 0)
		s.manager = block.NewManager(database)
	}
}

// Close stops the service, waiting to end the current operations, and closes
// the database manager.
func (s *blockService) Close(ctx context.Context) {
	s.service.Close(ctx)
	s.manager.Close()
}

// GetBlockByHash searches for the block associated with the given block hash.
// If the block does not exist on the database, then returns nil.
func (s *blockService) GetBlockByHash(blockHash []byte) *block.Block {
	s.AddProcess()
	defer s.DoneProcess()

	s.logger.
		With("blockHash", blockHash).
		Debug("GetBlockByHash")

	return s.manager.GetBlockByHash(blockHash)
}

// GetBlockByNumber searches for the block associated with the given block
// number. If the block does not exist on the database, then returns nil.
func (s *blockService) GetBlockByNumber(blockNumber uint64) *block.Block {
	s.AddProcess()
	defer s.DoneProcess()

	s.logger.
		With("blockNumber", blockNumber).
		Debug("GetBlockByNumber")

	return s.manager.GetBlockByNumber(blockNumber)
}

// StoreBlock stores the given block into the database. The key used to map the
// block it's the hash of the block. If the database already has a block with
// the same key, then the value is overwritten.
func (s *blockService) StoreBlock(blockHash []byte, block *block.Block) {
	s.AddProcess()
	defer s.DoneProcess()

	s.logger.
		With("blockHash", blockHash).
		Debug("StoreBlock")

	s.manager.PutBlock(blockHash, block)
}
