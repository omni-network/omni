// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ERC721 } from "solady/src/tokens/ERC721.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";

import { MerkleProofLib } from "solady/src/utils/MerkleProofLib.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";

contract Tag is ERC721, XAppBase, Ownable {
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           ERRORS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    error NotAHolder();
    error ZeroAmount();
    error ZeroAddress();
    error TagCooldown();
    error InvalidProof();
    error TooManyMints();
    error DisabledMint();
    error AlreadyClaimed();
    error InsufficientPayment();

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           EVENTS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    event RootUpdated(bytes32 root);
    event MintEnabled();
    event MintDisabled();
    event MintFinalized();

    event TagInitiated(uint64 indexed srcChainId, address indexed tagger, uint256 indexed tokenId, uint40 timestamp);
    event TagProcessed(address indexed tagger, uint256 indexed tokenId, uint40 timestamp);
    event TagBroadcasted(uint64 indexed destChainId, address indexed tagger, uint256 indexed tokenId, uint40 timestamp);
    event CrosschainSend(uint64 indexed destChainId, address indexed from, address to, uint256 indexed tokenId);
    event CrosschainReceive(uint64 indexed srcChainId, address from, address indexed to, uint256 indexed tokenId);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    struct PendingTag {
        uint32 fromTokenId;
        uint40 timestamp;
        address tagger;
    }

    string private _name;
    string private _symbol;

    uint32 private _totalSupply;

    bool public mintEnabled;
    bool public mintFinalized;
    uint8 public maxMintsPerWallet;
    uint16 public tagDelay;
    uint16 public tagCooldown;
    uint64 public price;

    bytes32 public root;
    mapping(uint256 tokenId => PendingTag[]) public tagQueue;
    mapping(uint256 tokenId => uint24 points) public accruedPoints;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        CONSTRUCTOR                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    constructor(address owner_, address omni_, string memory name_, string memory symbol_) payable {
        _initializeOwner(owner_);
        _setOmniPortal(omni_);

        _name = name_;
        _symbol = symbol_;
        maxMintsPerWallet = 5;
        tagDelay = 2 hours;
        tagCooldown = 1 minutes;
        price = 1 ether;

        if (block.chainid != 166) {
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

    function hasClaimedWhitelist(address minter) public view returns (bool) {
        return (_getAux(minter) & (1 << 223)) != 0;
    }

    function getMintCount(address minter) public view returns (uint224) {
        // Mask off the highest bit (whitelist) to get just the count
        return uint224(_getAux(minter) & ((1 << 223) - 1));
    }

    function getTagCursor(uint256 tokenId) public view returns (uint16) {
        // Shift right by 80 bits to get just the cursor
        return uint16(_getExtraData(tokenId) >> 80);
    }

    function getLastTagTimestamp(uint256 tokenId) public view returns (uint40) {
        // Shift right by 40 bits and mask to get just the timestamp
        return uint40((_getExtraData(tokenId) >> 40) & ((1 << 40) - 1));
    }

    function getTagCount(uint256 tokenId) public view returns (uint40) {
        // Mask off the top 56 bits (cursor + timestamp) to get just the count
        return uint40(_getExtraData(tokenId) & ((1 << 40) - 1));
    }

    function getAuxData(address addr) public view returns (bool whitelistClaimed, uint224 mintCount) {
        uint224 auxData = _getAux(addr);
        whitelistClaimed = (auxData & (1 << 223)) != 0;
        mintCount = uint224(auxData & ((1 << 223) - 1));
    }

    function getExtraData(uint256 tokenId) public view returns (uint16 cursor, uint40 timestamp, uint40 tagCount) {
        uint96 extraData = _getExtraData(tokenId);
        cursor = uint16(extraData >> 80);
        timestamp = uint40((extraData >> 40) & ((1 << 40) - 1));
        tagCount = uint40(extraData & ((1 << 40) - 1));
    }

    function crosschainTransferFee(uint64 destChainId) public view returns (uint256) {
        bytes memory data = abi.encodeCall(
            this.processCrosschainTransfer,
            (address(type(uint160).max), address(type(uint160).max), type(uint256).max, type(uint96).max)
        );
        return feeFor(destChainId, data, 100_000);
    }

    function crosschainTagFee(uint64 destChainId) public view returns (uint256) {
        bytes memory data = abi.encodeCall(
            this.processCrosschainTag,
            (
                type(uint256).max,
                PendingTag({
                    fromTokenId: type(uint32).max,
                    timestamp: uint40(block.timestamp),
                    tagger: address(type(uint160).max)
                })
            )
        );
        return feeFor(destChainId, data, 100_000);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       MINT FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function mint(address to, uint256 amount) public payable {
        _validateMint(to, amount, false);
        _processMint(to, amount, false);
    }

    function mintWhitelist(address to, bytes32[] calldata proof) public {
        _validateWhitelist(msg.sender, proof);
        _validateMint(to, 1, true);
        _processMint(to, 1, true);
    }

    function mintWhitelistWithExtra(address to, uint256 amount, bytes32[] calldata proof) public {
        _validateWhitelist(msg.sender, proof);
        _validateMint(to, amount, true);
        _processMint(to, amount, true);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                   INTERACTIVE FUNCTIONS                    */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function updateStats(uint256[] calldata tokenIds) public {
        for (uint256 i; i < tokenIds.length; ++i) {
            _processTagQueue(tokenIds[i]);
        }
    }

    function tag(uint256 fromTokenId, uint256 toTokenId) public {
        if (_ownerOf(fromTokenId) != msg.sender) revert Unauthorized();
        if (!_exists(toTokenId)) revert TokenDoesNotExist();
        if (_getLastTagTimestamp(fromTokenId) + tagCooldown > block.timestamp) revert TagCooldown();

        PendingTag memory pendingTag =
            PendingTag({ fromTokenId: uint32(fromTokenId), timestamp: uint40(block.timestamp), tagger: msg.sender });
        tagQueue[toTokenId].push(pendingTag);

        _setLastTagTimestamp(fromTokenId);

        emit TagInitiated(uint64(block.chainid), msg.sender, toTokenId, uint40(block.timestamp));
    }

    function crosschainTag(uint256 fromTokenId, uint256 toTokenId, uint64 destChainId) public payable {
        if (_ownerOf(fromTokenId) != msg.sender) revert Unauthorized();
        if (_getLastTagTimestamp(fromTokenId) + tagCooldown > block.timestamp) revert TagCooldown();

        PendingTag memory pendingTag =
            PendingTag({ fromTokenId: uint32(fromTokenId), timestamp: uint40(block.timestamp), tagger: msg.sender });

        bytes memory data = abi.encodeCall(this.processCrosschainTag, (toTokenId, pendingTag));
        uint256 fee = xcall(destChainId, ConfLevel.Latest, address(this), data, 100_000);

        _setLastTagTimestamp(fromTokenId);

        if (msg.value < fee) revert InsufficientPayment();
        if (msg.value > fee) payable(msg.sender).transfer(msg.value - fee);

        emit TagBroadcasted(destChainId, msg.sender, toTokenId, uint40(block.timestamp));
    }

    function crosschainTransfer(uint64 destChainId, address to, uint256 tokenId) public payable {
        _burn(msg.sender, tokenId);

        // Get the extraData but strip only the cursor
        uint96 extraData = _getExtraData(tokenId);
        uint96 preservedCursor = extraData & (~uint96((1 << 80) - 1)); // Keep cursor (top 16 bits)
        uint80 transferData = uint80(extraData & ((1 << 80) - 1));

        bytes memory data = abi.encodeCall(this.processCrosschainTransfer, (msg.sender, to, tokenId, transferData));
        uint256 fee = xcall(destChainId, ConfLevel.Latest, address(this), data, 100_000);

        _setExtraData(tokenId, preservedCursor);

        if (msg.value < fee) revert InsufficientPayment();
        if (msg.value > fee) payable(msg.sender).transfer(msg.value - fee);

        emit CrosschainSend(destChainId, msg.sender, to, tokenId);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      PORTAL FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function processCrosschainTag(uint256 tokenId, PendingTag memory pendingTag) public xrecv {
        if (xmsg.sender != address(this)) revert Unauthorized();

        if (_exists(tokenId)) {
            pendingTag.timestamp = uint40(block.timestamp);
            tagQueue[tokenId].push(pendingTag);
            emit TagInitiated(xmsg.sourceChainId, pendingTag.tagger, tokenId, pendingTag.timestamp);
        } else {
            revert TokenDoesNotExist();
        }
    }

    function processCrosschainTransfer(address from, address to, uint256 tokenId, uint96 extraData) public xrecv {
        if (xmsg.sender != address(this)) revert Unauthorized();

        // Get our local extraData to preserve the cursor
        uint96 localExtraData = _getExtraData(tokenId);
        uint96 preservedCursor = localExtraData & (~uint96((1 << 80) - 1));

        // Apply the incoming timestamp and tagCount while preserving our local cursor
        _setExtraData(tokenId, preservedCursor | extraData);
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

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     INTERNAL FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function _getLastTagTimestamp(uint256 fromTokenId) private view returns (uint40) {
        uint96 extraData = _getExtraData(fromTokenId);
        return uint40((extraData >> 40) & ((1 << 40) - 1));
    }

    function _setLastTagTimestamp(uint256 fromTokenId) private {
        uint96 extraData = _getExtraData(fromTokenId);
        uint16 cursor = uint16(extraData >> 80);
        uint40 tagCount = uint40(extraData & ((1 << 40) - 1));

        // Update timestamp while preserving cursor and tagCount
        _setExtraData(fromTokenId, (uint96(cursor) << 80) | (uint96(block.timestamp) << 40) | tagCount);
    }

    function _hasClaimedWhitelist(address minter) public view returns (bool) {
        // Check if the highest bit is set
        return (_getAux(minter) & (1 << 223)) != 0;
    }

    function _setWhitelistClaimed(address minter) private {
        uint224 auxData = _getAux(minter);
        // Set the highest bit while preserving the count
        _setAux(minter, auxData | (1 << 223));
    }

    function _validateWhitelist(address minter, bytes32[] calldata proof) private view {
        bytes32 leaf = keccak256(abi.encode(minter));

        if (_hasClaimedWhitelist(minter)) revert AlreadyClaimed();
        if (!MerkleProofLib.verifyCalldata(proof, root, leaf)) revert InvalidProof();
    }

    function _validateMint(address to, uint256 amount, bool whitelistClaimed) private view {
        uint224 mintCount = getMintCount(msg.sender);

        if (!mintEnabled || mintFinalized) revert DisabledMint();
        if (to == address(0)) revert ZeroAddress();
        if (amount == 0) revert ZeroAmount();
        if (mintCount + amount > maxMintsPerWallet) revert TooManyMints();
        if (whitelistClaimed) --amount;
        if (msg.value < price * amount) revert InsufficientPayment();
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
    }

    function _processTagQueue(uint256 tokenId) private {
        uint96 extraData = _getExtraData(tokenId);
        uint40 tagCount = uint40(extraData & ((1 << 40) - 1));
        uint40 timestamp = uint40((extraData >> 40) & ((1 << 40) - 1));
        uint16 cursor = uint16(extraData >> 80);

        uint256 length = tagQueue[tokenId].length;
        uint16 newCursor = cursor;

        for (uint256 i = cursor; i < length; ++i) {
            PendingTag memory pendingTag = tagQueue[tokenId][i];
            if (pendingTag.timestamp + tagDelay < block.timestamp) {
                unchecked {
                    ++tagCount;
                    ++newCursor;
                    ++accruedPoints[pendingTag.fromTokenId];
                }
                emit TagProcessed(pendingTag.tagger, tokenId, pendingTag.timestamp);
            } else {
                break;
            }
        }

        _setExtraData(tokenId, (uint96(newCursor) << 80) | (uint96(timestamp) << 40) | tagCount);
    }

    function _beforeTokenTransfer(address, address, uint256 tokenId) internal override {
        _processTagQueue(tokenId);
    }
}
