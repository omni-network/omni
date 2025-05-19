#![allow(unexpected_cfgs)]
mod error;
mod helpers;
mod types;

use anchor_lang::prelude::*;
use anchor_spl::token::*;
use error::*;
use helpers::*;
use types::*;

declare_id!("AwminMpVyPSX86m3w9KWcxjovtnjwxiKZUNTDgDqrctv");

#[program]
pub mod solver_inbox {
    use super::*;

    /// Initialize the inbox state
    /// This should be called only once, preferably by the upgrade authority.
    pub fn init(ctx: Context<Init>, chain_id: u64, close_buffer: i64) -> Result<()> {
        let state = &mut ctx.accounts.inbox_state;
        state.admin = ctx.accounts.admin.key();
        state.deployed_at = Clock::get()?.slot;
        state.bump = ctx.bumps.inbox_state;
        state.close_buffer_secs = close_buffer;
        state.chain_id = chain_id;

        Ok(())
    }

    /// Open a new order
    pub fn open(ctx: Context<Open>, params: OpenParams) -> Result<()> {
        let order_id = order_id(ctx.accounts.owner.key(), params.nonce);
        require_eq!(order_id, params.order_id, InboxError::InvalidID);

        transfer(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                Transfer {
                    from: ctx.accounts.owner_token_account.to_account_info(),
                    to: ctx.accounts.order_token_account.to_account_info(),
                    authority: ctx.accounts.owner.to_account_info(),
                },
            ),
            params.deposit_amount,
        )?;

        let state = &mut ctx.accounts.order_state;
        state.order_id = order_id;
        state.status = Status::Pending;
        state.owner = ctx.accounts.owner.key();
        state.created_at = Clock::get()?.unix_timestamp;
        state.closable_at = state.created_at + ctx.accounts.inbox_state.close_buffer_secs;
        state.bump = ctx.bumps.order_state;
        state.deposit_amount = params.deposit_amount;
        state.deposit_mint = ctx.accounts.mint_account.key();
        state.dest_chain_id = params.dest_chain_id;
        state.dest_call = params.call.clone();
        state.dest_expense = params.expense.clone();
        state.fill_hash = hash_fill(
            order_id,
            ctx.accounts.inbox_state.chain_id,
            state.dest_chain_id,
            state.closable_at,
            state.dest_call.clone(),
            state.dest_expense.clone(),
        );

        emit!(EventUpdated {
            order_id: state.order_id,
            status: state.status.clone(),
        });

        Ok(())
    }

    /// Reject an order, refunding owner closing accounts.
    /// Only admin can reject orders.
    pub fn reject(ctx: Context<Reject>, order_id: Pubkey, reason: u8) -> Result<()> {
        let state = &mut ctx.accounts.order_state;
        state.status = Status::Rejected;
        state.reject_reason = reason;

        // Sign transfer and close_account with order token PDA
        let order_token_seeds: &[&[&[u8]]] = &[&[
            ORDER_TOKEN_SEED_PREFIX,
            order_id.as_ref(),
            &[ctx.bumps.order_token_account],
        ]];

        // Transfer the deposit to the owner account
        transfer(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                Transfer {
                    from: ctx.accounts.order_token_account.to_account_info(),
                    to: ctx.accounts.owner_token_account.to_account_info(),
                    authority: ctx.accounts.order_token_account.to_account_info(),
                },
            )
            .with_signer(order_token_seeds),
            state.deposit_amount,
        )?;

        // Close the order token account, returning rent to owner
        close_account(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                CloseAccount {
                    account: ctx.accounts.order_token_account.to_account_info(),
                    destination: ctx.accounts.owner_token_account.to_account_info(),
                    authority: ctx.accounts.order_token_account.to_account_info(),
                },
            )
            .with_signer(order_token_seeds),
        )?;

        emit!(EventUpdated {
            order_id: state.order_id,
            status: state.status.clone(),
        });

        Ok(())
    }

    /// Mark an order as filled, and set the claimable_by account.
    /// This may only be called by the inbox admin.
    pub fn mark_filled(
        ctx: Context<MarkFilled>,
        _order_id: Pubkey,
        fill_hash: Pubkey,
        claimable_by: Pubkey,
    ) -> Result<()> {
        let state = &mut ctx.accounts.order_state;
        require_eq!(state.fill_hash, fill_hash, InboxError::InvalidFillHash);

        state.status = Status::Filled;
        state.claimable_by = claimable_by;

        emit!(EventUpdated {
            order_id: state.order_id,
            status: state.status.clone(),
        });

        Ok(())
    }

    // Claims the deposit of an filled order
    pub fn claim(ctx: Context<Claim>, order_id: Pubkey) -> Result<()> {
        let state = &mut ctx.accounts.order_state;
        state.status = Status::Claimed;

        // Sign transfer with order token PDA
        let order_token_seeds: &[&[&[u8]]] = &[&[
            ORDER_TOKEN_SEED_PREFIX,
            order_id.as_ref(),
            &[ctx.bumps.order_token_account],
        ]];

        // Transfer the deposit to the claimer account
        transfer(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                Transfer {
                    from: ctx.accounts.order_token_account.to_account_info(),
                    to: ctx.accounts.claimer_token_account.to_account_info(),
                    authority: ctx.accounts.order_token_account.to_account_info(),
                },
            )
            .with_signer(order_token_seeds),
            state.deposit_amount,
        )?;

        // Close the order token account and rent to owner
        close_account(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                CloseAccount {
                    account: ctx.accounts.order_token_account.to_account_info(),
                    destination: ctx.accounts.owner_token_account.to_account_info(),
                    authority: ctx.accounts.order_token_account.to_account_info(),
                },
            )
            .with_signer(order_token_seeds),
        )?;

        emit!(EventUpdated {
            order_id: state.order_id,
            status: state.status.clone(),
        });

        Ok(())
    }

    // Close an order
    pub fn close(ctx: Context<Close>, order_id: Pubkey) -> Result<()> {
        let state = &mut ctx.accounts.order_state;
        require_gte!(
            Clock::get()?.unix_timestamp,
            state.closable_at,
            InboxError::NotClosable
        );

        state.status = Status::Closed;

        // Sign transfer and close_account with order token PDA
        let order_token_seeds: &[&[&[u8]]] = &[&[
            ORDER_TOKEN_SEED_PREFIX,
            order_id.as_ref(),
            &[ctx.bumps.order_token_account],
        ]];

        // Transfer the deposit to the owner account
        transfer(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                Transfer {
                    from: ctx.accounts.order_token_account.to_account_info(),
                    to: ctx.accounts.owner_token_account.to_account_info(),
                    authority: ctx.accounts.order_token_account.to_account_info(),
                },
            )
            .with_signer(order_token_seeds),
            state.deposit_amount,
        )?;

        // Close the order token account, returning rent to owner
        close_account(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                CloseAccount {
                    account: ctx.accounts.order_token_account.to_account_info(),
                    destination: ctx.accounts.owner_token_account.to_account_info(),
                    authority: ctx.accounts.order_token_account.to_account_info(),
                },
            )
            .with_signer(order_token_seeds),
        )?;

        emit!(EventUpdated {
            order_id: state.order_id,
            status: state.status.clone(),
        });

        Ok(())
    }
}

#[derive(Accounts)]
pub struct Init<'info> {
    #[account(
        init,
        payer = admin,
        space = 8 + InboxState::INIT_SPACE,
        seeds = [InboxState::SEED_PREFIX],
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
        space = 8 + OrderState::INIT_SPACE,
        seeds = [OrderState::SEED_PREFIX, params.order_id.as_ref()],
        bump,
    )]
    pub order_state: Account<'info, OrderState>,
    #[account(mut)]
    pub owner: Signer<'info>,

    // Deposit token transfer
    // ======================
    // The mint (token) being transferred
    #[account(mut)]
    pub mint_account: Account<'info, Mint>,
    // The owner's token account
    #[account(
        mut,
        token::mint = mint_account,
        token::authority = owner,
    )]
    pub owner_token_account: Account<'info, TokenAccount>,
    // The order's (escrow) token account.
    #[account(
        init,
        payer = owner,
        seeds = [ORDER_TOKEN_SEED_PREFIX, params.order_id.as_ref()],
        bump,
        token::mint = mint_account,
        token::authority = order_token_account,
    )]
    pub order_token_account: Account<'info, TokenAccount>,

    // The global token program
    pub token_program: Program<'info, Token>,
    #[account(
        seeds = [InboxState::SEED_PREFIX],
        bump,
    )]
    pub inbox_state: Account<'info, InboxState>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(order_id: Pubkey)]
pub struct Reject<'info> {
    #[account(
        mut,
        seeds = [OrderState::SEED_PREFIX, order_id.as_ref()],
        bump,
        constraint = order_state.status == Status::Pending,
    )]
    pub order_state: Account<'info, OrderState>,
    #[account(
        mut,
        seeds = [ORDER_TOKEN_SEED_PREFIX, order_id.as_ref()],
        bump,
        token::mint = order_state.deposit_mint,
        token::authority = order_token_account,
    )]
    pub order_token_account: Account<'info, TokenAccount>,
    #[account(
        mut,
        token::mint = order_state.deposit_mint,
        token::authority = order_state.owner,
    )]
    pub owner_token_account: Account<'info, TokenAccount>,
    #[account(
        seeds = [InboxState::SEED_PREFIX],
        bump,
        constraint = inbox_state.admin == admin.key(),
    )]
    pub inbox_state: Account<'info, InboxState>,
    #[account(mut)]
    pub admin: Signer<'info>,
    pub token_program: Program<'info, Token>,
}

#[derive(Accounts)]
#[instruction(order_id: Pubkey)]
pub struct MarkFilled<'info> {
    #[account(
        mut,
        seeds = [OrderState::SEED_PREFIX, order_id.as_ref()],
        bump,
        constraint = order_state.status == Status::Pending,
    )]
    pub order_state: Account<'info, OrderState>,
    #[account(
        seeds = [InboxState::SEED_PREFIX],
        bump,
        constraint = inbox_state.admin == admin.key(),
    )]
    pub inbox_state: Account<'info, InboxState>,
    #[account(mut)]
    pub admin: Signer<'info>,
}

#[derive(Accounts)]
#[instruction(order_id: Pubkey)]
pub struct Claim<'info> {
    #[account(
        mut,
        seeds = [OrderState::SEED_PREFIX, order_id.as_ref()],
        bump,
        constraint = order_state.claimable_by == claimer.key(),
        constraint = order_state.status == Status::Filled,
    )]
    pub order_state: Account<'info, OrderState>,
    #[account(
        mut,
        seeds = [ORDER_TOKEN_SEED_PREFIX, order_id.as_ref()],
        bump,
        token::mint = order_state.deposit_mint,
        token::authority = order_token_account,
    )]
    pub order_token_account: Account<'info, TokenAccount>,
    #[account(
        mut,
        token::mint = order_state.deposit_mint,
        token::authority = order_state.owner,
    )]
    pub owner_token_account: Account<'info, TokenAccount>,
    #[account(mut)]
    pub claimer: Signer<'info>,
    #[account(
        mut,
        token::mint = order_state.deposit_mint,
        token::authority = claimer,
    )]
    pub claimer_token_account: Account<'info, TokenAccount>,
    pub token_program: Program<'info, Token>,
}

#[derive(Accounts)]
#[instruction(order_id: Pubkey)]
pub struct Close<'info> {
    #[account(
        mut,
        seeds = [OrderState::SEED_PREFIX, order_id.as_ref()],
        bump,
        constraint = order_state.owner == owner.key(),
    constraint = order_state.status == Status::Pending,
    )]
    pub order_state: Account<'info, OrderState>,
    #[account(
        mut,
        seeds = [ORDER_TOKEN_SEED_PREFIX, order_id.as_ref()],
        bump,
        token::mint = order_state.deposit_mint,
        token::authority = order_token_account,
    )]
    pub order_token_account: Account<'info, TokenAccount>,
    #[account(
        mut,
        token::mint = order_state.deposit_mint,
        token::authority = owner,
    )]
    pub owner_token_account: Account<'info, TokenAccount>,

    #[account(mut)]
    pub owner: Signer<'info>,
    pub token_program: Program<'info, Token>,
}
