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

func createNLike(keeper keeper.Keeper, ctx context.Context, n int) []types.Like {
	items := make([]types.Like, n)
	for i := range items {
		items[i].Id = keeper.AppendLike(ctx, items[i])
	}
	return items
}

func TestLikeGet(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNLike(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLike(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLikeRemove(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNLike(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLike(ctx, item.Id)
		_, found := keeper.GetLike(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLikeGetAll(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNLike(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLike(ctx)),
	)
}

func TestLikeCount(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNLike(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLikeCount(ctx))
}
