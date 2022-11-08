package metazone_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "metazone/testutil/keeper"
	"metazone/testutil/nullify"
	"metazone/x/metazone"
	"metazone/x/metazone/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MetazoneKeeper(t)
	metazone.InitGenesis(ctx, *k, genesisState)
	got := metazone.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
