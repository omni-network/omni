use crate::types::Status;
use anchor_lang::prelude::*;

#[event]
pub struct EventOpened {
    pub order_id: Pubkey,
    pub order_state: Pubkey,
    pub status: Status,
}

#[event]
pub struct EventMarkFilled {
    pub order_id: Pubkey,
    pub order_state: Pubkey,
    pub status: Status,
}

#[event]
pub struct EventClaimed {
    pub order_id: Pubkey,
    pub order_state: Pubkey,
    pub status: Status,
}
