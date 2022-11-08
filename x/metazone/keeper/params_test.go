package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "metazone/testutil/keeper"
	"metazone/x/metazone/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MetazoneKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
