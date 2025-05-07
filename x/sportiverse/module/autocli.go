package sportiverse

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "sportiverse/api/sportiverse/sportiverse"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "PostAll",
					Use:       "list-post",
					Short:     "List all post",
				},
				{
					RpcMethod:      "Post",
					Use:            "show-post [id]",
					Short:          "Shows a post by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "CommentAll",
					Use:       "list-comment",
					Short:     "List all comment",
				},
				{
					RpcMethod:      "Comment",
					Use:            "show-comment [id]",
					Short:          "Shows a comment by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "SubscriptionAll",
					Use:       "list-subscription",
					Short:     "List all subscription",
				},
				{
					RpcMethod:      "Subscription",
					Use:            "show-subscription [id]",
					Short:          "Shows a subscription by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "LikeAll",
					Use:       "list-like",
					Short:     "List all like",
				},
				{
					RpcMethod:      "Like",
					Use:            "show-like [id]",
					Short:          "Shows a like by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePost",
					Use:            "create-post [title] [body] [timestamp]",
					Short:          "Create post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "body"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "UpdatePost",
					Use:            "update-post [id] [title] [body] [timestamp]",
					Short:          "Update post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "title"}, {ProtoField: "body"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "DeletePost",
					Use:            "delete-post [id]",
					Short:          "Delete post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateComment",
					Use:            "create-comment [postId] [body] [timestamp]",
					Short:          "Create comment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "postId"}, {ProtoField: "body"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "UpdateComment",
					Use:            "update-comment [id] [postId] [body] [timestamp]",
					Short:          "Update comment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "postId"}, {ProtoField: "body"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "DeleteComment",
					Use:            "delete-comment [id]",
					Short:          "Delete comment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateSubscription",
					Use:            "create-subscription [userId]",
					Short:          "Create subscription",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "userId"}},
				},
				{
					RpcMethod:      "DeleteSubscription",
					Use:            "delete-subscription [id]",
					Short:          "Delete subscription",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateLike",
					Use:            "create-like [postId]",
					Short:          "Create like",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "postId"}},
				},
				{
					RpcMethod:      "DeleteLike",
					Use:            "delete-like [id]",
					Short:          "Delete like",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
