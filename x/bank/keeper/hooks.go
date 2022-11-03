package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Implements SendHooks interface
var _ types.SendHooks = BaseSendKeeper{}

func (keeper BaseSendKeeper) AfterSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) {
	if keeper.hooks != nil {
		keeper.hooks.AfterSend(ctx, fromAddr, toAddr, amt)
	}
}

func (keeper BaseSendKeeper) BeforeSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) {
	if keeper.hooks != nil {
		keeper.hooks.BeforeSend(ctx, fromAddr, toAddr, amt)
	}
}
