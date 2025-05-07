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

func createNSubscription(keeper keeper.Keeper, ctx context.Context, n int) []types.Subscription {
	items := make([]types.Subscription, n)
	for i := range items {
		items[i].Id = keeper.AppendSubscription(ctx, items[i])
	}
	return items
}

func TestSubscriptionGet(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNSubscription(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSubscription(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSubscriptionRemove(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNSubscription(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSubscription(ctx, item.Id)
		_, found := keeper.GetSubscription(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSubscriptionGetAll(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNSubscription(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSubscription(ctx)),
	)
}

func TestSubscriptionCount(t *testing.T) {
	keeper, ctx := keepertest.SportiverseKeeper(t)
	items := createNSubscription(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSubscriptionCount(ctx))
}
