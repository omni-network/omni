use crate::types::{EVMCall, EVMTokenExpense};
use anchor_lang::prelude::*;
use ethabi_solana::*;
use sha2::Digest;

// orderID returns the order ID (32 byte array Pubkey),
//  by hashing the account (pubkey) and nonce.
pub fn order_id(account: Pubkey, nonce: u64) -> Pubkey {
    let mut hasher = sha2::Sha256::new();
    hasher.update(account.to_bytes());
    hasher.update(nonce.to_le_bytes());

    Pubkey::new_from_array(hasher.finalize().into())
}

/// returns the order fill hash (32 byte array Pubkey).
pub fn hash_fill(
    order_id: Pubkey,
    src_chain_id: u64,
    dest_chain_id: u64,
    closable_at: i64,
    dest_call: EVMCall,
    dest_expense: EVMTokenExpense,
) -> Pubkey {
    // abi encode `[orderId, struct FillOriginData]` equivalent
    let encoded = encode(&[
        Token::FixedBytes(order_id.to_bytes().to_vec()), // bytes32 orderId
        Token::Tuple(vec![
            Token::Uint(src_chain_id.into()),  // uint64 srcChainId
            Token::Uint(dest_chain_id.into()), // uint64 destChainId
            Token::Uint(closable_at.into()),   // uint32 fillDeadline
            Token::Array(vec![Token::Tuple(vec![
                Token::Address(Address::from(dest_call.target)), // address target
                Token::FixedBytes(dest_call.selector.to_vec()),  // bytes4 selector
                Token::Uint(Uint::from(dest_call.value)),        // uint256 value
                Token::Bytes(dest_call.params),                  // bytes params
            ])]), // Call[] calls
            Token::Array(vec![Token::Tuple(vec![
                Token::Address(Address::from(dest_expense.spender)), // address spender
                Token::Address(Address::from(dest_expense.token)),   // address token
                Token::Uint(Uint::from(dest_expense.amount)),        // uint96 amount
            ])]), // TokenExpense[] expenses;
        ]),
    ]);

    // hash the encoded data
    let mut hasher = sha2::Sha256::new();
    hasher.update(encoded);

    // return the hash as a Pubkey
    Pubkey::new_from_array(hasher.finalize().into())
}
