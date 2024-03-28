// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.20;

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

// Followed patterns for EnumerableSet.UintSet
// See https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/structs/EnumerableSet.sol
library EnumerableUint64Set {
    struct Set {
        EnumerableSet.Bytes32Set _inner;
    }

    function add(Set storage set, uint64 value) internal returns (bool) {
        return EnumerableSet.add(set._inner, bytes32(uint256(value)));
    }

    function remove(Set storage set, uint64 value) internal returns (bool) {
        return EnumerableSet.remove(set._inner, bytes32(uint256(value)));
    }

    function contains(Set storage set, uint256 value) internal view returns (bool) {
        return EnumerableSet.contains(set._inner, bytes32(value));
    }

    function length(Set storage set) internal view returns (uint256) {
        return EnumerableSet.length(set._inner);
    }

    function at(Set storage set, uint256 index) internal view returns (uint64) {
        return uint64(uint256(EnumerableSet.at(set._inner, index)));
    }

    function values(Set storage set) internal view returns (uint64[] memory) {
        bytes32[] memory store = EnumerableSet.values(set._inner);
        uint64[] memory result;

        assembly {
            result := store
        }

        return result;
    }
}
