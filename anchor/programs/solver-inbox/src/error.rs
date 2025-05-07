use anchor_lang::prelude::*;

// =====================================================================
// ERRORS
// =====================================================================

#[error_code]
pub enum InboxError {
    #[msg("Invalid order ID")]
    InvalidID,
}
