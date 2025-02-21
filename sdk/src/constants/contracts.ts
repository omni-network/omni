/**
 * NOTE:
 *
 *  Addresses here are from staging and may change until SolverNet is stable.
 *  If you have any issues please reach out to us.
 *
 */

export const inbox = {
  abi: [
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
  ],
  // TODO: this will be fetched via the solver API
  address: '0x71D48a88600A790FB1E94efdBF6c284fC08F99B8',
} as const

export const outbox = {
  abi: [
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
  ],
  // TODO: this will be fetched via the solver API
  address: '0xD98068374C233e4d4081D24154489551DcD6FDF5',
} as const

export const middleman = {
  abi: [
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
  ],
  // TODO: this will be fetched via the solver API
  address: '0xF72eA80Edec895CABE430f266135DD0267e18cb2',
} as const
