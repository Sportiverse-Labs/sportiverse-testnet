package keeper_test

import (
	"context"
	"testing"

	keepertest "sportiverse/testutil/keeper"
	"sportiverse/testutil/nullify"
	"sportiverse/x/sportiverse/keeper"
	"sportiverse/x/sportiverse/types"

	"github.com/stretchr/testify/require"
)

func createNAccount(keeper keeper.Keeper, ctx context.Context, n int) []types.Account {
	items := make([]types.Account, n)
	for i := range items {
		items[i].Id = keeper.AppendAccount(ctx, items[i])
	}
	return items
}

func TestAccountGet(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAccount(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccount(ctx, item.Id)
		_, found := keeper.GetAccount(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccount(ctx)),
	)
}

func TestAccountCount(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAccountCount(ctx))
}
