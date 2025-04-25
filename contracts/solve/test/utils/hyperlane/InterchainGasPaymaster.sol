// SPDX-License-Identifier: MIT OR Apache-2.0
pragma solidity >=0.8.0;

/*@@@@@@@       @@@@@@@@@
 @@@@@@@@@       @@@@@@@@@
  @@@@@@@@@       @@@@@@@@@
   @@@@@@@@@       @@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@@@@
     @@@@@  HYPERLANE  @@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@@@@
   @@@@@@@@@       @@@@@@@@@
  @@@@@@@@@       @@@@@@@@@
 @@@@@@@@@       @@@@@@@@@
@@@@@@@@@       @@@@@@@@*/

// ============ Internal Imports ============
import { Message } from "@hyperlane-xyz/core/contracts/libs/Message.sol";
import { StandardHookMetadata } from "@hyperlane-xyz/core/contracts/hooks/libs/StandardHookMetadata.sol";
import { IGasOracle } from "@hyperlane-xyz/core/contracts/interfaces/IGasOracle.sol";
import { IInterchainGasPaymaster } from "@hyperlane-xyz/core/contracts/interfaces/IInterchainGasPaymaster.sol";
import { IPostDispatchHook } from "@hyperlane-xyz/core/contracts/interfaces/hooks/IPostDispatchHook.sol";
import { AbstractPostDispatchHook } from "@hyperlane-xyz/core/contracts/hooks/libs/AbstractPostDispatchHook.sol";
import { Indexed } from "@hyperlane-xyz/core/contracts/libs/Indexed.sol";

// ============ External Imports ============
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title InterchainGasPaymaster
 * @notice Manages payments on a source chain to cover gas costs of relaying
 * messages to destination chains and includes the gas overhead per destination
 * @dev The intended use of this contract is to store overhead gas amounts for destination
 * domains, e.g. Mailbox and ISM gas usage, such that users of this IGP are only required
 * to specify the gas amount used by their own applications.
 */
contract InterchainGasPaymaster is
    IInterchainGasPaymaster,
    AbstractPostDispatchHook,
    IGasOracle,
    Indexed,
    OwnableUpgradeable
{
    using Message for bytes;
    using SafeTransferLib for address;
    using StandardHookMetadata for bytes;
    // ============ Constants ============

    /// @notice The scale of gas oracle token exchange rates.
    uint256 internal constant TOKEN_EXCHANGE_RATE_SCALE = 1e10;
    /// @notice default for user call if metadata not provided
    uint256 internal immutable DEFAULT_GAS_USAGE = 50_000;

    // ============ Public Storage ============

    /// @notice Destination domain => gas oracle and overhead gas amount.
    mapping(uint32 => DomainGasConfig) public destinationGasConfigs;

    /// @notice The benficiary that can receive native tokens paid into this contract.
    address public beneficiary;

    // ============ Events ============

    /**
     * @notice Emitted when the gas oracle for a remote domain is set.
     * @param remoteDomain The remote domain.
     * @param gasOracle The gas oracle.
     * @param gasOverhead The destination gas overhead.
     */
    event DestinationGasConfigSet(uint32 remoteDomain, address gasOracle, uint96 gasOverhead);

    /**
     * @notice Emitted when the beneficiary is set.
     * @param beneficiary The new beneficiary.
     */
    event BeneficiarySet(address beneficiary);

    struct DomainGasConfig {
        IGasOracle gasOracle;
        uint96 gasOverhead;
    }

    struct GasParam {
        uint32 remoteDomain;
        DomainGasConfig config;
    }

    // ============ External Functions ============

    /// @inheritdoc IPostDispatchHook
    function hookType() external pure override returns (uint8) {
        return uint8(IPostDispatchHook.Types.INTERCHAIN_GAS_PAYMASTER);
    }

    /**
     * @param _owner The owner of the contract.
     * @param _beneficiary The beneficiary.
     */
    function initialize(address _owner, address _beneficiary) public initializer {
        __Ownable_init(_owner);
        _setBeneficiary(_beneficiary);
    }

    /**
     * @notice Transfers the entire native token balance to the beneficiary.
     * @dev The beneficiary must be able to receive native tokens.
     */
    function claim() external {
        // Transfer the entire balance to the beneficiary.
        (bool success,) = beneficiary.call{ value: address(this).balance }("");
        require(success, "IGP: claim failed");
    }

    /**
     * @notice Sets the gas oracles for remote domains specified in the config array.
     * @param _configs An array of configs including the remote domain and gas oracles to set.
     */
    function setDestinationGasConfigs(GasParam[] calldata _configs) external onlyOwner {
        uint256 _len = _configs.length;
        for (uint256 i = 0; i < _len; i++) {
            _setDestinationGasConfig(
                _configs[i].remoteDomain, _configs[i].config.gasOracle, _configs[i].config.gasOverhead
            );
        }
    }

    /**
     * @notice Sets the beneficiary.
     * @param _beneficiary The new beneficiary.
     */
    function setBeneficiary(address _beneficiary) external onlyOwner {
        _setBeneficiary(_beneficiary);
    }

    // ============ Public Functions ============

    /**
     * @notice Deposits msg.value as a payment for the relaying of a message
     * to its destination chain.
     * @dev Overpayment will result in a refund of native tokens to the _refundAddress.
     * Callers should be aware that this may present reentrancy issues.
     * @param _messageId The ID of the message to pay for.
     * @param _destinationDomain The domain of the message's destination chain.
     * @param _gasLimit The amount of destination gas to pay for.
     * @param _refundAddress The address to refund any overpayment to.
     */
    function payForGas(bytes32 _messageId, uint32 _destinationDomain, uint256 _gasLimit, address _refundAddress)
        public
        payable
        override
    {
        uint256 _requiredPayment = quoteGasPayment(_destinationDomain, _gasLimit);
        require(msg.value >= _requiredPayment, "IGP: insufficient interchain gas payment");
        uint256 _overpayment = msg.value - _requiredPayment;
        if (_overpayment > 0) {
            require(_refundAddress != address(0), "no refund address");
            _refundAddress.safeTransferETH(_overpayment);
        }

        emit GasPayment(_messageId, _destinationDomain, _gasLimit, _requiredPayment);
    }

    /**
     * @notice Quotes the amount of native tokens to pay for interchain gas.
     * @param _destinationDomain The domain of the message's destination chain.
     * @param _gasLimit The amount of destination gas to pay for.
     * @return The amount of native tokens required to pay for interchain gas.
     */
    function quoteGasPayment(uint32 _destinationDomain, uint256 _gasLimit)
        public
        view
        virtual
        override
        returns (uint256)
    {
        // Get the gas data for the destination domain.
        (uint128 _tokenExchangeRate, uint128 _gasPrice) = getExchangeRateAndGasPrice(_destinationDomain);

        // The total cost quoted in destination chain's native token.
        uint256 _destinationGasCost = _gasLimit * uint256(_gasPrice);

        // Convert to the local native token.
        return (_destinationGasCost * _tokenExchangeRate) / TOKEN_EXCHANGE_RATE_SCALE;
    }

    /**
     * @notice Gets the token exchange rate and gas price from the configured gas oracle
     * for a given destination domain.
     * @param _destinationDomain The destination domain.
     * @return tokenExchangeRate The exchange rate of the remote native token quoted in the local native token.
     * @return gasPrice The gas price on the remote chain.
     */
    function getExchangeRateAndGasPrice(uint32 _destinationDomain)
        public
        view
        override
        returns (uint128 tokenExchangeRate, uint128 gasPrice)
    {
        IGasOracle _gasOracle = destinationGasConfigs[_destinationDomain].gasOracle;

        if (address(_gasOracle) == address(0)) {
            revert(string.concat("Configured IGP doesn't support domain ", Strings.toString(_destinationDomain)));
        }
        return _gasOracle.getExchangeRateAndGasPrice(_destinationDomain);
    }

    /**
     * @notice Returns the stored destinationGasOverhead added to the _gasLimit.
     * @dev If there is no stored destinationGasOverhead, 0 is used. This is useful in the case
     *      the ISM deployer wants to subsidize the overhead gas cost. Then, can specify the gas oracle
     *      they want to use with the destination domain, but set the overhead to 0.
     * @param _destinationDomain The domain of the message's destination chain.
     * @param _gasLimit The amount of destination gas to pay for. This is only for application gas usage as
     *      the gas usage for the mailbox and the ISM is already accounted in the DomainGasConfig.gasOverhead
     */
    function destinationGasLimit(uint32 _destinationDomain, uint256 _gasLimit) public view returns (uint256) {
        return uint256(destinationGasConfigs[_destinationDomain].gasOverhead) + _gasLimit;
    }

    // ============ Internal Functions ============

    /// @inheritdoc AbstractPostDispatchHook
    function _postDispatch(bytes calldata metadata, bytes calldata message) internal override {
        payForGas(
            message.id(),
            message.destination(),
            destinationGasLimit(message.destination(), metadata.gasLimit(DEFAULT_GAS_USAGE)),
            metadata.refundAddress(message.senderAddress())
        );
    }

    /// @inheritdoc AbstractPostDispatchHook
    function _quoteDispatch(bytes calldata metadata, bytes calldata message) internal view override returns (uint256) {
        return quoteGasPayment(
            message.destination(), destinationGasLimit(message.destination(), metadata.gasLimit(DEFAULT_GAS_USAGE))
        );
    }

    /**
     * @notice Sets the beneficiary.
     * @param _beneficiary The new beneficiary.
     */
    function _setBeneficiary(address _beneficiary) internal {
        beneficiary = _beneficiary;
        emit BeneficiarySet(_beneficiary);
    }

    /**
     * @notice Sets the gas oracle and destination gas overhead for a remote domain.
     * @param _remoteDomain The remote domain.
     * @param _gasOracle The gas oracle.
     * @param _gasOverhead The destination gas overhead.
     */
    function _setDestinationGasConfig(uint32 _remoteDomain, IGasOracle _gasOracle, uint96 _gasOverhead) internal {
        destinationGasConfigs[_remoteDomain] = DomainGasConfig(_gasOracle, _gasOverhead);
        emit DestinationGasConfigSet(_remoteDomain, address(_gasOracle), _gasOverhead);
    }
}
