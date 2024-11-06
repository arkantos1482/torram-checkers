package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"github.com/alice/checkers"
	"github.com/alice/checkers/rules"
)

type torramMsgServer struct {
	k Keeper
}

var _ checkers.CheckersTorramServer = torramMsgServer{}

// NewTorramMsgServerImpl returns an implementation of the module MsgServer interface.
func NewTorramMsgServerImpl(keeper Keeper) checkers.CheckersTorramServer {
	return &torramMsgServer{k: keeper}
}

// CheckersCreateGm defines the handler for the ReqCheckersTorram message.
func (ms torramMsgServer) CheckersCreateGm(ctx context.Context, msg *checkers.ReqCheckersTorram) (*checkers.ResCheckersTorram, error) {
	if length := len([]byte(msg.Index)); checkers.MaxIndexLength < length || length < 1 {
		return nil, checkers.ErrIndexTooLong
	}
	if _, err := ms.k.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
	}

	newBoard := rules.New()
	storedGame := checkers.StoredGame{
		Board:       newBoard.String(),
		Turn:        rules.PieceStrings[newBoard.Turn],
		Black:       msg.Black,
		Red:         msg.Red,
		StartTime:   msg.StartTime,
		EndTime:     msg.EndTime,
		Description: msg.Description,
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.ResCheckersTorram{}, nil
}
