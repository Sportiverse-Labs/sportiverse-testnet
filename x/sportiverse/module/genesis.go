package sportiverse

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"sportiverse/x/sportiverse/keeper"
	"sportiverse/x/sportiverse/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the post
	for _, elem := range genState.PostList {
		k.SetPost(ctx, elem)
	}

	// Set post count
	k.SetPostCount(ctx, genState.PostCount)
	// Set all the comment
	for _, elem := range genState.CommentList {
		k.SetComment(ctx, elem)
	}

	// Set comment count
	k.SetCommentCount(ctx, genState.CommentCount)
	// Set all the subscription
	for _, elem := range genState.SubscriptionList {
		k.SetSubscription(ctx, elem)
	}

	// Set subscription count
	k.SetSubscriptionCount(ctx, genState.SubscriptionCount)
	// Set all the like
	for _, elem := range genState.LikeList {
		k.SetLike(ctx, elem)
	}

	// Set like count
	k.SetLikeCount(ctx, genState.LikeCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PostList = k.GetAllPost(ctx)
	genesis.PostCount = k.GetPostCount(ctx)
	genesis.CommentList = k.GetAllComment(ctx)
	genesis.CommentCount = k.GetCommentCount(ctx)
	genesis.SubscriptionList = k.GetAllSubscription(ctx)
	genesis.SubscriptionCount = k.GetSubscriptionCount(ctx)
	genesis.LikeList = k.GetAllLike(ctx)
	genesis.LikeCount = k.GetLikeCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
