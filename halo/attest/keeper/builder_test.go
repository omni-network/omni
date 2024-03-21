package keeper_test

import (
	"github.com/omni-network/omni/halo/attest/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/k1util"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
)

//nolint:gochecknoinits // this data is required by all tests
func init() {
	// make sure that the first test which uses block hashes has data populated
	fuzz.New().NilChance(0).NumElements(3, 3).Fuzz(&blockHashes)
}

//nolint:gochecknoglobals // Hard-coded test data.
var (
	blockHashes []common.Hash
	vals        = []k1.PrivKey{k1.GenPrivKey(), k1.GenPrivKey(), k1.GenPrivKey()}
	val1        = newValidator(vals[0].PubKey(), 10)
	val2        = newValidator(vals[1].PubKey(), 15)
	val3        = newValidator(vals[2].PubKey(), 15)
	attRoot     = []byte("test attestation root")
)

func newValSet(id uint64, vals ...*vtypes.Validator) *vtypes.ValidatorSetResponse {
	return &vtypes.ValidatorSetResponse{
		Id:              id,
		CreatedHeight:   0,
		ActivatedHeight: 0,
		Validators:      vals,
	}
}

func newValidator(key crypto.PubKey, power int64) *vtypes.Validator {
	addr, _ := k1util.PubKeyToAddress(key)

	return &vtypes.Validator{
		Address: addr.Bytes(),
		Power:   power,
	}
}

type MsgBuilder struct {
	msg *types.MsgAddVotes
}

func (b *MsgBuilder) WithAuthority(a string) *MsgBuilder {
	if b.msg == nil {
		b.msg = &types.MsgAddVotes{}
	}
	b.msg.Authority = a

	return b
}

func (b *MsgBuilder) WithVotes(votes ...*types.AggVote) *MsgBuilder {
	if b.msg == nil {
		b.msg = &types.MsgAddVotes{}
	}
	b.msg.Votes = votes

	return b
}

func (b *MsgBuilder) WithAppendVotes(votes ...*types.AggVote) *MsgBuilder {
	if b.msg == nil {
		b.msg = &types.MsgAddVotes{}
	}
	b.msg.Votes = append(b.msg.Votes, votes...)

	return b
}

func (b *MsgBuilder) Default() *MsgBuilder {
	b.msg = &types.MsgAddVotes{
		Authority: "test-authority",
		Votes: []*types.AggVote{
			new(AggVoteBuilder).Default().Vote(),
		},
	}

	return b
}

func (b *MsgBuilder) Msg() *types.MsgAddVotes {
	return b.msg
}

type AggVoteBuilder struct {
	vote *types.AggVote
}

func (b *AggVoteBuilder) Default() *AggVoteBuilder {
	b.vote = &types.AggVote{
		BlockHeader: &types.BlockHeader{
			ChainId: 1,
			Height:  500,
			Hash:    blockHashes[0].Bytes(),
		},
		AttestationRoot: attRoot,
		Signatures:      sigsTuples(val1, val2),
	}

	return b
}

func (b *AggVoteBuilder) WithChainID(id uint64) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.ChainId = id

	return b
}

func (b *AggVoteBuilder) WithBlockHeight(h uint64) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.Height = h

	return b
}

func (b *AggVoteBuilder) WithBlockHash(h common.Hash) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.Hash = h.Bytes()

	return b
}

func (b *AggVoteBuilder) WithBlockHeader(chainID uint64, height uint64, hash common.Hash) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.ChainId = chainID
	b.vote.BlockHeader.Height = height
	b.vote.BlockHeader.Hash = hash.Bytes()

	return b
}

func (b *AggVoteBuilder) WithAttestationRoot(r []byte) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{}
	}
	b.vote.AttestationRoot = r

	return b
}

func (b *AggVoteBuilder) WithSignatures(s ...*types.SigTuple) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{}
	}
	b.vote.Signatures = s

	return b
}

func (b *AggVoteBuilder) WithAppendSignatures(s ...*types.SigTuple) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{}
	}
	b.vote.Signatures = append(b.vote.Signatures, s...)

	return b
}

func (b *AggVoteBuilder) Vote() *types.AggVote {
	return b.vote
}

func sigsTuples(vals ...*vtypes.Validator) []*types.SigTuple {
	var sigs []*types.SigTuple
	for _, v := range vals {
		sigs = append(sigs, &types.SigTuple{
			ValidatorAddress: v.Address,
			Signature:        v.Address, // Just make it non-nil for now
		})
	}

	return sigs
}

func defaultMsg() *MsgBuilder {
	return new(MsgBuilder).Default()
}

func defaultAggVote() *AggVoteBuilder {
	return new(AggVoteBuilder).Default()
}
