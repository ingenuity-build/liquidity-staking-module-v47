package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/iqlusioninc/liquidity-staking-module/x/staking/types"
)

// SetTotalLiquidStakedTokens stores the total outstanding tokens owned by a liquid staking provider
func (k Keeper) SetTotalLiquidStakedTokens(ctx sdk.Context, tokens sdk.Int) {
	store := ctx.KVStore(k.storeKey)

	tokensBz, err := tokens.Marshal()
	if err != nil {
		panic(err)
	}

	store.Set(types.TotalLiquidStakedTokensKey, tokensBz)
}

// GetTotalLiquidStakedTokens returns the total outstanding tokens owned by a liquid staking provider
// Returns zero if the total liquid stake amount has not been initialized
func (k Keeper) GetTotalLiquidStakedTokens(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	tokensBz := store.Get(types.TotalLiquidStakedTokensKey)

	if tokensBz == nil {
		panic("total liquid staked was never initialized")
	}

	var tokens sdk.Int
	if err := tokens.Unmarshal(tokensBz); err != nil {
		panic(err)
	}

	return tokens
}

// Check if an account is a owned by a liquid staking provider
// This is determined by checking if the account is a 32-length module account
func (k Keeper) AccountIsLiquidStakingProvider(ctx sdk.Context, address sdk.AccAddress) bool {
	account := k.authKeeper.GetAccount(ctx, address)
	_, isModuleAccount := account.(*authtypes.ModuleAccount)
	return isModuleAccount && len(address) == 32
}

// ExceedsGlobalLiquidStakingCap checks if a liquid delegation would cause the
// global liquid staking cap to be exceeded
// A liquid delegation is defined as either tokenized shares, or a delegation from an ICA Account
// The total stake is determined by the balance of the bonded pool
// Returns true if the cap is exceeded
func (k Keeper) CheckExceedsGlobalLiquidStakingCap(ctx sdk.Context, tokens sdk.Int) bool {
	liquidStakingCap := k.GlobalLiquidStakingCap(ctx)
	liquidStakedAmount := k.GetTotalLiquidStakedTokens(ctx)

	// Determine the total stake from the balance of the bonded pools
	bondedPoolAddress := k.authKeeper.GetModuleAddress(types.BondedPoolName)
	totalStakedAmount := k.bankKeeper.GetBalance(ctx, bondedPoolAddress, k.BondDenom(ctx)).Amount

	// Calculate the percentage of stake that is liquid
	updatedTotalStaked := sdk.NewDecFromInt(totalStakedAmount.Add(tokens))
	updatedLiquidStaked := sdk.NewDecFromInt(liquidStakedAmount.Add(tokens))
	liquidStakePercent := updatedLiquidStaked.Quo(updatedTotalStaked)

	return liquidStakePercent.GT(liquidStakingCap)
}

// ExceedsValidatorBondCap checks if a liquid delegation to a validator would cause
// the liquid shares to exceed the validator bond factor
// A liquid delegation is defined as either tokenized shares, or a delegation from an ICA Account
// Returns true if the cap is exceeded
func (k Keeper) CheckExceedsValidatorBondCap(ctx sdk.Context, validator types.Validator, shares sdk.Dec) bool {
	validatorBondFactor := k.ValidatorBondFactor(ctx)
	maxValLiquidShares := validator.TotalValidatorBondShares.Mul(validatorBondFactor)
	return validator.TotalLiquidShares.Add(shares).GT(maxValLiquidShares)
}

// SafelyIncreaseTotalLiquidStakedTokens increments the total liquid staked tokens
// if the global cap is enabled and is not surpassed by this delegation
func (k Keeper) SafelyIncreaseTotalLiquidStakedTokens(ctx sdk.Context, amount sdk.Int) error {
	// If the cap is disabled, do nothing
	if !k.GlobalLiquidStakingCapEnabled(ctx) {
		return nil
	}

	// Confirm the cap will not be exceeded
	if k.CheckExceedsGlobalLiquidStakingCap(ctx, amount) {
		return types.ErrGlobalLiquidStakingCapExceeded
	}

	// Increment the global total liquid staked
	k.SetTotalLiquidStakedTokens(ctx, k.GetTotalLiquidStakedTokens(ctx).Add(amount))

	return nil
}

// DecreaseTotalLiquidStakedTokens decrements the total liquid staked tokens
// if the global cap is enabled
func (k Keeper) DecreaseTotalLiquidStakedTokens(ctx sdk.Context, amount sdk.Int) {
	if k.GlobalLiquidStakingCapEnabled(ctx) {
		k.SetTotalLiquidStakedTokens(ctx, k.GetTotalLiquidStakedTokens(ctx).Sub(amount))
	}
}

// SafelyIncreaseValidatorTotalLiquidShares increments the total liquid shares on a validator
// if the validator bond factor is enabled and is not surpassed by this delegation
func (k Keeper) SafelyIncreaseValidatorTotalLiquidShares(ctx sdk.Context, validator types.Validator, shares sdk.Dec) error {
	// If the cap is disabled, do nothing
	if !k.ValidatorBondFactorEnabled(ctx) {
		return nil
	}

	// Confirm the validator bond factor will be not exceeded
	if k.CheckExceedsValidatorBondCap(ctx, validator, shares) {
		return types.ErrInsufficientValidatorBondShares
	}

	// Increment the validator's total liquid shares
	validator.TotalLiquidShares = validator.TotalLiquidShares.Add(shares)
	k.SetValidator(ctx, validator)

	return nil
}

// DecreaseValidatorTotalLiquidShares decrements the total liquid shares on a validator
// if the validator bond factor is enabled
func (k Keeper) DecreaseValidatorTotalLiquidShares(ctx sdk.Context, validator types.Validator, shares sdk.Dec) {
	if k.ValidatorBondFactorEnabled(ctx) {
		validator.TotalLiquidShares = validator.TotalLiquidShares.Sub(shares)
		k.SetValidator(ctx, validator)
	}
}

// SafelyDecreaseValidatorBond decrements the total validator's self bond
// so long as it will not cause the current delegations to exceed the threshold
// set by validator bond factor
func (k Keeper) SafelyDecreaseValidatorBond(ctx sdk.Context, validator types.Validator, shares sdk.Dec) error {
	// If the cap is disabled, do nothing
	if !k.ValidatorBondFactorEnabled(ctx) {
		return nil
	}

	// Check if the decreased self bond will cause the validator bond threshold to be exceeded
	validatorBondFactor := k.ValidatorBondFactor(ctx)
	maxValTotalShare := validator.TotalValidatorBondShares.Sub(shares).Mul(validatorBondFactor)
	if validator.TotalLiquidShares.GT(maxValTotalShare) {
		return types.ErrInsufficientValidatorBondShares
	}

	// Decrement the validator's total self bond
	validator.TotalValidatorBondShares = validator.TotalValidatorBondShares.Sub(shares)
	k.SetValidator(ctx, validator)

	return nil
}

// AddTokenizeSharesLock adds a lock that prevents tokenizing shares for an account
// The tokenize share lock store is implemented by keying on the account address
// with a fixed value for each entry
func (k Keeper) AddTokenizeSharesLock(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetTokenizeSharesLockKey(address)
	store.Set(key, []byte{types.TokenizeSharesLockValue})
}

// Removes the tokenize share lock for an account to enable tokenizing shares
func (k Keeper) RemoveTokenizeSharesLock(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetTokenizeSharesLockKey(address)
	store.Delete(key)
}

// Checks if there is currently a tokenize share lock for a given account
// Returns true if the account is locked
func (k Keeper) IsTokenizeSharesDisabled(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.GetTokenizeSharesLockKey(address)
	bz := store.Get(key)
	if len(bz) == 0 {
		return false
	}
	return len(bz) == 1 && bz[0] == types.TokenizeSharesLockValue
}

// Returns a list of addresses pending tokenize share unlocking at the same time
func (k Keeper) GetPendingTokenizeShareAuthorizations(ctx sdk.Context, completionTime time.Time) types.PendingTokenizeShareAuthorizations {
	store := ctx.KVStore(k.storeKey)

	timeKey := types.GetTokenizeShareAuthorizationTimeKey(completionTime)
	bz := store.Get(timeKey)

	authorizations := types.PendingTokenizeShareAuthorizations{Addresses: []string{}}
	if len(bz) == 0 {
		return authorizations
	}
	k.cdc.MustUnmarshal(bz, &authorizations)

	return authorizations
}

// Stores a list of addresses pending tokenize share unlocking at the same time
func (k Keeper) SetPendingTokenizeShareAuthorizations(ctx sdk.Context, completionTime time.Time, authorizations types.PendingTokenizeShareAuthorizations) {
	store := ctx.KVStore(k.storeKey)
	timeKey := types.GetTokenizeShareAuthorizationTimeKey(completionTime)
	bz := k.cdc.MustMarshal(&authorizations)
	store.Set(timeKey, bz)
}

// Inserts thea ddress into a queue where it will sit for 1 unbonding period
// before the tokenize share lock is removed
// Returns the completion time
func (k Keeper) QueueTokenizeSharesAuthorization(ctx sdk.Context, address sdk.AccAddress) time.Time {
	params := k.GetParams(ctx)
	completionTime := ctx.BlockTime().Add(params.UnbondingTime)

	authorizations := k.GetPendingTokenizeShareAuthorizations(ctx, completionTime)
	authorizations.Addresses = append(authorizations.Addresses, address.String())

	k.SetPendingTokenizeShareAuthorizations(ctx, completionTime, authorizations)

	return completionTime
}

// Unlocks all queued tokenize share authorizations that have matured
// (i.e. have waited the full unbonding period)
func (k Keeper) RemoveExpiredTokenizeShareLocks(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	// iterators all time slices from time 0 until the current block time
	prefixEnd := sdk.InclusiveEndBytes(types.GetTokenizeShareAuthorizationTimeKey(ctx.BlockTime()))
	iterator := store.Iterator(types.TokenizeSharesUnlockQueueKey, prefixEnd)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		authorizations := types.PendingTokenizeShareAuthorizations{}
		k.cdc.MustUnmarshal(iterator.Value(), &authorizations)

		for _, address := range authorizations.Addresses {
			k.RemoveTokenizeSharesLock(ctx, sdk.MustAccAddressFromBech32(address))
		}
		store.Delete(iterator.Key())
	}
}
