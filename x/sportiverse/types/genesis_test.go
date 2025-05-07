package types_test

import (
	"testing"

	"sportiverse/x/sportiverse/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated post",
			genState: &types.GenesisState{
				PostList: []types.Post{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid post count",
			genState: &types.GenesisState{
				PostList: []types.Post{
					{
						Id: 1,
					},
				},
				PostCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated comment",
			genState: &types.GenesisState{
				CommentList: []types.Comment{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid comment count",
			genState: &types.GenesisState{
				CommentList: []types.Comment{
					{
						Id: 1,
					},
				},
				CommentCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated subscription",
			genState: &types.GenesisState{
				SubscriptionList: []types.Subscription{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid subscription count",
			genState: &types.GenesisState{
				SubscriptionList: []types.Subscription{
					{
						Id: 1,
					},
				},
				SubscriptionCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated like",
			genState: &types.GenesisState{
				LikeList: []types.Like{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid like count",
			genState: &types.GenesisState{
				LikeList: []types.Like{
					{
						Id: 1,
					},
				},
				LikeCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
