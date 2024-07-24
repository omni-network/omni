/**
 * Abi parameters for xtype structs. Allowing us to encode / decode to
 * and from bytes, matching Solidity's abi encoding. Struct fields must match
 * the order of the corresponding struct definition in Solidity.
 */

// ABI parameters for XTypes.BlockHeader (src/libraries/XTypes.sol)
export const xBlockHeaderParams = {
  name: 'blockHeader',
  type: 'tuple',
  components: [
    { name: 'sourceChainId', type: 'uint64' },
    { name: 'consensusChainId', type: 'uint64' },
    { name: 'confLevel', type: 'uint8' },
    { name: 'offset', type: 'uint64' },
    { name: 'sourceBlockHeight', type: 'uint64'},
    { name: 'sourceBlockHash', type: 'bytes32' },
  ],
} as const

// ABI parameters for XTypes.Msg (src/libraries/XTypes.sol)
export const xMsgParams = {
  name: 'msgs',
  type: 'tuple',
  components: [
    { name: 'destChainId', type: 'uint64' },
    { name: 'shardId', type: 'uint64' },
    { name: 'offset', type: 'uint64' },
    { name: 'sender', type: 'address' },
    { name: 'to', type: 'address' },
    { name: 'data', type: 'bytes' },
    { name: 'gasLimit', type: 'uint64' },
  ],
} as const

// ABI parameters for XTypes.Msg[]
export const xMsgsParams = { ...xMsgParams, type: 'tuple[]' } as const

// ABI parameters for XTypes.Submission (src/libraries/XTypes.sol)
export const xSubParams = {
  name: 'xsubs',
  type: 'tuple',
  components: [
    { name: 'attestationRoot', type: 'bytes32' },
    { name: 'validatorSetId', type: 'uint64' },
    xBlockHeaderParams,
    xMsgsParams,
    { name: 'proof', type: 'bytes32[]' },
    { name: 'proofFlags', type: 'bool[]' },
    {
      name: 'signatures',
      type: 'tuple[]',
      components: [
        { name: 'validatorPubKey', type: 'bytes' },
        { name: 'signature', type: 'bytes' },
      ],
    },
  ],
} as const

// ABI parameters for TestXTypes.Block (test/common/TestXTypes.sol)
export const xBlockParams = {
  name: 'block',
  type: 'tuple',
  components: [xBlockHeaderParams, xMsgsParams],
} as const
