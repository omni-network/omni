// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ERC721 } from "solady/src/tokens/ERC721.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";

import { MerkleProofLib } from "solady/src/utils/MerkleProofLib.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { IOmniGasPump } from "core/src/interfaces/IOmniGasPump.sol";

contract Tag is ERC721, XAppBase, Ownable {
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           ERRORS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    error NoSelfTag();
    error NotAHolder();
    error ZeroAmount();
    error ZeroAddress();
    error TagCooldown();
    error TagDisabled();
    error InvalidProof();
    error TooManyMints();
    error DisabledMint();
    error CrosschainOnly();
    error AlreadyClaimed();
    error TransferFailed();
    error InsufficientPayment();

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           EVENTS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    event GameEnded();
    event RootUpdated(bytes32 root);
    event MintEnabled();
    event MintDisabled();
    event MintFinalized();
    event TagDelaySet(uint16 newDelay);
    event TagCooldownSet(uint16 newCooldown);

    event TagInitiated(uint64 indexed srcChainId, address indexed tagger, uint256 indexed tokenId, uint32 timestamp);
    event TagProcessed(address indexed tagger, uint256 indexed tokenId, uint32 timestamp);
    event TagBroadcasted(uint64 indexed destChainId, address indexed tagger, uint256 indexed tokenId, uint32 timestamp);
    event CrosschainSend(uint64 indexed destChainId, address indexed from, address to, uint256 indexed tokenId);
    event CrosschainReceive(uint64 indexed srcChainId, address from, address indexed to, uint256 indexed tokenId);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    struct PendingTag {
        uint32 fromTokenId;
        uint32 timestamp;
        address tagger;
    }

    uint64 private constant OMNI_CHAIN_ID = 166;
    uint64 private constant TAG_GAS_LIMIT = 100_000;
    uint64 private constant MINT_GAS_LIMIT = 100_000;
    uint64 private constant TRANSFER_GAS_LIMIT = 100_000;
    uint64 private constant WHITELIST_MINT_GAS_LIMIT = 100_000;

    IOmniGasPump public immutable omniGasPump;

    string private _name;
    string private _symbol;

    uint32 private _totalSupply;

    bool public tagEnabled;
    bool public mintEnabled;
    bool public mintFinalized;
    uint8 public maxMintsPerWallet;
    uint16 public tagDelay;
    uint16 public tagCooldown;
    uint64 public price;

    bytes32 public root;
    mapping(uint256 tokenId => PendingTag[]) public tagQueue;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        CONSTRUCTOR                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    constructor(address owner_, address omni_, address omniGasPump_, string memory name_, string memory symbol_)
        payable
    {
        if (owner_ == address(0) || omni_ == address(0) || omniGasPump_ == address(0)) revert ZeroAddress();

        _initializeOwner(owner_);
        _setOmniPortal(omni_);
        omniGasPump = IOmniGasPump(omniGasPump_);

        _name = name_;
        _symbol = symbol_;
        tagEnabled = true;
        maxMintsPerWallet = 5;
        tagDelay = 2 hours;
        tagCooldown = 1 minutes;
        price = 1 ether;

        if (block.chainid != OMNI_CHAIN_ID) {
            mintFinalized = true;
            emit MintFinalized();
        }
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       VIEW FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function name() public view override returns (string memory) {
        return _name;
    }

    function symbol() public view override returns (string memory) {
        return _symbol;
    }

    function tokenURI(uint256 /* tokenId */ ) public pure override returns (string memory) {
        return ""; // TODO: Implement SVG renderer
    }

    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    function getWhitelistClaimStatus(address minter) public view returns (bool) {
        return _hasClaimedWhitelist(minter);
    }

    function getMintCount(address minter) public view returns (uint224) {
        return _getMintCount(minter);
    }

    function getTagCursor(uint256 tokenId) public view returns (uint16) {
        return _getTagCursor(tokenId);
    }

    function getLastTagTimestamp(uint256 tokenId) public view returns (uint32) {
        return _getLastTagTimestamp(tokenId);
    }

    function getTagPoints(uint256 tokenId) public view returns (uint16) {
        return _getTagPoints(tokenId);
    }

    function getTagCount(uint256 tokenId) public view returns (uint32) {
        return _getTagCount(tokenId);
    }

    function getAuxData(address addr) public view returns (bool whitelistClaimed, uint224 mintCount) {
        // Current layout (224 bits total):
        // [whitelistClaimed (1 bit)][mintCount (223 bits)]
        uint224 auxData = _getAux(addr);
        whitelistClaimed = (auxData & (1 << 223)) != 0;
        mintCount = uint224(auxData & ((1 << 223) - 1));
    }

    function getExtraData(uint256 tokenId)
        public
        view
        returns (uint16 cursor, uint32 timestamp, uint16 tagPoints, uint32 tagCount)
    {
        // Current layout (96 bits total):
        // [cursor (16 bits)][timestamp (32 bits)][tagPoints (16 bits)][tagCount (32 bits)]
        uint96 extraData = _getExtraData(tokenId);
        cursor = uint16(extraData >> 80);
        timestamp = uint32((extraData >> 48) & ((1 << 32) - 1));
        tagPoints = uint16((extraData >> 32) & ((1 << 16) - 1));
        tagCount = uint32(extraData & ((1 << 32) - 1));
    }

    function getPendingTagQueueLength(uint256 tokenId) public view returns (uint256) {
        uint96 cursor = uint16(_getExtraData(tokenId) >> 80);
        return tagQueue[tokenId].length - cursor;
    }

    function feeForCrosschainMint(uint256 amount) public view returns (uint256) {
        uint256 requiredEth = omniGasPump.quote(price * amount);
        uint256 pumpFee = omniGasPump.xfee();
        bytes memory data = abi.encodeCall(
            this.processCrosschainMint, (address(type(uint160).max), type(uint256).max, address(type(uint160).max))
        );
        return requiredEth + pumpFee + feeFor(OMNI_CHAIN_ID, data, MINT_GAS_LIMIT);
    }

    function feeForCrosschainMintWhitelist(uint256 amount, bytes32[] calldata proof) public view returns (uint256) {
        uint256 requiredEth = omniGasPump.quote(price * (amount - 1));
        uint256 pumpFee = omniGasPump.xfee();
        bytes memory data = abi.encodeCall(
            this.processCrosschainMintWhitelist,
            (address(type(uint160).max), type(uint256).max, proof, address(type(uint160).max))
        );
        return requiredEth + pumpFee + feeFor(OMNI_CHAIN_ID, data, WHITELIST_MINT_GAS_LIMIT);
    }

    function feeForCrosschainTransfer(uint64 destChainId) public view returns (uint256) {
        bytes memory data = abi.encodeCall(
            this.processCrosschainTransfer,
            (address(type(uint160).max), address(type(uint160).max), type(uint256).max, type(uint96).max)
        );
        return feeFor(destChainId, data, TRANSFER_GAS_LIMIT);
    }

    function feeForCrosschainTag(uint64 destChainId) public view returns (uint256) {
        bytes memory data = abi.encodeCall(
            this.processCrosschainTag,
            (
                type(uint256).max,
                PendingTag({
                    fromTokenId: type(uint32).max,
                    timestamp: uint32(block.timestamp),
                    tagger: address(type(uint160).max)
                })
            )
        );
        return feeFor(destChainId, data, TAG_GAS_LIMIT);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       MINT FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function mint(address to, uint256 amount) public payable {
        _validateMint(to, amount, false, false);
        _processMint(to, amount, false);
    }

    function mintWhitelist(address to, uint256 amount, bytes32[] calldata proof) public {
        _validateWhitelist(msg.sender, proof, false);
        _validateMint(to, amount, true, false);
        _processMint(to, amount, true);
    }

    function crosschainMint(address to, uint256 amount, address refundTo) public payable {
        if (block.chainid == OMNI_CHAIN_ID) revert CrosschainOnly();
        if (to == address(0) || refundTo == address(0)) revert ZeroAddress();
        if (amount == 0) revert ZeroAmount();

        // Handle gas pump fill first
        uint256 requiredEth = omniGasPump.quote(price * amount);
        uint256 pumpFee = omniGasPump.xfee();
        omniGasPump.fillUp{ value: requiredEth + pumpFee }(address(this));

        // Process the mint afterwards to ensure contract can refund if needed
        bytes memory data = abi.encodeCall(this.processCrosschainMint, (to, amount, refundTo));
        uint256 fee = xcall(OMNI_CHAIN_ID, ConfLevel.Latest, address(this), data, MINT_GAS_LIMIT);

        // Refund if necessary
        uint256 totalRequired = requiredEth + pumpFee + fee;
        if (msg.value < totalRequired) revert InsufficientPayment();
        if (msg.value > totalRequired) {
            (bool success,) = payable(msg.sender).call{ value: msg.value - totalRequired }("");
            if (!success) revert TransferFailed();
        }
    }

    function crosschainMintWhitelist(address to, uint256 amount, bytes32[] calldata proof, address refundTo)
        public
        payable
    {
        if (block.chainid == OMNI_CHAIN_ID) revert CrosschainOnly();
        if (to == address(0) || refundTo == address(0)) revert ZeroAddress();
        if (amount == 0) revert ZeroAmount();

        uint256 totalRequired;
        if (amount > 1) {
            // Handle gas pump fill first
            uint256 requiredEth = omniGasPump.quote(price * (amount - 1));
            uint256 pumpFee = omniGasPump.xfee();
            omniGasPump.fillUp{ value: requiredEth + pumpFee }(address(this));

            unchecked {
                totalRequired += requiredEth + pumpFee;
            }
        }

        bytes memory data = abi.encodeCall(this.processCrosschainMintWhitelist, (to, amount, proof, refundTo));
        uint256 fee = xcall(OMNI_CHAIN_ID, ConfLevel.Latest, address(this), data, WHITELIST_MINT_GAS_LIMIT);

        // Refund if necessary
        unchecked {
            totalRequired = totalRequired + fee;
        }
        if (msg.value < totalRequired) revert InsufficientPayment();
        if (msg.value > totalRequired) {
            (bool success,) = payable(msg.sender).call{ value: msg.value - totalRequired }("");
            if (!success) revert TransferFailed();
        }
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                   INTERACTIVE FUNCTIONS                    */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function updateStats(uint256[] calldata tokenIds, uint256 iterations) public {
        for (uint256 i; i < tokenIds.length; ++i) {
            _processTagQueue(tokenIds[i], iterations);
        }
    }

    function tag(uint256 fromTokenId, uint256 toTokenId) public {
        if (!tagEnabled) revert TagDisabled();
        if (_ownerOf(fromTokenId) != msg.sender) revert Unauthorized();
        if (!_exists(toTokenId)) revert TokenDoesNotExist();
        if (fromTokenId == toTokenId) revert NoSelfTag();
        if (_getLastTagTimestamp(fromTokenId) + tagCooldown > block.timestamp) revert TagCooldown();

        PendingTag memory pendingTag =
            PendingTag({ fromTokenId: uint32(fromTokenId), timestamp: uint32(block.timestamp), tagger: msg.sender });
        tagQueue[toTokenId].push(pendingTag);

        _setLastTagTimestamp(fromTokenId);

        emit TagInitiated(uint64(block.chainid), msg.sender, toTokenId, uint32(block.timestamp));
    }

    function crosschainTag(uint256 fromTokenId, uint256 toTokenId, uint64 destChainId) public payable {
        if (!tagEnabled) revert TagDisabled();
        if (_ownerOf(fromTokenId) != msg.sender) revert Unauthorized();
        if (fromTokenId == toTokenId) revert NoSelfTag();
        if (_getLastTagTimestamp(fromTokenId) + tagCooldown > block.timestamp) revert TagCooldown();

        PendingTag memory pendingTag =
            PendingTag({ fromTokenId: uint32(fromTokenId), timestamp: uint32(block.timestamp), tagger: msg.sender });

        bytes memory data = abi.encodeCall(this.processCrosschainTag, (toTokenId, pendingTag));
        uint256 fee = xcall(destChainId, ConfLevel.Latest, address(this), data, TAG_GAS_LIMIT);

        _setLastTagTimestamp(fromTokenId);

        if (msg.value < fee) revert InsufficientPayment();
        if (msg.value > fee) {
            (bool success,) = payable(msg.sender).call{ value: msg.value - fee }("");
            if (!success) revert TransferFailed();
        }

        emit TagBroadcasted(destChainId, msg.sender, toTokenId, uint32(block.timestamp));
    }

    function crosschainTransfer(uint64 destChainId, address to, uint256 tokenId) public payable {
        _burn(msg.sender, tokenId);

        // Get the extraData but strip only the cursor
        uint96 extraData = _getExtraData(tokenId);
        uint96 preservedCursor = extraData & (~uint96((1 << 80) - 1)); // Keep cursor (top 16 bits)
        uint96 transferData = extraData & ((1 << 80) - 1); // Send everything else (timestamp, points, count)

        bytes memory data = abi.encodeCall(this.processCrosschainTransfer, (msg.sender, to, tokenId, transferData));
        uint256 fee = xcall(destChainId, ConfLevel.Latest, address(this), data, TRANSFER_GAS_LIMIT);

        // Reset everything except the cursor
        _setExtraData(tokenId, preservedCursor);

        if (msg.value < fee) revert InsufficientPayment();
        if (msg.value > fee) {
            (bool success,) = payable(msg.sender).call{ value: msg.value - fee }("");
            if (!success) revert TransferFailed();
        }

        emit CrosschainSend(destChainId, msg.sender, to, tokenId);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      PORTAL FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function processCrosschainMint(address to, uint256 amount, address refundTo) public xrecv {
        if (xmsg.sender != address(this)) revert Unauthorized();

        // Determine how many NFTs can be minted
        uint256 validAmount;
        bool validMint;
        for (uint256 i; i < amount; ++i) {
            validMint = _validateMint(to, amount - i, false, true);
            if (validMint) {
                validAmount = amount - i;
                break;
            }
        }

        if (!validMint) {
            // Refund if no NFTs can be minted
            (bool success,) = payable(refundTo).call{ value: amount * price }("");
            if (!success) revert TransferFailed();
        } else if (validAmount < amount) {
            // Refund for NFTs that couldn't be minted
            (bool success,) = payable(refundTo).call{ value: (amount - validAmount) * price }("");
            if (!success) revert TransferFailed();
        }

        // Only mint the valid amount
        if (validMint) _processMint(to, validAmount, false);
    }

    function processCrosschainMintWhitelist(address to, uint256 amount, bytes32[] calldata proof, address refundTo)
        public
        xrecv
    {
        if (xmsg.sender != address(this)) revert Unauthorized();

        bool whitelistAvailable = _validateWhitelist(msg.sender, proof, true);
        if (!whitelistAvailable) {
            // If whitelist claim isn't available, process as normal mint
            processCrosschainMint(to, amount - 1, refundTo);
            return;
        }

        // Determine how many NFTs can be minted
        uint256 validAmount;
        bool validMint;
        for (uint256 i; i < amount; ++i) {
            validMint = _validateMint(to, amount - i, true, true);
            if (validMint) {
                validAmount = amount - i;
                break;
            }
        }

        if (!validMint) {
            // Refund if no paid for NFTs can be minted
            if (amount > 1) {
                (bool success,) = payable(refundTo).call{ value: (amount - 1) * price }("");
                if (!success) revert TransferFailed();
            }
        } else if (validAmount < amount && validAmount > 1) {
            // Refund for NFTs that couldn't be minted
            (bool success,) = payable(refundTo).call{ value: (amount - validAmount - 1) * price }("");
            if (!success) revert TransferFailed();
        }

        // Only mint the valid amount
        if (validMint) _processMint(to, validAmount, true);
    }

    function processCrosschainTag(uint256 tokenId, PendingTag memory pendingTag) public xrecv {
        if (xmsg.sender != address(this)) revert Unauthorized();

        if (_exists(tokenId)) {
            pendingTag.timestamp = uint32(block.timestamp);
            tagQueue[tokenId].push(pendingTag);
            emit TagInitiated(xmsg.sourceChainId, pendingTag.tagger, tokenId, pendingTag.timestamp);
        } else {
            revert TokenDoesNotExist();
        }
    }

    function processCrosschainTransfer(address from, address to, uint256 tokenId, uint96 extraData) public xrecv {
        if (xmsg.sender != address(this)) revert Unauthorized();

        _updateExtraData(tokenId, extraData);
        _mint(to, tokenId);

        emit CrosschainReceive(xmsg.sourceChainId, from, to, tokenId);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ADMIN FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function enableMint() public onlyOwner {
        if (mintFinalized) revert DisabledMint();
        mintEnabled = true;
        emit MintEnabled();
    }

    function disableMint(bool finalize) public onlyOwner {
        if (mintFinalized) revert DisabledMint();
        mintEnabled = false;
        if (finalize) {
            mintFinalized = true;
            emit MintFinalized();
        } else {
            emit MintDisabled();
        }
    }

    function updateRoot(bytes32 newRoot) public onlyOwner {
        root = newRoot;
        emit RootUpdated(newRoot);
    }

    function setTagCooldown(uint16 newCooldown) public onlyOwner {
        tagCooldown = newCooldown;
        emit TagCooldownSet(newCooldown);
    }

    function setTagDelay(uint16 newDelay) public onlyOwner {
        tagDelay = newDelay;
        emit TagDelaySet(newDelay);
    }

    function endGame() public onlyOwner {
        tagEnabled = false;
        emit GameEnded();
    }

    function withdraw(address to) public onlyOwner {
        (bool success,) = to.call{ value: address(this).balance }("");
        if (!success) revert TransferFailed();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     INTERNAL FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function _hasClaimedWhitelist(address minter) public view returns (bool) {
        // Check if the highest bit is set
        return (_getAux(minter) & (1 << 223)) != 0;
    }

    function _setWhitelistClaimed(address minter) private {
        uint224 auxData = _getAux(minter);
        // Set the highest bit while preserving the count
        _setAux(minter, auxData | (1 << 223));
    }

    function _getMintCount(address minter) private view returns (uint224) {
        // Mask off the highest bit (whitelist) to get just the count
        return uint224(_getAux(minter) & ((1 << 223) - 1));
    }

    function _getTagCursor(uint256 tokenId) private view returns (uint16) {
        // Shift right by 80 bits to get just the cursor
        return uint16(_getExtraData(tokenId) >> 80);
    }

    function _getLastTagTimestamp(uint256 fromTokenId) private view returns (uint32) {
        // Shift right by 48 bits and mask to get just the timestamp (32 bits)
        return uint32((_getExtraData(fromTokenId) >> 48) & ((1 << 32) - 1));
    }

    function _setLastTagTimestamp(uint256 fromTokenId) private {
        uint96 extraData = _getExtraData(fromTokenId);
        uint16 cursor = uint16(extraData >> 80);
        uint16 tagPoints = uint16((extraData >> 32) & ((1 << 16) - 1));
        uint32 tagCount = uint32(extraData & ((1 << 32) - 1));

        // Update timestamp while preserving cursor, points, and count
        _setExtraData(
            fromTokenId, (uint96(cursor) << 80) | (uint96(block.timestamp) << 48) | (uint96(tagPoints) << 32) | tagCount
        );
    }

    function _getTagPoints(uint256 tokenId) private view returns (uint16) {
        // Shift right by 32 and mask to get the points (16 bits)
        return uint16((_getExtraData(tokenId) >> 32) & ((1 << 16) - 1));
    }

    function _incrementTagPoints(uint256 tokenId) private {
        uint96 extraData = _getExtraData(tokenId);
        uint16 cursor = uint16(extraData >> 80);
        uint32 timestamp = uint32((extraData >> 48) & ((1 << 32) - 1));
        uint16 tagPoints = uint16((extraData >> 32) & ((1 << 16) - 1));
        uint32 tagCount = uint32(extraData & ((1 << 32) - 1));

        if (tagPoints == type(uint16).max) return;

        unchecked {
            ++tagPoints;
        }

        _setExtraData(
            tokenId, (uint96(cursor) << 80) | (uint96(timestamp) << 48) | (uint96(tagPoints) << 32) | tagCount
        );
    }

    function _getTagCount(uint256 tokenId) private view returns (uint32) {
        // Mask off all but the bottom 32 bits
        return uint32(_getExtraData(tokenId) & ((1 << 32) - 1));
    }

    function _validateWhitelist(address minter, bytes32[] calldata proof, bool softFail) private view returns (bool) {
        bytes32 leaf = keccak256(abi.encode(minter));

        if (softFail) {
            if (_hasClaimedWhitelist(minter)) return false;
            if (!MerkleProofLib.verifyCalldata(proof, root, leaf)) return false;
        } else {
            if (_hasClaimedWhitelist(minter)) revert AlreadyClaimed();
            if (!MerkleProofLib.verifyCalldata(proof, root, leaf)) revert InvalidProof();
        }
        return true;
    }

    function _validateMint(address to, uint256 amount, bool whitelistClaimed, bool softFail)
        private
        view
        returns (bool)
    {
        uint224 mintCount = getMintCount(msg.sender);

        if (softFail) {
            if (!mintEnabled || mintFinalized) return false;
            if (to == address(0)) return false;
            if (amount == 0) return false;
            if (mintCount + amount > maxMintsPerWallet) return false;
            if (whitelistClaimed) --amount;
            if (msg.value < price * amount) return false;
        } else {
            if (!mintEnabled || mintFinalized) revert DisabledMint();
            if (to == address(0)) revert ZeroAddress();
            if (amount == 0) revert ZeroAmount();
            if (mintCount + amount > maxMintsPerWallet) revert TooManyMints();
            if (whitelistClaimed) --amount;
            if (msg.value < price * amount) revert InsufficientPayment();
        }
        return true;
    }

    function _processMint(address to, uint256 amount, bool whitelistClaimed) private {
        uint256 tokenId = _totalSupply;
        uint224 auxData = _getAux(msg.sender);
        uint224 mintCount = auxData & ((1 << 223) - 1); // Use 223 bits for mintCount
        uint224 whitelistStatus = auxData & (1 << 223);

        for (uint256 i; i < amount; ++i) {
            unchecked {
                ++tokenId;
                ++mintCount;
            }
            _mint(to, tokenId);
        }

        unchecked {
            _totalSupply += uint32(amount);
            _setAux(msg.sender, mintCount | whitelistStatus);
        }

        if (whitelistClaimed) {
            _setWhitelistClaimed(msg.sender);
        }

        uint256 billableAmount = price * (whitelistClaimed ? amount - 1 : amount);
        if (msg.value > billableAmount) {
            (bool success,) = payable(msg.sender).call{ value: msg.value - billableAmount }("");
            if (!success) revert TransferFailed();
        }
    }

    function _processTagQueue(uint256 tokenId, uint256 iterations) private {
        uint96 extraData = _getExtraData(tokenId);
        uint32 tagCount = uint32(extraData & ((1 << 32) - 1));
        uint16 tagPoints = uint16((extraData >> 32) & ((1 << 16) - 1));
        uint32 timestamp = uint32((extraData >> 48) & ((1 << 32) - 1));
        uint16 cursor = uint16(extraData >> 80);

        if (cursor == type(uint16).max) return;

        // Decide how many we want to process
        uint256 batchSize = iterations == 0 ? 100 : iterations;
        uint256 end = cursor + batchSize;
        uint256 totalLength = tagQueue[tokenId].length;
        uint16 newCursor = cursor;

        // Make sure 'end' does not exceed queue length
        if (end > totalLength) end = totalLength;

        for (uint256 i = cursor; i < end; ++i) {
            PendingTag memory pendingTag = tagQueue[tokenId][i];
            if (pendingTag.timestamp + tagDelay < block.timestamp) {
                unchecked {
                    ++tagCount;
                    ++newCursor;
                    _incrementTagPoints(pendingTag.fromTokenId);
                }
                emit TagProcessed(pendingTag.tagger, tokenId, pendingTag.timestamp);
            } else {
                break;
            }
        }

        _setExtraData(
            tokenId, (uint96(newCursor) << 80) | (uint96(timestamp) << 48) | (uint96(tagPoints) << 32) | tagCount
        );
    }

    function _updateExtraData(uint256 tokenId, uint96 inboundExtraData) private {
        // Get our local extraData to preserve the cursor and add points
        uint96 extraData = _getExtraData(tokenId);
        uint96 preservedCursor = extraData & (~uint96((1 << 80) - 1));
        uint16 localPoints = uint16((extraData >> 32) & ((1 << 16) - 1));

        // Extract points from incoming data
        uint16 incomingPoints = uint16((inboundExtraData >> 32) & ((1 << 16) - 1));

        // Add points together (unchecked to allow overflow)
        uint16 combinedPoints;
        unchecked {
            combinedPoints = localPoints + incomingPoints;
        }
        if (combinedPoints < incomingPoints) combinedPoints = type(uint16).max;

        // Reconstruct extraData with combined points
        uint96 newExtraData = (extraData & ~uint96((1 << 48) - (1 << 32))) // Clear points section
            | (uint96(combinedPoints) << 32); // Insert combined points

        _setExtraData(tokenId, preservedCursor | newExtraData);
    }

    function _beforeTokenTransfer(address, address, uint256 tokenId) internal override {
        _processTagQueue(tokenId, 0);
    }

    receive() external payable { }
}
