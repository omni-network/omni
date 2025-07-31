// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.30;

import { ERC20 } from "solady/src/tokens/ERC20.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract Nomina is ERC20 {
    using SafeTransferLib for address;

    /**
     * @notice Thrown when an address parameter is zero.
     */
    error ZeroAddress();

    /**
     * @notice Thrown when the sender is not the mint authority.
     * @dev This is to prevent unauthorized minting.
     */
    error Unauthorized();

    /**
     * @notice Thrown when the conversion is disabled.
     * @dev This would only take place if deployed to chains where there isn't an OMNI ERC20 token.
     */
    error ConversionDisabled();

    /**
     * @notice Emitted when the mint authority is set.
     * @param mintAuthority The new mint authority.
     */
    event MintAuthoritySet(address indexed mintAuthority);

    /**
     * @notice Emitted when the minter is set.
     * @param minter The new minter.
     */
    event MinterSet(address indexed minter);

    /**
     * @notice The name of the token.
     */
    string private constant _NAME = "Nomina";

    /**
     * @notice The symbol of the token.
     */
    string private constant _SYMBOL = "NOM";

    /**
     * @notice The hash of the name, used in permit EIP-712 hashes.
     */
    bytes32 private constant _NAME_HASH = keccak256(bytes(_NAME));

    /**
     * @notice The address OMNI tokens are sent to on conversion as they cannot be sent to the zero address or burned.
     */
    address private constant _DEAD_ADDRESS = address(0xdead);

    /**
     * @notice The conversion rate from OMNI to NOM.
     */
    uint8 public constant CONVERSION_RATE = 75;

    /**
     * @notice The OMNI token contract.
     */
    address public immutable omni;

    /**
     * @notice The mint authority authorized to set the minter.
     */
    address public mintAuthority;

    /**
     * @notice The address authorized to mint NOM tokens.
     */
    address public minter;

    /**
     * @notice Modifier to check if the sender is the mint authority.
     */
    modifier onlyMintAuthority() {
        if (msg.sender != mintAuthority) revert Unauthorized();
        _;
    }

    /**
     * @notice Modifier to check if the sender is the minter.
     */
    modifier onlyMinter() {
        if (msg.sender != minter) revert Unauthorized();
        _;
    }

    /**
     * @notice Contract constructor.
     * @param _omni The OMNI token contract.
     * @param _mintAuthority The mint authority.
     * @param _minter The minter.
     */
    constructor(address _omni, address _mintAuthority, address _minter) {
        omni = _omni;
        mintAuthority = _mintAuthority;
        minter = _minter;

        emit MintAuthoritySet(_mintAuthority);
        emit MinterSet(_minter);
    }

    /**
     * @notice Returns the name of the token.
     * @return _ The name of the token.
     */
    function name() public pure override returns (string memory) {
        return _NAME;
    }

    /**
     * @notice Returns the symbol of the token.
     * @return _ The symbol of the token.
     */
    function symbol() public pure override returns (string memory) {
        return _SYMBOL;
    }

    /**
     * @notice Mints new tokens.
     * @dev Only the minter can mint new tokens. No OMNI tokens are utilized.
     * @param to The address to mint the tokens to.
     * @param amount The amount of tokens to mint.
     */
    function mint(address to, uint256 amount) public onlyMinter {
        _mint(to, amount);
    }

    /**
     * @notice Burns tokens.
     * @dev Only the sender can burn tokens.
     * @param amount The amount of tokens to burn.
     */
    function burn(uint256 amount) public {
        if (amount == 0) return;
        _burn(msg.sender, amount);
    }

    /**
     * @notice Converts OMNI tokens to NOM tokens.
     * @dev The sender must have approved the contract to spend their OMNI tokens.
     * @param to The address to send the NOM tokens to.
     * @param amount The amount of OMNI tokens to convert.
     */
    function convert(address to, uint256 amount) public {
        address _omni = omni;
        if (amount == 0) return;
        if (to == address(0)) revert ZeroAddress();
        if (_omni == address(0)) revert ConversionDisabled();

        _omni.safeTransferFrom(msg.sender, _DEAD_ADDRESS, amount);
        _mint(to, amount * CONVERSION_RATE);
    }

    /**
     * @notice Sets the mint authority.
     * @dev Only the mint authority can set the mint authority.
     * @param _mintAuthority The new mint authority.
     */
    function setMintAuthority(address _mintAuthority) public onlyMintAuthority {
        mintAuthority = _mintAuthority;
        emit MintAuthoritySet(_mintAuthority);
    }

    /**
     * @notice Sets the minter.
     * @dev Only the mint authority can set the minter.
     * @param _minter The new minter.
     */
    function setMinter(address _minter) public onlyMintAuthority {
        minter = _minter;
        emit MinterSet(_minter);
    }

    /**
     * @notice Sets a constant name hash in Solady ERC20 to optimize permit gas costs.
     * @return _ The hash of the name.
     */
    function _constantNameHash() internal pure override returns (bytes32) {
        return _NAME_HASH;
    }
}
