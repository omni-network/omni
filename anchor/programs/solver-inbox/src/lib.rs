mod error;
mod event;
mod helpers;
mod types;

use anchor_lang::prelude::*;
use error::*;
use event::*;
use helpers::*;
use types::*;

declare_id!("AwminMpVyPSX86m3w9KWcxjovtnjwxiKZUNTDgDqrctv");

#[program]
pub mod solver_inbox {
    use super::*;

    pub fn open(ctx: Context<Open>, params: OpenParams) -> Result<()> {
        let order_id = order_id(ctx.accounts.owner.key(), params.nonce);
        require_eq!(order_id, params.order_id, InboxError::InvalidID);

        let status = Status::Pending;

        let state = &mut ctx.accounts.order_state;
        state.order_id = order_id;
        state.status = status.clone();
        state.authority = ctx.accounts.owner.key();
        state.bump = ctx.bumps.order_state;
        state.deposit = params.deposit;
        state.expense = params.expense;

        emit!(EventOpened {
            order_id,
            status,
            order_state: ctx.accounts.order_state.key(),
        });

        Ok(())
    }
}

#[derive(Accounts)]
#[instruction(params: OpenParams)]
pub struct Open<'info> {
    #[account(
        init,
        payer = owner,
        space = 8 + OrderState::SIZE,
        seeds = [b"order-state", params.order_id.as_ref()],
        bump,
    )]
    pub order_state: Account<'info, OrderState>,
    #[account(mut)]
    pub owner: Signer<'info>,
    pub system_program: Program<'info, System>,
}
