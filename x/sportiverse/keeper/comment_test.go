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

func createNComment(keeper keeper.Keeper, ctx context.Context, n int) []types.Comment {
	items := make([]types.Comment, n)
	for i := range items {
		items[i].Id = keeper.AppendComment(ctx, items[i])
	}
	return items
}

func TestCommentGet(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNComment(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetComment(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCommentRemove(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNComment(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveComment(ctx, item.Id)
		_, found := keeper.GetComment(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCommentGetAll(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNComment(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllComment(ctx)),
	)
}

func TestCommentCount(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNComment(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCommentCount(ctx))
}
