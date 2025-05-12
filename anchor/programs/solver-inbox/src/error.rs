use anchor_lang::prelude::*;

// =====================================================================
// ERRORS
// =====================================================================

#[error_code]
pub enum InboxError {
    #[msg("Invalid order ID")]
    InvalidID,
    #[msg("Invalid status")]
    InvalidStatus,
    #[msg("Invalid mint")]
    InvalidMint,
    #[msg("Order not closable yet")]
    NotClosable,
}
