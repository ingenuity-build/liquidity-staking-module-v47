// Package v034 is used for legacy migration scripts. Actual migration scripts
// for v034 have been removed, but the v039->v042 migration script still
// references types from this file, so we're keeping it for now.
// DONTCOVER
package v034

import (
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"
)

const (
	ModuleName = "staking"
)

// staking constants
const (
	Unbonded  BondStatus = 0x00
	Unbonding BondStatus = 0x01
	Bonded    BondStatus = 0x02

	BondStatusUnbonded  = "Unbonded"
	BondStatusUnbonding = "Unbonding"
	BondStatusBonded    = "Bonded"
)

type (
	// BondStatus is the status of a validator
	BondStatus byte

	Pool struct {
		NotBondedTokens math.Int `json:"not_bonded_tokens"`
		BondedTokens    math.Int `json:"bonded_tokens"`
	}

	Params struct {
		UnbondingTime time.Duration `json:"unbonding_time"`
		MaxValidators uint16        `json:"max_validators"`
		MaxEntries    uint16        `json:"max_entries"`
		BondDenom     string        `json:"bond_denom"`
	}

	LastValidatorPower struct {
		Address sdk.ValAddress
		Power   int64
	}

	Description struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	}

	Commission struct {
		Rate          sdk.Dec   `json:"rate"`
		MaxRate       sdk.Dec   `json:"max_rate"`
		MaxChangeRate sdk.Dec   `json:"max_change_rate"`
		UpdateTime    time.Time `json:"update_time"`
	}

	bechValidator struct {
		OperatorAddress         sdk.ValAddress `json:"operator_address"`    // the bech32 address of the validator's operator
		ConsPubKey              string         `json:"consensus_pubkey"`    // the bech32 consensus public key of the validator
		Jailed                  bool           `json:"jailed"`              // has the validator been jailed from bonded status?
		Status                  BondStatus     `json:"status"`              // validator status (bonded/unbonding/unbonded)
		Tokens                  math.Int       `json:"tokens"`              // delegated tokens (incl. self-delegation)
		DelegatorShares         sdk.Dec        `json:"delegator_shares"`    // total shares issued to a validator's delegators
		Description             Description    `json:"description"`         // description terms for the validator
		UnbondingHeight         int64          `json:"unbonding_height"`    // if unbonding, height at which this validator has begun unbonding
		UnbondingCompletionTime time.Time      `json:"unbonding_time"`      // if unbonding, min time for the validator to complete unbonding
		Commission              Commission     `json:"commission"`          // commission parameters
		MinSelfDelegation       math.Int       `json:"min_self_delegation"` // minimum self delegation
	}

	Validator struct {
		OperatorAddress         sdk.ValAddress     `json:"operator_address"`
		ConsPubKey              cryptotypes.PubKey `json:"consensus_pubkey"`
		Jailed                  bool               `json:"jailed"`
		Status                  BondStatus         `json:"status"`
		Tokens                  math.Int           `json:"tokens"`
		DelegatorShares         sdk.Dec            `json:"delegator_shares"`
		Description             Description        `json:"description"`
		UnbondingHeight         int64              `json:"unbonding_height"`
		UnbondingCompletionTime time.Time          `json:"unbonding_time"`
		Commission              Commission         `json:"commission"`
		MinSelfDelegation       math.Int           `json:"min_self_delegation"`
	}

	Validators []Validator

	Delegation struct {
		DelegatorAddress sdk.AccAddress `json:"delegator_address"`
		ValidatorAddress sdk.ValAddress `json:"validator_address"`
		Shares           sdk.Dec        `json:"shares"`
	}

	Delegations []Delegation

	UnbondingDelegationEntry struct {
		CreationHeight int64     `json:"creation_height"`
		CompletionTime time.Time `json:"completion_time"`
		InitialBalance math.Int  `json:"initial_balance"`
		Balance        math.Int  `json:"balance"`
	}

	UnbondingDelegation struct {
		DelegatorAddress sdk.AccAddress             `json:"delegator_address"`
		ValidatorAddress sdk.ValAddress             `json:"validator_address"`
		Entries          []UnbondingDelegationEntry `json:"entries"`
	}

	RedelegationEntry struct {
		CreationHeight int64     `json:"creation_height"`
		CompletionTime time.Time `json:"completion_time"`
		InitialBalance math.Int  `json:"initial_balance"`
		SharesDst      sdk.Dec   `json:"shares_dst"`
	}

	Redelegation struct {
		DelegatorAddress    sdk.AccAddress      `json:"delegator_address"`
		ValidatorSrcAddress sdk.ValAddress      `json:"validator_src_address"`
		ValidatorDstAddress sdk.ValAddress      `json:"validator_dst_address"`
		Entries             []RedelegationEntry `json:"entries"`
	}

	GenesisState struct {
		Pool                 Pool                  `json:"pool"`
		Params               Params                `json:"params"`
		LastTotalPower       math.Int              `json:"last_total_power"`
		LastValidatorPowers  []LastValidatorPower  `json:"last_validator_powers"`
		Validators           Validators            `json:"validators"`
		Delegations          Delegations           `json:"delegations"`
		UnbondingDelegations []UnbondingDelegation `json:"unbonding_delegations"`
		Redelegations        []Redelegation        `json:"redelegations"`
		Exported             bool                  `json:"exported"`
	}
)

func (v Validator) MarshalJSON() ([]byte, error) {
	bechConsPubKey, err := legacybech32.MarshalPubKey(legacybech32.ConsPK, v.ConsPubKey)
	if err != nil {
		return nil, err
	}

	return legacy.Cdc.MarshalJSON(bechValidator{
		OperatorAddress:         v.OperatorAddress,
		ConsPubKey:              bechConsPubKey,
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		Tokens:                  v.Tokens,
		DelegatorShares:         v.DelegatorShares,
		Description:             v.Description,
		UnbondingHeight:         v.UnbondingHeight,
		UnbondingCompletionTime: v.UnbondingCompletionTime,
		MinSelfDelegation:       v.MinSelfDelegation,
		Commission:              v.Commission,
	})
}

// UnmarshalJSON unmarshals the validator from JSON using Bech32
func (v *Validator) UnmarshalJSON(data []byte) error {
	bv := &bechValidator{}
	if err := legacy.Cdc.UnmarshalJSON(data, bv); err != nil {
		return err
	}
	consPubKey, err := legacybech32.UnmarshalPubKey(legacybech32.ConsPK, bv.ConsPubKey)
	if err != nil {
		return err
	}

	*v = Validator{
		OperatorAddress:         bv.OperatorAddress,
		ConsPubKey:              consPubKey,
		Jailed:                  bv.Jailed,
		Tokens:                  bv.Tokens,
		Status:                  bv.Status,
		DelegatorShares:         bv.DelegatorShares,
		Description:             bv.Description,
		UnbondingHeight:         bv.UnbondingHeight,
		UnbondingCompletionTime: bv.UnbondingCompletionTime,
		Commission:              bv.Commission,
		MinSelfDelegation:       bv.MinSelfDelegation,
	}
	return nil
}
