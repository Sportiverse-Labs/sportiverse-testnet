package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "sportiverse/testutil/keeper"
	"sportiverse/x/sportiverse/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SportiverseKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
