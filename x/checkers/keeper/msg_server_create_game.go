package keeper

import (
	"checkers/x/checkers/rules"
	"context"
	"strconv"

	"checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)
	newGame := rules.New()
	storedGame := types.StoredGame{
		Index:     newIndex,
		Board:     newGame.String(),
		Turn:      rules.PieceStrings[newGame.Turn],
		Black:     msg.Black,
		Red:       msg.Red,
		MoveCount: 0,
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	k.Keeper.SetStoredGame(ctx, storedGame)
	systemInfo.NextId++
	k.SetSystemInfo(ctx, systemInfo)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.GameCreatedEventType,
		sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator)))

	return &types.MsgCreateGameResponse{GameIndex: newIndex}, nil
}
