/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/solver_inbox.json`.
 */
export type SolverInbox = {
  address: 'AwminMpVyPSX86m3w9KWcxjovtnjwxiKZUNTDgDqrctv'
  metadata: {
    name: 'solverInbox'
    version: '0.0.1'
    spec: '0.1.0'
    description: 'Created with Anchor'
  }
  instructions: [
    {
      name: 'claim'
      discriminator: [62, 198, 214, 193, 213, 159, 108, 210]
      accounts: [
        {
          name: 'orderState'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 115, 116, 97, 116, 101]
              },
              {
                kind: 'arg'
                path: 'orderId'
              },
            ]
          }
        },
        {
          name: 'orderTokenAccount'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 116, 111, 107, 101, 110]
              },
              {
                kind: 'arg'
                path: 'orderId'
              },
            ]
          }
        },
        {
          name: 'ownerTokenAccount'
          writable: true
        },
        {
          name: 'claimer'
          writable: true
          signer: true
        },
        {
          name: 'claimerTokenAccount'
          writable: true
        },
        {
          name: 'tokenProgram'
          address: 'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA'
        },
      ]
      args: [
        {
          name: 'orderId'
          type: 'pubkey'
        },
      ]
    },
    {
      name: 'close'
      discriminator: [98, 165, 201, 177, 108, 65, 206, 96]
      accounts: [
        {
          name: 'orderState'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 115, 116, 97, 116, 101]
              },
              {
                kind: 'arg'
                path: 'orderId'
              },
            ]
          }
        },
        {
          name: 'orderTokenAccount'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 116, 111, 107, 101, 110]
              },
              {
                kind: 'arg'
                path: 'orderId'
              },
            ]
          }
        },
        {
          name: 'ownerTokenAccount'
          writable: true
        },
        {
          name: 'owner'
          writable: true
          signer: true
        },
        {
          name: 'tokenProgram'
          address: 'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA'
        },
      ]
      args: [
        {
          name: 'orderId'
          type: 'pubkey'
        },
      ]
    },
    {
      name: 'init'
      docs: [
        'Initialize the inbox state',
        'This should be called only once, preferably by the upgrade authority.',
      ]
      discriminator: [220, 59, 207, 236, 108, 250, 47, 100]
      accounts: [
        {
          name: 'inboxState'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [105, 110, 98, 111, 120, 95, 115, 116, 97, 116, 101]
              },
            ]
          }
        },
        {
          name: 'admin'
          writable: true
          signer: true
        },
        {
          name: 'systemProgram'
          address: '11111111111111111111111111111111'
        },
      ]
      args: [
        {
          name: 'chainId'
          type: 'u64'
        },
        {
          name: 'closeBuffer'
          type: 'i64'
        },
      ]
    },
    {
      name: 'markFilled'
      docs: [
        'Mark an order as filled, and set the claimable_by account.',
        'This may only be called by the inbox admin.',
      ]
      discriminator: [192, 137, 170, 0, 70, 5, 127, 160]
      accounts: [
        {
          name: 'orderState'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 115, 116, 97, 116, 101]
              },
              {
                kind: 'arg'
                path: 'orderId'
              },
            ]
          }
        },
        {
          name: 'inboxState'
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [105, 110, 98, 111, 120, 95, 115, 116, 97, 116, 101]
              },
            ]
          }
        },
        {
          name: 'admin'
          writable: true
          signer: true
        },
      ]
      args: [
        {
          name: 'orderId'
          type: 'pubkey'
        },
        {
          name: 'fillHash'
          type: 'pubkey'
        },
        {
          name: 'claimableBy'
          type: 'pubkey'
        },
      ]
    },
    {
      name: 'open'
      docs: ['Open a new order']
      discriminator: [228, 220, 155, 71, 199, 189, 60, 45]
      accounts: [
        {
          name: 'orderState'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 115, 116, 97, 116, 101]
              },
              {
                kind: 'arg'
                path: 'params.order_id'
              },
            ]
          }
        },
        {
          name: 'owner'
          writable: true
          signer: true
        },
        {
          name: 'mintAccount'
          writable: true
        },
        {
          name: 'ownerTokenAccount'
          writable: true
        },
        {
          name: 'orderTokenAccount'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 116, 111, 107, 101, 110]
              },
              {
                kind: 'arg'
                path: 'params.order_id'
              },
            ]
          }
        },
        {
          name: 'tokenProgram'
          address: 'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA'
        },
        {
          name: 'inboxState'
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [105, 110, 98, 111, 120, 95, 115, 116, 97, 116, 101]
              },
            ]
          }
        },
        {
          name: 'systemProgram'
          address: '11111111111111111111111111111111'
        },
      ]
      args: [
        {
          name: 'params'
          type: {
            defined: {
              name: 'openParams'
            }
          }
        },
      ]
    },
    {
      name: 'reject'
      docs: [
        'Reject an order, refunding owner closing accounts.',
        'Only admin can reject orders.',
      ]
      discriminator: [135, 7, 63, 85, 131, 114, 111, 224]
      accounts: [
        {
          name: 'orderState'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 115, 116, 97, 116, 101]
              },
              {
                kind: 'arg'
                path: 'orderId'
              },
            ]
          }
        },
        {
          name: 'orderTokenAccount'
          writable: true
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [111, 114, 100, 101, 114, 95, 116, 111, 107, 101, 110]
              },
              {
                kind: 'arg'
                path: 'orderId'
              },
            ]
          }
        },
        {
          name: 'ownerTokenAccount'
          writable: true
        },
        {
          name: 'inboxState'
          pda: {
            seeds: [
              {
                kind: 'const'
                value: [105, 110, 98, 111, 120, 95, 115, 116, 97, 116, 101]
              },
            ]
          }
        },
        {
          name: 'admin'
          writable: true
          signer: true
        },
        {
          name: 'tokenProgram'
          address: 'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA'
        },
      ]
      args: [
        {
          name: 'orderId'
          type: 'pubkey'
        },
        {
          name: 'reason'
          type: 'u8'
        },
      ]
    },
  ]
  accounts: [
    {
      name: 'inboxState'
      discriminator: [161, 5, 9, 33, 125, 185, 63, 116]
    },
    {
      name: 'orderState'
      discriminator: [60, 123, 67, 162, 96, 43, 173, 225]
    },
  ]
  events: [
    {
      name: 'eventUpdated'
      discriminator: [238, 86, 17, 103, 12, 182, 141, 61]
    },
  ]
  errors: [
    {
      code: 6000
      name: 'invalidId'
      msg: 'Invalid order ID'
    },
    {
      code: 6001
      name: 'invalidStatus'
      msg: 'Invalid status'
    },
    {
      code: 6002
      name: 'invalidMint'
      msg: 'Invalid mint'
    },
    {
      code: 6003
      name: 'notClosable'
      msg: 'Order not closable yet'
    },
    {
      code: 6004
      name: 'invalidFillHash'
      msg: 'Invalid fill hash'
    },
  ]
  types: [
    {
      name: 'evmCall'
      docs: [
        'EVM call to execute on destination chain',
        'If the call is a native transfer, `target` is the recipient address, and `selector` / `params` are empty.',
      ]
      type: {
        kind: 'struct'
        fields: [
          {
            name: 'target'
            type: {
              array: ['u8', 20]
            }
          },
          {
            name: 'selector'
            type: {
              array: ['u8', 4]
            }
          },
          {
            name: 'value'
            type: 'u128'
          },
          {
            name: 'params'
            type: 'bytes'
          },
        ]
      }
    },
    {
      name: 'evmTokenExpense'
      docs: [
        'TokenExpense describes an ERC20 expense to be paid by the solver on destination chain when filling an',
        'order. Native expenses are inferred from the calls, and are not included in the order data.',
      ]
      type: {
        kind: 'struct'
        fields: [
          {
            name: 'spender'
            docs: [
              'The address that will do token.transferFrom(...) on fill. Required to set allowance',
            ]
            type: {
              array: ['u8', 20]
            }
          },
          {
            name: 'token'
            docs: ['The address of the token on the destination chain']
            type: {
              array: ['u8', 20]
            }
          },
          {
            name: 'amount'
            docs: ['The amount of the token to spend (max == uint96)']
            type: 'u128'
          },
        ]
      }
    },
    {
      name: 'eventUpdated'
      docs: ['Event emitted when an order is opened']
      type: {
        kind: 'struct'
        fields: [
          {
            name: 'orderId'
            type: 'pubkey'
          },
          {
            name: 'status'
            type: {
              defined: {
                name: 'status'
              }
            }
          },
        ]
      }
    },
    {
      name: 'inboxState'
      type: {
        kind: 'struct'
        fields: [
          {
            name: 'admin'
            type: 'pubkey'
          },
          {
            name: 'chainId'
            type: 'u64'
          },
          {
            name: 'deployedAt'
            type: 'u64'
          },
          {
            name: 'bump'
            type: 'u8'
          },
          {
            name: 'closeBufferSecs'
            type: 'i64'
          },
        ]
      }
    },
    {
      name: 'openParams'
      type: {
        kind: 'struct'
        fields: [
          {
            name: 'orderId'
            type: 'pubkey'
          },
          {
            name: 'nonce'
            type: 'u64'
          },
          {
            name: 'depositAmount'
            type: 'u64'
          },
          {
            name: 'destChainId'
            type: 'u64'
          },
          {
            name: 'call'
            type: {
              defined: {
                name: 'evmCall'
              }
            }
          },
          {
            name: 'expense'
            type: {
              defined: {
                name: 'evmTokenExpense'
              }
            }
          },
        ]
      }
    },
    {
      name: 'orderState'
      type: {
        kind: 'struct'
        fields: [
          {
            name: 'orderId'
            type: 'pubkey'
          },
          {
            name: 'status'
            type: {
              defined: {
                name: 'status'
              }
            }
          },
          {
            name: 'owner'
            type: 'pubkey'
          },
          {
            name: 'createdAt'
            type: 'i64'
          },
          {
            name: 'closableAt'
            type: 'i64'
          },
          {
            name: 'claimableBy'
            type: 'pubkey'
          },
          {
            name: 'bump'
            type: 'u8'
          },
          {
            name: 'depositAmount'
            type: 'u64'
          },
          {
            name: 'depositMint'
            type: 'pubkey'
          },
          {
            name: 'destChainId'
            type: 'u64'
          },
          {
            name: 'destCall'
            type: {
              defined: {
                name: 'evmCall'
              }
            }
          },
          {
            name: 'destExpense'
            type: {
              defined: {
                name: 'evmTokenExpense'
              }
            }
          },
          {
            name: 'fillHash'
            type: 'pubkey'
          },
          {
            name: 'rejectReason'
            type: 'u8'
          },
        ]
      }
    },
    {
      name: 'status'
      type: {
        kind: 'enum'
        variants: [
          {
            name: 'invalid'
          },
          {
            name: 'pending'
          },
          {
            name: 'rejected'
          },
          {
            name: 'closed'
          },
          {
            name: 'filled'
          },
          {
            name: 'claimed'
          },
        ]
      }
    },
  ]
}
