// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleGen } from "multiproof/src/MerkleGen.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { MerkleProof } from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import { XBlockMerkleProof } from "src/libraries/XBlockMerkleProof.sol";
import { PortalHarness } from "test/xchain/nomina/common/PortalHarness.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title XSubGen
 * @dev A utility contract for generating XSubmissions and validator sets.
 */
contract XSubGen is Test {
    string constant mnemonic = "test test test test test test test test test test test junk";

    uint64 constant baseValPower = 100;
    uint64 constant genesisValSetId = 1;

    struct Validator {
        address addr;
        uint64 power;
        uint256 pk;
    }

    mapping(uint64 => Validator[]) public valset;

    uint32 public valsetCount;

    PortalHarness public portal;

    constructor() {
        _initGenesisVals();
    }

    /// @dev Set the NominaPortal contract address
    function setPortal(address portalAddr) public {
        portal = PortalHarness(portalAddr);
    }

    /// @dev Make a mock xheader from source chain id and conf level
    function makeXHeader(uint64 sourceChainId, uint8 confLevel) public view returns (XTypes.BlockHeader memory) {
        require(address(portal) != address(0), "portal not set");

        return XTypes.BlockHeader({
            sourceChainId: sourceChainId,
            consensusChainId: portal.nominaCChainId(),
            confLevel: confLevel,
            offset: 1,
            sourceBlockHeight: 1234,
            sourceBlockHash: keccak256("abc")
        });
    }

    /// @dev Make an xsubmission signed by the genesis valset, with the given xheader and xmsgs.
    function makeXSub(uint64 valSetId, XTypes.BlockHeader memory xheader, XTypes.Msg[] memory msgs)
        public
        view
        returns (XTypes.Submission memory)
    {
        bool[] memory msgFlags = new bool[](msgs.length);
        for (uint256 i = 0; i < msgs.length; i++) {
            msgFlags[i] = true;
        }

        return makeXSub(valSetId, xheader, msgs, msgFlags);
    }

    /// @dev Make an xsubmission signed by `valSetId`, with the given xheader and selected xmsgs.
    function makeXSub(
        uint64 valSetId,
        XTypes.BlockHeader memory xheader,
        XTypes.Msg[] memory msgs,
        bool[] memory msgFlags
    ) public view returns (XTypes.Submission memory) {
        require(msgs.length == msgFlags.length, "msg length must match msgFlags length");

        _sortXMsgs(msgs, msgFlags);

        (XTypes.Msg[] memory selectedMsgs, uint256[] memory selectedIndices) = _getSelected(msgs, msgFlags);

        (bytes32[] memory msgProof, bool[] memory msgProofFlags, bytes32 msgRoot) =
            MerkleGen.generateMultiproof(_msgLeaves(msgs), selectedIndices);

        bytes32[] memory rootProof = new bytes32[](1);
        rootProof[0] = msgRoot;

        bytes32 root = MerkleProof.processProof(rootProof, _blockHeaderLeaf(xheader));

        XTypes.SigTuple[] memory sigs = _getSignatures(valSetId, root);

        return XTypes.Submission({
            attestationRoot: root,
            validatorSetId: valSetId,
            blockHeader: xheader,
            msgs: selectedMsgs,
            proof: msgProof,
            proofFlags: msgProofFlags,
            signatures: sigs
        });
    }

    /// @dev Set msgFlags according to destChainId
    function msgFlagsForDest(XTypes.Msg[] memory msgs, uint64 destChainId) public pure returns (bool[] memory) {
        bool[] memory msgFlags = new bool[](msgs.length);
        for (uint256 i; i < msgs.length; ++i) {
            if (msgs[i].destChainId == destChainId) msgFlags[i] = true;
        }
        return msgFlags;
    }

    /// @dev Get the validator set for a given valSetId
    function getVals(uint64 valSetId) public view returns (XTypes.Validator[] memory) {
        XTypes.Validator[] memory vals = new XTypes.Validator[](valset[valSetId].length);

        for (uint256 i = 0; i < valset[valSetId].length; i++) {
            vals[i] = XTypes.Validator(valset[valSetId][i].addr, valset[valSetId][i].power);
        }

        return vals;
    }

    /// @dev Add validators to a given valSetId
    function addVals(uint64 valSetId, Validator[] memory vals) public {
        for (uint256 i = 0; i < vals.length; i++) {
            valset[valSetId].push(Validator(vals[i].addr, vals[i].power, vals[i].pk));
        }
    }

    /// @dev Sign a digest with a private key
    function _sign(bytes32 digest, uint256 pk) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    /// @dev Initialize genesis validators
    function _initGenesisVals() internal {
        Validator memory v1;
        Validator memory v2;
        Validator memory v3;
        Validator memory v4;
        Validator memory v5;

        {
            (address val1, uint256 val1PK) = deriveRememberKey(mnemonic, 0);
            (address val2, uint256 val2Pk) = deriveRememberKey(mnemonic, 1);
            (address val3, uint256 val3Pk) = deriveRememberKey(mnemonic, 2);
            (address val4, uint256 val4Pk) = deriveRememberKey(mnemonic, 3);
            (address val5, uint256 val5Pk) = deriveRememberKey(mnemonic, 4);

            v1 = Validator(val1, baseValPower, val1PK);
            v2 = Validator(val2, baseValPower, val2Pk);
            v3 = Validator(val3, baseValPower, val3Pk);
            v4 = Validator(val4, baseValPower, val4Pk);
            v5 = Validator(val5, baseValPower, val5Pk);
        }

        // only use 1-4 for val set 1
        Validator[] storage genVals = valset[genesisValSetId];
        genVals.push(v1);
        genVals.push(v2);
        genVals.push(v3);
        genVals.push(v4);
        ++valsetCount;

        // val set 2 adds one validator, and removes val2
        Validator[] storage valSet2 = valset[genesisValSetId + 1];
        valSet2.push(v1);
        valSet2.push(v3);
        valSet2.push(v4);
        valSet2.push(v5);
        ++valsetCount;
    }

    /// @dev Sort sigs by validator address. NominaPortal.xsubmit expects sigs to be sorted.
    function _sortSigs(XTypes.SigTuple[] memory sigs) internal pure {
        uint256 n = sigs.length;
        for (uint256 i = 0; i < n - 1; i++) {
            for (uint256 j = 0; j < n - i - 1; j++) {
                if (sigs[j].validatorAddr > sigs[j + 1].validatorAddr) {
                    (sigs[j], sigs[j + 1]) = (sigs[j + 1], sigs[j]);
                }
            }
        }
    }

    /// @dev Sort xmsgs by offset. XMsgs within an xblock are order by log index,
    ///      which implicity orders them by offset (ascending).
    function _sortXMsgs(XTypes.Msg[] memory msgs, bool[] memory msgFlags) internal pure {
        uint256 n = msgs.length;
        for (uint256 i = 0; i < n - 1; i++) {
            for (uint256 j = 0; j < n - i - 1; j++) {
                if (msgs[j].offset > msgs[j + 1].offset) {
                    (msgs[j], msgs[j + 1]) = (msgs[j + 1], msgs[j]);
                    (msgFlags[j], msgFlags[j + 1]) = (msgFlags[j + 1], msgFlags[j]);
                }
            }
        }
    }

    /// @dev Generate a SigTuple array for a given valSetId and digest
    function _getSignatures(uint64 valSetId, bytes32 digest) internal view returns (XTypes.SigTuple[] memory sigs) {
        Validator[] storage vals = valset[valSetId];
        sigs = new XTypes.SigTuple[](vals.length);

        for (uint256 i = 0; i < vals.length; i++) {
            sigs[i] = XTypes.SigTuple(vals[i].addr, _sign(digest, vals[i].pk));
        }

        _sortSigs(sigs);
    }

    /// @dev For given msgs and msgFlags, return selected msgs and their indices
    function _getSelected(XTypes.Msg[] memory msgs, bool[] memory msgFlags)
        internal
        pure
        returns (XTypes.Msg[] memory, uint256[] memory)
    {
        uint256 numSelected = 0;
        for (uint256 i = 0; i < msgFlags.length; i++) {
            if (msgFlags[i]) {
                numSelected++;
            }
        }

        XTypes.Msg[] memory selectedMsgs = new XTypes.Msg[](numSelected);
        uint256[] memory selectedIndices = new uint256[](numSelected);
        uint256 j = 0;
        for (uint256 i = 0; i < msgFlags.length; i++) {
            if (msgFlags[i]) {
                selectedMsgs[j] = msgs[i];
                selectedIndices[j] = i;
                j++;
            }
        }

        return (selectedMsgs, selectedIndices);
    }

    /// @dev Convert xmsgs to leaf hashes
    function _msgLeaves(XTypes.Msg[] memory msgs) private pure returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](msgs.length);

        for (uint256 i = 0; i < msgs.length; i++) {
            leaves[i] = _leafHash(XBlockMerkleProof.DST_XMSG, abi.encode(msgs[i]));
        }

        return leaves;
    }

    /// @dev Convert xblock header to leaf hash
    function _blockHeaderLeaf(XTypes.BlockHeader memory blockHeader) private pure returns (bytes32) {
        return _leafHash(XBlockMerkleProof.DST_XBLOCK_HEADER, abi.encode(blockHeader));
    }

    /// @dev Double hash leaves, as recommended by OpenZeppelin, to prevent second preimage attacks
    ///      Leaves must be double hashed in tree / proof construction
    ///      Callers must specify the domain separation tag of the leaf, which will be hashed in
    function _leafHash(uint8 dst, bytes memory leaf) private pure returns (bytes32) {
        return keccak256(bytes.concat(keccak256(abi.encodePacked(dst, leaf))));
    }
}
