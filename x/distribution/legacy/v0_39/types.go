package v0_39

import sdk "github.com/cosmos/cosmos-sdk/types"

type FeePool struct {
	CommunityPool sdk.DecCoins `json:"community_pool" yaml:"community_pool"` // pool for community funds yet to be spent
}

type Params struct {
	CommunityTax        sdk.Dec `json:"community_tax" yaml:"community_tax"`
	BaseProposerReward  sdk.Dec `json:"base_proposer_reward" yaml:"base_proposer_reward"`
	BonusProposerReward sdk.Dec `json:"bonus_proposer_reward" yaml:"bonus_proposer_reward"`
	WithdrawAddrEnabled bool    `json:"withdraw_addr_enabled" yaml:"withdraw_addr_enabled"`
}
