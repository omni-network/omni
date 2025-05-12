use anchor_lang::prelude::*;

#[account]
#[derive(InitSpace)]
pub struct InboxState {
    pub admin: Pubkey,    // Inbox admin
    pub deployed_at: u64, // Slot when the inbox was deployed
    pub bump: u8,         // Bump seed
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
    pub order_id: Pubkey,     // Order ID
    pub status: Status,       // Order status
    pub owner: Pubkey,        // Owner of this order
    pub claimable_by: Pubkey, // Account that may claim the order
    pub bump: u8,             // Bump seed
    pub deposit: TokenAmount,
    pub expense: TokenAmount,
}

impl OrderState {
    pub const SEED_PREFIX: &'static [u8; 11] = b"order_state";
}

pub const ORDER_TOKEN_SEED_PREFIX: &[u8; 11] = b"order_token";

#[derive(AnchorSerialize, AnchorDeserialize, InitSpace, Clone, Debug, PartialEq, Eq)]
pub struct TokenAmount {
    pub mint: Pubkey, // Mint (token) address
    pub amount: u64,  // Amount of tokens
}

#[derive(AnchorSerialize, AnchorDeserialize)]
pub struct OpenParams {
    pub order_id: Pubkey,
    pub nonce: u64,
    pub deposit_amount: u64,
    pub expense_mint: Pubkey,
    pub expense_amount: u64,
}
