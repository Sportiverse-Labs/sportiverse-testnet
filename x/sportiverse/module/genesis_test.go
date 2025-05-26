package sportiverse_test

import (
	"testing"

	keepertest "sportiverse/testutil/keeper"
	"sportiverse/testutil/nullify"
	sportiverse "sportiverse/x/sportiverse/module"
	"sportiverse/x/sportiverse/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PostList: []types.Post{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PostCount: 2,
		CommentList: []types.Comment{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CommentCount: 2,
		SubscriptionList: []types.Subscription{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SubscriptionCount: 2,
		LikeList: []types.Like{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LikeCount: 2,
		AccountList: []types.Account{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AccountCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SportiverseKeeper(t)
	sportiverse.InitGenesis(ctx, k, genesisState)
	got := sportiverse.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	require.Equal(t, genesisState.PostCount, got.PostCount)
	require.ElementsMatch(t, genesisState.CommentList, got.CommentList)
	require.Equal(t, genesisState.CommentCount, got.CommentCount)
	require.ElementsMatch(t, genesisState.SubscriptionList, got.SubscriptionList)
	require.Equal(t, genesisState.SubscriptionCount, got.SubscriptionCount)
	require.ElementsMatch(t, genesisState.LikeList, got.LikeList)
	require.Equal(t, genesisState.LikeCount, got.LikeCount)
	require.ElementsMatch(t, genesisState.AccountList, got.AccountList)
	require.Equal(t, genesisState.AccountCount, got.AccountCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
