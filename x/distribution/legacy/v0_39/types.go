package v0_39

import sdk "github.com/cosmos/cosmos-sdk/types"

type FeePool struct {
	CommunityPool sdk.DecCoins `json:"community_pool" yaml:"community_pool"` // pool for community funds yet to be spent
}
