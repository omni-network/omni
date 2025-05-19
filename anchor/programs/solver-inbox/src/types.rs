use anchor_lang::prelude::*;

#[account]
#[derive(InitSpace)]
pub struct InboxState {
    pub admin: Pubkey,          // Inbox admin
    pub chain_id: u64,          // Chain ID of the inbox
    pub deployed_at: u64,       // Slot when the inbox was deployed
    pub bump: u8,               // Bump seed
    pub close_buffer_secs: i64, // Duration in seconds after which the order may be closed by owner
}

impl InboxState {
    pub const SEED_PREFIX: &'static [u8; 11] = b"inbox_state";
}

#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Clone, Debug, PartialEq, Eq)]
pub enum Status {
    Invalid,
    Pending,
    Rejected,
    Closed,
    Filled,
    Claimed,
}

#[account]
#[derive(InitSpace)]
pub struct OrderState {
    pub order_id: Pubkey,              // Order ID
    pub status: Status,                // Order status
    pub owner: Pubkey,                 // Owner of this order
    pub created_at: i64,               // Unix timestamp when the order was created
    pub closable_at: i64,              // Unix timestamp when the order may be closed
    pub claimable_by: Pubkey,          // Account that may claim the order
    pub bump: u8,                      // Bump seed
    pub deposit_amount: u64,           // Amount of tokens deposited by the owner
    pub deposit_mint: Pubkey,          // Deposit mint (token).
    pub dest_chain_id: u64,            // Chain ID of the destination chain
    pub dest_call: EVMCall,            // EVM call to execute on destination chain
    pub dest_expense: EVMTokenExpense, // Description of EVM token expense (encoded in EVMCall)
    pub fill_hash: Pubkey,             // Hash of the order fill data
    pub reject_reason: u8,             // Reject reason enum
}

impl OrderState {
    pub const SEED_PREFIX: &'static [u8; 11] = b"order_state";
}

pub const ORDER_TOKEN_SEED_PREFIX: &[u8; 11] = b"order_token";

#[derive(AnchorSerialize, AnchorDeserialize)]
pub struct OpenParams {
    pub order_id: Pubkey,
    pub nonce: u64,
    pub deposit_amount: u64,
    pub dest_chain_id: u64,
    pub call: EVMCall,
    pub expense: EVMTokenExpense,
}

/// EVM call to execute on destination chain
/// If the call is a native transfer, `target` is the recipient address, and `selector` / `params` are empty.
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Clone, Debug, PartialEq, Eq)]
pub struct EVMCall {
    pub target: [u8; 20],  // Address of the contract to call
    pub selector: [u8; 4], // Function selector
    pub value: u128,       // Amount of native tokens to send
    #[max_len(256)] // Maximum length of the params is 256 bytes
    pub params: Vec<u8>, // ABI encoded parameters
}

/// TokenExpense describes an ERC20 expense to be paid by the solver on destination chain when filling an
/// order. Native expenses are inferred from the calls, and are not included in the order data.
#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Clone, Debug, PartialEq, Eq)]
pub struct EVMTokenExpense {
    /// The address that will do token.transferFrom(...) on fill. Required to set allowance
    pub spender: [u8; 20],
    /// The address of the token on the destination chain
    pub token: [u8; 20],
    /// The amount of the token to spend (max == uint96)
    pub amount: u128,
}

/// Event emitted when an order is opened
#[event]
pub struct EventUpdated {
    pub order_id: Pubkey,
    pub status: Status,
}
