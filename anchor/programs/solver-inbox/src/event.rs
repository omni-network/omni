use anchor_lang::prelude::*;
use crate::types::Status;

#[event]
pub struct EventOpened {
    pub order_id: Pubkey,
    pub order_state: Pubkey,
    pub status: Status,
}
