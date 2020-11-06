package types

import (
	"fmt"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	// default paramspace for params keeper
	DefaultParamspace = ModuleName
)

// Parameter keys
var (
	ParamStoreKeyCommunityTax            = []byte("communitytax")
	ParamStoreKeyBaseProposerReward      = []byte("baseproposerreward")
	ParamStoreKeyBonusProposerReward     = []byte("bonusproposerreward")
	ParamStoreKeyLiquidityProviderReward = []byte("liquidityproviderreward")
	ParamStoreKeyWithdrawAddrEnabled     = []byte("withdrawaddrenabled")
)

// Params defines the set of distribution parameters.
type Params struct {
	CommunityTax            sdk.Dec `json:"community_tax" yaml:"community_tax"`
	BaseProposerReward      sdk.Dec `json:"base_proposer_reward" yaml:"base_proposer_reward"`
	BonusProposerReward     sdk.Dec `json:"bonus_proposer_reward" yaml:"bonus_proposer_reward"`
	LiquidityProviderReward sdk.Dec `json:"liquidity_provider_reward" yaml:"liquidity_provider_reward"`
	WithdrawAddrEnabled     bool    `json:"withdraw_addr_enabled" yaml:"withdraw_addr_enabled"`
}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		CommunityTax:            sdk.NewDecWithPrec(2, 2), // 2%
		BaseProposerReward:      sdk.NewDecWithPrec(1, 2), // 1%
		BonusProposerReward:     sdk.NewDecWithPrec(4, 2), // 4%
		LiquidityProviderReward: sdk.ZeroDec(),            // 0%
		WithdrawAddrEnabled:     true,
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(ParamStoreKeyCommunityTax, &p.CommunityTax, validateCommunityTax),
		params.NewParamSetPair(ParamStoreKeyBaseProposerReward, &p.BaseProposerReward, validateBaseProposerReward),
		params.NewParamSetPair(ParamStoreKeyBonusProposerReward, &p.BonusProposerReward, validateBonusProposerReward),
		params.NewParamSetPair(ParamStoreKeyLiquidityProviderReward, &p.LiquidityProviderReward, validateLiquidityProviderReward),
		params.NewParamSetPair(ParamStoreKeyWithdrawAddrEnabled, &p.WithdrawAddrEnabled, validateWithdrawAddrEnabled),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.CommunityTax.IsNegative() {
		return fmt.Errorf(
			"community tax should be positive: %s", p.CommunityTax,
		)
	}
	if p.BaseProposerReward.IsNegative() {
		return fmt.Errorf(
			"base proposer reward should be positive: %s", p.BaseProposerReward,
		)
	}
	if p.BonusProposerReward.IsNegative() {
		return fmt.Errorf(
			"bonus proposer reward should be positive: %s", p.BonusProposerReward,
		)
	}
	if p.LiquidityProviderReward.IsNegative() {
		return fmt.Errorf(
			"liquidity provider reward should be positive: %s", p.CommunityTax,
		)
	}
	if v := p.BaseProposerReward.Add(p.BonusProposerReward).Add(p.CommunityTax).Add(p.LiquidityProviderReward); v.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"sum of all rewards cannot greater than one: %s", v,
		)
	}

	return nil
}

func validateCommunityTax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("community tax must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("community tax must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("community tax too large: %s", v)
	}

	return nil
}

func validateBaseProposerReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("base proposer reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("base proposer reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("base proposer reward too large: %s", v)
	}

	return nil
}

func validateBonusProposerReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("bonus proposer reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("bonus proposer reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("bonus proposer reward too large: %s", v)
	}

	return nil
}

func validateLiquidityProviderReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("liquidity provider reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("liquidity provider reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("liquidity provider reward too large: %s", v)
	}

	return nil
}

func validateWithdrawAddrEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
