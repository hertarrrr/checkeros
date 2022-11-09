package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RejectGame(goCtx context.Context, msg *types.MsgRejectGame) (*types.MsgRejectGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	game, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
	}
	//if game.Black == msg.Creator {
	//	if 0 < game.MoveCount { // Notice the use of the new field
	//		return nil, types.ErrBlackAlreadyPlayed
	//	}
	//} else if game.Red == msg.Creator {
	//	if 1 < game.MoveCount { // Notice the use of the new field
	//		return nil, types.ErrRedAlreadyPlayed
	//	}
	//} else {
	//	return nil, sdkerrors.Wrapf(types.ErrCreatorNotPlayer, "%s", msg.Creator)
	//}
	_ = game
	k.Keeper.RemoveStoredGame(ctx, msg.GameIndex)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.GameRejectedEventType,
			sdk.NewAttribute(types.GameRejectedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameRejectedEventGameIndex, msg.GameIndex)))
	return &types.MsgRejectGameResponse{}, nil
}
