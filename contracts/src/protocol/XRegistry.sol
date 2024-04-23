// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { XRegistryReplica } from "./XRegistryReplica.sol";

/**
 * @title XRegistry
 * @notice Registry for cross-chain contract deployments. Replicated across each supported chain.
 * @dev Using Ownable, rather than OwnableUpgradeable, to because predeploys don't require initalization.
 */
contract XRegistry is Ownable, XRegistryBase {
    /**
     * @notice Register a contract deployment.
     */
    event ContractRegistered(uint64 indexed chainId, string indexed name, address indexed registrant, address addr);

    /**
     * @notice xcall gas limit for XRegistryReplica.set
     */
    uint64 public constant XSET_GAS_LIMIT = 100_000;

    /**
     * @notice OmniPortal contract.
     */
    IOmniPortal public omni;

    /**
     * @notice Mapping of chain IDs to the XRegistryReplica address on that chain.
     */
    mapping(uint64 => address) public replicas;

    /**
     * @notice List of supported chain IDs (ie. chain ids with XRegistry replicas)
     */
    uint64[] internal _chainIds;

    /**
     * @notice Register a contract deployment.
     */
    function register(uint64 chainId, string calldata name, address addr) external payable {
        require(isSupportedChain(chainId), "XRegistry: chain not supported");
        require(msg.value >= _syncFee(chainId, name, msg.sender, addr), "XRegistry: insufficient fee");

        _set(chainId, name, msg.sender, addr);
        _sync(chainId, name, msg.sender, addr);

        emit ContractRegistered(chainId, name, msg.sender, addr);
    }

    /**
     * @notice Calculate the fee to register a contract.
     */
    function registrationFee(uint64 chainId, string calldata name, address addr) external view returns (uint256) {
        return _syncFee(chainId, name, msg.sender, addr);
    }

    /**
     * @notice Syncs replicas with a new registration for contract `name` by `registrant`.
     *          - Adds all existing registrations of `name` by `registrant` to the replica on `chainId`
     *          - Adds the new deployment on `chainId` to all replicas with existing registrations
     *            of `name` by `registrant`.
     *
     *          Ex. Consider contract with name "MyContract" and registrant "0x1234". It has 2
     *          deployments.
     *              - Chain 1, address 0x1111
     *              - Chain 2, address 0x2222
     *
     *          When registering a new deployment on Chain 3, address 0x3333, the sync process is as
     *          follows.
     *              - Add (Chain 1 -> 0x1111) and (Chain 2 -> 0x2222) to Chain 3
     *              - Add (Chain 3 -> 0x3333) to Chain 1 and Chain 2
     */
    function _sync(uint64 chainId, string calldata name, address registrant, address addr) internal {
        for (uint256 i = 0; i < _chainIds.length; i++) {
            uint64 otherChainId = _chainIds[i];

            // don't sync with self
            if (otherChainId == chainId) continue;

            address otherAddr = _get(otherChainId, name, registrant);

            // if this chain does not have a registration for `name` by `registrant`, do nothing
            if (otherAddr == address(0)) continue;

            // set "other" registration on chain of new registration
            _xset(chainId, otherChainId, name, registrant, otherAddr);

            // set new registration on "other" chain
            _xset(otherChainId, chainId, name, registrant, addr);
        }
    }

    /**
     * @notice Calculate the fee to sync replicas with a new registration for contract `name` by `registrant`.
     */
    function _syncFee(uint64 chainId, string calldata name, address registrant, address addr)
        internal
        view
        returns (uint256)
    {
        uint256 fee;

        for (uint256 i = 0; i < _chainIds.length; i++) {
            uint64 otherChainId = _chainIds[i];

            // don't sync with self
            if (otherChainId == chainId) continue;

            address otherAddr = _get(otherChainId, name, registrant);

            // if this chain does not have a registration for `name` by `registrant`, do nothing
            if (otherAddr == address(0)) continue;

            // fee to set set "other" registration on chain of new registration
            fee += _xsetFee(chainId, otherChainId, name, registrant, otherAddr);

            // fee to set new registration on "other" chain
            fee += _xsetFee(otherChainId, chainId, name, registrant, addr);
        }

        return fee;
    }

    function _xset(uint64 destChainId, uint64 chainId, string calldata name, address registrant, address addr)
        internal
    {
        // don't xset to self
        if (destChainId == omni.chainId()) return;

        address replica = replicas[destChainId];
        require(replica != address(0), "XRegistry: unknown chain");

        // encode XRegistryReplica.set(chainId, name, registrant, addr)
        bytes memory xcalldata = abi.encodeWithSelector(XRegistryReplica.set.selector, chainId, name, registrant, addr);

        // make xcall, paying fee
        omni.xcall{ value: omni.feeFor(destChainId, xcalldata, XSET_GAS_LIMIT) }(
            destChainId, replica, xcalldata, XSET_GAS_LIMIT
        );
    }

    function _xsetFee(uint64 destChainId, uint64 chainId, string calldata name, address registrant, address addr)
        internal
        view
        returns (uint256)
    {
        // don't xset to self
        if (destChainId == omni.chainId()) return 0;

        address replica = replicas[destChainId];
        require(replica != address(0), "XRegistry: unknown chain");

        // encode XRegistryReplica.set(chainId, name, registrant, addr)
        bytes memory xcalldata = abi.encodeWithSelector(XRegistryReplica.set.selector, chainId, name, registrant, addr);

        return omni.feeFor(destChainId, xcalldata, XSET_GAS_LIMIT);
    }

    /**
     * @notice Return the list of supported chain IDs.
     */
    function chainIds() public view returns (uint64[] memory) {
        return _chainIds;
    }

    /**
     * @notice Return true if the given chain is supported (i.e. has a replica)
     */
    function isSupportedChain(uint64 chainId) public view returns (bool) {
        return replicas[chainId] != address(0);
    }

    /**
     * @notice Set the address of the replica contract on the given chain.
     */
    function setReplica(uint64 chainId, address replica) public onlyOwner {
        require(replicas[chainId] == address(0), "XRegistry: replica already set");
        replicas[chainId] = replica;
        _chainIds.push(chainId);
    }

    /**
     * @notice Set the address of the OmniPortal contract.
     */
    function setPortal(address _omni) public onlyOwner {
        omni = IOmniPortal(_omni);
    }
}
