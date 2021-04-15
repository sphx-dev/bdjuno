package distribution

import (
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"google.golang.org/grpc"

	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/types"
	"github.com/go-co-op/gocron"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/forbole/bdjuno/database"
)

var _ modules.Module = &Module{}

// Module represents the x/distr module
type Module struct {
	db          *database.BigDipperDb
	distrClient distrtypes.QueryClient
}

// NewModule returns a new Module instance
func NewModule(grpConnection *grpc.ClientConn, db *database.BigDipperDb) *Module {
	return &Module{
		distrClient: distrtypes.NewQueryClient(grpConnection),
		db:          db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "distribution"
}

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	return RegisterPeriodicOps(scheduler, m.distrClient, m.db)
}

// HandleBlock implements modules.Module
func (m *Module) HandleBlock(b *tmctypes.ResultBlock, _ []*types.Tx, vals *tmctypes.ResultValidators) error {
	return HandleBlock(b, m.distrClient, m.db)
}