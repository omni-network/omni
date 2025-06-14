// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package anchorinbox

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Claim is the `claim` instruction.
type ClaimInstruction struct {
	OrderId *ag_solanago.PublicKey

	// [0] = [WRITE] order_state
	//
	// [1] = [WRITE] order_token_account
	//
	// [2] = [WRITE] owner_token_account
	//
	// [3] = [WRITE, SIGNER] claimer
	//
	// [4] = [WRITE] claimer_token_account
	//
	// [5] = [] token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClaimInstructionBuilder creates a new `ClaimInstruction` instruction builder.
func NewClaimInstructionBuilder() *ClaimInstruction {
	nd := &ClaimInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[5] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	return nd
}

// SetOrderId sets the "order_id" parameter.
func (inst *ClaimInstruction) SetOrderId(order_id ag_solanago.PublicKey) *ClaimInstruction {
	inst.OrderId = &order_id
	return inst
}

// SetOrderStateAccount sets the "order_state" account.
func (inst *ClaimInstruction) SetOrderStateAccount(orderState ag_solanago.PublicKey) *ClaimInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(orderState).WRITE()
	return inst
}

func (inst *ClaimInstruction) findFindOrderStateAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: order_state
	seeds = append(seeds, []byte{byte(0x6f), byte(0x72), byte(0x64), byte(0x65), byte(0x72), byte(0x5f), byte(0x73), byte(0x74), byte(0x61), byte(0x74), byte(0x65)})
	// arg: OrderId
	orderIdSeed, err := ag_binary.MarshalBorsh(inst.OrderId)
	if err != nil {
		return
	}
	seeds = append(seeds, orderIdSeed)

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindOrderStateAddressWithBumpSeed calculates OrderState account address with given seeds and a known bump seed.
func (inst *ClaimInstruction) FindOrderStateAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindOrderStateAddress(bumpSeed)
	return
}

func (inst *ClaimInstruction) MustFindOrderStateAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindOrderStateAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindOrderStateAddress finds OrderState account address with given seeds.
func (inst *ClaimInstruction) FindOrderStateAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindOrderStateAddress(0)
	return
}

func (inst *ClaimInstruction) MustFindOrderStateAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindOrderStateAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetOrderStateAccount gets the "order_state" account.
func (inst *ClaimInstruction) GetOrderStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOrderTokenAccount sets the "order_token_account" account.
func (inst *ClaimInstruction) SetOrderTokenAccount(orderTokenAccount ag_solanago.PublicKey) *ClaimInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(orderTokenAccount).WRITE()
	return inst
}

func (inst *ClaimInstruction) findFindOrderTokenAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: order_token
	seeds = append(seeds, []byte{byte(0x6f), byte(0x72), byte(0x64), byte(0x65), byte(0x72), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e)})
	// arg: OrderId
	orderIdSeed, err := ag_binary.MarshalBorsh(inst.OrderId)
	if err != nil {
		return
	}
	seeds = append(seeds, orderIdSeed)

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindOrderTokenAddressWithBumpSeed calculates OrderTokenAccount account address with given seeds and a known bump seed.
func (inst *ClaimInstruction) FindOrderTokenAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindOrderTokenAddress(bumpSeed)
	return
}

func (inst *ClaimInstruction) MustFindOrderTokenAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindOrderTokenAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindOrderTokenAddress finds OrderTokenAccount account address with given seeds.
func (inst *ClaimInstruction) FindOrderTokenAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindOrderTokenAddress(0)
	return
}

func (inst *ClaimInstruction) MustFindOrderTokenAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindOrderTokenAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetOrderTokenAccount gets the "order_token_account" account.
func (inst *ClaimInstruction) GetOrderTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerTokenAccount sets the "owner_token_account" account.
func (inst *ClaimInstruction) SetOwnerTokenAccount(ownerTokenAccount ag_solanago.PublicKey) *ClaimInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ownerTokenAccount).WRITE()
	return inst
}

// GetOwnerTokenAccount gets the "owner_token_account" account.
func (inst *ClaimInstruction) GetOwnerTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetClaimerAccount sets the "claimer" account.
func (inst *ClaimInstruction) SetClaimerAccount(claimer ag_solanago.PublicKey) *ClaimInstruction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(claimer).WRITE().SIGNER()
	return inst
}

// GetClaimerAccount gets the "claimer" account.
func (inst *ClaimInstruction) GetClaimerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetClaimerTokenAccount sets the "claimer_token_account" account.
func (inst *ClaimInstruction) SetClaimerTokenAccount(claimerTokenAccount ag_solanago.PublicKey) *ClaimInstruction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(claimerTokenAccount).WRITE()
	return inst
}

// GetClaimerTokenAccount gets the "claimer_token_account" account.
func (inst *ClaimInstruction) GetClaimerTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *ClaimInstruction) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClaimInstruction {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *ClaimInstruction) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst ClaimInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Claim,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimInstruction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.OrderId == nil {
			return errors.New("OrderId parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.OrderState is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OrderTokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.OwnerTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Claimer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ClaimerTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *ClaimInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Claim")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" OrderId", *inst.OrderId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   order_state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  order_token_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  owner_token_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       claimer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("claimer_token_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta(" token_program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj ClaimInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OrderId` param:
	err = encoder.Encode(obj.OrderId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ClaimInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OrderId`:
	err = decoder.Decode(&obj.OrderId)
	if err != nil {
		return err
	}
	return nil
}

// NewClaimInstruction declares a new Claim instruction with the provided parameters and accounts.
func NewClaimInstruction(
	// Parameters:
	order_id ag_solanago.PublicKey,
	// Accounts:
	orderState ag_solanago.PublicKey,
	orderTokenAccount ag_solanago.PublicKey,
	ownerTokenAccount ag_solanago.PublicKey,
	claimer ag_solanago.PublicKey,
	claimerTokenAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *ClaimInstruction {
	return NewClaimInstructionBuilder().
		SetOrderId(order_id).
		SetOrderStateAccount(orderState).
		SetOrderTokenAccount(orderTokenAccount).
		SetOwnerTokenAccount(ownerTokenAccount).
		SetClaimerAccount(claimer).
		SetClaimerTokenAccount(claimerTokenAccount).
		SetTokenProgramAccount(tokenProgram)
}
