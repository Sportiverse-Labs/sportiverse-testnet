package sportiverse

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"sportiverse/testutil/sample"
	sportiversesimulation "sportiverse/x/sportiverse/simulation"
	"sportiverse/x/sportiverse/types"
)

// avoid unused import issue
var (
	_ = sportiversesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePost int = 100

	opWeightMsgUpdatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePost int = 100

	opWeightMsgDeletePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePost int = 100

	opWeightMsgCreateComment = "op_weight_msg_comment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateComment int = 100

	opWeightMsgUpdateComment = "op_weight_msg_comment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateComment int = 100

	opWeightMsgDeleteComment = "op_weight_msg_comment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteComment int = 100

	opWeightMsgCreateSubscription = "op_weight_msg_subscription"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSubscription int = 100

	opWeightMsgDeleteSubscription = "op_weight_msg_subscription"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSubscription int = 100

	opWeightMsgCreateLike = "op_weight_msg_like"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLike int = 100

	opWeightMsgDeleteLike = "op_weight_msg_like"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteLike int = 100

	opWeightMsgCreateAccount = "op_weight_msg_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAccount int = 100

	opWeightMsgDeleteAccount = "op_weight_msg_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAccount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	sportiverseGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PostList: []types.Post{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PostCount: 2,
		CommentList: []types.Comment{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CommentCount: 2,
		SubscriptionList: []types.Subscription{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SubscriptionCount: 2,
		LikeList: []types.Like{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		LikeCount: 2,
		AccountList: []types.Account{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AccountCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&sportiverseGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		sportiversesimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePost int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePost, &weightMsgUpdatePost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePost = defaultWeightMsgUpdatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePost,
		sportiversesimulation.SimulateMsgUpdatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePost int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = defaultWeightMsgDeletePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePost,
		sportiversesimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateComment int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateComment, &weightMsgCreateComment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateComment = defaultWeightMsgCreateComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateComment,
		sportiversesimulation.SimulateMsgCreateComment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateComment int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateComment, &weightMsgUpdateComment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateComment = defaultWeightMsgUpdateComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateComment,
		sportiversesimulation.SimulateMsgUpdateComment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteComment int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteComment, &weightMsgDeleteComment, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteComment = defaultWeightMsgDeleteComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteComment,
		sportiversesimulation.SimulateMsgDeleteComment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSubscription int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSubscription, &weightMsgCreateSubscription, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSubscription = defaultWeightMsgCreateSubscription
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSubscription,
		sportiversesimulation.SimulateMsgCreateSubscription(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSubscription int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSubscription, &weightMsgDeleteSubscription, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSubscription = defaultWeightMsgDeleteSubscription
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSubscription,
		sportiversesimulation.SimulateMsgDeleteSubscription(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateLike int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateLike, &weightMsgCreateLike, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLike = defaultWeightMsgCreateLike
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLike,
		sportiversesimulation.SimulateMsgCreateLike(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteLike int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteLike, &weightMsgDeleteLike, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLike = defaultWeightMsgDeleteLike
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLike,
		sportiversesimulation.SimulateMsgDeleteLike(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateAccount, &weightMsgCreateAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAccount = defaultWeightMsgCreateAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAccount,
		sportiversesimulation.SimulateMsgCreateAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteAccount, &weightMsgDeleteAccount, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAccount = defaultWeightMsgDeleteAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAccount,
		sportiversesimulation.SimulateMsgDeleteAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePost,
			defaultWeightMsgCreatePost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePost,
			defaultWeightMsgUpdatePost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgUpdatePost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePost,
			defaultWeightMsgDeletePost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateComment,
			defaultWeightMsgCreateComment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgCreateComment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateComment,
			defaultWeightMsgUpdateComment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgUpdateComment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteComment,
			defaultWeightMsgDeleteComment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgDeleteComment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSubscription,
			defaultWeightMsgCreateSubscription,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgCreateSubscription(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteSubscription,
			defaultWeightMsgDeleteSubscription,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgDeleteSubscription(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateLike,
			defaultWeightMsgCreateLike,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgCreateLike(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteLike,
			defaultWeightMsgDeleteLike,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgDeleteLike(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAccount,
			defaultWeightMsgCreateAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgCreateAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAccount,
			defaultWeightMsgDeleteAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				sportiversesimulation.SimulateMsgDeleteAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
