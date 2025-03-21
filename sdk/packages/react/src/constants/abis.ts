export const inboxABI = [
  {
    type: 'function',
    inputs: [
      {
        name: 'order',
        internalType: 'struct IERC7683.OnchainCrossChainOrder',
        type: 'tuple',
        components: [
          { name: 'fillDeadline', internalType: 'uint32', type: 'uint32' },
          { name: 'orderDataType', internalType: 'bytes32', type: 'bytes32' },
          { name: 'orderData', internalType: 'bytes', type: 'bytes' },
        ],
      },
    ],
    name: 'open',
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    inputs: [{ name: 'id', internalType: 'bytes32', type: 'bytes32' }],
    name: 'getOrder',
    outputs: [
      {
        name: 'resolved',
        internalType: 'struct IERC7683.ResolvedCrossChainOrder',
        type: 'tuple',
        components: [
          { name: 'user', internalType: 'address', type: 'address' },
          { name: 'originChainId', internalType: 'uint256', type: 'uint256' },
          { name: 'openDeadline', internalType: 'uint32', type: 'uint32' },
          { name: 'fillDeadline', internalType: 'uint32', type: 'uint32' },
          { name: 'orderId', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'maxSpent',
            internalType: 'struct IERC7683.Output[]',
            type: 'tuple[]',
            components: [
              { name: 'token', internalType: 'bytes32', type: 'bytes32' },
              { name: 'amount', internalType: 'uint256', type: 'uint256' },
              { name: 'recipient', internalType: 'bytes32', type: 'bytes32' },
              { name: 'chainId', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'minReceived',
            internalType: 'struct IERC7683.Output[]',
            type: 'tuple[]',
            components: [
              { name: 'token', internalType: 'bytes32', type: 'bytes32' },
              { name: 'amount', internalType: 'uint256', type: 'uint256' },
              { name: 'recipient', internalType: 'bytes32', type: 'bytes32' },
              { name: 'chainId', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'fillInstructions',
            internalType: 'struct IERC7683.FillInstruction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'destinationChainId',
                internalType: 'uint64',
                type: 'uint64',
              },
              {
                name: 'destinationSettler',
                internalType: 'bytes32',
                type: 'bytes32',
              },
              { name: 'originData', internalType: 'bytes', type: 'bytes' },
            ],
          },
        ],
      },
      {
        name: 'state',
        internalType: 'struct ISolverNetInbox.OrderState',
        type: 'tuple',
        components: [
          {
            name: 'status',
            internalType: 'enum ISolverNetInbox.Status',
            type: 'uint8',
          },
          { name: 'timestamp', internalType: 'uint32', type: 'uint32' },
          { name: 'claimant', internalType: 'address', type: 'address' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'orderId',
        internalType: 'bytes32',
        type: 'bytes32',
        indexed: true,
      },
      {
        name: 'resolvedOrder',
        internalType: 'struct IERC7683.ResolvedCrossChainOrder',
        type: 'tuple',
        components: [
          { name: 'user', internalType: 'address', type: 'address' },
          { name: 'originChainId', internalType: 'uint256', type: 'uint256' },
          { name: 'openDeadline', internalType: 'uint32', type: 'uint32' },
          { name: 'fillDeadline', internalType: 'uint32', type: 'uint32' },
          { name: 'orderId', internalType: 'bytes32', type: 'bytes32' },
          {
            name: 'maxSpent',
            internalType: 'struct IERC7683.Output[]',
            type: 'tuple[]',
            components: [
              { name: 'token', internalType: 'bytes32', type: 'bytes32' },
              { name: 'amount', internalType: 'uint256', type: 'uint256' },
              { name: 'recipient', internalType: 'bytes32', type: 'bytes32' },
              { name: 'chainId', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'minReceived',
            internalType: 'struct IERC7683.Output[]',
            type: 'tuple[]',
            components: [
              { name: 'token', internalType: 'bytes32', type: 'bytes32' },
              { name: 'amount', internalType: 'uint256', type: 'uint256' },
              { name: 'recipient', internalType: 'bytes32', type: 'bytes32' },
              { name: 'chainId', internalType: 'uint256', type: 'uint256' },
            ],
          },
          {
            name: 'fillInstructions',
            internalType: 'struct IERC7683.FillInstruction[]',
            type: 'tuple[]',
            components: [
              {
                name: 'destinationChainId',
                internalType: 'uint64',
                type: 'uint64',
              },
              {
                name: 'destinationSettler',
                internalType: 'bytes32',
                type: 'bytes32',
              },
              { name: 'originData', internalType: 'bytes', type: 'bytes' },
            ],
          },
        ],
        indexed: false,
      },
    ],
    name: 'Open',
  },
] as const

export const outboxABI = [
  {
    type: 'function',
    inputs: [
      { name: 'orderId', internalType: 'bytes32', type: 'bytes32' },
      { name: 'originData', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'didFill',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
] as const

export const middlemanABI = [
  {
    type: 'function',
    inputs: [
      { name: 'token', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'target', internalType: 'address', type: 'address' },
      { name: 'data', internalType: 'bytes', type: 'bytes' },
    ],
    name: 'executeAndTransfer',
    outputs: [],
    stateMutability: 'payable',
  },
  { type: 'error', inputs: [], name: 'CallFailed' },
] as const
