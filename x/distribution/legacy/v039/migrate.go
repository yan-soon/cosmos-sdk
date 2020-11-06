package v039

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v036distr "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v036"
)

// DONTCOVER
// nolint

const (
	ModuleName = "distribution"
)

// Migrate accepts exported genesis state from v0.38 and migrates it to
// v0.39 genesis state. All entries are identical except for parameters.
func Migrate(oldGenState v036distr.GenesisState) GenesisState {
	params := Params{
		CommunityTax:            oldGenState.CommunityTax,
		BaseProposerReward:      oldGenState.BaseProposerReward,
		BonusProposerReward:     oldGenState.BonusProposerReward,
		LiquidityProviderReward: sdk.ZeroDec(),
		WithdrawAddrEnabled:     oldGenState.WithdrawAddrEnabled,
	}

	feePool := FeePool{
		CommunityPool:         oldGenState.FeePool.CommunityPool,
		LiquidityProviderPool: sdk.NewDecCoins(),
	}

	return NewGenesisState(
		params, feePool,
		oldGenState.DelegatorWithdrawInfos, oldGenState.PreviousProposer,
		oldGenState.OutstandingRewards, oldGenState.ValidatorAccumulatedCommissions,
		oldGenState.ValidatorHistoricalRewards, oldGenState.ValidatorCurrentRewards,
		oldGenState.DelegatorStartingInfos, oldGenState.ValidatorSlashEvents,
	)
}
