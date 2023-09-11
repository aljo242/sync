package keeper

import (
	"github.com/aljo242/sync/x/sync/types"
)

var _ types.QueryServer = Keeper{}
