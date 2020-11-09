package v0_39

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// Parameter keys
var (
	ParamStoreKeyCommunityTax        = []byte("communitytax")
	ParamStoreKeyBaseProposerReward  = []byte("baseproposerreward")
	ParamStoreKeyBonusProposerReward = []byte("bonusproposerreward")
	ParamStoreKeyWithdrawAddrEnabled = []byte("withdrawaddrenabled")
)

// FeePool is the global fee pool for distribution in v0.39.
type FeePool struct {
	CommunityPool sdk.DecCoins `json:"community_pool" yaml:"community_pool"` // pool for community funds yet to be spent
}

// Params defines the set of distribution parameters in v0.39.
type Params struct {
	CommunityTax        sdk.Dec `json:"community_tax" yaml:"community_tax"`
	BaseProposerReward  sdk.Dec `json:"base_proposer_reward" yaml:"base_proposer_reward"`
	BonusProposerReward sdk.Dec `json:"bonus_proposer_reward" yaml:"bonus_proposer_reward"`
	WithdrawAddrEnabled bool    `json:"withdraw_addr_enabled" yaml:"withdraw_addr_enabled"`
}

// ParamSetPairs returns the parameter set pairs. Validation functions are removed.
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(ParamStoreKeyCommunityTax, &p.CommunityTax, noValidation),
		params.NewParamSetPair(ParamStoreKeyBaseProposerReward, &p.BaseProposerReward, noValidation),
		params.NewParamSetPair(ParamStoreKeyBonusProposerReward, &p.BonusProposerReward, noValidation),
		params.NewParamSetPair(ParamStoreKeyWithdrawAddrEnabled, &p.WithdrawAddrEnabled, noValidation),
	}
}

func noValidation(i interface{}) error {
	return nil
}
