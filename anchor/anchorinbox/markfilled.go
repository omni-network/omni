// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package anchorinbox

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Mark an order as filled, and set the claimable_by account.
// This may only be called by the inbox admin.
type MarkFilledInstruction struct {
	OrderId     *ag_solanago.PublicKey
	FillHash    *ag_solanago.PublicKey
	ClaimableBy *ag_solanago.PublicKey

	// [0] = [WRITE] order_state
	//
	// [1] = [] inbox_state
	//
	// [2] = [WRITE, SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMarkFilledInstructionBuilder creates a new `MarkFilledInstruction` instruction builder.
func NewMarkFilledInstructionBuilder() *MarkFilledInstruction {
	nd := &MarkFilledInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetOrderId sets the "_order_id" parameter.
func (inst *MarkFilledInstruction) SetOrderId(_order_id ag_solanago.PublicKey) *MarkFilledInstruction {
	inst.OrderId = &_order_id
	return inst
}

// SetFillHash sets the "fill_hash" parameter.
func (inst *MarkFilledInstruction) SetFillHash(fill_hash ag_solanago.PublicKey) *MarkFilledInstruction {
	inst.FillHash = &fill_hash
	return inst
}

// SetClaimableBy sets the "claimable_by" parameter.
func (inst *MarkFilledInstruction) SetClaimableBy(claimable_by ag_solanago.PublicKey) *MarkFilledInstruction {
	inst.ClaimableBy = &claimable_by
	return inst
}

// SetOrderStateAccount sets the "order_state" account.
func (inst *MarkFilledInstruction) SetOrderStateAccount(orderState ag_solanago.PublicKey) *MarkFilledInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(orderState).WRITE()
	return inst
}

func (inst *MarkFilledInstruction) findFindOrderStateAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *MarkFilledInstruction) FindOrderStateAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindOrderStateAddress(bumpSeed)
	return
}

func (inst *MarkFilledInstruction) MustFindOrderStateAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindOrderStateAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindOrderStateAddress finds OrderState account address with given seeds.
func (inst *MarkFilledInstruction) FindOrderStateAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindOrderStateAddress(0)
	return
}

func (inst *MarkFilledInstruction) MustFindOrderStateAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindOrderStateAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetOrderStateAccount gets the "order_state" account.
func (inst *MarkFilledInstruction) GetOrderStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetInboxStateAccount sets the "inbox_state" account.
func (inst *MarkFilledInstruction) SetInboxStateAccount(inboxState ag_solanago.PublicKey) *MarkFilledInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(inboxState)
	return inst
}

func (inst *MarkFilledInstruction) findFindInboxStateAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: inbox_state
	seeds = append(seeds, []byte{byte(0x69), byte(0x6e), byte(0x62), byte(0x6f), byte(0x78), byte(0x5f), byte(0x73), byte(0x74), byte(0x61), byte(0x74), byte(0x65)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindInboxStateAddressWithBumpSeed calculates InboxState account address with given seeds and a known bump seed.
func (inst *MarkFilledInstruction) FindInboxStateAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindInboxStateAddress(bumpSeed)
	return
}

func (inst *MarkFilledInstruction) MustFindInboxStateAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindInboxStateAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindInboxStateAddress finds InboxState account address with given seeds.
func (inst *MarkFilledInstruction) FindInboxStateAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindInboxStateAddress(0)
	return
}

func (inst *MarkFilledInstruction) MustFindInboxStateAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindInboxStateAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetInboxStateAccount gets the "inbox_state" account.
func (inst *MarkFilledInstruction) GetInboxStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAdminAccount sets the "admin" account.
func (inst *MarkFilledInstruction) SetAdminAccount(admin ag_solanago.PublicKey) *MarkFilledInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *MarkFilledInstruction) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst MarkFilledInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MarkFilled,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MarkFilledInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MarkFilledInstruction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.OrderId == nil {
			return errors.New("OrderId parameter is not set")
		}
		if inst.FillHash == nil {
			return errors.New("FillHash parameter is not set")
		}
		if inst.ClaimableBy == nil {
			return errors.New("ClaimableBy parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.OrderState is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.InboxState is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *MarkFilledInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MarkFilled")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     OrderId", *inst.OrderId))
						paramsBranch.Child(ag_format.Param("    FillHash", *inst.FillHash))
						paramsBranch.Child(ag_format.Param(" ClaimableBy", *inst.ClaimableBy))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("order_state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("inbox_state", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      admin", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj MarkFilledInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OrderId` param:
	err = encoder.Encode(obj.OrderId)
	if err != nil {
		return err
	}
	// Serialize `FillHash` param:
	err = encoder.Encode(obj.FillHash)
	if err != nil {
		return err
	}
	// Serialize `ClaimableBy` param:
	err = encoder.Encode(obj.ClaimableBy)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MarkFilledInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OrderId`:
	err = decoder.Decode(&obj.OrderId)
	if err != nil {
		return err
	}
	// Deserialize `FillHash`:
	err = decoder.Decode(&obj.FillHash)
	if err != nil {
		return err
	}
	// Deserialize `ClaimableBy`:
	err = decoder.Decode(&obj.ClaimableBy)
	if err != nil {
		return err
	}
	return nil
}

// NewMarkFilledInstruction declares a new MarkFilled instruction with the provided parameters and accounts.
func NewMarkFilledInstruction(
	// Parameters:
	_order_id ag_solanago.PublicKey,
	fill_hash ag_solanago.PublicKey,
	claimable_by ag_solanago.PublicKey,
	// Accounts:
	orderState ag_solanago.PublicKey,
	inboxState ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *MarkFilledInstruction {
	return NewMarkFilledInstructionBuilder().
		SetOrderId(_order_id).
		SetFillHash(fill_hash).
		SetClaimableBy(claimable_by).
		SetOrderStateAccount(orderState).
		SetInboxStateAccount(inboxState).
		SetAdminAccount(admin)
}
