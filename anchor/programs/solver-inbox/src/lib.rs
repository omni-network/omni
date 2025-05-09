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

    /// Initialize the inbox state
    /// This should be called only once, preferably by the upgrade authority.
    pub fn init(ctx: Context<Init>) -> Result<()> {
        let state = &mut ctx.accounts.inbox_state;
        state.admin = ctx.accounts.admin.key();
        state.deployed_at = Clock::get()?.slot;
        state.bump = ctx.bumps.inbox_state;

        Ok(())
    }

    /// Open a new order
    pub fn open(ctx: Context<Open>, params: OpenParams) -> Result<()> {
        let order_id = order_id(ctx.accounts.owner.key(), params.nonce);
        require_eq!(order_id, params.order_id, InboxError::InvalidID);

        let state = &mut ctx.accounts.order_state;
        state.order_id = order_id;
        state.status = Status::Pending;
        state.owner = ctx.accounts.owner.key();
        state.bump = ctx.bumps.order_state;
        state.deposit = params.deposit;
        state.expense = params.expense;

        emit!(EventOpened {
            order_id: state.order_id,
            status: state.status.clone(),
            order_state: state.key(),
        });

        Ok(())
    }

    /// Mark an order as filled
    /// This may only be called by the inbox admin.
    pub fn mark_filled(ctx: Context<MarkFilled>, _order_id: Pubkey) -> Result<()> {
        let state = &mut ctx.accounts.order_state;
        require!(state.status == Status::Pending, InboxError::InvalidID);

        state.status = Status::Filled;

        emit!(EventMarkFilled {
            order_id: state.order_id,
            status: state.status.clone(),
            order_state: state.key(),
        });

        Ok(())
    }
}

#[derive(Accounts)]
pub struct Init<'info> {
    #[account(
        init,
        payer = admin,
        space = 8 + InboxState::SIZE,
        seeds = [b"inbox-state"],
        bump,
    )]
    pub inbox_state: Account<'info, InboxState>,
    #[account(mut)]
    pub admin: Signer<'info>,
    pub system_program: Program<'info, System>,
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

#[derive(Accounts)]
#[instruction(order_id: Pubkey)]
pub struct MarkFilled<'info> {
    #[account(
        mut,
        seeds = [b"order-state", order_id.as_ref()],
        bump,
    )]
    pub order_state: Account<'info, OrderState>,
    #[account(
        seeds = [b"inbox-state"],
        bump,
        constraint = inbox_state.admin == admin.key(),
    )]
    pub inbox_state: Account<'info, InboxState>,
    #[account(mut)]
    pub admin: Signer<'info>,
}
