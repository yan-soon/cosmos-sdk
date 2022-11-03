package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// combine multiple send hooks, all hook functions are run in array sequence
type MutiSendHooks []SendHooks

func NewMutiSendHooks(hooks ...SendHooks) MutiSendHooks {
	return hooks
}

func (h MutiSendHooks) AfterSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) {
	for i := range h {
		h[i].AfterSend(ctx, fromAddr, toAddr, amt)
	}
}

func (h MutiSendHooks) BeforeSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) {
	for i := range h {
		h[i].BeforeSend(ctx, fromAddr, toAddr, amt)
	}
}
