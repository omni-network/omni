use anchor_lang::{prelude::*, AnchorDeserialize, AnchorSerialize};

#[account]
pub struct InboxState {
    pub admin: Pubkey,  // Inbox admin
    pub deployed_at: u64, // Slot when the inbox was deployed
    pub bump: u8,          // Bump seed
}

impl InboxState {
    pub const SIZE: usize = 32 // Admin
        + 8 // Deployed at
        + 1; // Bump seed
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug, PartialEq, Eq)]
pub enum Status {
    Invalid,
    Pending,
    Rejected,
    Closed,
    Filled,
    Claimed,
}

#[account]
pub struct OrderState {
    pub order_id: Pubkey,  // Order ID
    pub status: Status,    // Order status
    pub owner: Pubkey, // Owner of this order
    pub bump: u8,          // Bump seed
    pub deposit: TokenAmount,
    pub expense: TokenAmount,
}

impl OrderState {
    pub const SIZE: usize = 32 // Order ID
        + 1 // Status
        + 32 // Authority
        + 1 // Bump seed
        + (TokenAmount::SIZE * 2) // Deposit and expense
    ;
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug, PartialEq, Eq)]
pub struct TokenAmount {
    pub token: Pubkey, // Token address
    pub amount: u64,   // Amount of tokens
}

impl TokenAmount {
    pub const SIZE: usize = 32 // Token address
        + 8; // Amount of tokens
}

#[derive(AnchorSerialize, AnchorDeserialize)]
pub struct OpenParams {
    pub order_id: Pubkey,
    pub nonce: u64,
    pub deposit: TokenAmount,
    pub expense: TokenAmount,
}
